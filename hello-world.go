package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {

	timesToRun :=
		flag.Int("timesToRun", 1, "Number of times you want to display a message.")
	const messagePrefix string = "hello world"
	var message = messagePrefix + " marty in variable, with constant"

	flag.Parse()

	for i := 0; i < *timesToRun; i++ {
		fmt.Println(message + ", ran " + strconv.Itoa(i+1) + " times")
	}

}
