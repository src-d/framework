package utils

import "unicode"

func CamelToSnake(s string) string {
	ret := ""
	for i, r := range s {
		if i > 0 && unicode.IsUpper(r) {
			ret += "_"
		}
		ret += string(unicode.ToLower(r))
	}
	return ret
}

func SnakeToCamel(s string, capitalize bool) string {
	ret := ""
	nextIsUpper := false
	for i, r := range s {
		if nextIsUpper || (i == 0 && capitalize) {
			r = unicode.ToUpper(r)
			nextIsUpper = false
		}
		if r != '_' {
			ret += string(r)
		} else {
			nextIsUpper = true
		}
	}
	return ret
}
