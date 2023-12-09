package main

import "fmt"

func main() {

	// Declering array and setting array and setting the array type
	var ages [3]int = [3]int{20, 25, 30}
	var ages2 = [3]int{20, 25, 30}

	names := [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	names[1] = "Luigi"

	fmt.Println(ages, ages2, len(ages), names, "\n")

	//Slices -> Can change length, use array under the hood
	var scores = []int{100, 43, 223, 194, 389} //^ Dont put number in to have unlimted numbers
	scores[2] = 444
	fmt.Printf("Scores before append %v \n", scores)

	scores = append(scores, 69420) //returns new slice, have to make it equal to the orginal slice to append it

	fmt.Printf("Scores after append %v \n", scores)

	//Slice ranges
	rangeOne := names[1:3]

	fmt.Println(rangeOne)

}
