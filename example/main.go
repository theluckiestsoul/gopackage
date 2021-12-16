package main

import (
	"fmt"

	greet "github.com/theluckiestsoul/gopackage"
)

func main() {
	greet.Hello()
	fmt.Println(greet.Shark)

	oct := greet.Octopus{
		Name:  "Jesse",
		Color: "orange",
	}
	fmt.Println(oct)
	oct.Reset()
	fmt.Println(oct)
}
