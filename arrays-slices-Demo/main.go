package main

import "fmt"

type Product struct {
	id          string
	title       string
	Description string
}

func main() {
	var newStaticArr = [5]string{"ali", "manahal", "sarah", "tara", "ayesha"}

	sliceFromNewArray := newStaticArr[1:] // carving slices
	slice2FromNewArray := newStaticArr[:3]

	sliceFromNewArray[0] = "manahal motto"
	fmt.Println(sliceFromNewArray, len(sliceFromNewArray), cap(sliceFromNewArray))
	fmt.Println(slice2FromNewArray, len(slice2FromNewArray), cap(slice2FromNewArray))

	dynamicSlice := []Product{}
	// here we must understand that here we have slice formed which is at the BE handled by go
	// the creates a new array and then returns a slice

	fmt.Println("the values of dynamic array= ", dynamicSlice, len(dynamicSlice), cap(dynamicSlice))

	dynamicSlice = append(dynamicSlice, Product{id: "13", title: "carpet", Description: "a beautiful carpet"})
	dynamicSlice = append(dynamicSlice, Product{id: "12", title: "Screw Driver", Description: "for car engines and chairs"})

	fmt.Println(dynamicSlice)
	// tempSlice := dynamicSlice[:1]
	dynamicSlice = dynamicSlice[1:]
	// fmt.Println(tempSlice)
	dynamicSlice = dynamicSlice[0:]
	fmt.Println(dynamicSlice)
	// fmt.Println(tempSlice[0:]) //[{13 carpet a beautiful carpet}]
	// fmt.Println(tempSlice[0:2]) //[{13 carpet a beautiful carpet} {12 Screw Driver for car engines and chairs}]

	// concatenation using spread operator just like in javascript
	var strList1 []string = []string{"Ali", "Manahal", "tara", "sara"}
	var strList2 []string = []string{"ayesha", "umer", "usman", "rabia"}

	fmt.Println(append(strList1, strList2...))

}
