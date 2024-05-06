package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNotFoundHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/notfound", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    NotFoundHandler(rr, req)

    if status := rr.Code; status != http.StatusNotFound {
        t.Errorf("Handler returned wrong status code: got %v", status)
    }

    if rr.Body.String() != "No policy found." {
        t.Error("Response body is malformed")
    }
}
