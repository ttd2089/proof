package proof

import "testing"

func TestExpectNil(t *testing.T) {

	t.Run("returns true if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectNil(mt, "", nil) {
			t.Errorf("expected ExpectNil to return true")
		}
	})

	t.Run("calls Error and returns false if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newNilFailure("someExpr()", 2))
		if ExpectNil(mt, "someExpr()", 2) {
			t.Errorf("expected ExpectNil to return false")
		}
	})
}

func TestAssertNil(t *testing.T) {

	t.Run("does not fail the test if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertNil(mt, "", nil)
	})

	t.Run("calls Fatalf if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newNilFailure("someExpr()", 2))
		AssertNil(mt, "someExpr()", 2)
	})
}

func TestPreconditionNil(t *testing.T) {

	t.Run("does not fail the test if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionNil(mt, "", nil)
	})

	t.Run("calls Fatalf if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newNilFailure("someExpr()", 2))
		PreconditionNil(mt, "someExpr()", 2)
	})
}

func TestExpectNotNil(t *testing.T) {

	t.Run("returns true if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectNotNil(mt, "", 2) {
			t.Errorf("expected ExpectNotNil to return true")
		}
	})

	t.Run("calls Error and returns false if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newNotNilFailure("someExpr()"))
		if ExpectNotNil(mt, "someExpr()", nil) {
			t.Errorf("expected ExpectNotNil to return false")
		}
	})
}

func TestAssertNotNil(t *testing.T) {

	t.Run("does not fail the test if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertNotNil(mt, "", 2)
	})

	t.Run("calls Fatalf if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newNotNilFailure("someExpr()"))
		AssertNotNil(mt, "someExpr()", nil)
	})
}

func TestPreconditionNotNil(t *testing.T) {

	t.Run("does not fail the test if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionNotNil(mt, "", 2)
	})

	t.Run("calls Fatalf if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newNotNilFailure("someExpr()"))
		PreconditionNotNil(mt, "someExpr()", nil)
	})
}
