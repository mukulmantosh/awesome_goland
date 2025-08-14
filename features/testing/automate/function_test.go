package automate

import (
	"testing"
)

func TestEmployee_Greet(t *testing.T) {
	type fields struct {
		Name   string
		Age    int
		Salary int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{
			name:   "Test1",
			fields: fields{Name: "Bob", Age: 22},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Employee{
				Name:   tt.fields.Name,
				Age:    tt.fields.Age,
				Salary: tt.fields.Salary,
			}
			expected := "Hello Bob"
			result := e.Greet()
			if expected != result {
				t.Errorf("Expected: %s || Got: %s", expected, result)
			}
		})
	}
}
