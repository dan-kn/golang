package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "...",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_addVals(t *testing.T) {
	type args struct {
		xVal int
		yVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "...",
			args: args {
				xVal: 32,
				yVal: 10,
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addVals(tt.args.xVal, tt.args.yVal); got != tt.want {
				t.Errorf("addVals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subVals(t *testing.T) {
	type args struct {
		xVal int
		yVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "...",
			args: args {
				xVal: 32,
				yVal: 10,
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subVals(tt.args.xVal, tt.args.yVal); got != tt.want {
				t.Errorf("subVals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_multVals(t *testing.T) {
	type args struct {
		xVal int
		yVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "...",
			args: args {
				xVal: 32,
				yVal: 10,
			},
			want: 320,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := multVals(tt.args.xVal, tt.args.yVal); got != tt.want {
				t.Errorf("multVals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isEven(t *testing.T) {
	type args struct {
		xVal int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{
			name: "...",
			args: args {
				xVal: 32,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isEven(tt.args.xVal); got != tt.want {
				t.Errorf("isEven() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_addThreeVals(t *testing.T) {
	type args struct {
		xVal int
		yVal int
		zVal int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "...",
			args: args {
				xVal: 32,
				yVal: 10,
				zVal: 128,
			},
			want: 170,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addThreeVals(tt.args.xVal, tt.args.yVal, tt.args.zVal); got != tt.want {
				t.Errorf("addThreeVals() = %v, want %v", got, tt.want)
			}
		})
	}
}
