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

type teststep struct {
	Url        string
	Method     string
	Body       string
	TestCode   int
	TestHeader map[string]string
	TestBody   string
}

func testRunner(test teststep, t *testing.T) {
	testbody := bytes.NewBufferString(test.Body)
	var body []byte
	router := NewRouter()
	resp := httptest.NewRecorder()
	var req *http.Request

	t.Logf("%s %s", test.Method, test.Url)

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
		t.Errorf(strconv.Itoa(resp.Code), "not equal", strconv.Itoa(test.TestCode))
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
			t.Errorf("key %s not found", key)
		}
	}
	if string(body) == test.TestBody {
		t.Logf("Body: %s", body)
	} else {
		t.Errorf(string(body), "not equal", test.Body)
	}
}
