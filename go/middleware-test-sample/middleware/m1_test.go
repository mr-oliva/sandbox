package middleware

import (
	"net/http"
	"reflect"
	"testing"
)

func TestM1(t *testing.T) {
	type args struct {
		next http.HandlerFunc
	}
	tests := []struct {
		name string
		args args
		want http.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := M1(tt.args.next); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("M1() = %v, want %v", got, tt.want)
			}
		})
	}
}
