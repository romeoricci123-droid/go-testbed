package main

import "testing"

func TestGreeting(t *testing.T) {
	got := Greeting()
	want := "hello from go-testbed"
	if got != want {
		t.Errorf("Greeting() = %q, want %q", got, want)
	}
}

func TestVersion(t *testing.T) {
	if Version != "0.1.0" {
		t.Errorf("Version = %q, want %q", Version, "0.1.0")
	}
}

func TestNewIDLength(t *testing.T) {
	id := NewID()
	if len(id) != 36 {
		t.Errorf("NewID() length = %d, want 36", len(id))
	}
}
