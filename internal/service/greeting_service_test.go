package service

import "testing"

func TestGreetingService_Greet(t *testing.T) {
	s := NewGreetingService()
	got := s.Greet("Codex")
	if got != "Hello, Codex!" {
		t.Errorf("unexpected greeting %q", got)
	}
	got = s.Greet("")
	if got != "Hello, World!" {
		t.Errorf("unexpected greeting for empty name: %q", got)
	}
}
