package hello

import "fmt"

const (
	spanish               = "Spanish"
	french                = "French"
	english               = "English"
	englishGreetingPrefix = "Hello, "
	spanishGreetingPrefix = "Hola, "
	frenchGreetingPrefix  = "Bonjour, "
)

func Hello(person, lang string) string {
	if person == "" {
		person = "World"
	}

	prefix := englishGreetingPrefix

	switch lang {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	}

	return prefix + person
}

func main() {
	fmt.Println(Hello("", ""))
}
