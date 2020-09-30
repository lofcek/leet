package number

import (
	"testing"
)

func number(t *testing.T, s string) {
	t.Helper()
	if !isNumber(s) {
		t.Fatalf("%q should be recognized as number", s)
	}
}

func notNumber(t *testing.T, s string) {
	t.Helper()
	if isNumber(s) {
		t.Fatalf("%q should not be recognized as number", s)
	}
}

func TestNumber(t *testing.T) {
	tests := []struct {
		name string
		expr string
		test func(t *testing.T, s string)
	}{
		{"only digits", "0", number},
		{"decimal number", "0.1", number},
		{"decimal number - missing leading zero", ".32", number},
		{"decimal number - missing end", "3.", number},
		{"only letters", "abc", notNumber},
		{"wrong end", "1 a", notNumber},
		{"simple exponent", "2e10", number},
		{"using sign", "-90e3", number},
		{"using spaces", " 1.43", number},
		{"using spaces with e", " -90e3   ", number},
		{"wrong exponent", " 1e", notNumber},
		{"negative exponent", " 6e-1", number},
		{"confusing exponent", " 99e2.5 ", notNumber},
		{"big number", "53.5e9332", number},
		{"two signs", " --6 ", notNumber},
		{"two signs2", "-+3", notNumber},
		{"two sign in exponent", " 3.114e--6 ", notNumber},
		{"start with point", ".32", number},
		{"only point", ".", notNumber},
		{"neither mantissa nor exponent", ".e23 ", notNumber},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tc.test(t, tc.expr)
		})
	}
}
