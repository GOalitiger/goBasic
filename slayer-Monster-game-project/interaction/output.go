package interaction

import "fmt"

func ShowGreetings() {
	fmt.Println("Welcome to MONSTER SLAYER")
	fmt.Println("------------------------")
	fmt.Println("Starting the game!")
}

func PlayerChooseActionOutput(isSpecialRound bool) {
	fmt.Println("Choose the action from the following")
	fmt.Println("1) Heal")
	fmt.Println("2) Attack")

	if isSpecialRound {
		fmt.Println("3) SpecialAttack")
	}
}

func OutputResult(result string) {
	fmt.Println("-----------------------")
	fmt.Println("Result = ", result)
	fmt.Println("-----------------------")
}
