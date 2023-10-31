package proof

import "fmt"

// ExpectSetEq logs a test failure if the given set does not equal the expected set.
func ExpectSetEq[T comparable](t TestingT, expr string, actual, expected []T) bool {
	t.Helper()
	return Expect(t, setSetEqCheck(expr, actual, expected))
}

// AssertSetEq fails the test immediately if the given set does not equal the expected set.
func AssertSetEq[T comparable](t TestingT, expr string, actual, expected []T) {
	t.Helper()
	Assert(t, setSetEqCheck(expr, actual, expected))
}

// PreconditionSetEq sets a precondition that if the given set equals the expected set.
func PreconditionSetEq[T comparable](t TestingT, expr string, actual, expected []T) {
	t.Helper()
	Precondition(t, setSetEqCheck(expr, actual, expected))
}

func setSetEqCheck[T comparable](expr string, actual, expected []T) ExpectationCheck {
	return func() *ExpectationFailure {
		expectedOnly := map[T]struct{}{}
		for _, v := range expected {
			expectedOnly[v] = struct{}{}
		}
		actualOnly := map[T]struct{}{}
		for _, v := range actual {
			actualOnly[v] = struct{}{}
		}
		for _, actual := range actual {
			if _, ok := expectedOnly[actual]; !ok {
				continue
			}
			delete(expectedOnly, actual)
			delete(actualOnly, actual)
		}
		if len(expectedOnly) > 0 || len(actualOnly) > 0 {
			expectedOnlyKeys := make([]T, 0, len(expectedOnly))
			for k := range expectedOnly {
				expectedOnlyKeys = append(expectedOnlyKeys, k)
			}
			actualOnlyKeys := make([]T, 0, len(expectedOnly))
			for k := range actualOnly {
				actualOnlyKeys = append(actualOnlyKeys, k)
			}
			return newSetEqFailure(expr, expected, expectedOnlyKeys, actualOnlyKeys)
		}
		return nil
	}
}

func newSetEqFailure[T comparable](expr string, expected, expectedOnly, actualOnly []T) *ExpectationFailure {
	return &ExpectationFailure{
		Expectation: fmt.Sprintf("%s to be set-equal to %v", expr, expected),
		Reason:      fmt.Sprintf("set difference: expected=%v actual=%v", expectedOnly, actualOnly),
	}
}
