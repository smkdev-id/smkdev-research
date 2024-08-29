package main

import (
	"fmt"
	"os/exec"
)

// ! DON'T CHANGE OR REMOVE THIS CODE BELOW. FOR TESTING PURPOSES
func main() {

	cmd := exec.Command(
		"go",
		"test",
		"arithmetics_calculator.go",
		"arithmetics_calculator_test.go",
		"-v",
		"-coverprofile=coverage.out",
	)

	out, err := cmd.Output()

	if err != nil {
		fmt.Println("could not run command: ", err)
	}

	fmt.Println("================OUTPUT==================")
	fmt.Println(string(out))
}

// ! DON'T CHANGE OR REMOVE THIS CODE BELOW. FOR TESTING PURPOSES
