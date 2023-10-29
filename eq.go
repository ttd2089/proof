package proof

import (
	"fmt"
)

// ExpectNoError logs a test failure if the given expression does not have the expected value.
func ExpectEq[T comparable](t TestingT, expr string, actual, expected T) bool {
	return Expect(t, eqCheck(expr, actual, expected))
}

// AssertEq fails the test immediately if the given expression does not have the expected value.
func AssertEq[T comparable](t TestingT, expr string, actual, expected T) {
	Assert(t, eqCheck(expr, actual, expected))
}

// PreconditionEq sets a precondition that the given expression has the expected value.
func PreconditionEq[T comparable](t TestingT, expr string, actual, expected T) {
	Precondition(t, eqCheck(expr, actual, expected))
}

func eqCheck[T comparable](expr string, actual, expected T) ExpectationCheck {
	return func() *ExpectationFailure {
		if expected == actual {
			return nil
		}
		return newEqFailure(expr, actual, expected)
	}
}

func newEqFailure[T comparable](expr string, actual, expected T) *ExpectationFailure {
	return &ExpectationFailure{
		Expectation: fmt.Sprintf("%s to equal %v", expr, expected),
		Reason:      fmt.Sprintf("got %v", actual),
	}
}