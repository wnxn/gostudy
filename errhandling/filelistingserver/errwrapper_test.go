package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"os"
	"errors"
)

func errPanic(_ http.ResponseWriter,
	_ *http.Request) error {
	panic(123)
}

type testingUserError string

func (t testingUserError) Error() string {
	return t.Message()
}

func (t testingUserError) Message() string {
	return string(t)
}

func errUser(_ http.ResponseWriter,
	_ *http.Request) error {
	return testingUserError("User Error")
}

func errNotExist(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrNotExist
}

func errForbidden(_ http.ResponseWriter,
	_ *http.Request) error {
	return os.ErrPermission
}

func errUnknown(_ http.ResponseWriter,
	_ *http.Request) error {
	return errors.New("unknown error")
}

func errNo(writer http.ResponseWriter,
	_ *http.Request) error {
		fmt.Fprint(writer, "no error")
	return nil
}

var tests = []struct {
	h       appHandler
	code    int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUser, 400, "User Error"},
	{errNotExist, 404, "Not Found"},
	{errForbidden, 403, "Forbidden"},
	{errUnknown, 500, "Internal Server Error"},
	{errNo, 200, "no error"},
}

func TestErrWrapper(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.imooc.com",
			nil,
		)
		f(response, request)
		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

func TestErrWrapperInServer(t *testing.T){
	for _,tt := range tests{
		f:=errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		resp, _ := http.Get(server.URL)
		verifyResponse(resp, tt.code, tt.message, t)
	}
}

func verifyResponse(
	resp *http.Response,
	expectedCode int,
	expectedMsg string,
	t *testing.T) {
	b, _ := ioutil.ReadAll(resp.Body)
	body := strings.Trim(string(b), "\n")
	if resp.StatusCode != expectedCode || body != expectedMsg {
		t.Errorf("expect (%d, %s); "+
			"got (%d, %s)",
			expectedCode, expectedMsg,
			resp.StatusCode, body)
	}
}
