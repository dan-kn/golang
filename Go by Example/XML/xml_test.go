package main

import (
	"encoding/xml"
	"testing"
)

func TestPlant_String(t *testing.T) {
	type fields struct {
		XMLName xml.Name
		Id      int
		Name    string
		Origin  []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Plant{
				XMLName: tt.fields.XMLName,
				Id:      tt.fields.Id,
				Name:    tt.fields.Name,
				Origin:  tt.fields.Origin,
			}
			if got := p.String(); got != tt.want {
				t.Errorf("Plant.String() = %v, want %v", got, tt.want)
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
