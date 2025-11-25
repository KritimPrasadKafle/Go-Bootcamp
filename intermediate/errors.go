// package main

// import (
// 	"errors"
// 	"fmt"
// )

// func sqrt(x float64) (float64, error) {
// 	if x > 0 {
// 		return 0, errors.New("Math Error: square root of negative number")
// 	}
// 	return 1, nil
// }

// func process(data []byte) error {
// 	if len(data) == 0 {
// 		return errors.New("Empty data")
// 	}
// 	// Process data
// 	return nil // Success - no error
// }

// type myError struct {
// 	message string
// }

// func (m *myError) Error() string {
// 	return fmt.Sprintf("Error: %s", m.message)
// }

// func main() {

// }

// package main
// import (
//     "errors"
//     "fmt"
// )

// func openFile(filename string) error {
//     if filename == "" {
//         return errors.New("filename cannot be empty")
//     }

//     // Simulate file not found
//     if filename != "config.txt" {
//         return fmt.Errorf("file not found: %s", filename)
//     }

//     return nil
// }

// func main() {
//     files := []string{"", "data.txt", "config.txt"}

//     for _, file := range files {
//         err := openFile(file)
//         if err != nil {
//             fmt.Printf("Error opening '%s': %v\n", file, err)
//             continue
//         }
//         fmt.Printf("Successfully opened '%s'\n", file)
//     }
// }
// ```

// **Output:**
// ```
// Error opening '': filename cannot be empty
// Error opening 'data.txt': file not found: data.txt
// Successfully opened 'config.txt'

// package main
// import "fmt"

// type ValidationError struct {
//     Field   string
//     Message string
//     Value   interface{}
// }

// func (v *ValidationError) Error() string {
//     return fmt.Sprintf("validation error: field '%s' %s (got: %v)",
//         v.Field, v.Message, v.Value)
// }

// func validateAge(age int) error {
//     if age < 0 {
//         return &ValidationError{
//             Field:   "age",
//             Message: "cannot be negative",
//             Value:   age,
//         }
//     }
//     if age > 150 {
//         return &ValidationError{
//             Field:   "age",
//             Message: "is unrealistic",
//             Value:   age,
//         }
//     }
//     return nil
// }

// func validateEmail(email string) error {
//     if email == "" {
//         return &ValidationError{
//             Field:   "email",
//             Message: "is required",
//             Value:   email,
//         }
//     }
//     return nil
// }

// func main() {
//     ages := []int{-5, 25, 200}

//     for _, age := range ages {
//         err := validateAge(age)
//         if err != nil {
//             fmt.Println(err)
//         } else {
//             fmt.Printf("Age %d is valid\n", age)
//         }
//     }

//     fmt.Println()

//     err := validateEmail("")
//     if err != nil {
//         fmt.Println(err)
//     }
// }
// ```

// **Output:**
// ```
// validation error: field 'age' cannot be negative (got: -5)
// Age 25 is valid
// validation error: field 'age' is unrealistic (got: 200)

// validation error: field 'email' is required (got: )

// package main
// import (
//     "errors"
//     "fmt"
// )

// func connectDatabase() error {
//     return errors.New("connection timeout")
// }

// func initializeDatabase() error {
//     err := connectDatabase()
//     if err != nil {
//         return fmt.Errorf("initializeDatabase: %w", err)
//     }
//     return nil
// }

// func startServer() error {
//     err := initializeDatabase()
//     if err != nil {
//         return fmt.Errorf("startServer: %w", err)
//     }
//     return nil
// }

// func main() {
//     err := startServer()
//     if err != nil {
//         fmt.Println("Application error:", err)

//         // Unwrap to get original error
//         originalErr := errors.Unwrap(err)
//         fmt.Println("Unwrapped once:", originalErr)

//         originalErr = errors.Unwrap(originalErr)
//         fmt.Println("Unwrapped twice:", originalErr)
//     }
// }
// ```

// **Output:**
// ```
// Application error: startServer: initializeDatabase: connection timeout
// Unwrapped once: initializeDatabase: connection timeout
// Unwrapped twice: connection timeout

// package main
// import (
//     "errors"
//     "fmt"
// )

// var ErrNotFound = errors.New("not found")
// var ErrPermissionDenied = errors.New("permission denied")

// type CustomError struct {
//     Code    int
//     Message string
// }

// func (c *CustomError) Error() string {
//     return fmt.Sprintf("error %d: %s", c.Code, c.Message)
// }

// func getUser(id int) error {
//     if id < 0 {
//         return fmt.Errorf("invalid id: %w", ErrNotFound)
//     }
//     if id == 999 {
//         return &CustomError{Code: 403, Message: "admin access required"}
//     }
//     return nil
// }

// func main() {
//     // Test errors.Is()
//     err := getUser(-1)
//     if errors.Is(err, ErrNotFound) {
//         fmt.Println("User not found!")
//     }

//     // Test errors.As()
//     err = getUser(999)
//     var customErr *CustomError
//     if errors.As(err, &customErr) {
//         fmt.Printf("Custom error with code: %d\n", customErr.Code)
//     }
// }
// ```

// **Output:**
// ```
// User not found!
// Custom error with code: 403

// package main
// import (
//     "errors"
//     "fmt"
// )

// func validateUser(name string, age int, email string) error {
//     var errs []error

//     if name == "" {
//         errs = append(errs, errors.New("name is required"))
//     }

//     if age < 18 {
//         errs = append(errs, errors.New("must be 18 or older"))
//     }

//     if email == "" {
//         errs = append(errs, errors.New("email is required"))
//     }

//     if len(errs) > 0 {
//         return errors.Join(errs...)  // Go 1.20+ - combine multiple errors
//     }

//     return nil
// }

// func main() {
//     err := validateUser("", 15, "")
//     if err != nil {
//         fmt.Println("Validation errors:")
//         fmt.Println(err)
//     }
// }
// ```

// **Output:**
// ```
// Validation errors:
// name is required
// must be 18 or older
// email is required