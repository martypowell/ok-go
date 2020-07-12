package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	const messagePrefix string = "hello world"

	animals := []string{"Dog", "Cat", "Elephant"}
	pets := []pet{pet{name: "Bob", age: 1, animalType: "dog"}, pet{name: "Grumpy", age: 10, animalType: "cat"}}

	timesToRun :=
		flag.Int("timesToRun", 1, "Number of times you want to display a message.")
	message := messagePrefix + " marty in variable, with constant"

	flag.Parse()

	// Print message based on the number of times to run
	for i := 0; i < *timesToRun; i++ {
		fmt.Println(message + ", ran " + strconv.Itoa(i+1) + " times")
	}

	logAnimals(animals)
	logPets(pets)

	// Test the multiple return value logic
	fmt.Println(getBasicArithmetic(2, 2))

	// Test Struct Method
	fmt.Println(pets[1].getGreeting())

}

// Write a list of animals to the terminal
func logPets(pets []pet) {
	// Loop through animals
	for _, pet := range pets {
		fmt.Println(pet.name + " is a " + strconv.Itoa(pet.age) + " year old " + pet.animalType)
	}
}

// Write a message to the terminal that describes a pet's age
func logAnimals(animals []string) {
	// Loop through animals
	for _, animal := range animals {
		fmt.Println(animal)
	}
}
