package main

import (
	"fmt"
)

func main() {
	var operator string
	var num1 float64
	var num2 float64
	var answer float64


fmt.Println("This is a calculator")
	sum := 1
	for sum < 2 {
	fmt.Println("What operator would you like to use? (+, -, *, /)")
	fmt.Scan(&operator)
	fmt.Println("Good, now choose the two numbers you want to use (num num)")
	fmt.Scan(&num1, &num2)

	switch operator {
	case "+":
		answer = num1 + num2
	case "-":
		answer = num1 - num2
	case "*":
		answer = num1 * num2
	case "/":
		answer = num1 / num2
	default:
		fmt.Println("Invalid operator")
		continue
	}
	fmt.Println("The answer is: ", answer)
	sum++
}
}