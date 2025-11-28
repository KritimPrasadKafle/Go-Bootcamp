// package main

// import "fmt"

// func main() {

// 	// 1. ZERO-PADDED INTEGERS
// 	num := 424
// 	fmt.Printf("%05d\n", num) // Output: 00424
// 	// %05d means: print integer with minimum 5 digits, pad with zeros

// 	// 2. RIGHT-ALIGNED STRING (default)
// 	message := "Hello"
// 	fmt.Printf("|%10s|\n", message) // Output: |     Hello|
// 	// %10s means: print string in 10-character width, right-aligned

// 	// 3. LEFT-ALIGNED STRING
// 	fmt.Printf("|%-10s|\n", message) // Output: |Hello     |
// 	// %-10s means: print string in 10-character width, left-aligned (- flag)

// 	// 4. REGULAR STRING vs RAW STRING
// 	message1 := "Hello \nWorld!" // \n is interpreted as newline
// 	message2 := `Hello \nWorld!` // \n is literal (raw string)

// 	fmt.Println(message1)
// 	// Output:
// 	// Hello
// 	// World!

// 	fmt.Println(message2)
// 	// Output: Hello \nWorld!

// }

// num := 42

// // Basic integer
// fmt.Printf("%d\n", num)        // 42

// // With sign
// fmt.Printf("%+d\n", num)       // +42
// fmt.Printf("%+d\n", -num)      // -42

// // Width (right-aligned)
// fmt.Printf("|%5d|\n", num)     // |   42|

// // Width with zero padding
// fmt.Printf("%05d\n", num)      // 00042

// // Left-aligned
// fmt.Printf("|%-5d|\n", num)    // |42   |

// // With thousands separator (Go 1.13+)
// bigNum := 1234567
// fmt.Printf("%d\n", bigNum)     // 1234567
// // Note: Go doesn't have built-in thousands separator in printf

// // Binary, Octal, Hex
// fmt.Printf("Binary: %b\n", num)      // 101010
// fmt.Printf("Octal: %o\n", num)       // 52
// fmt.Printf("Hex: %x\n", num)         // 2a (lowercase)
// fmt.Printf("Hex: %X\n", num)         // 2A (uppercase)
// fmt.Printf("Hex with 0x: %#x\n", num) // 0x2a

// // Order numbers with leading zeros
// orderID := 42
// fmt.Printf("Order #%06d\n", orderID)  // Order #000042

// // Phone numbers
// area := 555
// exchange := 123
// number := 4567
// fmt.Printf("(%03d) %03d-%04d\n", area, exchange, number)
// // Output: (555) 123-4567

// // Progress indicators
// for i := 0; i <= 100; i += 10 {
//     fmt.Printf("Progress: %3d%%\n", i)
// }
// // Output:
// // Progress:   0%
// // Progress:  10%
// // Progress:  20%
// // ...
// // Progress: 100%

// pi := 3.14159265359

// // Default precision
// fmt.Printf("%f\n", pi)         // 3.141593

// // Custom precision
// fmt.Printf("%.2f\n", pi)       // 3.14
// fmt.Printf("%.4f\n", pi)       // 3.1416

// // Width and precision
// fmt.Printf("%8.2f\n", pi)      // "    3.14" (8 total width, 2 decimals)
// fmt.Printf("%08.2f\n", pi)     // "00003.14" (zero padding)

// // Scientific notation
// fmt.Printf("%e\n", pi)         // 3.141593e+00
// fmt.Printf("%.2e\n", pi)       // 3.14e+00

// // Compact form (shorter of %e or %f)
// fmt.Printf("%g\n", pi)         // 3.14159265359

// // Always show sign
// fmt.Printf("%+.2f\n", pi)      // +3.14
// fmt.Printf("%+.2f\n", -pi)     // -3.14

// // Currency formatting
// price := 1234.56
// fmt.Printf("$%.2f\n", price)   // $1234.56
// fmt.Printf("$%8.2f\n", price)  // $ 1234.56

// // Aligning prices in a table
// prices := []float64{9.99, 123.45, 5.00}
// for _, p := range prices {
//     fmt.Printf("$%7.2f\n", p)
// }
// // Output:
// //   $  9.99
// //   $123.45
// //   $  5.00

// // Temperature
// temp := 98.6
// fmt.Printf("Temperature: %.1f°F\n", temp)  // Temperature: 98.6°F

// // Percentage
// ratio := 0.8547
// fmt.Printf("Success rate: %.2f%%\n", ratio*100)  // Success rate: 85.47%

// name := "Alice"

// // Basic string
// fmt.Printf("%s\n", name)       // Alice

// // Width (right-aligned)
// fmt.Printf("|%10s|\n", name)   // |     Alice|

// // Width (left-aligned)
// fmt.Printf("|%-10s|\n", name)  // |Alice     |

