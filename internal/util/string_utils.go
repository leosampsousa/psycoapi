package util

import (
	"regexp"
	"strings"
	"unicode"
)

type StringUtils struct {}

func (su StringUtils) HasUpperCase(value string) bool {
	characters := []rune(value);

	for i := 0; i < len(characters); i++  {
		if (unicode.IsLetter(rune(characters[i])) && unicode.ToUpper(characters[i]) == characters[i]) {
			return true
		}
	}
	return false
}

func (su StringUtils) HasLowerCase(value string) bool {
	characters := []rune(value);

	for i := 0; i < len(characters); i++  {
		if (unicode.IsLetter(rune(characters[i])) && unicode.ToLower(characters[i]) == characters[i]) {
			return true
		}
	}
	return false
}

func (su StringUtils) ContainsWhitespace(value string) bool {
	return strings.ContainsRune(value, rune(' '));
}

func (su StringUtils) ContainsSymbols(value string) bool {
	return !regexp.MustCompile(`^[a-z0-9]*$`).MatchString(value)
}