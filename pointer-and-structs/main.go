package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var reader = bufio.NewReader(os.Stdin)

type User struct {
	firstName   string
	lastName    string
	fatherName  string
	dateOfBirth string
	createdDate time.Time
}

// in Structs we can also call the properties without using * in the start
// we can also add recieveing function in the start of the func to attach this fucntion to the struct
func (user User) OutputUser() {
	fmt.Printf("user details , first name = %v , lastname = %v ,father name =%v ,dateofbirth = %v,createdDate =%v",
		user.firstName, user.lastName, user.fatherName, user.dateOfBirth, user.createdDate)
	// (*user).firstName   // this is the orignal way but lexer does it for us
}

func main() {

	value := 18
	tempPtr := &value
	var newTempPtr *int = &value

	fmt.Println(tempPtr)
	fmt.Println(*newTempPtr)

	firstName := replaceInputValues(promptMsg("Add First Name:"))
	lastName := replaceInputValues(promptMsg("Add Last Name:"))
	fatherName := replaceInputValues(promptMsg("Add Father Name:"))
	dob := replaceInputValues(promptMsg("Add DOB in dd/mm/yyyy:"))

	newUser := mapUserDetails(firstName, lastName, fatherName, dob)

	newUser.OutputUser()
	fmt.Println('\n')
	//this is the direct way to get the values in the struct instance
	outputUserBypointer(newUser)
}

func promptMsg(msg string) string {
	fmt.Print(msg)
	inputMsg, _ := reader.ReadString('\n')
	return inputMsg
}

func replaceInputValues(value string) string {
	return strings.Replace(value, "\r\n", "", -1)
}

// we can create struct like this and we can also create struct by passing values in the same sequence
func mapUserDetails(fName string, lName string, fatherName string, dob string) *User {
	createdDate := time.Now()
	user := User{
		firstName:   fName,
		lastName:    lName,
		fatherName:  fatherName,
		dateOfBirth: dob,
		createdDate: createdDate,
	}

	return &user
}

func outputUserBypointer(user *User) {
	fmt.Printf("user details , first name = %v , lastname = %v ,father name =%v ,dateofbirth = %v,createdDate =%v",
		(*user).firstName, user.lastName, user.fatherName, user.dateOfBirth, user.createdDate)
	// (*user).firstName   // this is the orignal way but lexer does it for us it dereferences in the case of struct
}
