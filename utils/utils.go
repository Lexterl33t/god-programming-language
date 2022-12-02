package utils

import (
	"regexp"
)

func IsAlpha(word string) bool {
	r, _ := regexp.Compile("^[a-zA-Z_$][a-zA-Z_$0-9]*$")

	return r.MatchString(word)
}
