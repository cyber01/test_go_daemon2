package controllers

import (
	"github.com/dabfleming/go-martini-example/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOne(t *testing.T) {
	recorder := httptest.NewRecorder()

	req, err := http.NewRequest("GET", "/user/one", nil)
	if err != nil {
		t.Fatalf("NewRequest returned error: %v", err)
	}

	server.Server.ServeHTTP(recorder, req)

	if recorder.Code != 200 {
		t.Errorf("Expected response code 200, got %v", recorder.Code)
	}
	if recorder.Body.String() != "one" {
		t.Errorf("Unexpected body, got `%v`", recorder.Body.String())
	}
}
