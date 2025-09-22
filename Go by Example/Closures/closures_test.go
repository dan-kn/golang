package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_intSeq(t *testing.T) {
	tests := []struct {
		name string
		want func() int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intSeq(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("intSeq() = %v, want %v", got, tt.want)
			}
		})
	}
}
