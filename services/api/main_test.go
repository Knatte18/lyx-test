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
