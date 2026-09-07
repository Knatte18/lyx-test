package main

import "fmt"

// Dummy subpath fixture for weft relpath-mirroring tests.
func main() {
	fmt.Println(ComposeGreeting("lyx"))
}

func defaultName(name string) string {
	if name == "" {
		return "world"
	}
	return name
}

// ComposeGreeting returns a greeting for name, defaulting to "world" when name is empty.
func ComposeGreeting(name string) string {
	return fmt.Sprintf("Hello, %s!", defaultName(name))
}

// ComposeFarewell returns a farewell for name, defaulting to "world" when name is empty.
func ComposeFarewell(name string) string {
	return fmt.Sprintf("Goodbye, %s!", defaultName(name))
}
