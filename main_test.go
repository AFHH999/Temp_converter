package main

import (
	"math"
	"strings"
	"testing"
)

func TestFToC(t *testing.T) {
	tests := []struct {
		name string
		temp float64
		want float64
	}{
		{"Cold", 20, -6.666667},
		{"Medium", 120, 48.88889},
		{"Hot", 200, 93.33333},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := fToC(tt.temp)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("fToC(%v) = %v, want %v", tt.temp, got, tt.want)
			}
		})
	}
}

func TestCToF(t *testing.T) {
	tests := []struct {
		name string
		temp float64
		want float64
	}{
		{"Zero", 0, 32},
		{"Same", -40, -40},
		{"Boiling", 100, 212},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := cToF(tt.temp)
			if math.Abs(got-tt.want) > 0.001 {
				t.Errorf("cToF(%v) = %v, want %v", tt.temp, got, tt.want)
			}
		})
	}
}

func TestGetUserInput(t *testing.T) {
	input := strings.NewReader("bad\nTerrible\n25\n")
	got := getUserInput(input)
	if got != 25 {
		t.Errorf("got %v, want 25", got)
	}
}

func TestGetUserUnit(t *testing.T) {
	input := strings.NewReader("C\n")
	got := getUserUnit(input)
	if got != "C" {
		t.Errorf("getUserUnit() = %q, want \"C\"", got)
	}
}
