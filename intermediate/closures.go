// package main

// import "fmt"

// func main() {

// 	sequence := adder()

// 	fmt.Println(sequence())
// 	fmt.Println(sequence())
// 	fmt.Println(sequence())

// 	subtracter := func() func(int) int {
// 		countdown := 99

// 		return func(i int) int {
// 			countdown -= i
// 			return countdown

// 		}

// 	}()

// 	fmt.Println(subtracter(1))
// 	fmt.Println(subtracter(2))
// 	fmt.Println(subtracter(3))
// 	fmt.Println(subtracter(4))

// }

// func adder() func() int {
// 	i := 0
// 	fmt.Println("Previous value of i:", i)
// 	return func() int {
// 		fmt.Println("added 1 to i")
// 		return i
// 	}
// }
