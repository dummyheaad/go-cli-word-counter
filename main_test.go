package main

import (
	"bytes"
	"testing"
)

// TestCountWords tests the count function set to count words
func TestCountWords(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3 word4\n")

	exp := 4

	res := count(b, false, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead.\n", exp, res)
	}
}

// TestCountLines tests the count function set to count lines
func TestCountLines(t *testing.T) {
	b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

	exp := 3

	res := count(b, true, false)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

// TestCountBytes tests the count function set to count bytes
func TestCountBytes(t *testing.T) {
	b := bytes.NewBufferString("word1\nword2 word 3\nwordd")

	exp := 24

	res := count(b, false, true)

	if res != exp {
		t.Errorf("Expected %d, got %d instead. \n", exp, res)
	}
}

// TestCountFromFile tests the count function from a file
func TestCountFromFile(t *testing.T) {
	filename := []string{"test.txt"}

	countWords, err := countFromFile(filename, false, false)
	if err != nil {
		t.Fatal(err)
	}

	countBytes, err := countFromFile(filename, false, true)
	if err != nil {
		t.Fatal(err)
	}

	countLines, err := countFromFile(filename, true, false)
	if err != nil {
		t.Fatal(err)
	}

	expWords := 4
	expBytes := 14
	expLines := 2

	if expWords != countWords["test.txt"] {
		t.Errorf("Expected %d, got %d instead\n", expWords, countWords["test.txt"])
	}

	if expBytes != countBytes["test.txt"] {
		t.Errorf("Expected %d, got %d instead\n", expBytes, countBytes["test.txt"])
	}

	if expLines != countLines["test.txt"] {
		t.Errorf("Expected %d, got %d instead\n", expLines, countLines["test.txt"])
	}
}
