package number

import (
	"strings"
)

/*
Simple solution based on regexp
var pattern = regexp.MustCompile(`^\s*[+-]?(\d+\.?\d*|\.\d+)([eE][+-]?\d+)?\s*$`)

func isNumber(s string) bool {
	return pattern.MatchString(s)
}*/

func isNumber(s string) bool {
	b := bytes(s)

	for b.startWithAny(space)  {
		b.dropFirstByte()
	}

	if b.startWithAny(sign) {
		b.dropFirstByte()
	}

	digitsBeforeDot := false
	for ; b.startWithAny(digit); b.dropFirstByte() {
		digitsBeforeDot = true
	}

	if b.startWithAny(dot) {
		b.dropFirstByte()
	}

	digitsAfterDot := false
	for ; b.startWithAny(digit); b.dropFirstByte() {
		digitsAfterDot = true
	}

	if !digitsBeforeDot && !digitsAfterDot {
		return false
	}

	if b.startWithAny(e) {
		b.dropFirstByte()

		if b.startWithAny(sign) {
			b.dropFirstByte()
		}
	
		if !b.startWithAny(digit) {
			return false
		}
		for b.startWithAny(digit)  {
			b.dropFirstByte()
		}
	}

	for b.startWithAny(space)  {
		b.dropFirstByte()
	}
	return len(b) == 0
}

type bytes []byte

func (b bytes) startWithAny(s string) bool {
	return len(b) > 0 && strings.IndexByte(s, b[0]) >= 0
}

func (b *bytes) dropFirstByte() {
	*b = (*b)[1:]
}

const (
	sign  = "+-"
	digit = "0123456789"
	e     = "eE"
	dot   = "."
	space = " \t\n"
)

// func digits(c byte) bool { return c >= '0' && c <= '9' }
// func sign(c byte) bool   { return c == '+' || c == '-' }
// func dot(c byte) bool    { return c == '.' }
// func e(c byte) bool      { return c == 'e' || c == 'E' }
// func spaces(c byte) bool { return c == ' ' || c == '\t' }
