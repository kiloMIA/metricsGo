package main

import (
	"testing"

	api "github.com/kiloMIA/metricsGo/server/api"
)

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
		{"unknown", "unknown", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := api.chooseCity(tt.city); got != tt.want {
				t.Errorf("chooseCity() = %v, want %v", got, tt.want)
			}
		})
	}
}

func chooseCity(s string) {
	panic("unimplemented")
}
