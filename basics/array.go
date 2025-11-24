package main

import "fmt"

// func main() {
// 	var numbers [5]int
// 	fmt.Println(numbers)

// 	numbers[4] = 20
// 	fmt.Println(numbers)

// 	numbers[1] = 0
// 	fmt.Println(numbers)

// 	fruits := [4]string{"Apple", "Banana", "Orange", "Grapes"}

// 	fmt.Println("Fruits array:", fruits)
// 	fmt.Println("Third element:", fruits[2])

// 	originalArray := [3]int{1, 2, 3}
// 	copiedArray := originalArray

// 	copiedArray[0] = 100

// 	fmt.Println("Original array:", originalArray)
// 	fmt.Println("Copied array:", copiedArray)

// 	for i := 0; i < len(numbers); i++ {
// 		fmt.Println("Element at index", i, ":", numbers[i])

// 	}

// 	for index, value := range numbers {
// 		fmt.Printf("Index: %d, value: %d\n", index, value)
// 	}

// check1 := [4]int{1, 2, 3, 4}
// s = array(check1)
// fmt.Println(s)
// }

// func someFunction() (int, int) {
// 	return 1, 2
// }

// func array(check [4]int) {

// 	for i := 0; i < len(check); i++ {
// 		fmt.Println(check[i])
// 	}

// }

// func main() {

// 	check1 := [4]int{1, 2, 3, 4}
// 	array(check1)
// }

// Define a custom type
type Student struct {
	Name   string
	RollNo int
	Absent bool
}

func main() {
	// Create an array of 3 students
	var students [3]Student

	students[0] = Student{Name: "Alice", RollNo: 1, Absent: false}
	students[1] = Student{Name: "Bob", RollNo: 2, Absent: true}
	students[2] = Student{Name: "Charlie", RollNo: 3, Absent: false}

	// Loop through and print
	for _, s := range students {
		status := "Present"
		if s.Absent {
			status = "Absent"
		}
		fmt.Printf("Name: %s, Roll No: %d, Status: %s\n", s.Name, s.RollNo, status)
	}
}
