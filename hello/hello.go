package hello

import "fmt"

const greetingPrefix = "Hello, "

func Hello(person string) string {
	if person == "" {
		person = "World"
	}
	return greetingPrefix + person
}

func main() {
	fmt.Println(Hello(""))
}
