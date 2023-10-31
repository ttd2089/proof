package proof

import "testing"

func TestExpectSetEq(t *testing.T) {

	t.Run("returns true if the given slice is set-equal to the expected set", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectSetEq(mt, "", []int{1, 2, 3}, []int{1, 2, 3}) {
			t.Errorf("expected ExpectSetEq to return true")
		}
	})

	t.Run("calls Error and returns false if the given slice contains values not in the expected slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newSetEqFailure("someExpr()", []int{1, 2, 3}, []int{}, []int{4}))
		if ExpectSetEq(mt, "someExpr()", []int{1, 2, 3, 4}, []int{1, 2, 3}) {
			t.Errorf("expected ExpectSetEq to return false")
		}
	})

	t.Run("calls Error and returns false if the expected slice contains values not in the given slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newSetEqFailure("someExpr()", []int{1, 2, 3, 4}, []int{4}, []int{}))
		if ExpectSetEq(mt, "someExpr()", []int{1, 2, 3}, []int{1, 2, 3, 4}) {
			t.Errorf("expected ExpectSetEq to return false")
		}
	})

	t.Run("calls Error and returns false if the expected and actuals slices both contains unique values", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newSetEqFailure("someExpr()", []int{1, 2, 3, 5}, []int{5}, []int{4}))
		if ExpectSetEq(mt, "someExpr()", []int{1, 2, 3, 4}, []int{1, 2, 3, 5}) {
			t.Errorf("expected ExpectSetEq to return false")
		}
	})
}

func TestAssertSetEq(t *testing.T) {

	t.Run("does not fail if the given slice is set-equal to the expected set", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertSetEq(mt, "", []int{1, 2, 3}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the given slice contains values not in the expected slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newSetEqFailure("someExpr()", []int{1, 2, 3}, []int{}, []int{4}))
		AssertSetEq(mt, "someExpr()", []int{1, 2, 3, 4}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the expected slice contains values not in the given slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newSetEqFailure("someExpr()", []int{1, 2, 3, 4}, []int{4}, []int{}))
		AssertSetEq(mt, "someExpr()", []int{1, 2, 3}, []int{1, 2, 3, 4})
	})

	t.Run("calls Fatalf and returns false if the expected and actuals slices both contains unique values", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newSetEqFailure("someExpr()", []int{1, 2, 3, 5}, []int{5}, []int{4}))
		AssertSetEq(mt, "someExpr()", []int{1, 2, 3, 4}, []int{1, 2, 3, 5})
	})
}

func TestPreconditionSetEq(t *testing.T) {

	t.Run("does not fail if the given slice is set-equal to the expected set", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionSetEq(mt, "", []int{1, 2, 3}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the given slice contains values not in the expected slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newSetEqFailure("someExpr()", []int{1, 2, 3}, []int{}, []int{4}))
		PreconditionSetEq(mt, "someExpr()", []int{1, 2, 3, 4}, []int{1, 2, 3})
	})

	t.Run("calls Fatalf and returns false if the expected slice contains values not in the given slice", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newSetEqFailure("someExpr()", []int{1, 2, 3, 4}, []int{4}, []int{}))
		PreconditionSetEq(mt, "someExpr()", []int{1, 2, 3}, []int{1, 2, 3, 4})
	})

	t.Run("calls Fatalf and returns false if the expected and actuals slices both contains unique values", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newSetEqFailure("someExpr()", []int{1, 2, 3, 5}, []int{5}, []int{4}))
		PreconditionSetEq(mt, "someExpr()", []int{1, 2, 3, 4}, []int{1, 2, 3, 5})
	})
}
