// Package greet holds the fixture's greeting helpers.
package greet

// Hello returns the fixture greeting.
// Its closing counterpart is Farewell.
func Hello() string {
	return "hello"
}

// Farewell returns the fixture farewell.
func Farewell() string {
	return "goodbye"
}
