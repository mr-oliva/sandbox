package sample_test

import (
	"testing"

	"github.com/bookun/sandbox/go/error-handling-iota/sample"
)

func TestSample(t *testing.T) {
	cases := []struct {
		name              string
		arg               int
		expectedErrorCode sample.ErrorCode
	}{
		{name: "1", arg: 1, expectedErrorCode: sample.FooError},
		{name: "2", arg: 2, expectedErrorCode: sample.BarError},
		{name: "3", arg: 3, expectedErrorCode: sample.OK},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			resultErrCode, _ := sample.Sample(c.arg)
			if c.expectedErrorCode != resultErrCode {
				t.Errorf("got %d, want %d\n", resultErrCode, c.expectedErrorCode)
			}
		})
	}
}
