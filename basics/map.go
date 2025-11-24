// package main

// import (
// 	"fmt"
// 	"maps"
// )

// func main() {

// 	// mapVariable = map[KeyType]valueType{
// 	// 	key1: value1,
// 	// 	key2: value2
// 	// }
// 	myMap := make(map[string]int)

// 	fmt.Println(myMap)
// 	myMap["key1"] = 9
// 	myMap["code"] = 18

// 	fmt.Println(myMap)
// 	fmt.Println(myMap["key1"])
// 	fmt.Println(myMap["key"])

// 	myMap["code"] = 21
// 	fmt.Println(myMap)

// 	// delete(myMap, "key1")

// 	_, unknown := myMap["key1"]
// 	fmt.Println(unknown)

// 	myMap2 := map[string]int{"a": 1, "b": 2}

// 	fmt.Println(myMap2)

// 	if maps.Equal(myMap, myMap2) {
// 		fmt.Println("myMap and myMap2 are equal")
// 	} else {
// 		fmt.Println("No it is not")
// 	}

// 	for _, v := range myMap2 {
// 		fmt.Println(v)
// 	}

// 	var myMap4 map[string]string

// 	if myMap4 == nil {
// 		fmt.Println("The map is initialized to nil value.")

// 	} else {
// 		fmt.Println("The map is not initialized to nil value.")

// 	}

// 	val := myMap4["key"]
// 	fmt.Println(val)

// 	myMap = make(map[string]int)
// 	// myarray := [4]int{1, 2, 3, 4}

// 	// var array [5]int

// }
