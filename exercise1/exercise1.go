package exercise1

import (
	"fmt"
	"rsc.io/quote"
)

func glassdoor() string {
	something := quote.Glass()
	greeting := "This is my window"
	fmt.Println(something)
	fmt.Println(greeting)
	return something
}
