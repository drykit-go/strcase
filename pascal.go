package strcase

import (
	"unicode"
)

// upperExeptions is a map of lowercase strings to their upper result
// that do not follow the PascalCase rules.
// It is far from exhaustive, entries being be added on the need.
var upperExceptions = map[string]string{
	"id":   "ID",
	"http": "HTTP",
	"json": "JSON",
	"yaml": "YAML",
	"toml": "TOML",
	"csv":  "CSV",
}

func pascal(in string) string {
	firstWord, rest := splitAfterFirstWord(in)
	if unicode.IsUpper(rune(firstWord[0])) {
		return in
	}
	if upper, isException := upperExceptions[firstWord]; isException {
		return upper + rest
	}
	return capitalize(in)
}

func splitAfterFirstWord(in string) (firstWord, rest string) {
	if in == "" {
		return "", ""
	}

	splitIndex := 0
	chars := []rune{}
	for i, r := range in {
		if i != 0 && newWordStartsAtIndex(i, r) {
			splitIndex = i
			break
		}
		chars = append(chars, r)
	}
	return string(chars), in[splitIndex:]
}

func newWordStartsAtIndex(i int, r rune) bool {
	return unicode.IsUpper(r)
}

func capitalize(in string) string {
	if in == "" {
		return ""
	}
	first := string(unicode.ToUpper(rune(in[0])))
	if len(in) == 1 {
		return first
	}
	return first + in[1:]
}
