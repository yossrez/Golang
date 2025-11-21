package math_test

import (
	"testing"

	"github.com/bellshade/Golang/math"
)

func TestAngkaArmstrong(t *testing.T) {
	testingData := []struct {
		name     string
		number   int
		expected bool
	}{
		{"Armstrong: 0", 0, true},
		{"Armstrong: 1", 1, true},
		{"Armstrong: 371", 371, true},
		{"Armstrong: 9474", 9474, true},

		{"Bukan Armstrong: 10", 10, false},
		{"Bukan Armstrong: 123", 123, false},

		{"Angka Negatif: -153", -153, false},
		{"Angka Negatif: -1", -1, false},
	}

	for _, tc := range testingData {
		t.Run(tc.name, func(t *testing.T) {
			got := math.AngkaArmstrong(tc.number)
			if got != tc.expected {
				t.Errorf("AngkaArmstrong(%d) = %v; ekspetasi %v", tc.number, got, tc.expected)
			}
		})
	}
}
