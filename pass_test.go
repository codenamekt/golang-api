package main

import (
	"testing"
)

// All these tests will be passing.
func TestDocPass(t *testing.T) {

	teststeps := []teststep{
		{
			Url:      "http://localhost/foo/bar",
			Method:   "POST",
			Body:     "{\"_id\":\"559768cca92da80f7e000002\",\"password\":\"xyz\",\"username\":\"xyz\"}",
			TestCode: 201,
			TestHeader: map[string]string{
				"Content-Length": "[68]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "{\"_id\":\"559768cca92da80f7e000002\",\"password\":\"xyz\",\"username\":\"xyz\"}",
		},
		{
			Url:      "http://localhost/",
			Method:   "GET",
			Body:     "",
			TestCode: 200,
			TestHeader: map[string]string{
				"Content-Length": "[16]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "[foo,local,user]",
		},
		{
			Url:      "http://localhost/foo",
			Method:   "GET",
			Body:     "",
			TestCode: 200,
			TestHeader: map[string]string{
				"Content-Length": "[20]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "[bar,system.indexes]",
		},
		{
			Url:      "http://localhost/foo/bar",
			Method:   "GET",
			Body:     "",
			TestCode: 200,
			TestHeader: map[string]string{
				"Content-Length": "[70]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "[{\"_id\":\"559768cca92da80f7e000002\",\"password\":\"xyz\",\"username\":\"xyz\"}]",
		},
		{
			Url:      "http://localhost/foo/bar/559768cca92da80f7e000002",
			Method:   "GET",
			Body:     "",
			TestCode: 200,
			TestHeader: map[string]string{
				"Content-Length": "[68]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "{\"_id\":\"559768cca92da80f7e000002\",\"password\":\"xyz\",\"username\":\"xyz\"}",
		},
		{
			Url:      "http://localhost/foo/bar/559768cca92da80f7e000002",
			Method:   "DELETE",
			Body:     "",
			TestCode: 204,
			TestHeader: map[string]string{
				"Content-Length": "[0]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "",
		},
	}

	for _, test := range teststeps {
		testRunner(test, t)
	}
}
