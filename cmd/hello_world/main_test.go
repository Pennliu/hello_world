package main

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"time"
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

	testStartTime := time.Now()
	PrintHello("Hello, World!")

	_ = w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)
	output := buf.String()

	// 1. Check for the greeting
	expectedGreeting := "Hello, World!"
	if !strings.Contains(output, expectedGreeting) {
		t.Errorf("expected output to contain %q, got %q", expectedGreeting, output)
	}

	// 2. Check for and validate the timestamp
	timePrefix := "Current time: "
	lines := strings.Split(strings.TrimSpace(output), "\n")
	var timeLine string
	for _, line := range lines {
		if strings.HasPrefix(line, timePrefix) {
			timeLine = line
			break
		}
	}

	if timeLine == "" {
		t.Fatalf("Could not find line with time prefix 'Current time: '")
	}

	timeStr := strings.TrimPrefix(timeLine, timePrefix)
	parsedTime, err := time.Parse(time.RFC1123, timeStr)
	if err != nil {
		t.Fatalf("failed to parse time string %q: %v", timeStr, err)
	}

	// Allow a small delta for execution delay.
	// The parsed time has second-level precision due to RFC1123, so it might be
	// slightly before the high-precision testStartTime if the test crosses a second boundary.
	// We allow a small negative difference to account for this truncation.
	diff := parsedTime.Sub(testStartTime)
	if diff < -1*time.Second || diff > 2*time.Second {
		t.Errorf("parsed time %v is not within the allowed range of test start time %v (difference: %v)", parsedTime, testStartTime, diff)
	}
}

func BenchmarkPrintHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PrintHello("Hello, World!")
	}
}
