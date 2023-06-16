package main

import "testing"

func Test_chooseCity(t *testing.T) {
	tests := []struct {
		name string
		city string
		want int64
	}{
		{"good case", "astana", 1},
		{"good case", "oskemen", 2},
		{"good case", "atyrau", 4},
		{"good case", "semey", 3},
		{"unknown", "unknown", 0}, // Testing for an unknown city
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chooseCity(tt.city); got != tt.want {
				t.Errorf("chooseCity() = %v, want %v", got, tt.want)
			}
		})
	}
}
