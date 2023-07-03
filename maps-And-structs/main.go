package main

import "fmt"

type WebServices struct {
	google string
	aws    string
}

func main() {

	newWebservicesMap := map[string]string{
		"google": "www.google.com",
		"aws":    "www.aws.com",
	}

	fmt.Println(newWebservicesMap)
}
