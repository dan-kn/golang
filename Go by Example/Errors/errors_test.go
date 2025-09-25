package main

import "testing"

func Test_f(t *testing.T) {
	type args struct {
		arg int
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := f(tt.args.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("f() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("f() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeTea(t *testing.T) {
	type args struct {
		arg int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := makeTea(tt.args.arg); (err != nil) != tt.wantErr {
				t.Errorf("makeTea() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

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
