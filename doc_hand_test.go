package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type test struct {
	Url        string
	Method     string
	Body       string
	TestCode   int
	TestHeader map[string]string
	TestBody   string
}

func TestDoc(t *testing.T) {

	tests := []test{
		{
			Url:      "http://localhost/todo/todo",
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
			Url:      "http://localhost/todo/todo",
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
			Url:      "http://localhost/todo/todo/559768cca92da80f7e000002",
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
			Url:      "http://localhost/todo/todo/559768cca92da80f7e000002",
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

	for _, test := range tests {
		testbody := bytes.NewBufferString(test.Body)
		var body []byte
		router := NewRouter()
		resp := httptest.NewRecorder()
		var req *http.Request

		req, err := http.NewRequest(test.Method, test.Url, testbody)
		if err != nil {
			t.Fail()
		}

		router.ServeHTTP(resp, req)
		if err == nil {
			body, err = ioutil.ReadAll(resp.Body)
		}

		if resp.Code == test.TestCode {
			t.Logf("Code: %s", strconv.Itoa(resp.Code))
		} else {
			t.Fail()
		}

		for key, value := range resp.Header() {
			testvalue := strings.Join(value, ",")
			testvalue = strings.Join([]string{"[", testvalue, "]"}, "")

			if i, ok := test.TestHeader[key]; ok {
				if i == testvalue {
					t.Logf("%s: %s", key, value)
				} else {
					t.Errorf(i, "not equal", testvalue)
				}
			} else {
				t.Fail()
			}
		}

		if string(body) == test.TestBody {
			t.Logf("Body: %s", body)
		} else {
			t.Errorf(string(body), "not equal", test.Body)
		}
	}
}
