package proof

import "testing"

func TestExpectNil(t *testing.T) {

	t.Run("returns true if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectNil(mt, "nilExpr()", nil) {
			t.Errorf("expected ExpectNil to return true")
		}
	})

	t.Run("calls Error and returns false if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newNilFailure("nilExpr()", 2))
		if ExpectNil(mt, "nilExpr()", 2) {
			t.Errorf("expected ExpectNil to return false")
		}
	})
}

func TestAssertNil(t *testing.T) {

	t.Run("does not fail the test if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertNil(mt, "nilExpr()", nil)
	})

	t.Run("calls Fatalf if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newNilFailure("nilExpr()", 2))
		AssertNil(mt, "nilExpr()", 2)
	})
}

func TestPreconditionNil(t *testing.T) {

	t.Run("does not fail the test if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionNil(mt, "nilExpr()", nil)
	})

	t.Run("calls Fatalf if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newNilFailure("nilExpr()", 2))
		PreconditionNil(mt, "nilExpr()", 2)
	})
}

func TestExpectNotNil(t *testing.T) {

	t.Run("returns true if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectNotNil(mt, "nonNilExpr", 2) {
			t.Errorf("expected ExpectNotNil to return true")
		}
	})

	t.Run("calls Error and returns false if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newNotNilFailure("nonNilExpr()"))
		if ExpectNotNil(mt, "nonNilExpr()", nil) {
			t.Errorf("expected ExpectNotNil to return false")
		}
	})
}

func TestAssertNotNil(t *testing.T) {

	t.Run("does not fail the test if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertNotNil(mt, "nonNilExpr", 2)
	})

	t.Run("calls Fatalf if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newNotNilFailure("nonNilExpr()"))
		AssertNotNil(mt, "nonNilExpr()", nil)
	})
}

func TestPreconditionNotNil(t *testing.T) {

	t.Run("does not fail the test if the given value is not nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionNotNil(mt, "nonNilExpr", 2)
	})

	t.Run("calls Fatalf if the given value is nil", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newNotNilFailure("nonNilExpr()"))
		PreconditionNotNil(mt, "nonNilExpr()", nil)
	})
}
