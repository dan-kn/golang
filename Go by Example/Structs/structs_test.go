package main

import (
	"reflect"
	"testing"
)

func Test_newPerson(t *testing.T) {
	type args struct {
		firstName string
		lastName  string
		age       int
	}
	tests := []struct {
		name string
		args args
		want *Person
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newPerson(tt.args.firstName, tt.args.lastName, tt.args.age); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPerson() = %v, want %v", got, tt.want)
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
