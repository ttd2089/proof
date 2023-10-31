package proof

import (
	"errors"
	"fmt"
)

// ExpectError logs a test failure if the given error is not the target error.
func ExpectError(t TestingT, expr string, err, target error) bool {
	t.Helper()
	return Expect(t, errorCheck(expr, err, target))
}

// AssertError fails the test immediately if the given error is not the target error.
func AssertError(t TestingT, expr string, err, target error) {
	t.Helper()
	Assert(t, errorCheck(expr, err, target))
}

// PreconditionError sets a precondition that the given error is the target error.
func PreconditionError(t TestingT, expr string, err, target error) {
	t.Helper()
	Precondition(t, errorCheck(expr, err, target))
}

func errorCheck(expr string, err, target error) ExpectationCheck {
	return func() *ExpectationFailure {
		if errors.Is(err, target) {
			return nil
		}
		return newErrorFailure(expr, err, target)
	}
}

func newErrorFailure(expr string, err, target error) *ExpectationFailure {
	return &ExpectationFailure{
		Expectation: fmt.Sprintf("%s to return error %v", expr, target),
		Reason:      fmt.Sprintf("got %v", err),
	}
}
