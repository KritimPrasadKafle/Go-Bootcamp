// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func main() {

// 	//func functionName(parameter1 type1, parameter2 type2,...) (returnType1, returnType2, ...){
// 	//code block
// 	//reurn retturnvalue1, returnvalue2,...
// 	//}

// 	q, r := divide(10, 3)
// 	fmt.Println("Quotient: %v, Remainder: %v\n", q, r)

// 	result, err := compare(4, 5)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 	} else {
// 		fmt.Println(result)
// 	}
// }

// func divide(a, b int) (int, int) {
// 	quotient := a / b
// 	remainder := a % b

// 	return quotient, remainder
// }

// func compare(a, b int) (string, error) {
// 	if a > b {
// 		return "a is greater than b", nil
// 	} else if b > a {
// 		return "b is greater than a", nil
// 	} else {
// 		return "", errors.New("Unnable to compare which s greater")
// 	}
// }
