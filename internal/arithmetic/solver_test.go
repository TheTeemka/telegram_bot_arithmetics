package arithmetic

import (
	"testing"
)

func TestSolveExpr(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"10 * 24 - 24", 216},
		{"-(10 * 5) + 23", -27},
		{"(43 + 27) * (5 - 3)", 140},
		{"-(23) * 12", -276},
		{"-20 * -10", 200},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := SolveExpr(tt.input)
			if err != nil {
				t.Error(err.Error())
				return
			}
			if result != tt.expected {
				t.Errorf("SolveExpr(%ss) = %d; want %d", tt.input, result, tt.expected)
			}
		})
	}
}
