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
			Method:   "GET",
			Body:     "",
			TestCode: 200,
			TestHeader: map[string]string{
				"Content-Length": "[509]",
				"Content-Type":   "[application/json; charset=UTF-8]",
			},
			TestBody: "[{\"_id\":\"559768cca92da80f7e000002\",\"password\":\"xyz\",\"username\":\"xyz\"},{\"_id\":\"55977226a92da81351000001\",\"password\":\"batman\",\"username\":\"zzz\"},{\"_id\":\"5599a720a92da80ebb000001\",\"password\":\"batman\",\"username\":\"zzz\"},{\"_id\":\"5599ed80fdac67b86b370bc7\",\"password\":\"batman\",\"username\":\"zzz\"},{\"_id\":\"5599ee72a92da8138e000001\",\"password\":\"batman\",\"username\":\"robin\"},{\"_id\":\"5599f751a92da81755000001\",\"password\":\"batman\",\"username\":\"robin\"},{\"_id\":\"559768cca92da80f7e000004\",\"password\":\"batman5\",\"username\":\"robin\"}]",
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
