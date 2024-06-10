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
	passedCount := 0
	totalCount := 0
	for _, entry := range dir {

		file, err := os.Open(fmt.Sprintf("%s/%s", directory, entry.Name()))
		if err != nil {
			fmt.Printf("error opening file: %s, %s \n", entry.Name(), err)
			return
		}
		result, err := validator.Validate(file)

		if entry.Name()[0:4] == "fail" && err != nil {
			fmt.Printf("Filename : %s, result %d: PASSED\n", entry.Name(), result)
			passedCount++
		} else if err == nil {
			fmt.Printf("Filename : %s, result %d: PASSED\n", entry.Name(), result)
			passedCount++
		} else {
			fmt.Printf("Filename : %s, result %d: FAILED, error is %s \n", entry.Name(), result, err.Error())
		}
		totalCount++
	}
	fmt.Printf("Total: %d, Passed: %d,\n", totalCount, passedCount)
	return
}
