package convert

import (
	"fmt"
	"strings"

	"github.com/mztlive/go-pkgs/reflect_utils"
	"github.com/samber/lo"
)

func ConvertMapToSetStr(fields map[string]interface{}) (string, []interface{}) {
	var setStr string
	var args []interface{} = make([]interface{}, 0, len(fields))

	for k, v := range fields {
		setStr += k + "=" + "?,"
		args = append(args, v)
	}

	return setStr[:len(setStr)-1], args
}

// ConvertEntityToUpdateStmt converts entity to update statement
// e.g. UPDATE table SET field1=?, field2=? WHERE id=?
//
// ignoreFields: id, created_at, updated_at, deleted_at
func ConvertEntityToUpdateStmt(entity interface{}) (string, []interface{}) {
	var (
		table    string
		setStr   string
		whereStr string
		args     []interface{}
	)

	ignoreFields := []string{"id", "created_at", "updated_at", "deleted_at"}

	table = reflect_utils.GetSnakeNameFromStruct(entity)
	fieldsMap := reflect_utils.StructToMapStr(entity)

	// 构造SET部分
	setStrings := []string{}
	for k, v := range fieldsMap {
		if !lo.Contains(ignoreFields, k) {
			setStrings = append(setStrings, "`"+k+"`"+"=?")
			args = append(args, v)
		}
	}
	setStr = strings.Join(setStrings, ", ")

	// 构造WHERE部分
	id, ok := fieldsMap["id"]
	if ok {
		whereStr = "WHERE id=?"
		args = append(args, id)
	}

	sql := fmt.Sprintf("UPDATE %s SET %s %s", table, setStr, whereStr)
	return sql, args
}
