package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}

	Greet(&buffer, "beb")

	got := buffer.String()
	want := "Hello beb"

	if got != want {
		t.Errorf("got %s wat %s", got, want)
	}
}
