package proof

import (
	"fmt"
)

// ExpectNil logs a test failure if the given value is not nil.
func ExpectNil(t TestingT, expr string, value any) bool {
	t.Helper()
	return Expect(t, nilCheck(expr, value))
}

// AssertNil fails the test immediately if the given value is not nil.
func AssertNil(t TestingT, expr string, value any) {
	t.Helper()
	Assert(t, nilCheck(expr, value))
}

// PreconditionNil sets a precondition that the given value is nil.
func PreconditionNil(t TestingT, expr string, value any) {
	t.Helper()
	Precondition(t, nilCheck(expr, value))
}

func nilCheck(expr string, value any) ExpectationCheck {
	return func() *ExpectationFailure {
		if value == nil {
			return nil
		}
		return newNilFailure(expr, value)
	}
}

func newNilFailure(expr string, value any) *ExpectationFailure {
	return &ExpectationFailure{
		Expectation: fmt.Sprintf("%s to be nil", expr),
		Reason:      fmt.Sprintf("got %v", value),
	}
}

// ExpectNotNil logs a test failure if the given value is nil.
func ExpectNotNil(t TestingT, expr string, value any) bool {
	t.Helper()
	return Expect(t, notNilCheck(expr, value))
}

// AssertNotNil fails the test immediately if the given value is nil.
func AssertNotNil(t TestingT, expr string, value any) {
	t.Helper()
	Assert(t, notNilCheck(expr, value))
}

// PreconditionNotNil sets a precondition that the given value is not nil.
func PreconditionNotNil(t TestingT, expr string, value any) {
	t.Helper()
	Precondition(t, notNilCheck(expr, value))
}

func notNilCheck(expr string, value any) ExpectationCheck {
	return func() *ExpectationFailure {
		if value != nil {
			return nil
		}
		return newNotNilFailure(expr)
	}
}

func newNotNilFailure(expr string) *ExpectationFailure {
	return &ExpectationFailure{
		Expectation: fmt.Sprintf("%s not to be nil", expr),
	}
}
