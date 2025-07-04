package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainWithoutArguments(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Save original args and set no arguments
	oldArgs := os.Args
	os.Args = []string{"hello-world"}

	// Run main function
	main()

	// Restore stdout and args
	w.Close()
	os.Stdout = old
	os.Args = oldArgs

	// Read captured output
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestMainWithArgument(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Save original args and set argument
	oldArgs := os.Args
	os.Args = []string{"hello-world", "bob"}

	// Run main function
	main()

	// Restore stdout and args
	w.Close()
	os.Stdout = old
	os.Args = oldArgs

	// Read captured output
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello, bob!\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}

func TestMainWithMultipleArguments(t *testing.T) {
	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Save original args and set multiple arguments (should use first one)
	oldArgs := os.Args
	os.Args = []string{"hello-world", "alice", "bob"}

	// Run main function
	main()

	// Restore stdout and args
	w.Close()
	os.Stdout = old
	os.Args = oldArgs

	// Read captured output
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	output := buf.String()

	expected := "Hello, alice!\n"
	if output != expected {
		t.Errorf("Expected '%s' but got '%s'", expected, output)
	}
}
