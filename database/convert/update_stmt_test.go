// BEGIN: 9f8c7e6d4b3a
package convert

import (
	"reflect"
	"testing"
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
}

func TestConvertEntityToUpdateStmt(t *testing.T) {

	user := User{
		ID:        1,
		Name:      "John",
		Age:       30,
		CreatedAt: time.Now(),
	}

	expectedSQL := "UPDATE user SET name=?, age=? WHERE id=?"

	sql, args := ConvertEntityToUpdateStmt(user)

	if sql != expectedSQL {
		t.Errorf("Expected sql: %q, got %q", expectedSQL, sql)
	}

	expectedArgs := []interface{}{"John", "30", "1"}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("Expected args: %v, got %v", expectedArgs, args)
	}
}
