package main

import (
	"fmt"
	"github.com/rjturek/go-phrase-util/phrase"
)

var phrase1 = "Starbringizizzle"

func GetPhrase() string {
	return phrase1
}

func main() {
	fmt.Println(GetPhrase())
	fmt.Println(phrase.GetPhrase())
}
