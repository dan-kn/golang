package main

import "testing"

func TestRectangle_area(t *testing.T) {
	type fields struct {
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.area(); got != tt.want {
				t.Errorf("Rectangle.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_perimeter(t *testing.T) {
	type fields struct {
		width  int
		height int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
				width:  tt.fields.width,
				height: tt.fields.height,
			}
			if got := r.perimeter(); got != tt.want {
				t.Errorf("Rectangle.perimeter() = %v, want %v", got, tt.want)
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
