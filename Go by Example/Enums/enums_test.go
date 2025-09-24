package main

import "testing"

func TestServerState_String(t *testing.T) {
	tests := []struct {
		name string
		ss   ServerState
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ss.String(); got != tt.want {
				t.Errorf("ServerState.String() = %v, want %v", got, tt.want)
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

func Test_transition(t *testing.T) {
	type args struct {
		s ServerState
	}
	tests := []struct {
		name string
		args args
		want ServerState
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := transition(tt.args.s); got != tt.want {
				t.Errorf("transition() = %v, want %v", got, tt.want)
			}
		})
	}
}
