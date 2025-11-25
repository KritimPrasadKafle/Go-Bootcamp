// package main

// import "fmt"

// type Person struct {
// 	name string
// 	age  int
// }
// type Employee struct {
// 	Person
// 	EmployeeId int
// 	address    string
// }

// func (e Employee) checkAddress() string {
// 	return e.address
// }

// func main() {

// 	e := Employee{
// 		Person:     Person{name: "Kritim", age: 20},
// 		EmployeeId: 123,
// 		address:    "Hetauda",
// 	}

// 	result := e.checkAddress()
// 	fmt.Println(result)

// }

// package main
// import "fmt"

// // Base struct
// type Vehicle struct {
// 	brand string
// 	year  int
// }

// func (v Vehicle) info() {
// 	fmt.Printf("Brand: %s, Year: %d\n", v.brand, v.year)
// }

// func (v Vehicle) start() {
// 	fmt.Println("Vehicle starting...")
// }

// // Car embeds Vehicle (anonymous field)
// type Car struct {
// 	Vehicle  // Embedding
// 	doors   int
// }

// // Override start method
// func (c Car) start() {
// 	fmt.Printf("%s car with %d doors starting...\n", c.brand, c.doors)
// }

// func (c Car) honk() {
// 	fmt.Println("Beep beep!")
// }

// // Motorcycle embeds Vehicle
// type Motorcycle struct {
// 	Vehicle  // Embedding
// 	hasWindshield bool
// }

// func (m Motorcycle) wheelie() {
// 	fmt.Println("Doing a wheelie!")
// }

// func main() {
// 	// Create Car
// 	car := Car{
// 		Vehicle: Vehicle{brand: "Toyota", year: 2022},
// 		doors:   4,
// 	}

// 	// Access Vehicle fields directly
// 	fmt.Println("Car brand:", car.brand)     // Direct access
// 	fmt.Println("Car year:", car.year)       // Direct access
// 	fmt.Println("Car doors:", car.doors)

// 	// Call methods
// 	car.info()      // Vehicle's method
// 	car.start()     // Car's overridden method
// 	car.honk()      // Car's own method

// 	fmt.Println()

// 	// Create Motorcycle
// 	bike := Motorcycle{
// 		Vehicle:       Vehicle{brand: "Harley", year: 2023},
// 		hasWindshield: true,
// 	}

// 	fmt.Println("Bike brand:", bike.brand)  // Direct access
// 	bike.info()     // Vehicle's method
// 	bike.start()    // Vehicle's method (not overridden)
// 	bike.wheelie()  // Motorcycle's own method
// }
// ```

// **Output:**
// ```
// Car brand: Toyota
// Car year: 2022
// Car doors: 4
// Brand: Toyota, Year: 2022
// Toyota car with 4 doors starting...
// Beep beep!

// Bike brand: Harley
// Brand: Harley, Year: 2023
// Vehicle starting...
// Doing a wheelie!

// package main
// import "fmt"

// type User struct {
// 	username string
// 	email    string
// }

// func (u User) login() {
// 	fmt.Printf("%s logged in\n", u.username)
// }

// type Admin struct {
// 	User  // Embedding User
// 	permissions []string
// }

// func (a Admin) deleteUser(username string) {
// 	fmt.Printf("Admin %s deleted user %s\n", a.username, username)
// }

// type Moderator struct {
// 	User  // Embedding User
// 	canBan bool
// }

// func (m Moderator) banUser(username string) {
// 	if m.canBan {
// 		fmt.Printf("Moderator %s banned user %s\n", m.username, username)
// 	}
// }

// func main() {
// 	// Create admin
// 	admin := Admin{
// 		User: User{username: "admin1", email: "admin@example.com"},
// 		permissions: []string{"delete", "ban", "edit"},
// 	}

// 	// Access User fields directly
// 	fmt.Println("Admin username:", admin.username)
// 	fmt.Println("Admin email:", admin.email)
// 	fmt.Println("Permissions:", admin.permissions)

// 	// Call methods
// 	admin.login()              // User's method
// 	admin.deleteUser("spammer") // Admin's method

// 	fmt.Println()

// 	// Create moderator
// 	mod := Moderator{
// 		User:   User{username: "mod1", email: "mod@example.com"},
// 		canBan: true,
// 	}

// 	fmt.Println("Mod username:", mod.username)
// 	mod.login()           // User's method
// 	mod.banUser("troll")  // Moderator's method
// }
// ```

// **Output:**
// ```
// Admin username: admin1
// Admin email: admin@example.com
// Permissions: [delete ban edit]
// admin1 logged in
// Admin admin1 deleted user spammer

