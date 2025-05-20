package http

import (
	"net/http/httptest"
	"testing"

	"hexademo/internal/service"
)

func TestHandler_ServeHTTP(t *testing.T) {
	h := NewHandler(service.NewGreetingService())

	req := httptest.NewRequest("GET", "/?name=Codex", nil)
	rr := httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if body := rr.Body.String(); body != "Hello, Codex!" {
		t.Errorf("unexpected body %q", body)
	}

	req = httptest.NewRequest("GET", "/", nil)
	rr = httptest.NewRecorder()
	h.ServeHTTP(rr, req)
	if body := rr.Body.String(); body != "Hello, World!" {
		t.Errorf("unexpected body for empty name %q", body)
	}
}
