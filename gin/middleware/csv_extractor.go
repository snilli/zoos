package middleware

import (
	"encoding/csv"
	"io"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/mitchellh/mapstructure"
)

func CsvParserMiddleware[T interface{}](formDataFieldName string, shape T) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileBin, getFileErr := c.FormFile(formDataFieldName)
		if getFileErr != nil {
			c.Set("csv.error", gin.H{"error": getFileErr.Error()})
			c.Next()
			return
		}

		file, openFileBinErr := fileBin.Open()
		if openFileBinErr != nil {
			c.Set("csv.error", gin.H{"error": openFileBinErr.Error()})
			c.Next()
			return
		}
		defer file.Close()

		reader := csv.NewReader(file)
		columns, readColErr := reader.Read()
		if readColErr != nil {
			c.Set("csv.error", gin.H{"error": readColErr.Error()})
			c.Next()
			return
		}

		v := validator.New()
		res := make([]interface{}, 0)

		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				c.Set("csv.error", gin.H{"error": err.Error()})
				c.Next()
				return
			}

			row := make(map[string]interface{})
			for i, value := range record {
				row[columns[i]] = value
			}

			var data T
			if err := mapstructure.Decode(row, &data); err != nil {
				c.Set("csv.error", gin.H{"error": err.Error()})
				c.Next()
				return
			}

			if err := v.Struct(&data); err != nil {
				c.Set("csv.error", gin.H{"error": err.Error()})
				c.Next()
				return
			}

			res = append(res, data)
		}

		c.Set("csv.content", res)
		c.Next()
	}
}
