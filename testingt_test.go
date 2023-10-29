package proof

import (
	"encoding/json"
	"reflect"
	"testing"
)

type errorCall struct {
	Args []any
}

type fatalfCall struct {
	Format string
	Args   []any
}

type mockTestingT struct {
	expectedErrorCalls  []errorCall
	errorCalls          []errorCall
	expectedFatalfCalls []fatalfCall
	fatalfCalls         []fatalfCall
}

func (m *mockTestingT) ExpectError(args ...any) {
	m.expectedErrorCalls = append(m.expectedErrorCalls, errorCall{Args: args})
}

func (m *mockTestingT) Error(args ...any) {
	m.errorCalls = append(m.errorCalls, errorCall{Args: args})
}

func (m *mockTestingT) ExpectFatalf(format string, args ...any) {
	m.expectedFatalfCalls = append(m.expectedFatalfCalls, fatalfCall{Format: format, Args: args})
}

func (m *mockTestingT) Fatalf(format string, args ...any) {
	m.fatalfCalls = append(m.fatalfCalls, fatalfCall{Format: format, Args: args})
}

func (m *mockTestingT) Verify(t *testing.T) {
	verifyCalls(t, "Error", m.expectedErrorCalls, m.errorCalls)
	verifyCalls(t, "Fatalf", m.expectedFatalfCalls, m.fatalfCalls)
}

func verifyCalls[T any](t *testing.T, name string, expected, actual []T) {
	if expected, actual := len(expected), len(actual); expected != actual {
		t.Errorf("expected %d calls to %s, got %d", expected, name, actual)
		return
	}
	for i := range expected {
		expected, actual := expected[i], actual[i]
		if !reflect.DeepEqual(actual, expected) {
			expected, err := json.Marshal(expected)
			if err != nil {
				panic("failed to marshall JSON")
			}
			actual, err := json.Marshal(actual)
			if err != nil {
				panic("failed to marshall JSON")
			}
			t.Errorf("expected call #%d to %s to receive %s, got %s", i, name, expected, actual)
			return
		}
	}
}

var _ TestingT = &mockTestingT{}
