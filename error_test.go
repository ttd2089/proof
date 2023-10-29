package proof

import (
	"errors"
	"io"
	"testing"
)

func TestExpectError(t *testing.T) {

	t.Run("returns true if the given error is the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectError(mt, "", io.EOF, io.EOF) {
			t.Errorf("expected ExpectError to return true")
		}
	})

	t.Run("calls Error and returns false if the given error is not the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		err := errors.New("actual of error")
		target := errors.New("target error")
		mt.ExpectError(newErrorFailure("someExpr()", err, target))
		if ExpectError(mt, "someExpr()", err, target) {
			t.Errorf("expected ExpectError to return false")
		}
	})
}

func TestAssertError(t *testing.T) {

	t.Run("does not fail the test if the given error is the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertError(mt, "", io.EOF, io.EOF)
	})

	t.Run("calls Fatalf if the given error is not the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		err := errors.New("actual of error")
		target := errors.New("target error")
		mt.ExpectFatalf(assertionFailedFmt, newErrorFailure("someExpr()", err, target))
		AssertError(mt, "someExpr()", err, target)
	})
}

func TestPreconditionError(t *testing.T) {

	t.Run("does not fail the test if the given error is the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionError(mt, "", io.EOF, io.EOF)
	})

	t.Run("calls Fatalf if the given error is not the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		err := errors.New("actual of error")
		target := errors.New("target error")
		mt.ExpectFatalf(preconditionFailedFmt, newErrorFailure("someExpr()", err, target))
		PreconditionError(mt, "someExpr()", err, target)
	})
}
