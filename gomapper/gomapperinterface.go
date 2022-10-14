package gomapper

import (
	"fmt"
	"unicode"
)

// problem - 2

type skipString struct {
	pos int    // position of the value to be capitalized
	arr []rune // Holds rune values for all letters of the string,
	// as we Transform these values will be updated
	mapper map[int]rune // holds the positions to capitalize
}

func NewSkipString(pos int, str string) *skipString {
	arr := []rune(str)
	return &skipString{
		pos:    pos,
		arr:    arr,
		mapper: getPositionMapper(pos, arr),
	}
}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// GetValueAsRuneSlice returns a slice of runes.
func (s *skipString) GetValueAsRuneSlice() []rune {
	return s.arr
}

// This function checks if a rune should be capitalized or not based on the position
// and condition that the rune is a alphanumeric.
func (s *skipString) TransformRune(pos int) {
	// If the position is found in mapper then replace arr[pos] with upper case rune value
	if upper_case_val, ok := s.mapper[pos]; ok {
		s.arr[pos] = upper_case_val
	}
}

// This function returns a map with all the elements position as key and corresponding values
// that should be capitalized.
func getPositionMapper(pos int, arr []rune) map[int]rune {
	m := make(map[int]rune)
	j := 0
	for i := 0; i < len(arr); i++ {
		//check if the rune is either a letter or a number
		// skip iteration if it is not alpha numeric
		if !unicode.IsLetter(arr[i]) && !unicode.IsNumber(arr[i]) {
			continue
		}

		if j == pos-1 {
			// transform rune to upper case
			m[i] = unicode.ToUpper(arr[i])
			j = 0
			continue
		}
		j++
	}
	return m
}

// Implement stringer interface
func (s *skipString) String() string {
	return fmt.Sprintf("capitalized string is: %s", string(s.arr))
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}
