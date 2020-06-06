package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	const messagePrefix string = "hello world"

	animals := []string{"Dog", "Cat", "Elephant"}
	petAges := map[string]int{"Bob": 1, "Sparky": 10, "Bailey": 2}

	timesToRun :=
		flag.Int("timesToRun", 1, "Number of times you want to display a message.")
	message := messagePrefix + " marty in variable, with constant"

	flag.Parse()

	// Print message based on the number of times to run
	for i := 0; i < *timesToRun; i++ {
		fmt.Println(message + ", ran " + strconv.Itoa(i+1) + " times")
	}

	logAnimals(animals)
	logPets(petAges)
}

// Write a list of animals to the terminal
func logPets(petAges map[string]int) {
	// Loop through animals
	for name, age := range petAges {
		fmt.Println(name + " is " + strconv.Itoa(age) + " years old")
	}
}

// Write a message to the terminal that describes a pet's age
func logAnimals(animals []string) {
	// Loop through animals
	for _, animal := range animals {
		fmt.Println(animal)
	}
}
