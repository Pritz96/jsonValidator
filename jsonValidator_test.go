package main

import (
	"encoding/json"
	"testing"
)

func TestValidateJSON(t *testing.T) {
	tests := []struct {
		json         string
		expectedBool bool
	}{
		{
			json: `{
				"name":"John",
				"age":30,
				"cars": {
				  "car1":"Ford",
				  "car2":"BMW",
				  "car3":"Fiat"
				}
			   } `,
			expectedBool: true,
		},
		{
			json: `
				"name":"John",
				"age":30,
				"cars": {
				  "car1":"Ford",
				  "car2":"BMW",
				  "car3":"Fiat"
				}
			   } `,
			expectedBool: false,
		},
		{
			json:         `{}`,
			expectedBool: true,
		},
		{
			json: `{
				"name":"John",
				"age":30,
				"cars": {
				  "car1":"Ford",
				  "car2":"BMW",
				  "car3":["Fiat"
				}
			   } `,
			expectedBool: false,
		},
		{
			json: `{
				"name":"John",
				"age":30,
				"cars": {
				  "car1":"Ford",
				  "car2":"BMW",
				  "car3":["Fiat","Bugatti"]
				}
			   } `,
			expectedBool: true,
		},
		{
			json:         `{{}}`,
			expectedBool: false,
		},
		{
			json:         `{1:2}`,
			expectedBool: false,
		},
		{
			json:         `{1:"2"}`,
			expectedBool: false,
		},
		{
			json:         `{"1":"2"}`,
			expectedBool: true,
		},
	}
	for _, test := range tests {
		actualBool := json.Valid([]byte(test.json))
		if actualBool != test.expectedBool {
			t.Errorf("Actual bool %v , expected bool %v", actualBool, test.expectedBool)
		}
	}
}
