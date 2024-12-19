package calc

import (
	"testing"
)

func TestCalc(t *testing.T) {
	cases := []struct {
		name      string
		exp       string
		expected  float64
		wantError bool
	}{
		{
			name:      "big numbers",
			exp:       "999999999 + 888888888",
			expected:  1888888887.0,
			wantError: false,
		},
		{
			name:      "combined operations",
			exp:       "(5 + 3) * (10 / 2)",
			expected:  40.0,
			wantError: false,
		},
		{
			name:      "division with remainder",
			exp:       "17 / 4",
			expected:  4.25,
			wantError: false,
		},
		{
			name:      "long sequence",
			exp:       "1 + 2 * 3 - 4 / 2 + 5",
			expected:  10.0,
			wantError: false,
		},
		{
			name:      "negative with operations",
			exp:       "(-3 + 7) * (-2)",
			expected:  -8.0,
			wantError: false,
		},
		{
			name:      "long fraction",
			exp:       "10/3",
			expected:  3.33,
			wantError: false,
		},
		{
			name:      "multiply decimal",
			exp:       "12.5 * 4.4",
			expected:  55.0,
			wantError: false,
		},
		{
			name:      "negative subtraction",
			exp:       "-1500 - 2000",
			expected:  -3500.0,
			wantError: false,
		},
		{
			name:      "empty expression",
			exp:       "",
			expected:  0,
			wantError: true,
		},
		{
			name:      "invalid expression",
			exp:       "12b + 12a",
			expected:  0,
			wantError: true,
		},
		{
			name:      "invalid parentheses",
			exp:       "2 * (10+5",
			expected:  0,
			wantError: true,
		},
		{
			name:      "division by zero",
			exp:       "100 / 0",
			expected:  0,
			wantError: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			res, err := Calc(tc.exp)
			if err != nil && !tc.wantError {
				t.Error("Expected no error, but got: ", err)
			}
			if err == nil && tc.wantError {
				t.Error("Expected error, but got none")
			}
			if int(res*100) != int(tc.expected*100) {
				t.Errorf("Expected result %f, but got %f", tc.expected, res)
			}
		})
	}
}
