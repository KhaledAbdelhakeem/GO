package main

import (
	"fmt"
)

const (
	Spanish = "Spanish"
	French  = "French"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case Spanish:
		prefix = spanishHelloPrefix
	case French:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return

}
func main() {
	fmt.Println(Hello("World"))

}
