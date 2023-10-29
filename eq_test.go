package proof

import "testing"

func TestExpectEq(t *testing.T) {

	t.Run("returns true if the given values are equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectEq(mt, "", 1, 1) {
			t.Errorf("expected ExpectEq to return true")
		}
	})

	t.Run("calls Error and returns false if the given values are not equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newEqFailure("one", 2, 1))
		if ExpectEq(mt, "one", 2, 1) {
			t.Errorf("expected ExpectEq to return false")
		}
	})
}

func TestAssertEq(t *testing.T) {

	t.Run("does not fail the test if the given values are equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertEq(mt, "", 1, 1)
	})

	t.Run("calls Fatalf if the given values are not equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newEqFailure("one", 2, 1))
		AssertEq(mt, "one", 2, 1)
	})
}

func TestPreconditionEq(t *testing.T) {

	t.Run("does not fail the test if the given values are equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionEq(mt, "", 1, 1)
	})

	t.Run("calls Fatalf if the given values are not equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newEqFailure("one", 2, 1))
		PreconditionEq(mt, "one", 2, 1)
	})
}
