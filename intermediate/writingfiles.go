// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Method 1: Write bytes
// 	file1, err := os.Create("output.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file1.Close()

// 	file1.Write([]byte("Written with Write()\n"))
// 	fmt.Println("output.txt created")

// 	// Method 2: Write string
// 	file2, err := os.Create("writeString.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	defer file2.Close()

// 	file2.WriteString("Written with WriteString()\n")
// 	fmt.Println("writeString.txt created")
// }
