package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := getUserInput()

	//nil is just like null or no value
	//errror is handled in this way we need to check the error
	// if there is error we can use the return nothing statement to break the main function

	if err != nil {
		// we can also generate our own error

		// fmt.Println(errors.New("Invalid Syntax"))
		fmt.Println(err)
		return
	}

	// execution code conditionally
	if input >= 30 && input <= 60 {
		fmt.Println("You are eligible for our senior job")
	} else if input >= 18 && input < 30 {
		fmt.Println("You are welcome to club")
	} else {
		fmt.Println("you are not old enough")
	}

}

// here we can also return the error and the int value we get from the user
func getUserInput() (int, error) {
	var reader = bufio.NewReader(os.Stdin)
	fmt.Print("Enter your age :")
	inputAge, _ := reader.ReadString('\n')
	inputAge = strings.Replace(inputAge, "\r\n", "", -1)
	intInputAge, err := strconv.ParseInt(inputAge, 0, 64)

	return int(intInputAge), err
}
