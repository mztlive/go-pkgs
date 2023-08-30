package convert

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/mztlive/go-pkgs/string_utils"
)

// ConvertEntityToInsertNamedStmt 将实体转换为插入语句(Named方式)
func ConvertEntityToInsertNamedStmt(i interface{}) (string, error) {
	// Add type check to avoid panic
	val := reflect.ValueOf(i)
	if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
		return "", errors.New("input should be a pointer to struct")
	}
	val = val.Elem()

	typeOfEntity := val.Type()
	var columns, placeholders []string
	for i := 0; i < val.NumField(); i++ {
		fieldType := typeOfEntity.Field(i)
		// Skip unexported fields
		if fieldType.PkgPath != "" {
			continue
		}

		// Skip fields without a db tag
		dbTag := typeOfEntity.Field(i).Tag.Get("db")
		if dbTag == "" {
			continue
		}
		switch dbTag {
		case "id", "deleted_at", "updated_at", "created_at":
			continue
		}
		columns = append(columns, "`"+dbTag+"`")
		placeholders = append(placeholders, ":"+dbTag)
	}

	if len(columns) == 0 {
		return "", errors.New("no valid db tags found")
	}

	tableName := string_utils.ToSnakeCase(typeOfEntity.Name())
	stmt := fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES (%s)",
		tableName,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)
	return stmt, nil
}
