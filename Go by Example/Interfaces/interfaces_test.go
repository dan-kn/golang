package main

import "testing"

func TestRectangle_area(t *testing.T) {
	type fields struct {
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Rectangle{
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
		width  float64
		height float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
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

func TestCircle_area(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			if got := c.area(); got != tt.want {
				t.Errorf("Circle.area() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCircle_perimeter(t *testing.T) {
	type fields struct {
		radius float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Circle{
				radius: tt.fields.radius,
			}
			if got := c.perimeter(); got != tt.want {
				t.Errorf("Circle.perimeter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_measure(t *testing.T) {
	type args struct {
		g Geometry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			measure(tt.args.g)
		})
	}
}

func Test_detectRectangle(t *testing.T) {
	type args struct {
		g Geometry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			detectRectangle(tt.args.g)
		})
	}
}

func Test_detectCircle(t *testing.T) {
	type args struct {
		g Geometry
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			detectCircle(tt.args.g)
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
