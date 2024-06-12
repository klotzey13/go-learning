package main

func Hello(name string) string {
	const englishHelloPrefix = "Hello, "
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}
