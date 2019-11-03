package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFormHandler(t *testing.T) {
	tests := []struct {
		method         string
		expectedStatus int
	}{
		{
			method:         "GET",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			method:         "POST",
			expectedStatus: http.StatusOK,
		},
		{
			method:         "DELETE",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}
	for _, test := range tests {
		req, err := http.NewRequest(test.method, "/formHandler", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		handler := http.HandlerFunc(formHandler)
		handler.ServeHTTP(resp, req)

		status := resp.Code
		if status != test.expectedStatus {
			t.Errorf("%s method returned wrong status code: got %v want %v",
				test.method, status, test.expectedStatus)
		}

		// expected := test.expectedResponse
		// if resp.Body.String() != expected {
		// 	t.Errorf("handler returned unexpected body: got %v want %v",
		// 		resp.Body.String(), expected)
		// }
	}
}

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
