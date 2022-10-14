package main

import (
	"bufio"
	"fmt"
	"gomappersolution/gomapper"
	"os"
)

func main() {
	var (
		capString string
		pos       int
	)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter string to be capitalized")
	// Reading the input string
	capString, _ = reader.ReadString('\n')
	fmt.Println("enter position of the letter to be capitalized")

	// Reading input position
	_, err := fmt.Scanf("%d", &pos)
	if err != nil {
		fmt.Println(err)
	}

	s := gomapper.NewSkipString(pos, capString)
	gomapper.MapString(s)
	fmt.Println(s)
}
