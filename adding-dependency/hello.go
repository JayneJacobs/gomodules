package helloadd

import (
	"fmt"
	"rsc.io/quote"
)

func greet() string {
	fmt.Println("Hello")
	return quote.Hello()
}
