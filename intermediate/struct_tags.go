// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// )

// type Person struct {
// 	XMLName xml.Name `xml:"person"`          // Root element name
// 	ID      int      `xml:"id,attr"`         // Attribute
// 	Name    string   `xml:"name"`            // Element
// 	City    string   `xml:"address>city"`    // Nested
// 	Country string   `xml:"address>country"` // Nested
// }

// func main() {
// 	person := Person{
// 		ID:      1,
// 		Name:    "John Doe",
// 		City:    "New York",
// 		Country: "USA",
// 	}

// 	xmlData, _ := xml.MarshalIndent(person, "", "  ")
// 	fmt.Println(string(xmlData))
// }

// package main

// import (
//     "fmt"
//     "github.com/go-playground/validator/v10"
// )

// type User struct {
//     Name     string `validate:"required,min=2,max=50"`
//     Email    string `validate:"required,email"`
//     Age      int    `validate:"gte=18,lte=120"`
//     Password string `validate:"required,min=8"`
// }

// func main() {
//     validate := validator.New()

//     // Invalid user
//     user := User{
//         Name:     "J",                    // Too short (min=2)
//         Email:    "not-an-email",         // Invalid email
//         Age:      15,                     // Too young (gte=18)
//         Password: "123",                  // Too short (min=8)
//     }

//     err := validate.Struct(user)
//     if err != nil {
//         for _, e := range err.(validator.ValidationErrors) {
//             fmt.Printf("Field '%s' failed on '%s' tag\n", e.Field(), e.Tag())
//         }
//     }
// }
// ```

// **Output:**
// ```
// Field 'Name' failed on 'min' tag
// Field 'Email' failed on 'email' tag
// Field 'Age' failed on 'gte' tag
// Field 'Password' failed on 'min' tag

// package main

// import (
//     "fmt"
//     "reflect"
// )

// type Person struct {
//     FirstName string `json:"first_name" db:"firstn" xml:"first"`
//     LastName  string `json:"last_name,omitempty"`
//     Age       int    `json:"-"`
// }

// func main() {
//     t := reflect.TypeOf(Person{})

//     for i := 0; i < t.NumField(); i++ {
//         field := t.Field(i)

//         fmt.Printf("Field: %s\n", field.Name)
//         fmt.Printf("  Full Tag: %s\n", field.Tag)
//         fmt.Printf("  json: %s\n", field.Tag.Get("json"))
//         fmt.Printf("  db:   %s\n", field.Tag.Get("db"))
//         fmt.Printf("  xml:  %s\n", field.Tag.Get("xml"))
//         fmt.Println()
//     }
// }
// ```

// **Output:**
// ```
// Field: FirstName
//   Full Tag: json:"first_name" db:"firstn" xml:"first"
//   json: first_name
//   db:   firstn
//   xml:  first

// Field: LastName
//   Full Tag: json:"last_name,omitempty"
//   json: last_name,omitempty
//   db:
//   xml:

// Field: Age
//   Full Tag: json:"-"
//   json: -
//   db:
//   xml:

//   package main

// import (
//     "encoding/json"
//     "fmt"
// )

// // Request from client
// type CreateUserRequest struct {
//     Username string `json:"username" validate:"required,min=3,max=20"`
//     Email    string `json:"email" validate:"required,email"`
//     Password string `json:"password" validate:"required,min=8"`
//     Age      int    `json:"age,omitempty" validate:"omitempty,gte=13"`
// }

// // Response to client (no password!)
// type UserResponse struct {
//     ID       int    `json:"id"`
//     Username string `json:"username"`
//     Email    string `json:"email"`
//     Age      int    `json:"age,omitempty"`
// }

// // Internal/Database model
// type User struct {
//     ID           int    `json:"-" db:"id" gorm:"primaryKey"`
//     Username     string `json:"username" db:"username"`
//     Email        string `json:"email" db:"email"`
//     PasswordHash string `json:"-" db:"password_hash"` // Never in JSON!
//     Age          int    `json:"age,omitempty" db:"age"`
// }

// func main() {
//     // Incoming JSON request
//     requestJSON := `{
//         "username": "john_doe",
//         "email": "john@example.com",
//         "password": "secret123",
//         "age": 25
//     }`

//     var request CreateUserRequest
//     json.Unmarshal([]byte(requestJSON), &request)

//     fmt.Println("Received request:")
//     fmt.Printf("  Username: %s\n", request.Username)
//     fmt.Printf("  Email: %s\n", request.Email)
//     fmt.Printf("  Password: %s\n", request.Password)

//     // Create internal user (hash password, etc.)
//     user := User{
//         ID:           1,
//         Username:     request.Username,
//         Email:        request.Email,
//         PasswordHash: "hashed_" + request.Password, // In reality, use bcrypt
//         Age:          request.Age,
//     }

//     // Return response (password excluded!)
//     response := UserResponse{
//         ID:       user.ID,
//         Username: user.Username,
//         Email:    user.Email,
//         Age:      user.Age,
//     }

//     responseJSON, _ := json.MarshalIndent(response, "", "  ")
//     fmt.Println("\nResponse to client:")
//     fmt.Println(string(responseJSON))
// }
// ```

// **Output:**
// ```
// Received request:
//   Username: john_doe
//   Email: john@example.com
//   Password: secret123

// Response to client:
// {
//   "id": 1,
//   "username": "john_doe",
//   "email": "john@example.com",
//   "age": 25
// }