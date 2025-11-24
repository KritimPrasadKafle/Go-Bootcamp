// package main

// import "fmt"

// func main() {

// 	//panic(interface{})

// 	//interface -> any in go

// 	process(10)
// 	process(-2)

// }

// func process(input int) {

// 	defer fmt.Println("Deferred 1")
// 	defer fmt.Println("Deferred 2")
// 	if input < 0 {

// 		fmt.Println("Before Panic")
// 		panic("Input must be a non-negative number")
// 	}

// 	fmt.Println("Processing input:", input)
// }
