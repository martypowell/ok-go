package main

import "fmt"

func main() {
	const messagePrefix string = "hello world"
	var message = messagePrefix + " marty in variable, with constant"

	fmt.Println(message)
}
