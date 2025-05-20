package main

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain_StartsServer(t *testing.T) {
	called := false
	originalLS := listenAndServe
	originalLF := logFatal
	defer func() {
		listenAndServe = originalLS
		logFatal = originalLF
	}()
	listenAndServe = func(addr string, h http.Handler) error {
		called = true
		if addr != ":8080" {
			t.Errorf("expected address :8080, got %s", addr)
		}
		// invoke handler once to ensure it works
		rr := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/?name=Test", nil)
		h.ServeHTTP(rr, req)
		if rr.Body.String() != "Hello, Test!" {
			t.Errorf("unexpected handler output %q", rr.Body.String())
		}
		return io.EOF
	}
	fatalCalled := false
	logFatal = func(v ...interface{}) {
		fatalCalled = true
	}
	log.SetOutput(io.Discard)
	main()
	if !called {
		t.Error("listenAndServe was not called")
	}
	if !fatalCalled {
		t.Error("logFatal should be called on error")
	}
}
