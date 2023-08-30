package reflect_utils

import (
	"reflect"
	"testing"
	"time"
)

func TestGetNameFromStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}
	name := GetNameFromStruct(p)

	if name != "Person" {
		t.Errorf("Expected name to be 'Person', but got '%s'", name)
	}
}

func TestGetSnakeNameFromStruct(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	p := Person{Name: "John", Age: 30}
	name := GetSnakeNameFromStruct(p)

	if name != "person" {
		t.Errorf("Expected name to be 'person', but got '%s'", name)
	}
}

type TestStruct struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at" db:"updated_at"`

	Int  int  `json:"int" db:"int"`
	IntP *int `json:"int_p" db:"int_p"`

	Float32  float32  `json:"float32" db:"float32"`
	Float32P *float32 `json:"float32_p" db:"float32_p"`

	Bool  bool  `json:"bool" db:"bool"`
	BoolP *bool `json:"bool_p" db:"bool_p"`

	Str  string  `json:"str" db:"str"`
	StrP *string `json:"str_p" db:"str_p"`
}

func TestStructToMapStr(t *testing.T) {

	now := time.Now()
	future := now.Add(24 * time.Hour)

	obj := TestStruct{
		ID:        1,
		Name:      "test",
		CreatedAt: now,
		UpdatedAt: &future,

		Int:  2,
		IntP: nil,

		Float32:  1.5,
		Float32P: nil,

		Bool:  true,
		BoolP: nil,

		Str:  "test",
		StrP: nil,
	}

	expected := map[string]string{
		"id":         "1",
		"name":       "test",
		"created_at": now.Format("2006-01-02 15:04:05"),
		"updated_at": future.Format("2006-01-02 15:04:05"),

		"int":     "2",
		"float32": "1.5",
		"bool":    "1",
		"str":     "test",
	}

	result := StructToMapStr(obj)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
