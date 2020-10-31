// Package twofer exposes a 2-fer func
package twofer

// ShareWith is a func that return a 2-fer string
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
