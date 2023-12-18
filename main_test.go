package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHomePage(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }
}

func TestSetAndGetKey(t *testing.T) {
    originalValue := "test_value"
    key, err := setKey(originalValue)
    if err != nil {
        t.Fatal(err)
    }

    retrievedValue, err := getKey(key)
    if err != nil {
        t.Fatal(err)
    }

    if retrievedValue != originalValue {
        t.Errorf("setKey and getKey returned different values: set %v, got %v", originalValue, retrievedValue)
    }
}

func TestIsUrl(t *testing.T) {
    testUrl := "http://example.com"
    if !isUrl(testUrl) {
        t.Errorf("isUrl failed to recognize a valid URL: %v", testUrl)
    }
}

//go test -v 