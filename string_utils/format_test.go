package string_utils

import "testing"

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"HelloWorld", "hello_world"},
		{"Members", "members"},
		{"Hello", "hello"},
		{"SQLStatement", "sql_statement"},
		{"longCamelCaseWithSeveralWords", "long_camel_case_with_several_words"},
	}

	for _, c := range cases {
		result := ToSnakeCase(c.input)
		if result != c.expected {
			t.Errorf("toSnakeCase(%s) == %s, expected %s", c.input, result, c.expected)
		}
	}
}
