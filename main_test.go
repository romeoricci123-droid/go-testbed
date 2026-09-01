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

func TestFormatGreetingUppercase(t *testing.T) {
	got := FormatGreeting(true)
	want := "HELLO FROM GO-TESTBED"
	if got != want {
		t.Errorf("FormatGreeting(true) = %q, want %q", got, want)
	}

	got = FormatGreeting(false)
	want = "hello from go-testbed"
	if got != want {
		t.Errorf("FormatGreeting(false) = %q, want %q", got, want)
	}
}
