package sql

import (
	"fmt"
	"strings"
)

type sqlOperations struct{}

// type sqlOperations interface {
// 	In(ref []string) string
// }

func NewSqlOperation() sqlOperations {
	return sqlOperations{}
}

func (op *sqlOperations) In(ref []string) string {
	return fmt.Sprintf("(%s)", strings.Join(ref, ","))
}
