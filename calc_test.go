package main

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		expression     string
		expectedResult string
		expectError    bool
	}{
		{
			expression:     "14 - 10 * (1 - 11)",
			expectedResult: "114.000000",
			expectError:    false,
		},
		{
			expression:     "  10 / 1",
			expectedResult: "10.000000",
			expectError:    false,
		},
		{
			expression:     "20 / 2",
			expectedResult: "10.000000",
			expectError:    false,
		},
		{
			expression:     "1 + b",
			expectedResult: "",
			expectError:    true,
		},
		{
			expression:     "a + b",
			expectedResult: "",
			expectError:    true,
		},
		{
			expression:     "0 / 0",
			expectedResult: "",
			expectError:    true,
		},
		{
			expression:     "17 + 1 + ()",
			expectedResult: "",
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		result, err := Calc(tc.expression)

		if (err != nil) != tc.expectError {
			t.Fatalf("Unexpected test outcome. Input: %s. Fail?: %v. Actual: %v", tc.expression, tc.expectError, err)
		}

		if !tc.expectError {
			resultStr := fmt.Sprintf("%f", result)

			if resultStr != tc.expectedResult {
				t.Fatalf("Input: %s. Expected: %s. Actual: %s", tc.expression, tc.expectedResult, resultStr)
			}
		}
	}
}
