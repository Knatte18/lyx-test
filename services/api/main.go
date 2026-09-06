package main

import "fmt"

// Dummy subpath fixture for weft relpath-mirroring tests.
func main() {
	fmt.Println(FormatGreeting("lyx"))
}

// FormatGreeting returns a greeting for name, defaulting to "world" when name is empty.
func FormatGreeting(name string) string {
	if name == "" {
		name = "world"
	}
	return fmt.Sprintf("Hello, %s!", name)
}
