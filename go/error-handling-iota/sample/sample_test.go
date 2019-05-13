package sample_test

import (
	"testing"

	"github.com/bookun/sandbox/go/error-handling-iota/sample"
)

func TestSample(t *testing.T) {
	cases := []struct {
		name          string
		arg           int
		expectedError int
	}{
		{name: "1", arg: 1, expectedError: sample.FooError},
		{name: "2", arg: 2, expectedError: sample.BarError},
		{name: "3", arg: 3, expectedError: sample.OK},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if c.expectedError != sample.Sample(c.arg) {
				t.Errorf("got %d, want %d\n", sample.Sample(c.arg), c.expectedError)
			}
		})
	}
}
