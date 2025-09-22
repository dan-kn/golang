package main

import (
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

func Test_swapVals(t *testing.T) {
	type args struct {
		xVal int
		yVal int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := swapVals(tt.args.xVal, tt.args.yVal)
			if got != tt.want {
				t.Errorf("swapVals() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("swapVals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_vals(t *testing.T) {
	tests := []struct {
		name  string
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := vals()
			if got != tt.want {
				t.Errorf("vals() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("vals() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
