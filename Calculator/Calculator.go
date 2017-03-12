package main

import "fmt"


func main(){
	var function int
	var input float32
	var input2 float32

	fmt.Print("Welcome ! For Addition Press 1\nFor Summation, Press 2\nFor Production, Press 3\nFor Division, Press 4.\n\n")
	fmt.Print("Function : ")
	fmt.Scan(&function)
	if(function ==1){
		fmt.Println("Function choosed : Addition\n")
	}
	if(function == 2){
		fmt.Println("Function choosed : Summation\n")
	}
	if(function == 3){
		fmt.Println("Function choosed : Production\n")
	}
	if(function == 4){
		fmt.Println("Function choosed : Division\n")
	}

	// check for valid function choices
	for(function > 4 || function <= 0){
		fmt.Println("ERROR ! For Addition Press 1\nFor Summation, Press 2\nFor Production, Press 3\nFor Division, Press 4.\n")
		fmt.Println("Please Try Again.")
		fmt.Print("Function : ")
		fmt.Scan(&function)
	}

	// Enter two number you want
	fmt.Println("Enter First number Please")
	fmt.Scan(&input)
	fmt.Println("Enter Second number Please")
	fmt.Scan(&input2)

	// Function 1: Addition
	if(function == 1){
		fmt.Println("The Result is ", input+input2)
	}

	// Function 2 : Summation
	if(function == 2){
		if(input >= input2){
			fmt.Println("The Result is ", input-input2)
		}
		if(input < input2){
			fmt.Println("The Result is ", input2-input)
		}
	}

	// Function 3 : Production
	if(function == 3){
		fmt.Println("The result is ", input*input2)
	}

	// Function 4 : Division
	if(function == 4){
		for(input2 == 0){
			fmt.Println("Please re-enter the second number. 0 is not allowed")
			fmt.Scan(&input2)
		}
		fmt.Println("The result is ", input/input2)
	}
 }
