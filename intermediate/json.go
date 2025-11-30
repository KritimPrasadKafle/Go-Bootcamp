// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	ID       int    `json:"id"`                 // Always included as "id"
// 	Username string `json:"username"`           // Always included
// 	Email    string `json:"email,omitempty"`    // Omit if empty string
// 	Age      int    `json:"age,omitempty"`      // Omit if 0
// 	Password string `json:"-"`                  // NEVER include (sensitive!)
// 	IsAdmin  bool   `json:"is_admin,omitempty"` // Omit if false
// 	Score    int    `json:"score,string"`       // Encode as "123" not 123
// }

// func main() {
// 	// Full user
// 	user1 := User{
// 		ID:       1,
// 		Username: "john_doe",
// 		Email:    "john@example.com",
// 		Age:      30,
// 		Password: "secret123", // Won't appear in JSON!
// 		IsAdmin:  true,
// 		Score:    100,
// 	}

// 	// Minimal user (many empty fields)
// 	user2 := User{
// 		ID:       2,
// 		Username: "jane_doe",
// 		Password: "password456",
// 	}

// 	json1, _ := json.MarshalIndent(user1, "", "  ")
// 	json2, _ := json.MarshalIndent(user2, "", "  ")

// 	fmt.Println("Full user:")
// 	fmt.Println(string(json1))

// 	fmt.Println("\nMinimal user:")
// 	fmt.Println(string(json2))
// }

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type Person struct {
//     FirstName string `json:"name"`
//     Age       int    `json:"age,omitempty"`
// }

// func main() {
//     person := Person{FirstName: "John"}

//     jsonData, err := json.Marshal(person)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println(string(jsonData))
// }

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type Address struct {
//     City  string `json:"city"`
//     State string `json:"state"`
// }

// type Person struct {
//     Name    string  `json:"name"`
//     Age     int     `json:"age"`
//     Address Address `json:"address"`
// }

// func main() {
//     person := Person{
//         Name: "Jane",
//         Age:  30,
//         Address: Address{
//             City:  "New York",
//             State: "NY",
//         },
//     }

//     // Regular Marshal — compact, single line
//     compact, _ := json.Marshal(person)
//     fmt.Println("Compact:")
//     fmt.Println(string(compact))

//     // MarshalIndent — pretty printed
//     pretty, _ := json.MarshalIndent(person, "", "  ")
//     fmt.Println("\nPretty:")
//     fmt.Println(string(pretty))
// }
// ```

// **Output:**
// ```
// Compact:
// {"name":"Jane","age":30,"address":{"city":"New York","state":"NY"}}

// Pretty:
// {
//   "name": "Jane",
//   "age": 30,
//   "address": {
//     "city": "New York",
//     "state": "NY"
//   }
// }

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type Address struct {
//     City  string `json:"city"`
//     State string `json:"state"`
// }

// func main() {
//     addresses := []Address{
//         {City: "New York", State: "NY"},
//         {City: "San Jose", State: "CA"},
//         {City: "Las Vegas", State: "NV"},
//     }

//     jsonData, _ := json.MarshalIndent(addresses, "", "  ")
//     fmt.Println(string(jsonData))
// }

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type Address struct {
//     City  string `json:"city"`
//     State string `json:"state"`
// }

// type Employee struct {
//     FullName string  `json:"full_name"`
//     EmpID    string  `json:"emp_id"`
//     Age      int     `json:"age"`
//     Address  Address `json:"address"`
// }

// func main() {
//     // JSON string (could come from API, file, etc.)
//     jsonString := `{
//         "full_name": "Jenny Doe",
//         "emp_id": "0009",
//         "age": 30,
//         "address": {
//             "city": "San Jose",
//             "state": "CA"
//         }
//     }`

//     // Create empty struct to hold the data
//     var employee Employee

//     // Unmarshal JSON into struct
//     err := json.Unmarshal([]byte(jsonString), &employee)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     // Now use the struct like normal Go code!
//     fmt.Printf("Name: %s\n", employee.FullName)
//     fmt.Printf("ID: %s\n", employee.EmpID)
//     fmt.Printf("Age: %d\n", employee.Age)
//     fmt.Printf("City: %s\n", employee.Address.City)
//     fmt.Printf("Age in 5 years: %d\n", employee.Age+5)
// }
// ```

// **Output:**
// ```
// Name: Jenny Doe
// ID: 0009
// Age: 30
// City: San Jose
// Age in 5 years: 35

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// type Product struct {
//     Name  string  `json:"name"`
//     Price float64 `json:"price"`
// }

// func main() {
//     jsonString := `[
//         {"name": "Laptop", "price": 999.99},
//         {"name": "Mouse", "price": 29.99},
//         {"name": "Keyboard", "price": 79.99}
//     ]`

//     var products []Product

//     err := json.Unmarshal([]byte(jsonString), &products)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     for i, p := range products {
//         fmt.Printf("%d. %s - $%.2f\n", i+1, p.Name, p.Price)
//     }
// }
// ```

// **Output:**
// ```
// 1. Laptop - $999.99
// 2. Mouse - $29.99
// 3. Keyboard - $79.99

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// func main() {
//     // Unknown JSON structure
//     jsonString := `{
//         "name": "John",
//         "age": 30,
//         "active": true,
//         "scores": [85, 90, 78],
//         "address": {
//             "city": "New York",
//             "state": "NY"
//         }
//     }`

//     // Use map with interface{} for any value type
//     var data map[string]interface{}

//     err := json.Unmarshal([]byte(jsonString), &data)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     // Access values (need type assertions)
//     fmt.Println("Name:", data["name"])
//     fmt.Println("Age:", data["age"])
//     fmt.Println("Active:", data["active"])
//     fmt.Println("Scores:", data["scores"])
//     fmt.Println("Address:", data["address"])
// }
// ```

// **Output:**
// ```
// Name: John
// Age: 30
// Active: true
// Scores: [85 90 78]
// Address: map[city:New York state:NY]

// package main

// import (
//     "encoding/json"
//     "fmt"
// )

// // API response structure
// type APIResponse struct {
//     Success bool        `json:"success"`
//     Data    UserData    `json:"data,omitempty"`
//     Error   string      `json:"error,omitempty"`
// }

// type UserData struct {
//     ID       int    `json:"id"`
//     Username string `json:"username"`
//     Email    string `json:"email"`
// }

// func main() {
//     // Successful response
//     successJSON := `{
//         "success": true,
//         "data": {
//             "id": 123,
//             "username": "john_doe",
//             "email": "john@example.com"
//         }
//     }`

//     // Error response
//     errorJSON := `{
//         "success": false,
//         "error": "User not found"
//     }`

//     // Handle success
//     var successResp APIResponse
//     json.Unmarshal([]byte(successJSON), &successResp)

//     if successResp.Success {
//         fmt.Printf("Found user: %s (%s)\n",
//             successResp.Data.Username,
//             successResp.Data.Email)
//     }

//     // Handle error
//     var errorResp APIResponse
//     json.Unmarshal([]byte(errorJSON), &errorResp)

//     if !errorResp.Success {
//         fmt.Printf("API Error: %s\n", errorResp.Error)
//     }
// }
// ```

// **Output:**
// ```
// Found user: john_doe (john@example.com)
// API Error: User not found