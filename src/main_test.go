package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestGreet(t *testing.T) {
	req, err := http.NewRequest("GET", "/greet", nil) //new get request to /greet with no body
	if err != nil {
		t.Fatal(err)
	}
	// response recorder
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(greet)

	//do request to greet function (which is in the handler object now)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned bad status code, expected 200 got %v", status)
	}

	expected := "hello world"
	if rr.Body.String() != expected {
		t.Errorf("got a wrong response, expected %v, got %v", expected, rr.Body.String())
	}
}


func TestFail(t *testing.T) {
	t.Fail()
}
