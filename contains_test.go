package proof

import "testing"

func TestExpectContains(t *testing.T) {

	t.Run("returns true if the given slice contains the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectContains(mt, "", []int{1, 2, 3}, 2) {
			t.Errorf("expected ExpectContains to return true")
		}
	})

	t.Run("calls Error and returns false if the given slice does not contain the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newContainsFailure("someExpr()", []int{1, 2, 3}, 4))
		if ExpectContains(mt, "someExpr()", []int{1, 2, 3}, 4) {
			t.Errorf("expected ExpectContains to return false")
		}
	})
}

func TestAssertContains(t *testing.T) {

	t.Run("does not fail the test if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertContains(mt, "", []int{1, 2, 3}, 2)
	})

	t.Run("calls Fatalf if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newContainsFailure("someExpr()", []int{1, 2, 3}, 4))
		AssertContains(mt, "someExpr()", []int{1, 2, 3}, 4)
	})
}

func TestPreconditionContains(t *testing.T) {

	t.Run("does not fail the test if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionContains(mt, "", []int{1, 2, 3}, 2)
	})

	t.Run("calls Fatalf if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newContainsFailure("someExpr()", []int{1, 2, 3}, 4))
		PreconditionContains(mt, "someExpr()", []int{1, 2, 3}, 4)
	})
}
