package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	const messagePrefix string = "hello world"

	animals := []string{"Dog", "Cat", "Elephant"}

	timesToRun :=
		flag.Int("timesToRun", 1, "Number of times you want to display a message.")
	message := messagePrefix + " marty in variable, with constant"

	flag.Parse()

	// Print message based on the number of times to run
	for i := 0; i < *timesToRun; i++ {
		fmt.Println(message + ", ran " + strconv.Itoa(i+1) + " times")
	}

	logAnimals(animals)
}

// Write a list of animals to the terminal
func logAnimals(animals []string) {
	// Loop through animals
	for _, animal := range animals {
		fmt.Println(animal)
	}
}
