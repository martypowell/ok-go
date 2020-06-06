package main

import "strconv"

type pet struct {
	animalType string
	name       string
	age        int
}

// Get Pet Age for a given pet
func (pet *pet) getGreeting() string {
	return "Hi, my name is " + pet.name + ". I am a " + strconv.Itoa(pet.age) + " year old, " + pet.animalType
}
