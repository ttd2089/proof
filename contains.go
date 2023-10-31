package proof

import "fmt"

// ExpectContains logs a test failure if the given slice does not contain the expected value.
func ExpectContains[T comparable](t TestingT, expr string, arr []T, expected T) bool {
	t.Helper()
	return Expect(t, containsCheck(expr, arr, expected))
}

// AssertContains fails the test immediately if the given slice does not contain the expected value.
func AssertContains[T comparable](t TestingT, expr string, arr []T, expected T) {
	t.Helper()
	Assert(t, containsCheck(expr, arr, expected))
}

// PreconditionContains sets a precondition that the given slice contains the expected value.
func PreconditionContains[T comparable](t TestingT, expr string, arr []T, expected T) {
	t.Helper()
	Precondition(t, containsCheck(expr, arr, expected))
}

func containsCheck[T comparable](expr string, arr []T, expected T) ExpectationCheck {
	return func() *ExpectationFailure {
		for _, item := range arr {
			if expected == item {
				return nil
			}
		}
		return newContainsFailure(expr, arr, expected)
	}
}

func newContainsFailure[T comparable](expr string, arr []T, expected T) *ExpectationFailure {
	return &ExpectationFailure{
		Expectation: fmt.Sprintf("%s to contain %v", expr, expected),
		Reason:      fmt.Sprintf("actual contents: %v", arr),
	}
}
