package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPayments(t *testing.T) {
	req, _ := http.NewRequest("GET", "/payments", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPayments)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected %v got %v", http.StatusOK, status)
	}
}
