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

// All these tests will be failing.
func TestDocFail(t *testing.T) {

	tests := []test{
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
