package middleware

import (
	"encoding/csv"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
)

func CsvParserMiddleware[T interface{}](fieldNmae string) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileBin, getFileErr := c.FormFile(fieldNmae)
		if getFileErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": getFileErr.Error()})
			c.Abort()
			return
		}

		file, openFileBinErr := fileBin.Open()
		if openFileBinErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": openFileBinErr.Error()})
			c.Abort()
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		columns, readColErr := reader.Read()
		if readColErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": readColErr.Error()})
			c.Abort()
			return
		}

		v := validator.New()
		res := make([]interface{}, 0)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			row := make(map[string]interface{})
			for i, value := range record {
				row[columns[i]] = value
			}

			var data T
			if err := mapstructure.Decode(row, &data); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			if err := v.Struct(&data); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				c.Abort()
				return
			}

			res = append(res, data)
		}

		c.Set("csv.content", res)
		c.Next()
	}
}
