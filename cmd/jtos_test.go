package cmd

import (
	"strings"
	"testing"
)

var testCases = []map[string]string{
	{
		"case":   "Empty objects",
		"input":  `{}`,
		"output": `type AutoStruct struct {`,
	},
	{
		"case":   "Object with multiple data type",
		"input":  `{"test": "string", "check": true}`,
		"output": "Test string `json:\"test\"`",
	},
	{
		"case":   "Array of objects",
		"input":  `[{"check": true}]`,
		"output": "Check bool `json:\"check\"`",
	},
	{
		"case":   "Array of objects with nexted array objects",
		"input":  `[{"check": true, "my_field": {"test": [{"some_other":["some_thing"], "test2": "something"}], "test3": {"some": "value"}}}]`,
		"output": "[]map[string]interface {} `json:\"test\"`",
	},
}

func TestGenerateStruct(t *testing.T) {

	t.Log("Error on generate struct")
	{
		ip := "some string"
		errStr := "invalid json err: invalid character 's' looking for beginning of value"

		_, err := generateStruct(ip)
		if err == nil {
			t.Errorf("expected error got nil")
			return
		}
		if err.Error() != errStr {
			t.Errorf("expected error %s got %s", errStr, err)
		}

	}
	t.Log("Error on generate struct")
	{
		ip := `["test"]`
		errStr := "array is not a map but string"

		_, err := generateStruct(ip)
		if err == nil {
			t.Errorf("expected error got nil")
			return
		}
		if err.Error() != errStr {
			t.Errorf("expected error %s got %s", errStr, err)
		}

	}

	t.Log("Generate struct for non error case")
	{
		for _, c := range testCases {
			ip := c["input"]
			op := c["output"]

			r, err := generateStruct(ip)
			if err != nil {
				t.Errorf("unexpected error %v", err)
				continue
			}

			if !strings.Contains(r, op) {
				t.Errorf("expected %s  to have included %s", r, op)
			}
		}
	}
}
