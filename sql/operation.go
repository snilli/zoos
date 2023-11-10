package operation

import (
	"fmt"
	"strings"
)

func In(ref []string) string {
	return fmt.Sprintf("(%s)", strings.Join(ref, ","))
}
