package gomapper

import (
	"unicode"
)

// problem - 1

// This function is to capitalize every third alphanumeric character.
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	// type casting a string into slice of rune
	arr := []rune(s)
	// check if length of the string is greater than 3
	if len(arr) < 2 {
		return s
	}

	var j int
	// iterate through string and transform if it is a third character
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			j = 0
		}

		// check if the rune is either a letter or a number.
		// if rune is not alphanumeric skip the iteration.
		if !unicode.IsLetter(arr[i]) && !unicode.IsNumber(arr[i]) {
			continue
		}

		// this is configurable but fixed as only third letter should be capitalized.
		if j == 2 {
			// transform rune to upper caßse
			arr[i] = unicode.ToUpper(arr[i])
			j = 0
			continue
		}
		arr[i] = unicode.ToLower(arr[i])
		j++
	}
	return string(arr)
}
