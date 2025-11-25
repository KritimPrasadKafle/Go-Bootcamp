// package main

// import "fmt"

// type Person struct {
// 	firstName string
// 	lastName  string
// 	age       int
// }

// func main() {

// 	p := Person{
// 		firstName: "Kritim",
// 		lastName:  "Doe",
// 		age:       30,
// 	}

// 	p1 := Person{
// 		firstName: "Jane",
// 		age:       10,
// 	}
// 	fmt.Println(p.lastName)
// 	fmt.Println(p1.firstName)

// 	user := struct {
// 		userName string
// 		email    string
// 	}{
// 		userName: "user123",
// 		email:    "hey@gmail.com",
// 	}

// }

// type Person struct{
// 	Name string
// 	Age int
// 	Email string
// }

// p1 := Person{
// 	Name: "Alice",
// 	Age: 30,
// 	Email: "alice@example.com"
// }

// p2 := Person{"Bob", 25, "bob@example.com"}

// var p3 Person

// p4 := &Person{Name: "Charlie", Age: 35}
// fmt.Println(p1.Name)

// p1.Age = 31

// func (p Person) Greet() string{
// 	return "Hello, " + p.Name
// }

// func (p *Person) HaveBirthday(){
// 	p.Age++
// }

// p1.Greet()
// p1.HaveBirthday()

// type Employee struct{
// 	Person
// 	EmployeeId int
// }

// e := Employee{
// 	Person: Person{Name: "Dave", Age: 28},
// 	EmployeeId: 12345,
// }

// e.Name

// type User struct{
// 	ID int `json:"id"`
// 	UserName string `json:"username"`
// 	Password string `json:"-"`
// }

