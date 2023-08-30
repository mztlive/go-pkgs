package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type EntityWithTags struct {
	Name  string `db:"name"`
	Email string `db:"email"`
	ID    int    `db:"id"`
}

type EntityWithoutTags struct {
	Name  string
	Email string
}

type EntityWithNonDbTags struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type IgnoredFieldsEntity struct {
	Name      string `db:"name"`
	Email     string `db:"email"`
	ID        int    `db:"id"`
	DeletedAt string `db:"deleted_at"`
	UpdatedAt string `db:"updated_at"`
	CreatedAt string `db:"created_at"`
}

type EntityWithUnexportedField struct {
	Name     string `db:"name"`
	Email    string `db:"email"`
	userName string `db:"username"`
}

func TestConvertEntityToInsertNamedStmt(t *testing.T) {
	// Test case with all required tags
	{
		entity := &EntityWithTags{Name: "test", Email: "test@test.com", ID: 1}
		stmt, err := ConvertEntityToInsertNamedStmt(entity)
		assert.Nil(t, err)
		assert.Equal(t, "INSERT INTO entity_with_tags (name, email) VALUES (:name, :email)", stmt)
	}

	// Test case without tags
	{
		entity := &EntityWithoutTags{Name: "test", Email: "test@test.com"}
		stmt, err := ConvertEntityToInsertNamedStmt(entity)
		assert.NotNil(t, err)
		assert.Equal(t, "", stmt)
	}

	// Test case with non-db tags
	{
		entity := &EntityWithNonDbTags{Name: "test", Email: "test@test.com"}
		stmt, err := ConvertEntityToInsertNamedStmt(entity)
		assert.NotNil(t, err)
		assert.Equal(t, "", stmt)
	}

	// Test case with non-struct type
	{
		entity := "NonStructType"
		stmt, err := ConvertEntityToInsertNamedStmt(&entity)
		assert.NotNil(t, err)
		assert.Equal(t, "", stmt)
	}

	// Test case with ignored fields
	{
		entity := &IgnoredFieldsEntity{Name: "test", Email: "test@test.com", ID: 1, DeletedAt: "2023-07-14", UpdatedAt: "2023-07-14", CreatedAt: "2023-07-14"}
		stmt, err := ConvertEntityToInsertNamedStmt(entity)
		assert.Nil(t, err)
		assert.Equal(t, "INSERT INTO ignored_fields_entity (name, email) VALUES (:name, :email)", stmt)
	}

	// Test case with unexported field
	{
		entity := &EntityWithUnexportedField{Name: "test", Email: "test@test.com", userName: "testusername"}
		stmt, err := ConvertEntityToInsertNamedStmt(entity)
		assert.Nil(t, err)
		assert.Equal(t, "INSERT INTO entity_with_unexported_field (name, email) VALUES (:name, :email)", stmt)
	}
}
