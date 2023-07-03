package main

import (
	"BMICalculator/info"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func main() {

	// show title and separator
	fmt.Println(info.MainTitle)
	fmt.Println(info.Separator)

	// take inputs from user store in a variable and convert to float64
	fmt.Println(info.HeightPrompt)
	heightInput, _ := Reader.ReadString('\n')

	fmt.Println(info.WeightPrompt)
	weightInput, _ := Reader.ReadString('\n')
	fmt.Println(info.Separator)
	fmt.Print(heightInput)
	fmt.Print(weightInput)

	weightInput = strings.Replace(weightInput, "\r\n", "", -1)
	heightInput = strings.Replace(heightInput, "\r\n", "", -1)
	fmt.Println(info.Separator)
	fmt.Println(heightInput, weightInput)
	fmt.Println(info.Separator)
	// fmt.Print(weightInput)

	weight := newFunction(weightInput)
	// height, _ := strconv.ParseFloat(heightInput, 64)
	height := newFunction(heightInput)

	fmt.Println(weight, height)
	fmt.Println(info.Separator)

	//calculate BMI and show
	BMI := calculateBMI(height, weight)
	fmt.Printf("\n The BMI calculated is  %f ", BMI)
}

func newFunction(Input string) float64 {
	inputFloatValue, _ := strconv.ParseFloat(Input, 64)
	return inputFloatValue
}

func calculateBMI(height float64, wieght float64) (bmi float64) {
	bmi = wieght / (height * height)
	return
}

func generateRandomNumber() (int, int) {
	num1 := rand.Intn(1000)
	num2 := rand.Intn(1000)
	return num1, num2
}

func summingTwo(a int, b int) (sum int) {
	sum = a + b
	return sum
}
