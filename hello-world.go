package main

import (
	"fmt"
	"strconv"
)

func main() {
	const numberOfTimesToMessage = 2
	const messagePrefix string = "hello world"
	var message = messagePrefix + " marty in variable, with constant"

	for i := 0; i < numberOfTimesToMessage; i++ {
		fmt.Println(message + ", ran " + strconv.Itoa(i) + " times")
	}

}
