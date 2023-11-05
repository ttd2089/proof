package proof

import "testing"

func TestExpectSetEq(t *testing.T) {

	t.Run("returns true if the given slice is set-equal to the expected set", func(t *testing.T) {
		for _, tc := range []struct {
			name     string
			actual   []int
			expected []int
		}{
			{name: "identical slices", actual: []int{1, 2, 3}, expected: []int{1, 2, 3}},
			{name: "same values different order", actual: []int{1, 2, 3}, expected: []int{2, 1, 3}},
			{name: "duplicate values", actual: []int{1, 2, 2, 3}, expected: []int{1, 2, 3, 3}},
		} {
			t.Run(tc.name, func(t *testing.T) {
				mt := &mockTestingT{}
				defer mt.Verify(t)
				if !ExpectSetEq(mt, "oneTwoThree()", tc.actual, tc.expected) {
					t.Errorf("expected ExpectSetEq to return true")
				}
			})
		}
	})

	t.Run("calls Error and returns false if the given slice contains values not in the expected slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{}, []int{4}))
		if ExpectSetEq(mt, "oneTwoThree()", []int{1, 2, 3, 4}, []int{1, 2, 3}) {
			t.Errorf("expected ExpectSetEq to return false")
		}
	})

	t.Run("calls Error and returns false if the expected slice contains values not in the given slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{3}, []int{}))
		if ExpectSetEq(mt, "oneTwoThree()", []int{1, 2}, []int{1, 2, 3}) {
			t.Errorf("expected ExpectSetEq to return false")
		}
	})

	t.Run("calls Error and returns false if the expected and actuals slices both contains unique values", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{3}, []int{4}))
		if ExpectSetEq(mt, "oneTwoThree()", []int{1, 2, 4}, []int{1, 2, 3}) {
			t.Errorf("expected ExpectSetEq to return false")
		}
	})
}

func TestAssertSetEq(t *testing.T) {

	t.Run("does not fail if the given slice is set-equal to the expected set", func(t *testing.T) {
		for _, tc := range []struct {
			name     string
			actual   []int
			expected []int
		}{
			{name: "identical slices", actual: []int{1, 2, 3}, expected: []int{1, 2, 3}},
			{name: "same values different order", actual: []int{1, 2, 3}, expected: []int{2, 1, 3}},
			{name: "duplicate values", actual: []int{1, 2, 2, 3}, expected: []int{1, 2, 3, 3}},
		} {
			t.Run(tc.name, func(t *testing.T) {
				mt := &mockTestingT{}
				defer mt.Verify(t)
				AssertSetEq(mt, "oneTwoThree()", []int{1, 2, 3}, []int{1, 2, 3})
			})
		}
	})

	t.Run("calls Fatalf and returns false if the given slice contains values not in the expected slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{}, []int{4}))
		AssertSetEq(mt, "oneTwoThree()", []int{1, 2, 3, 4}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the expected slice contains values not in the given slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{3}, []int{}))
		AssertSetEq(mt, "oneTwoThree()", []int{1, 2}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the expected and actuals slices both contains unique values", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{3}, []int{4}))
		AssertSetEq(mt, "oneTwoThree()", []int{1, 2, 4}, []int{1, 2, 3})
	})
}

func TestPreconditionSetEq(t *testing.T) {

	t.Run("does not fail if the given slice is set-equal to the expected set", func(t *testing.T) {
		for _, tc := range []struct {
			name     string
			actual   []int
			expected []int
		}{
			{name: "identical slices", actual: []int{1, 2, 3}, expected: []int{1, 2, 3}},
			{name: "same values different order", actual: []int{1, 2, 3}, expected: []int{2, 1, 3}},
			{name: "duplicate values", actual: []int{1, 2, 2, 3}, expected: []int{1, 2, 3, 3}},
		} {
			t.Run(tc.name, func(t *testing.T) {
				mt := &mockTestingT{}
				defer mt.Verify(t)
				PreconditionSetEq(mt, "oneTwoThree", []int{1, 2, 3}, []int{1, 2, 3})
			})
		}
	})

	t.Run("calls Fatalf and returns false if the given slice contains values not in the expected slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{}, []int{4}))
		PreconditionSetEq(mt, "oneTwoThree()", []int{1, 2, 3, 4}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the expected slice contains values not in the given slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{3}, []int{}))
		PreconditionSetEq(mt, "oneTwoThree()", []int{1, 2}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the expected and actuals slices both contains unique values", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newSetEqFailure("oneTwoThree()", []int{1, 2, 3}, []int{3}, []int{4}))
		PreconditionSetEq(mt, "oneTwoThree()", []int{1, 2, 4}, []int{1, 2, 3})
	})
}
