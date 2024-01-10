package main

import (
	"math/big"
	"testing"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		input    int
		expected *big.Int
		wantErr  bool
	}{
		{0, big.NewInt(1), false},
		{1, big.NewInt(1), false},
		{5, big.NewInt(120), false},
		{-1, nil, true}, // error test case for negative value
	}

	for _, test := range tests {
		result, err := fact(test.input)

		if (err != nil) != test.wantErr {
			t.Errorf("For input %d, expected error: %v, got error: %v", test.input, test.wantErr, err)
			continue
		}

		if !test.wantErr && result.Cmp(test.expected) != 0 {
			t.Errorf("For input %d, expected %s, got %s", test.input, test.expected, result)
		}
	}
}
