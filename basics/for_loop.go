// package main

// import "fmt"

// func main() {
// 	// for i := 1, i<=5, i++ {
// 	// 	fmt.Println(i)
// 	// }

// 	numbers := []int{1, 2, 3, 4, 5, 6}
// 	for index, value := range numbers {
// 		fmt.Printf("Index: %d, Value: %d\n", index, value)
// 	}

// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			continue
// 		}
// 		fmt.Println("Odd Number:", i)
// 		if i == 5 {
// 			break
// 		}
// 	}

// 	rows := 5

// 	//Outer loop

// 	for i := 1; i <= rows; i++ {
// 		for j := 1; j <= rows-i; j++ {
// 			fmt.Println(" ")
// 		}
// 		for k := 1; k <= 2*i-1; k++ {
// 			fmt.Println("*")
// 		}
// 		fmt.Println()
// 	}

// 	for i:= range 10{
// 		fmt.Println(10 - i)
// 	}
// 	fmt.Println("We have a lift off!s")
// }
