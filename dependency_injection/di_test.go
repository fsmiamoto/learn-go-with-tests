package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Anakin")

	got := buf.String()
	want := "Hello, Anakin"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
