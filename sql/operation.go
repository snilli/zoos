package operation

import (
	"fmt"
	"strconv"
	"strings"
)

func In(ref []string) string {
	var params = []string{}
	for index := range ref {
		params = append(params, fmt.Sprintf("$%s", strconv.Itoa(index+1)))
	}

	return fmt.Sprintf("(%s)", strings.Join(params, ","))
}
