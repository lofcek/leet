package number

import (
	"regexp"
)

var pattern = regexp.MustCompile(`^\s*[+-]?(\d+\.?\d*|\.\d+)([eE][+-]?\d+)?\s*$`)

func isNumber(s string) bool {
	return pattern.MatchString(s)
}
