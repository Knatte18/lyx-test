package main

import "fmt"

// Dummy subpath fixture for weft relpath-mirroring tests.
func main() {
	fmt.Println(shout(ComposeGreeting("lyx")))
	fmt.Println(ComposeFarewell("lyx"))
}

func orDefault(name string) string {
	if name == "" {
		return "world"
	}
	return name
}

func shout(s string) string {
	return s + "!"
}

// ComposeGreeting returns a greeting for name, defaulting to "world" when name is empty.
func ComposeGreeting(name string) string {
	return fmt.Sprintf("Hello, %s!", orDefault(name))
}

// ComposeFarewell returns a farewell for name, defaulting to "world" when name is empty.
// It delegates that default to the shared orDefault helper, the same one ComposeGreeting uses.
func ComposeFarewell(name string) string {
	return fmt.Sprintf("Goodbye, %s!", orDefault(name))
}
