package main

import (
	"sync"
	"testing"
)

func TestContainer_inc(t *testing.T) {
	type fields struct {
		mu       sync.Mutex
		counters map[string]int
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Container{
				mu:       tt.fields.mu,
				counters: tt.fields.counters,
			}
			c.inc(tt.args.name)
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
