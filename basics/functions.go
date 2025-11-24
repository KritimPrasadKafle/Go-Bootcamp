// package main

// import "fmt"

// func main() {

// 	//func <name>(parameters list) returnType{
// 	//code block
// 	//return value
// 	// }

// 	fmt.Println()

// 	sum := add(1, 2)
// 	fmt.Println(sum)

// 	greet := func() {
// 		fmt.Println("Hello Anonymous Function")
// 	}

// 	greet()

// 	operation := add

// 	result := operation(3, 5)

// 	fmt.Println(result)

// 	result1 := applyOperation(5, 3, add)
// 	fmt.Println("5 + 3 = ", result1)

// 	multiplyBy2 := createMultiplier(2)
// 	fmt.Println("6 * 2 = ", multiplyBy2(6))

// }
// func add(a, b int) int {
// 	return a + b
// }

// func applyOperation(x int, y int, operation func(int, int) int) int {
// 	return operation(x, y)
// }

// func createMultiplier(factor int) func(int) int {
// 	return func(i int) int {
// 		return i * factor
// 	}
// }
