package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Product struct {
	ID                 string
	ProductType        string
	ProductDescription string
	price              float64
}

func (prod Product) OutputProduct() {
	fmt.Printf(
		"Product ID= %v\n Product type = %v\n Product Name = %v\n Product PKR price = %v\n",
		prod.ID,
		prod.ProductType,
		prod.ProductDescription,
		prod.price)
}

func storeInFile(product *Product) {
	file, _ := os.Create(product.ID + ".txt")
	content := fmt.Sprintf(
		"Product ID : %v\nProduct type : %v\nProduct Name : %v\nProduct PKR price : %.2f\n",
		product.ID,
		product.ProductType,
		product.ProductDescription,
		product.price)

	file.WriteString(content)
	file.Close()
}

func main() {

	firstProd := Product{
		ID:                 "b-carpet",
		ProductType:        "carpets",
		ProductDescription: "A beautiful carpet",
		price:              8000,
	}
	storeInFile(&firstProd)
	//get values and fill in struct instance
	reader := bufio.NewReader(os.Stdin)

	id := getValueFromUser("Enter Product ID:", reader)
	prodType := getValueFromUser("Enter Product Type:", reader)
	prodDesc := getValueFromUser("Enter Product Description:", reader)
	prodPrice := getValueFromUser("Enter Product Price:", reader)

	prodPriceInFloat, _ := strconv.ParseFloat(prodPrice, 64)

	//replace \n
	//output the instance of struct
	newProduct := mapToStruct(id, prodType, prodDesc, prodPriceInFloat)
	newProduct.OutputProduct()
	storeInFile(newProduct)

	//create file
}

func getValueFromUser(promptMsg string, reader *bufio.Reader) string {
	fmt.Print(promptMsg)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\r\n", "", -1)
	return input
}

func mapToStruct(id string, pt string, pd string, pp float64) *Product {
	newProd := Product{id, pt, pd, pp}
	return &newProd
}
