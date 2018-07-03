// Package version provides a single location for the version
package version

// Semantic defines a semver string for aegis
const Semantic = "1.15.0"

// Current will return the current version
func Current() string {
	return Semantic
}