// // Max characters
// fmt.Printf("%.3s\n", name)     // Ali

// // Width + Max characters
// fmt.Printf("%10.3s\n", name)   // "       Ali"

// // Quoted string
// fmt.Printf("%q\n", name)       // "Alice"

// // Creating tables
// users := []struct {
//     Name  string
//     Age   int
//     City  string
// }{
//     {"Alice", 30, "NYC"},
//     {"Bob", 25, "LA"},
//     {"Charlie", 35, "Chicago"},
// }

// fmt.Printf("%-10s %-5s %-10s\n", "NAME", "AGE", "CITY")
// fmt.Println(strings.Repeat("-", 30))
// for _, u := range users {
//     fmt.Printf("%-10s %-5d %-10s\n", u.Name, u.Age, u.City)
// }
// // Output:
// // NAME       AGE   CITY
// // ------------------------------
// // Alice      30    NYC
// // Bob        25    LA
// // Charlie    35    Chicago

// // Menu with aligned prices
// menuItems := []struct {
//     Item  string
//     Price float64
// }{
//     {"Burger", 8.99},
//     {"Fries", 3.50},
//     {"Milkshake", 4.99},
// }

// for _, item := range menuItems {
//     fmt.Printf("%-15s $%5.2f\n", item.Item, item.Price)
// }
// // Output:
// // Burger          $ 8.99
// // Fries           $ 3.50
// // Milkshake       $ 4.99

// // REGULAR STRINGS - Escape sequences are processed
// regular := "Line 1\nLine 2\tTabbed"
// fmt.Println(regular)
// // Output:
// // Line 1
// // Line 2    Tabbed

// // RAW STRINGS - Everything is literal
// raw := `Line 1\nLine 2\tTabbed`
// fmt.Println(raw)
// // Output: Line 1\nLine 2\tTabbed

// // Practical use: SQL queries
// sqlQuery := `
//     SELECT id, name, email
//     FROM users
//     WHERE age > 30
//     AND status = 'active'
//     ORDER BY name
// `
// fmt.Println(sqlQuery)

// // Practical use: JSON templates
// jsonTemplate := `{
//     "name": "%s",
//     "age": %d,
//     "active": %t
// }`
// json := fmt.Sprintf(jsonTemplate, "Alice", 30, true)
// fmt.Println(json)

// // Practical use: Multi-line messages
// message := `
// Dear Customer,

// Thank you for your order!

// Order ID: %s
// Total: $%.2f

// Best regards,
// The Team
// `
// fmt.Printf(message, "ORD-12345", 99.99)

// // File paths (Windows)
// windowsPath := `C:\Users\Documents\file.txt`
// fmt.Println(windowsPath)  // C:\Users\Documents\file.txt

// // Regular expressions
// regex := `^\d{3}-\d{2}-\d{4}$`  // SSN pattern
// fmt.Println(regex)

// package main

// import (
// 	"fmt"
// 	"strings"
// 	"time"
// )

// type Product struct {
// 	ID    int
// 	Name  string
// 	Price float64
// 	Stock int
// }

// func main() {
// 	fmt.Println("=== RECEIPT GENERATOR ===\n")

// 	products := []Product{
// 		{1, "Laptop", 999.99, 5},
// 		{2, "Mouse", 29.99, 150},
// 		{3, "Keyboard", 79.99, 50},
// 	}

// 	// Header
// 	fmt.Printf("%-5s %-20s %10s %8s\n", "ID", "Product", "Price", "Stock")
// 	fmt.Println(strings.Repeat("-", 50))

// 	// Items
// 	total := 0.0
// 	for _, p := range products {
// 		fmt.Printf("%05d %-20s $%9.2f %8d\n", p.ID, p.Name, p.Price, p.Stock)
// 		total += p.Price
// 	}

// 	// Footer
// 	fmt.Println(strings.Repeat("-", 50))
// 	fmt.Printf("%-26s $%9.2f\n", "TOTAL:", total)

// 	// Date
// 	now := time.Now()
// 	fmt.Printf("\nDate: %s\n", now.Format("2006-01-02 15:04:05"))

// 	// Raw string example
// 	terms := `
// Terms & Conditions:
// - All sales are final
// - Returns within 30 days
// - Valid ID required
// `
// 	fmt.Println(terms)
// }
// ```

// **Output:**
// ```
// === RECEIPT GENERATOR ===

// ID    Product              Price    Stock
// --------------------------------------------------
// 00001 Laptop               $   999.99        5
// 00002 Mouse                $    29.99      150
// 00003 Keyboard             $    79.99       50
// --------------------------------------------------
// TOTAL:                     $  1109.97

// Date: 2024-01-15 14:30:45

// Terms & Conditions:
// - All sales are final
// - Returns within 30 days
// - Valid ID required