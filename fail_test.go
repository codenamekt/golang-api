package main

import (
	"testing"
)

// All these tests will be failing.
func TestDocFail(t *testing.T) {

	teststeps := []teststep{
		// Invalid Id
		{
			Url:      "http://localhost/foo/bar",
			Method:   "POST",
			Body:     "{\"_id\":\"22\",\"password\":\"xyz\",\"username\":\"xyz\"}",
			TestCode: 400,
			TestHeader: map[string]string{
				"Content-Length": "[22]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "{\"error\":\"Invalid id\"}",
		},
		// Invalid Id
		{
			Url:      "http://localhost/foo/bar/40000",
			Method:   "GET",
			Body:     "",
			TestCode: 400,
			TestHeader: map[string]string{
				"Content-Length": "[22]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "{\"error\":\"Invalid id\"}",
		},
		// Body invalid json
		{
			Url:      "http://localhost/foo/bar",
			Method:   "POST",
			Body:     "{\"_id\":\"22\",\"passwo\"xyz\",\"username\":\"xyz\"}",
			TestCode: 500,
			TestHeader: map[string]string{
				"Content-Length": "[29]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "{\"error\":\"Body invalid json\"}",
		},
	}

	for _, test := range teststeps {
		testRunner(test, t)
	}
}
