package main

import (
	"fmt"
	"gomappersolution/gomapper"
)

func main() {
	s := gomapper.NewSkipString(3, "Aspiration.com")
	gomapper.MapString(s)
	fmt.Println(s)
}
