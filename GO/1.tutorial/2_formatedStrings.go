package main

import "fmt"

func formatedStrings() {
	age := 35
	name := "Shuan"

	// Print
	fmt.Println("Hello")
	fmt.Println("World! \n")
	fmt.Println("World")

	fmt.Println("Hello")
	fmt.Println("goodbye")
	fmt.Println("My age is ", age, "and my name is", name)

	// Formating Strings %_ = format specifier
	fmt.Printf("my current age is %v and my name is %v \n", age, name)

	//v for int, q for string
	fmt.Printf("my current age is %v and my name is %q \n", age, name)

	//Getting type of variable
	fmt.Printf("age os of type %T \n", age)

	fmt.Printf("You scored points :(%0.2f) \n", 69.33)
	fmt.Printf("You scored points :(%0.2f) \n", 69.33)

	//Sprintf (save formatted string)
	var str = fmt.Sprintf("This is a sprinted string %v and go is pretty useful %v", age, name)

	fmt.Println(str)

}
