package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetInputChoiceFromUser(isSpecialRound bool) (string, error) {

	for {

		inputValue, err := readInputFromUser()

		if err != nil {
			fmt.Println(err)
			return "", err
		}
		inputValue = strings.Replace(inputValue, "\r\n", "", -1)
		// choiceInt,err = strconv.ParseInt(inputValue,0,64);
		if inputValue == "1" {
			return "HEAL", nil
		} else if inputValue == "2" {
			return "ATTACK", nil
		} else if isSpecialRound && inputValue == "3" {
			return "SPECIAL_ATTACK", nil
		} else {
			fmt.Println("You have added an invalid choice , try again ")
		}
	}
}

func readInputFromUser() (string, error) {
	input, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	} else {
		return input, nil
	}
}
