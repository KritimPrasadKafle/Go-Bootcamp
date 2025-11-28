// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	fmt.Println("=== BASIC RANDOM FUNCTIONS ===")
// 	fmt.Println()

// 	// Random integer from 0 to n-1
// 	fmt.Println("rand.Intn(10):", rand.Intn(10))     // 0-9
// 	fmt.Println("rand.Intn(100):", rand.Intn(100))   // 0-99
// 	fmt.Println("rand.Intn(1000):", rand.Intn(1000)) // 0-999

// 	// Random integer (can be negative, full int range)
// 	fmt.Println("\nrand.Int():", rand.Int())

// 	// Random int31 (0 to 2^31-1)
// 	fmt.Println("rand.Int31():", rand.Int31())

// 	// Random int63 (0 to 2^63-1)
// 	fmt.Println("rand.Int63():", rand.Int63())

// 	// Random uint32 (0 to 2^32-1)
// 	fmt.Println("rand.Uint32():", rand.Uint32())

// 	// Random uint64 (0 to 2^64-1)
// 	fmt.Println("rand.Uint64():", rand.Uint64())

// 	// Random float64 (0.0 to 1.0)
// 	fmt.Println("\nrand.Float64():", rand.Float64())

// 	// Random float32 (0.0 to 1.0)
// 	fmt.Println("rand.Float32():", rand.Float32())
// }

// func main() {
// 	fmt.Println("=== RANDOM IN RANGE ===")
// 	fmt.Println()

// 	// Integer range: min to max (inclusive)
// 	min, max := 10, 20
// 	randomInRange := rand.Intn(max-min+1) + min
// 	fmt.Printf("Random between %d and %d: %d\n", min, max, randomInRange)

// 	// Generate 10 random numbers between 1 and 100
// 	fmt.Println("\n10 random numbers (1-100):")
// 	for i := 0; i < 10; i++ {
// 		fmt.Printf("%d ", rand.Intn(100)+1)
// 	}
// 	fmt.Println()

// 	// Float range: min to max
// 	minF, maxF := 5.0, 15.0
// 	randomFloat := minF + rand.Float64()*(maxF-minF)
// 	fmt.Printf("\nRandom float between %.1f and %.1f: %.4f\n", minF, maxF, randomFloat)

// 	// Generate 5 random floats between 0 and 100
// 	fmt.Println("\n5 random floats (0-100):")
// 	for i := 0; i < 5; i++ {
// 		fmt.Printf("%.2f ", rand.Float64()*100)
// 	}
// 	fmt.Println()
// }

// // Helper function for random range
// func RandomRange(min, max int) int {
// 	return rand.Intn(max-min+1) + min
// }

// func RandomFloatRange(min, max float64) float64 {
// 	return min + rand.Float64()*(max-min)
// }
