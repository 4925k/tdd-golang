package main

import "fmt"

const (
	spanish       = "spanish"
	french        = "french"
	prefixEnglish = "Hello, "
	prefixSpanish = "Hola, "
	prefixFrench  = "bonsuir"
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greet(language) + name
}

func greet(language string) string {
	switch language {
	case french:
		return prefixFrench
	case spanish:
		return prefixSpanish
	default:
		return prefixEnglish
	}
}
func main() {
	fmt.Println(Hello("world", ""))
}
