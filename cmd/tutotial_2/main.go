// function and Control structures
package main

import (
	"errors"
	"fmt"
)

func main() {
	x := "You are good"
	printMe(x)
	var numerator, denominator int = 23, 0
	var result, remainder, err = intDivision(numerator, denominator)
	// if err != nil {
	// 	fmt.Printf(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("the remainder is %v",result)
	// }else {
	// 	fmt.Printf("the remainder is %v",result)
	// }
	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("the remainder is %v", result)
	default:
		fmt.Printf("the remainder is %v", result)
	}
}

// func printMe() {
// 	fmt.Println("hello world")
// }

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("Cannot Divide the num")
		return 0, 0, err
	}
	result := numerator / denominator
	remainder := numerator % denominator
	return result, remainder, err

}
