package main

import (
	"fmt"
	"unicode"
)

type skipString struct {
	pos int //position of the value to be capitalized
	str string
}

func NewSkipString(pos int, str string) *skipString {
	return &skipString{
		pos: pos,
		str: str,
	}
}

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

func (s *skipString) GetValueAsRuneSlice() []rune {
	return []rune(s.str)
}

var arr []rune
var m map[int]rune

// this function checks if a rune should be capitalized or not based on the position
// and condition that the rune is a alphanumeric
func (s *skipString) TransformRune(pos int) {
	if pos == 0 {
		arr = []rune(s.str)
		m = getEveryThirdElementPositionMap(s, arr)
	}

	// if the position is to be capitalized then
	// the value is fetched from map and if not then it is converted to lower case.
	if val, ok := m[pos]; ok {
		arr[pos] = val
	} else {
		arr[pos] = unicode.ToLower(arr[pos])
	}
}

// This function returns a map with all the elements position as key and corresponding values
// that should be capitalized.
func getEveryThirdElementPositionMap(s *skipString, arr []rune) map[int]rune {
	m := make(map[int]rune)
	var j int
	for i := 0; i < len(arr); i++ {
		if i == 0 {
			j = 0
		}

		//check if the rune is either a letter or a number
		// skip iteration if it is not alpha numeric
		if !unicode.IsLetter(arr[i]) && !unicode.IsNumber(arr[i]) {
			continue
		}

		if j == s.pos-1 {
			// transform rune to upper case
			m[i] = unicode.ToUpper(arr[i])
			j = 0
			continue
		}
		j++
	}
	return m
}

// implement stringer interface
func (s *skipString) String() string {
	return fmt.Sprintf("capitalized string is: %s", s.str)
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(s)
	s.str = string(arr)
	fmt.Println(s.String())
}
