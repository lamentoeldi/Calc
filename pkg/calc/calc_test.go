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
			name:      "1",
			exp:       "99*89+98",
			expected:  8909,
			wantError: false,
		},
		{
			name:      "2",
			exp:       "(17 * (99 + 135))/2",
			expected:  1989,
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
			if res != tc.expected {
				t.Errorf("Expected result %f, but got %f", tc.expected, res)
			}
		})
	}
}
