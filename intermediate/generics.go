// package main

// // Need separate functions for each type
// func swapInt(a, b int) (int, int) {
// 	return b, a
// }

// func swapString(a, b string) (string, string) {
// 	return b, a
// }

// func swapFloat(a, b float64) (float64, float64) {
// 	return b, a
// }

// func swap[T any](a, b T) (T, T) {
// 	return b, a
// }

// // Annoying! Same logic, different types 😫
// func main() {
// 	swap(1, 2)        // Works with int
// 	swap("hi", "bye") // Works with string
// 	swap(3.14, 2.71)  // Works with float64
// }

// package main
// import "fmt"

// func swap[T any](a, b T) (T, T) {
//     return b, a
// }

// func main() {
//     // Works with int
//     x, y := 1, 2
//     x, y = swap(x, y)
//     fmt.Println(x, y)  // 2 1

//     // Works with string
//     s1, s2 := "Hello", "World"
//     s1, s2 = swap(s1, s2)
//     fmt.Println(s1, s2)  // World Hello

//     // Works with float64
//     f1, f2 := 3.14, 2.71
//     f1, f2 = swap(f1, f2)
//     fmt.Println(f1, f2)  // 2.71 3.14
// }

// package main
// import "fmt"

// // Constraint: T must be a type that can be compared with < and >
// func min[T int | float64 | string](a, b T) T {
//     if a < b {
//         return a
//     }
//     return b
// }

// func main() {
//     fmt.Println(min(5, 3))           // 3
//     fmt.Println(min(3.14, 2.71))     // 2.71
//     fmt.Println(min("apple", "banana")) // apple
// }

// package main
// import "fmt"

// // Apply a function to each element
// func mapSlice[T any, U any](slice []T, fn func(T) U) []U {
//     result := make([]U, len(slice))
//     for i, v := range slice {
//         result[i] = fn(v)
//     }
//     return result
// }

// func main() {
//     // Double each number
//     numbers := []int{1, 2, 3, 4, 5}
//     doubled := mapSlice(numbers, func(n int) int {
//         return n * 2
//     })
//     fmt.Println(doubled)  // [2 4 6 8 10]

//     // Convert int to string
//     strings := mapSlice(numbers, func(n int) string {
//         return fmt.Sprintf("Number: %d", n)
//     })
//     fmt.Println(strings)  // [Number: 1 Number: 2 ...]

//     // Get string lengths
//     words := []string{"hello", "world", "go"}
//     lengths := mapSlice(words, func(s string) int {
//         return len(s)
//     })
//     fmt.Println(lengths)  // [5 5 2]
// }

// package main
// import "fmt"

// type Pair[T any, U any] struct {
//     First  T
//     Second U
// }

// func (p Pair[T, U]) String() string {
//     return fmt.Sprintf("(%v, %v)", p.First, p.Second)
// }

// func main() {
//     // int and string
//     p1 := Pair[int, string]{First: 1, Second: "one"}
//     fmt.Println(p1)  // (1, one)

//     // string and bool
//     p2 := Pair[string, bool]{First: "active", Second: true}
//     fmt.Println(p2)  // (active, true)

//     // float64 and int
//     p3 := Pair[float64, int]{First: 3.14, Second: 42}
//     fmt.Println(p3)  // (3.14, 42)
// }

// package main
// import "fmt"

// type Node[T any] struct {
//     value T
//     next  *Node[T]
// }

// type LinkedList[T any] struct {
//     head *Node[T]
// }

// func (ll *LinkedList[T]) Add(value T) {
//     newNode := &Node[T]{value: value}
//     if ll.head == nil {
//         ll.head = newNode
//         return
//     }

//     current := ll.head
//     for current.next != nil {
//         current = current.next
//     }
//     current.next = newNode
// }

// func (ll LinkedList[T]) Print() {
//     current := ll.head
//     for current != nil {
//         fmt.Print(current.value, " -> ")
//         current = current.next
//     }
//     fmt.Println("nil")
// }

// func main() {
//     // Int linked list
//     intList := LinkedList[int]{}
//     intList.Add(1)
//     intList.Add(2)
//     intList.Add(3)
//     intList.Print()  // 1 -> 2 -> 3 -> nil

//     // String linked list
//     strList := LinkedList[string]{}
//     strList.Add("Hello")
//     strList.Add("World")
//     strList.Print()  // Hello -> World -> nil
// }

// // any - accepts any type
// func print[T any](value T) {
//     fmt.Println(value)
// }

// // comparable - types that can be compared with == and !=
// func contains[T comparable](slice []T, value T) bool {
//     for _, v := range slice {
//         if v == value {
//             return true
//         }
//     }
//     return false
// }

// // Only numeric types
// type Number interface {
//     int | int8 | int16 | int32 | int64 |
//     float32 | float64
// }

// func sum[T Number](numbers []T) T {
//     var total T
//     for _, n := range numbers {
//         total += n
//     }
//     return total
// }

// func main() {
//     fmt.Println(sum([]int{1, 2, 3}))         // 6
//     fmt.Println(sum([]float64{1.1, 2.2}))    // 3.3
//     // sum([]string{"a", "b"})  // ❌ ERROR! string is not Number
// }