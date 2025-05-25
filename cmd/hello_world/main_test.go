package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMain(t *testing.T) {
	// This is a simple test to ensure the package builds
	t.Run("package builds", func(t *testing.T) {
		// If we get here, the package builds successfully
	})
}

func TestPrintHello(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintHello()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("expected %q, got %q", expected, output)
	}
}
