package proof

import "testing"

// TestingT is an interface containing the subset of testing.T functionality this package depends on.
type TestingT interface {
	Error(args ...any)
	Fatalf(format string, args ...any)
}

var _ TestingT = &testing.T{}
