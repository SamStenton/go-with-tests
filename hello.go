package main

import "fmt"

// Hello : Replies with a greeting when given a name
func Hello(name string) string {
    return "Hello, " + name
}

func main() {
    fmt.Println(Hello("Sam"))
}