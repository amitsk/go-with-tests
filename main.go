package main

import (
	"go.uber.org/zap"
)

const spanish = "Spanish"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const french = "French"

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
	return
}

func main() {
	var logger = zap.NewExample()
	defer logger.Sync()
	sugar := logger.Sugar()
	// plain := sugar.Desugar()
	sugar.Infow(Hello("Amit", "English"))
	sugar.Infow(Hello("Bruno", "Spanish"))
	sugar.Infow(Hello("Pierre", "French"))
}
