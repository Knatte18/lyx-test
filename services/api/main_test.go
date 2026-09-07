package main

import "testing"

func TestComposeGreeting(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"lyx", "Hello, lyx!"},
		{"", "Hello, world!"},
		{"Ada Lovelace", "Hello, Ada Lovelace!"},
	}

	for _, c := range cases {
		if got := ComposeGreeting(c.name); got != c.want {
			t.Errorf("ComposeGreeting(%q) = %q, want %q", c.name, got, c.want)
		}
	}
}

func TestComposeFarewell(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"lyx", "Goodbye, lyx!"},
		{"", "Goodbye, world!"},
		{"Ada Lovelace", "Goodbye, Ada Lovelace!"},
	}

	for _, c := range cases {
		if got := ComposeFarewell(c.name); got != c.want {
			t.Errorf("ComposeFarewell(%q) = %q, want %q", c.name, got, c.want)
		}
	}
}
