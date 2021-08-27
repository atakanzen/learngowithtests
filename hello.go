package main

import (
	"fmt"
)

const spanish = "Spanish"
const french = "French"
const polish = "Polish"
const englishHelloPrefix = "Hello"
const spanishHelloPrefix = "Hola"
const frenchHelloPrefix = "Bonjour"
const polishHelloPrefix = "Cześć"

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return fmt.Sprintf("%s, %s!", getPrefix(language), name)
}

func getPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case polish:
		prefix = polishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("World", ""))
}
