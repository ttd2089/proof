package proof

import "testing"

func TestExpectContains(t *testing.T) {

	t.Run("returns true if the given slice contains the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		if !ExpectContains(mt, "containsTwo()", []int{1, 2, 3}, 2) {
			t.Errorf("expected ExpectContains to return true")
		}
	})

	t.Run("calls Error and returns false if the given slice does not contain the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectError(newContainsFailure("containsTwo()", []int{1, 4, 3}, 2))
		if ExpectContains(mt, "containsTwo()", []int{1, 4, 3}, 2) {
			t.Errorf("expected ExpectContains to return false")
		}
	})
}

func TestAssertContains(t *testing.T) {

	t.Run("does not fail the test if the given slice contains the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		AssertContains(mt, "containsTwo()", []int{1, 2, 3}, 2)
	})

	t.Run("calls Fatalf if the given slice does not contain the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(assertionFailedFmt, newContainsFailure("containsTwo()", []int{1, 4, 3}, 2))
		AssertContains(mt, "containsTwo()", []int{1, 4, 3}, 2)
	})
}

func TestPreconditionContains(t *testing.T) {

	t.Run("does not fail the test if the given slice contains the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		PreconditionContains(mt, "containsTwo()", []int{1, 2, 3}, 2)
	})

	t.Run("calls Fatalf if the given slice does not contain the expected value", func(t *testing.T) {
		mt := &mockTestingT{}
		defer mt.Verify(t)
		mt.ExpectFatalf(preconditionFailedFmt, newContainsFailure("containsTwo()", []int{1, 4, 3}, 2))
		PreconditionContains(mt, "containsTwo()", []int{1, 4, 3}, 2)
	})
}
