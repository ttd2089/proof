package proof

import (
	"fmt"
)

const assertionFailedFmt = "assertion failed: %s"
const preconditionFailedFmt = "precondition failed: %s"

// An ExpectationFailure is an error describing an expectation and the reason it wasn't met.
type ExpectationFailure struct {

	// Expectation describes the expectation that was not met.
	Expectation string

	// Reason describes the reason the expectation was not met.
	Reason string
}

func (e ExpectationFailure) String() string {
	if e.Reason == "" {
		return fmt.Sprintf("expected %s", e.Expectation)
	}
	return fmt.Sprintf("expected %s; %s", e.Expectation, e.Reason)
}

// A ExpectationCheck check is function that checks whether an expectation is met and returns nil
// if it is, or a non-nil ExpectationFailure describing the failure if not.
type ExpectationCheck func() *ExpectationFailure

// Expect evaluates an ExpectationCheck and returns true if the expectation is met; otherwise, an
// error is logged, the test is failed, and the function returns false.
func Expect(t TestingT, check ExpectationCheck) bool {
	t.Helper()
	if failure := check(); failure != nil {
		t.Error(failure)
		return false
	}
	return true
}

// Assert evaluates an ExpectationCheck, logging the error and failing the test immediately if the
// expectation is not met.
func Assert(t TestingT, check ExpectationCheck) {
	t.Helper()
	if failure := check(); failure != nil {
		t.Fatalf(assertionFailedFmt, failure)
	}
}

// Precondition evaluates an ExpectationCheck, logging the error as a precondition failure and
// failing the test immediately if the expectation is not met.
func Precondition(t TestingT, check ExpectationCheck) {
	t.Helper()
	if failure := check(); failure != nil {
		t.Fatalf(preconditionFailedFmt, failure)
	}
}
