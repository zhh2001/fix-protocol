package fields

import "testing"

// TestSideEnumValues verifies that Side enum values comply with the official FIX specification
func TestSideEnumValues(t *testing.T) {
	// Define test cases: enum name → actual value → expected value
	testCases := []struct {
		name     string
		actual   string
		expected string
	}{
		{"SideBuy", SideBuy, "1"},
		{"SideSell", SideSell, "2"},
		{"SideBuyMinus", SideBuyMinus, "3"},
		{"SideSellPlus", SideSellPlus, "4"},
		{"SideSellShort", SideSellShort, "5"},
		{"SideSellShortExempt", SideSellShortExempt, "6"},
		{"SideUndisclosed", SideUndisclosed, "7"},
		{"SideCross", SideCross, "8"},
		{"SideCrossShort", SideCrossShort, "9"},
	}

	// Run each test case
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.actual != tc.expected {
				t.Errorf("Enum %s has an incorrect value: actual %q, expected %q", tc.name, tc.actual, tc.expected)
			}
		})
	}
}
