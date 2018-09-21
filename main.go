package phrases

import "fmt"

var phrase = "You can tune a piano but you can't tuna fish"

func GetPhrase() string {
	return phrase
}

func main() {
	fmt.Println(GetPhrase())
}
