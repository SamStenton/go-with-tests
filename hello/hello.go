package main

import "fmt"

const spanish = "Spanish"
const french = "French"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello : Replies with a greeting when given a name
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	
	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	// Returns the functions default return variable
	return
}

func main() {
    fmt.Println(Hello("Sam", ""))
}