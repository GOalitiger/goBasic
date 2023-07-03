package main

import (
	"fmt"
	"strconv"
	"variables/greetings"
)

func main() {

	// staticly-typed
	// var greetingText string = "My new varibale section"

	// implicit type
	// var greetingText = "My new varibale section"

	//inferred type // we use the inferred operator
	// greetingText := "My new varibale section" // by hovering we can see the type
	numberVariable := 17                     // by hovering we can see the type
	tempNumberVariable := numberVariable + 5 // by hovering we can see the type

	fmt.Println(greetings.GreetingsText)
	fmt.Println(numberVariable)
	fmt.Println(tempNumberVariable)

	tempNumberVariable = numberVariable * 3
	fmt.Println(tempNumberVariable)

	// casting int to float as int takes floor value only
	var newFloatVariable float64 = float64(numberVariable) / 3
	var newFloat32Variable float32 = float32(numberVariable) / 3

	fmt.Println(newFloatVariable)
	fmt.Println(newFloat32Variable)

	var runeTempvalue rune = '$' // we can also stores emojis in rune
	var byteTempValue byte = '@'

	// will get the asci or unicode value if we try to get there value;
	// we will need to convert it into string to make it work;
	fmt.Println(string(runeTempvalue))
	fmt.Println(string(byteTempValue))

	// checkNewNum := int("9") + 1
	newNumber, _ := strconv.ParseInt("9", 10, 64)
	fmt.Println(strconv.ParseInt("9", 10, 64))

	fmt.Println(newNumber)

	firstName := "Ali"
	lastName := "Muazzam"
	fullName := firstName + " " + lastName
	var age float32 = 29.3232
	formatedString := fmt.Sprintf("hi, my name is %v and my age is %.2f ", fullName, age)
	fmt.Println(fullName)
	fmt.Println(formatedString)

	fmt.Printf("My mother name is %v ,father name is %v , wife %v , two sister , %v , %v , and a niece %v \n",
		motherName, fatherName, wifeName, sisterName, elderSisterName, nieceName)
	// fmt.Println(motherName)

}
