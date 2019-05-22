package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoot(t *testing.T) {
	req, err := http.NewRequest("GET", "http://example.com/foo", nil)
	if err != nil {
		t.FailNow()
	}

	w := httptest.NewRecorder()
	HandleRoot(w, req)

	if w.Code != 200 {
		t.Fail()
	}
}
