package main

import (
	"fmt"
	"os"
)

func main() {
	directory := "./test"
	dir, err := os.ReadDir("./test")
	if err != nil {
		fmt.Printf("error opening directiony: %s, %s \n", directory, err)
		return
	}

	lexer := NewLexer()
	parser := NewParser()
	validator := NewJSONValidator(lexer, parser)
	for _, entry := range dir {
		file, err := os.Open(fmt.Sprintf("%s/%s", directory, entry.Name()))
		if err != nil {
			fmt.Printf("error opening file: %s, %s \n", entry.Name(), err)
			return
		}
		result, err := validator.Validate(file)
		if err != nil {
			fmt.Printf("error validating file: %s, %s \n", entry.Name(), err)
			return
		}
		fmt.Printf("Filename : %s, result %d \n", entry.Name(), result)
	}
	return
}
