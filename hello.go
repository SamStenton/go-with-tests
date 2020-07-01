package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello : Replies with a greeting when given a name
func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("Sam"))
}