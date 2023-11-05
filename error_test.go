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
		if !ExpectError(mt, "readPastEOF()", io.EOF, io.EOF) {
			t.Errorf("expected ExpectError to return true")
		}
	})

	t.Run("calls Error and returns false if the given error is not the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		err := errors.New("not EOF")
		mt.ExpectError(newErrorFailure("readPastEOF()", err, io.EOF))
		if ExpectError(mt, "readPastEOF()", err, io.EOF) {
			t.Errorf("expected ExpectError to return false")
		}
	})
}

func TestAssertError(t *testing.T) {

	t.Run("does not fail the test if the given error is the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertError(mt, "readPastEOF()", io.EOF, io.EOF)
	})

	t.Run("calls Fatalf if the given error is not the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		err := errors.New("not EOF")
		mt.ExpectFatalf(assertionFailedFmt, newErrorFailure("readPastEOF()", err, io.EOF))
		AssertError(mt, "readPastEOF()", err, io.EOF)
	})
}

func TestPreconditionError(t *testing.T) {

	t.Run("does not fail the test if the given error is the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionError(mt, "readPastEOF()", io.EOF, io.EOF)
	})

	t.Run("calls Fatalf if the given error is not the target error", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		err := errors.New("not EOF")
		mt.ExpectFatalf(preconditionFailedFmt, newErrorFailure("readPastEOF()", err, io.EOF))
		PreconditionError(mt, "readPastEOF()", err, io.EOF)
	})
}
