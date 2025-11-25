// package main

// import "fmt"

// type Rectangle struct {
// 	length float64
// 	width  float64
// }

// func (r Rectangle) Area() float64 {

// 	return r.length * r.width
// }

// func main() {
// 	rect := Rectangle{length: 20, width: 30}
// 	area := rect.Area()
// 	fmt.Println("Area of rectangle with width 9 and length 10 is", area)

// }

// type BankAccount struct {
// 	balance float64
// }

// func (b BankAccount) checkBalance() float64 {
// 	return b.balance
// }

// func (b *BankAccount) Deposit(amount float64) {
// 	b.balance += amount
// }

// func (b *BankAccount) Withdraw(amount float64) bool {
// 	if b.balance >= amount {
// 		b.balance -= amount
// 		return true
// 	}
// 	return false
// }

// account := BankAccount{balance : 100}
// account.checkBalance()
// account.Deposit(50)
// account.Withdraw(30)

// type Calculator struct{
// 	brand string
// }

// func (c Calculator) Add(a,b int) int{
// 	return a + b
// }

// func (c Calculator) Multiply(a,b int) int{
// 	return a * b
// }

// func (c Calculator) Power(base, exponent int) int{
// 	result := 1
// 	for i := 0; i < exponent; i++{
// 		result *= base
// 	}
// 	return result
// }

// calc := Calculator{brand: "Casio"}
// sum := calc.Add(5,3)
// product := calc.Multiply(4,7)
// power := calc.Power(2,3)

// type Circle struct {
//     radius float64
// }

// func (c Circle) AreaAndCircumference() (float64, float64) {
//     area := 3.14159 * c.radius * c.radius
//     circumference := 2 * 3.14159 * c.radius
//     return area, circumference
// }

// // Named return values
// func (c Circle) Dimensions() (area float64, perimeter float64) {
//     area = 3.14159 * c.radius * c.radius
//     perimeter = 2 * 3.14159 * c.radius
//     return  // Returns named values automatically
// }

// // Usage
// circle := Circle{radius: 5}
// area, circ := circle.AreaAndCircumference()
// fmt.Println(area, circ)  // 78.53975 31.4159

// type Temperature float64

// func (t Temperature) Celsius() float64 {
//     return float64(t)
// }

// func (t Temperature) Fahrenheit() float64 {
//     return float64(t)*9/5 + 32
// }

// func (t Temperature) Kelvin() float64 {
//     return float64(t) + 273.15
// }

// // Usage
// temp := Temperature(25)
// fmt.Println(temp.Celsius())     // 25
// fmt.Println(temp.Fahrenheit())  // 77
// fmt.Println(temp.Kelvin())      // 298.15

// type Person struct {
//     firstName string
//     lastName  string
//     age       int
// }

// // String() method is special - used by fmt package
// func (p Person) String() string {
//     return fmt.Sprintf("%s %s (age %d)", p.firstName, p.lastName, p.age)
// }

// func (p Person) FullName() string {
//     return p.firstName + " " + p.lastName
// }

// func (p *Person) HaveBirthday() {
//     p.age++
// }

// // Usage
// person := Person{firstName: "John", lastName: "Doe", age: 30}
// fmt.Println(person)           // Automatically calls String(): "John Doe (age 30)"
// fmt.Println(person.FullName()) // "John Doe"
// person.HaveBirthday()
// fmt.Println(person.age)       // 31

// type Email struct {
//     address string
// }

// func (e Email) IsValid() bool {
//     return len(e.address) > 0 &&
//            strings.Contains(e.address, "@") &&
//            strings.Contains(e.address, ".")
// }

// func (e Email) Domain() string {
//     parts := strings.Split(e.address, "@")
//     if len(parts) == 2 {
//         return parts[1]
//     }
//     return ""
// }

// // Usage
// email := Email{address: "user@example.com"}
// if email.IsValid() {
//     fmt.Println("Valid email, domain:", email.Domain())
// }

// type StringBuilder struct{
// 	text string
// }

// func (sb *StringBuilder) Append(s string) *StringBuilder{
// 	sb.text += s
// }

// func (sb *StringBuilder) Reset() *StringBuilder{
// 	sb.text = ""
// 	return sb
// }

// func (sb StringBuilder) string() string{
// 	return sb.text
// }

// sb := &StringBuilder{}
// result := sb.Append("Hello").Append("World").String()
// fmt.Println(result)