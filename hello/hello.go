package hello

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

	return greetingPrefix(lang) + person
}

// (prefix string) -- named return - this creates a variabled 'prefix' with a zero-value, in this case = ""
// which will be returned when the function returns
// this will also be present in the go-docs of your function, so the intent can be made clearer
func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishGreetingPrefix
	case french:
		prefix = frenchGreetingPrefix
	default:
		prefix = englishGreetingPrefix
	}
	return
}
