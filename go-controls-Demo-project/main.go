package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader = bufio.NewReader(os.Stdin)
	input, err := getUserInput(reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	switch input {
	case "1":
		getSumUptilNNumber(reader)
	case "2":
		getFactorialValueOfNumberN(reader)
	case "3":
		getSumOfManuallyAddedValues(reader)
	default:
		getSumOfListOfValues(reader)

	}

	// if input == "1" {
	// 	getSumUptilNNumber(reader)
	// } else if input == "2" {
	// 	getFactorialValueOfNumberN(reader)
	// } else if input == "3" {
	// 	getSumOfManuallyAddedValues(reader)
	// } else {
	// 	getSumOfListOfValues(reader)
	// }
}

func getUserInputGeneric(reader *bufio.Reader) (int, error) {
	numberInput, _ := reader.ReadString('\n')
	numberInput = strings.Replace(numberInput, "\r\n", "", -1)
	numInt, err := strconv.ParseInt(numberInput, 0, 64)

	if err != nil {
		// fmt.Println(err)
		return -1, err
	}

	return int(numInt), nil
}

func getSumUptilNNumber(reader *bufio.Reader) {
	fmt.Println("Enter the number N , for sum till that number:")
	numberInput, _ := reader.ReadString('\n')
	numberInput = strings.Replace(numberInput, "\r\n", "", -1)
	numInt, err := strconv.ParseInt(numberInput, 0, 64)

	if err != nil {
		fmt.Println(err)
		return
	}

	sum := 0
	for i := 1; i <= int(numInt); i++ {
		sum += i
	}

	fmt.Printf("The sum of number N = %v", sum)
}

func getFactorialValueOfNumberN(reader *bufio.Reader) {
	fmt.Println("Enter the number N , for factorial:")

	input, err := getUserInputGeneric(reader)

	if err != nil {
		fmt.Println(err)
		return
	}

	fact := 1
	for i := input; i >= 1; i-- {
		fact *= i

	}

	fmt.Printf("factorial = %v", fact)

}

func getSumOfManuallyAddedValues(reader *bufio.Reader) {
	isLoopRunning := false
	sum := 0
	fmt.Println("Enter values to add sum manually and enter any other value to stop entering and getting sum")
	for !isLoopRunning {
		input, err := getUserInputGeneric(reader)
		if err != nil {
			isLoopRunning = true
			fmt.Printf("Sum of manual values = %v", sum)
		} else {
			sum += input
		}

	}

}

func getSumOfListOfValues(reader *bufio.Reader) {
	fmt.Println("Pls Enter the values in a single list separated by single white space")
	input, err := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)

	if err != nil {
		fmt.Println(err)
		return
	}

	inputList := strings.Split(input, " ")
	sum := 0
	// this kind of loop can be used to iterate through maps(via key:values) and slices , arrays all
	for index, value := range inputList {
		fmt.Printf("At index = %v , value =%v \n", index, value)
		intVal, err := strconv.ParseInt(value, 0, 64)

		if err != nil {
			fmt.Println("Invalid value in the list")
			// fmt.Printf("Result for list ti = %v", sum)
			continue
			// break
		}
		sum += int(intVal)
	}
	fmt.Printf("Result for list = %v", sum)
}

func getUserInput(reader *bufio.Reader) (string, error) {

	fmt.Println("Pls Enter your choice from the following:")
	choicesMsg := "1) get the sum upto a number N \n2) get the factorial of number N \n3) get the sum of manually entered numbers by the users \n4) get the sum of the list of numbers entered"

	fmt.Println(choicesMsg)
	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\r\n", "", -1)
	// checkInt, err = strconv.ParseInt(userInput, 0, 64)

	// if err != nil {
	// 	return "", errors.New("Invalid selection")
	// }

	checkArr := "1234"

	if strings.Contains(checkArr, userInput) {
		return userInput, nil
	} else {
		return "", errors.New("INVALID SELECTION")
	}
}
