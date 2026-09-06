package main

import "testing"

func TestFormatGreeting(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"lyx", "Hello, lyx!"},
		{"", "Hello, world!"},
		{"Ada Lovelace", "Hello, Ada Lovelace!"},
	}

	for _, c := range cases {
		if got := FormatGreeting(c.name); got != c.want {
			t.Errorf("FormatGreeting(%q) = %q, want %q", c.name, got, c.want)
		}
	}
}
