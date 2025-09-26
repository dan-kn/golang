package main

import (
	"os"
	"reflect"
	"testing"
)

func Test_createFile(t *testing.T) {
	type args struct {
		p string
	}
	tests := []struct {
		name string
		args args
		want *os.File
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createFile(tt.args.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_writeFile(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writeFile(tt.args.f)
		})
	}
}

func Test_closeFile(t *testing.T) {
	type args struct {
		f *os.File
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closeFile(tt.args.f)
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
