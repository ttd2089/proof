package proof

import (
	"testing"
)

func TestExpectationFailure(t *testing.T) {

	t.Run("String()", func(t *testing.T) {

		t.Run("appends prefix to Expectation when Reason is empty", func(t *testing.T) {
			t.Parallel()
			expected := "expected expectation"
			underTest := ExpectationFailure{
				Expectation: "expectation",
			}
			if actual := underTest.String(); actual != expected {
				t.Errorf("expected %q' got %q", expected, actual)
			}
		})

		t.Run("appends prefix to Expectation and Reason when Reason is not empty", func(t *testing.T) {
			t.Parallel()
			expected := "expected expectation; reason"
			underTest := ExpectationFailure{
				Expectation: "expectation",
				Reason:      "reason",
			}
			if actual := underTest.String(); actual != expected {
				t.Errorf("expected %q' got %q", expected, actual)
			}
		})

	})

}

func TestExpect(t *testing.T) {

	t.Run("returns true if the expectation is met", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !Expect(mt, func() *ExpectationFailure { return nil }) {
			t.Errorf("expected Expect to return true")
		}
		mt.Verify(t)
	})

	t.Run("calls Error and returns false if the expectation is not met", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		failure := &ExpectationFailure{
			Expectation: "expectation",
			Reason:      "reason",
		}
		mt.ExpectError(failure)
		if Expect(mt, func() *ExpectationFailure { return failure }) {
			t.Errorf("expected Expect to return false")
		}
	})
}

func TestAssert(t *testing.T) {

	t.Run("does not fail test if the expectation is met", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		Assert(mt, func() *ExpectationFailure { return nil })
		mt.Verify(t)
	})

	t.Run("calls Fatalf if the expectation is not met", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		failure := &ExpectationFailure{
			Expectation: "expectation",
			Reason:      "reason",
		}
		mt.ExpectFatalf(assertionFailedFmt, failure)
		Assert(mt, func() *ExpectationFailure { return failure })
	})
}

func TestPrecondition(t *testing.T) {

	t.Run("does not fail test if the expectation is met", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		Precondition(mt, func() *ExpectationFailure { return nil })
		mt.Verify(t)
	})

	t.Run("calls Fatalf if the expectation is not met", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		failure := &ExpectationFailure{
			Expectation: "expectation",
			Reason:      "reason",
		}
		mt.ExpectFatalf(preconditionFailedFmt, failure)
		Precondition(mt, func() *ExpectationFailure { return failure })
	})
}
