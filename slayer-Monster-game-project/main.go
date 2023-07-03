package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"slayerMonster/actions"
	"slayerMonster/interaction"
)

var countingRound int = 0
var winnerDeclared bool = false

type StatsLogger struct {
	RoundNumber         int
	PlayerChoice        string
	PlayerAttackDamage  int
	MonsterAttackDamage int
	PlayerHealth        int
	MonsterHealth       int
	PlayerHealValue     int
}

var statLoggerArr = []StatsLogger{}

type name string

func main() {

	// fName := name("ali")
	// fmt.Println(fName)
	//---------------------------------
	//closure function example
	// double := tranformNumberFuctions(2)
	// triple := tranformNumberFuctions(3)

	// fmt.Println(double(15))
	// fmt.Println(triple(15))
	//---------------------------------

	startGame()
	executeRound()
	if winnerDeclared {
		endGame()
	}
}

func startGame() {
	interaction.ShowGreetings()
}

func executeRound() {

	for !winnerDeclared {
		countingRound++
		isSpecialRound := countingRound%3 == 0
		interaction.PlayerChooseActionOutput(isSpecialRound)
		input, err := interaction.GetInputChoiceFromUser(isSpecialRound)

		if err != nil {
			fmt.Println(err)
			return
		}
		var healVal int
		var playerAttackVal int
		var monsterAttackVal int
		var statsForRound StatsLogger
		// write code for input we got from user
		if input == "HEAL" {
			healVal = actions.HealPlayer()
		} else if input == "ATTACK" || input == "SPECIAL_ATTACK" {
			playerAttackVal = actions.AttackMonster(isSpecialRound)
		}
		monsterAttackVal = actions.AttackPlayer()

		currPlayerHealth, currMonsterHealth := actions.GetPlayerAndMonsterHealth()

		statsForRound = StatsLogger{countingRound,
			input,
			playerAttackVal,
			monsterAttackVal,
			currPlayerHealth,
			currMonsterHealth,
			healVal,
		}

		statLoggerArr = append(statLoggerArr, statsForRound)

		val, err := json.Marshal(statsForRound)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(val))
		}

		// decodedStruct := StatsLogger{}
		// err1 := json.Unmarshal(val, &decodedStruct)
		// if err1 != nil {
		// 	fmt.Println(err1)
		// } else {
		// 	fmt.Println(decodedStruct)
		// }

		// for i := 0; i < len(strForDisplayArr) ;i++ {
		// 	strForDisplayArr +=
		// }
		// execute function on the basis of return values
		// checkForWinner
		result := actions.CheckWinner()
		// if some winner then break the loop and declare winner and set winner declared
		if result != "" {
			interaction.OutputResult(result)
			winnerDeclared = true
		}
		// fmt.Println(input)

	}
}

func endGame() {

	// When we need to log in standalone executable mode after creating a binary by using go build
	expath, err := os.Executable()
	if err != nil {
		fmt.Println("Executable not found")
		return
	}
	expath = filepath.Dir(expath)
	file, err := os.Create(expath + "/loggerNew.txt")

	// file, err := os.Create("logger.txt") // when we need to run in dev mode using go run .

	if err != nil {
		fmt.Println("File creation failed")
		return
	}

	// this is a way in which we can get an error if the closing fails
	// defer will make this function run before the execution of this function
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println("File closing failed")
			return
		}
	}()

	// statLoggerArr
	for _, statsForRound := range statLoggerArr {
		// var resultMap = map[string]StatsLogger{}
		strForDisplayArr := map[string]interface{}{"Round No": statsForRound.RoundNumber,
			"PlayerChoice":        statsForRound.PlayerChoice,
			"PlayerAttackDamage":  statsForRound.PlayerAttackDamage,
			"MonsterAttackDamage": statsForRound.MonsterAttackDamage,
			"PlayerHealth":        statsForRound.PlayerHealth,
			"MonsterHealth":       statsForRound.MonsterHealth,
			"PlayerHealValue":     statsForRound.PlayerHealValue}
		file.WriteString(fmt.Sprintln(strForDisplayArr))
	}

}

// classic example of closure function
// this is factory design way of implementing function
// a way to generate function with configurations we give through parameter
// actually it locks the variable value for that function scope,
// for example the inner anonymous fucntion gets the value of factor once passed in the outer function
func tranformNumberFuctions(factor int) func(number int) int {
	return func(number int) int {
		return number * factor
	}
}

func checkType(value interface{}) (interface{}, bool) {
	val, ok := value.(string)
	return val, ok
}
