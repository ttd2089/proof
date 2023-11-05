package proof

import "testing"

func TestExpectEq(t *testing.T) {

	t.Run("returns true if the given values are equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectEq(mt, "one", 1, 1) {
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
		AssertEq(mt, "one", 1, 1)
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
		PreconditionEq(mt, "one", 1, 1)
	})

	t.Run("calls Fatalf if the given values are not equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newEqFailure("one", 2, 1))
		PreconditionEq(mt, "one", 2, 1)
	})
}

func TestExpectNeq(t *testing.T) {

	t.Run("returns true if the given values are not equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectNeq(mt, "not-one", 2, 1) {
			t.Errorf("expected ExpectNeq to return true")
		}
	})

	t.Run("calls Error and returns false if the given values are equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newNeqFailure("not-one", 1, 1))
		if ExpectNeq(mt, "not-one", 1, 1) {
			t.Errorf("expected ExpectNeq to return false")
		}
	})
}

func TestAssertNeq(t *testing.T) {

	t.Run("does not fail the test if the given values are not equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertNeq(mt, "not-one", 2, 1)
	})

	t.Run("calls Fatalf if the given values are equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newNeqFailure("not-one", 1, 1))
		AssertNeq(mt, "not-one", 1, 1)
	})
}

func TestPreconditionNeq(t *testing.T) {

	t.Run("does not fail the test if the given values are not equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionNeq(mt, "not-one", 2, 1)
	})

	t.Run("calls Fatalf if the given values are equal", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newNeqFailure("not-one", 1, 1))
		PreconditionNeq(mt, "not-one", 1, 1)
	})
}