// Mod username: mod1
// mod1 logged in
// Moderator mod1 banned user troll

// package main
// import "fmt"

// type Address struct {
// 	street  string
// 	city    string
// 	zipCode string
// }

// func (a Address) fullAddress() string {
// 	return fmt.Sprintf("%s, %s, %s", a.street, a.city, a.zipCode)
// }

// type Person struct {
// 	name    string
// 	age     int
// 	Address  // Embedding Address
// }

// type Company struct {
// 	name     string
// 	Address  // Embedding Address
// 	employees int
// }

// func main() {
// 	// Person with embedded Address
// 	person := Person{
// 		name: "Alice",
// 		age:  30,
// 		Address: Address{
// 			street:  "123 Main St",
// 			city:    "New York",
// 			zipCode: "10001",
// 		},
// 	}

// 	// Access Address fields directly
// 	fmt.Println("Person:", person.name)
// 	fmt.Println("City:", person.city)        // Direct access
// 	fmt.Println("Street:", person.street)    // Direct access
// 	fmt.Println("Full:", person.fullAddress()) // Address method

// 	fmt.Println()

// 	// Company with embedded Address
// 	company := Company{
// 		name: "TechCorp",
// 		Address: Address{
// 			street:  "456 Tech Ave",
// 			city:    "San Francisco",
// 			zipCode: "94102",
// 		},
// 		employees: 100,
// 	}

// 	fmt.Println("Company:", company.name)
// 	fmt.Println("City:", company.city)        // Direct access
// 	fmt.Println("Full:", company.fullAddress()) // Address method
// 	fmt.Println("Employees:", company.employees)
// }
// ```

// **Output:**
// ```
// Person: Alice
// City: New York
// Street: 123 Main St
// Full: 123 Main St, New York, 10001

// Company: TechCorp
// City: San Francisco
// Full: 456 Tech Ave, San Francisco, 94102
// Employees: 100

// package main
// import "fmt"

// type Base struct {
// 	name string
// 	id   int
// }

// func (b Base) display() {
// 	fmt.Printf("Base: name=%s, id=%d\n", b.name, b.id)
// }

// type Derived struct {
// 	Base  // Embedding
// 	name  string  // ← Same field name as Base!
// 	code  string
// }

// func (d Derived) display() {
// 	fmt.Printf("Derived: name=%s, code=%s\n", d.name, d.code)
// }

// func main() {
// 	d := Derived{
// 		Base: Base{name: "BaseName", id: 1},
// 		name: "DerivedName",
// 		code: "XYZ",
// 	}

// 	// Which "name" is accessed?
// 	fmt.Println("d.name:", d.name)           // DerivedName (outer takes precedence)
// 	fmt.Println("d.Base.name:", d.Base.name) // BaseName (explicit access)

// 	// Which display() is called?
// 	d.display()      // Derived's display (outer takes precedence)
// 	d.Base.display() // Base's display (explicit access)
// }
// ```

// **Output:**
// ```
// d.name: DerivedName
// d.Base.name: BaseName
// Derived: name=DerivedName, code=XYZ
// Base: name=BaseName, id=1

// package main
// import "fmt"

// type Logger struct {
// 	logLevel string
// }

// func (l Logger) log(message string) {
// 	fmt.Printf("[%s] %s\n", l.logLevel, message)
// }

// type Database struct {
// 	connectionString string
// }

// func (db Database) connect() {
// 	fmt.Println("Connecting to database:", db.connectionString)
// }

// // Service embeds BOTH Logger and Database
// type Service struct {
// 	Logger    // First embedding
// 	Database  // Second embedding
// 	name string
// }

// func (s Service) start() {
// 	s.log("Service starting...")  // Logger's method
// 	s.connect()                    // Database's method
// 	fmt.Println("Service", s.name, "is running")
// }

// func main() {
// 	service := Service{
// 		Logger:   Logger{logLevel: "INFO"},
// 		Database: Database{connectionString: "localhost:5432"},
// 		name:     "UserService",
// 	}

// 	// Access fields from both embedded structs
// 	fmt.Println("Log level:", service.logLevel)  // From Logger
// 	fmt.Println("DB:", service.connectionString) // From Database

// 	// Call methods from both embedded structs
// 	service.log("Hello")  // Logger's method
// 	service.connect()     // Database's method
// 	service.start()       // Service's own method
// }
// ```

// **Output:**
// ```
// Log level: INFO
// DB: localhost:5432
// [INFO] Hello
// Connecting to database: localhost:5432
// [INFO] Service starting...
// Connecting to database: localhost:5432
// Service UserService is running