package main

import "fmt"

func variable() {
	// fmt.Println("GOOSE_CPT")
	// var nameOne string = "This a string, dude"
	// var nameTwo = "Peaches"
	// var nameThree string //^ Declering variable

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameOne = "MTG"

	// fmt.Println(nameOne, nameTwo, nameThree)

	// nameFour := "This is a short hand to declare a variable and GO automatically infers type string"

	// fmt.Println(nameFour, nameTwo, nameThree)

	//^ INTs
	var ageOne int = 20
	var ageTwp = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwp, ageThree)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255 //Cant be negative, can go beyond int64

	fmt.Println(numOne, numTwo, numThree)

	// Floats
	var scoreOne float32 = -1.543
	var scoreTwo float64 = 646329384637284.33
	scoreThree := 65.33
	fmt.Println(scoreOne, scoreTwo, scoreThree)

}
