package hello

import (
	"fmt"
	"rsc.io/quote"
)

func greet(n string) (greeting string, err string) {
	err = ""

	if n == "" {
		err = "No Name was given"
	}
	greeting = fmt.Sprintf("hello: %v", n)
	fmt.Printf("hello: %v", n)
	return greeting, err
}

func Hello(n string) string {
	return quote.Hello()
}
