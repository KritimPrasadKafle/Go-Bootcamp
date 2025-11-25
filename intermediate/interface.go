// package main

// import "fmt"

// // Interface - defines what all payment methods must do
// type PaymentMethod interface {
// 	ProcessPayment(amount float64) bool
// 	GetTransactionFee() float64
// }

// // CreditCard implements PaymentMethod
// type CreditCard struct {
// 	cardNumber string
// 	cvv        string
// }

// func (c CreditCard) ProcessPayment(amount float64) bool {
// 	fmt.Printf("Processing $%.2f with Credit Card\n", amount)
// 	return true
// }

// func (c CreditCard) GetTransactionFee() float64 {
// 	return 2.5 // 2.5% fee
// }

// // PayPal implements PaymentMethod
// type PayPal struct {
// 	email string
// }

// func (p PayPal) ProcessPayment(amount float64) bool {
// 	fmt.Printf("Processing $%.2f with PayPal\n", amount)
// 	return true
// }

// func (p PayPal) GetTransactionFee() float64 {
// 	return 3.0 // 3% fee
// }

// // Bitcoin implements PaymentMethod
// type Bitcoin struct {
// 	walletAddress string
// }

// func (b Bitcoin) ProcessPayment(amount float64) bool {
// 	fmt.Printf("Processing $%.2f with Bitcoin\n", amount)
// 	return true
// }

// func (b Bitcoin) GetTransactionFee() float64 {
// 	return 1.0 // 1% fee
// }

// // Cash does NOT implement PaymentMethod (missing GetTransactionFee)
// type Cash struct {
// 	amount float64
// }

// func (c Cash) ProcessPayment(amount float64) bool {
// 	fmt.Printf("Processing $%.2f with Cash\n", amount)
// 	return true
// }

// // Missing GetTransactionFee() ❌

// // Function that works with ANY payment method
// func checkout(total float64, payment PaymentMethod) {
// 	fee := payment.GetTransactionFee()
// 	totalWithFee := total + (total * fee / 100)

// 	fmt.Printf("Total: $%.2f\n", total)
// 	fmt.Printf("Fee: %.1f%%\n", fee)
// 	fmt.Printf("Total with fee: $%.2f\n", totalWithFee)

// 	success := payment.ProcessPayment(totalWithFee)
// 	if success {
// 		fmt.Println("Payment successful!\n")
// 	}
// }

// func main() {
// 	card := CreditCard{cardNumber: "1234-5678", cvv: "123"}
// 	paypal := PayPal{email: "user@example.com"}
// 	bitcoin := Bitcoin{walletAddress: "1A2B3C"}
// 	// cash := Cash{amount: 100}

// 	checkout(100, card)   // ✅ Works
// 	checkout(50, paypal)  // ✅ Works
// 	checkout(75, bitcoin) // ✅ Works
// 	// checkout(25, cash)  // ❌ ERROR! Cash doesn't implement PaymentMethod
// }

// package main
// import "fmt"

// // Notifier interface - all notification methods must implement
// type Notifier interface {
// 	Send(message string) bool
// 	GetRecipient() string
// }

// // EmailNotifier implements Notifier
// type EmailNotifier struct {
// 	email string
// }

// func (e EmailNotifier) Send(message string) bool {
// 	fmt.Printf("📧 Sending email to %s: %s\n", e.email, message)
// 	return true
// }

// func (e EmailNotifier) GetRecipient() string {
// 	return e.email
// }

// // SMSNotifier implements Notifier
// type SMSNotifier struct {
// 	phoneNumber string
// }

// func (s SMSNotifier) Send(message string) bool {
// 	fmt.Printf("📱 Sending SMS to %s: %s\n", s.phoneNumber, message)
// 	return true
// }

// func (s SMSNotifier) GetRecipient() string {
// 	return s.phoneNumber
// }

// // PushNotifier implements Notifier
// type PushNotifier struct {
// 	deviceToken string
// }

// func (p PushNotifier) Send(message string) bool {
// 	fmt.Printf("🔔 Sending push to %s: %s\n", p.deviceToken, message)
// 	return true
// }

// func (p PushNotifier) GetRecipient() string {
// 	return p.deviceToken
// }

// // Logger does NOT implement Notifier (missing GetRecipient)
// type Logger struct {
// 	logFile string
// }

// func (l Logger) Send(message string) bool {
// 	fmt.Printf("📝 Logging to %s: %s\n", l.logFile, message)
// 	return true
// }
// // Missing GetRecipient() ❌

// // Function that sends notification using any method
// func notify(n Notifier, message string) {
// 	recipient := n.GetRecipient()
// 	fmt.Printf("Preparing to notify: %s\n", recipient)
// 	success := n.Send(message)
// 	if success {
// 		fmt.Println("✅ Notification sent!\n")
// 	}
// }

// func main() {
// 	email := EmailNotifier{email: "user@example.com"}
// 	sms := SMSNotifier{phoneNumber: "+1234567890"}
// 	push := PushNotifier{deviceToken: "abc123xyz"}
// 	// logger := Logger{logFile: "/var/log/app.log"}

// 	notify(email, "Your order has shipped!")  // ✅ Works
// 	notify(sms, "Your code is 123456")        // ✅ Works
// 	notify(push, "New message received")      // ✅ Works
// 	// notify(logger, "System started")       // ❌ ERROR! Logger doesn't implement Notifier
// }

// package main
// import "fmt"

// // Storage interface - all storage types must implement these
// type Storage interface {
// 	Save(data string) bool
// 	Load() string
// 	Delete() bool
// }

// // Database implements Storage
// type Database struct {
// 	connection string
// 	data       string
// }

// func (db *Database) Save(data string) bool {
// 	db.data = data
// 	fmt.Println("Saved to database:", data)
// 	return true
// }

// func (db Database) Load() string {
// 	fmt.Println("Loading from database")
// 	return db.data
// }

// func (db *Database) Delete() bool {
// 	db.data = ""
// 	fmt.Println("Deleted from database")
// 	return true
// }

// // FileSystem implements Storage
// type FileSystem struct {
// 	filePath string
// 	content  string
// }

// func (fs *FileSystem) Save(data string) bool {
// 	fs.content = data
// 	fmt.Println("Saved to file:", fs.filePath)
// 	return true
// }

// func (fs FileSystem) Load() string {
// 	fmt.Println("Loading from file:", fs.filePath)
// 	return fs.content
// }

// func (fs *FileSystem) Delete() bool {
// 	fs.content = ""
// 	fmt.Println("Deleted file:", fs.filePath)
// 	return true
// }

// // Cloud implements Storage
// type Cloud struct {
// 	bucket string
// 	data   string
// }

// func (c *Cloud) Save(data string) bool {
// 	c.data = data
// 	fmt.Println("Saved to cloud bucket:", c.bucket)
// 	return true
// }

// func (c Cloud) Load() string {
// 	fmt.Println("Loading from cloud bucket:", c.bucket)
// 	return c.data
// }

// func (c *Cloud) Delete() bool {
// 	c.data = ""
// 	fmt.Println("Deleted from cloud bucket:", c.bucket)
// 	return true
// }

// // Cache does NOT fully implement Storage (missing Delete)
// type Cache struct {
// 	memory map[string]string
// }

// func (c *Cache) Save(data string) bool {
// 	c.memory["key"] = data
// 	fmt.Println("Saved to cache")
// 	return true
// }

// func (c Cache) Load() string {
// 	return c.memory["key"]
// }
// // Missing Delete() ❌

// // Function that works with ANY storage
// func backupData(s Storage, data string) {
// 	fmt.Println("\n=== Backup Process ===")
// 	s.Save(data)
// 	retrieved := s.Load()
// 	fmt.Println("Retrieved:", retrieved)
// 	s.Delete()
// 	fmt.Println("=== Backup Complete ===\n")
// }

// func main() {
// 	db := &Database{connection: "localhost:5432"}
// 	fs := &FileSystem{filePath: "/data/backup.txt"}
// 	cloud := &Cloud{bucket: "my-backup-bucket"}
// 	// cache := &Cache{memory: make(map[string]string)}

// 	backupData(db, "User data")       // ✅ Works
// 	backupData(fs, "File content")    // ✅ Works
// 	backupData(cloud, "Cloud backup") // ✅ Works
// 	// backupData(cache, "Cache data") // ❌ ERROR! Cache doesn't implement Storage
// }

// package main
// import "fmt"

// // Animal interface - all animals must do these things
// type Animal interface {
// 	Speak() string
// 	Move() string
// }

// // Dog implements Animal
// type Dog struct {
// 	name string
// }

// func (d Dog) Speak() string {
// 	return "Woof! Woof!"
// }

// func (d Dog) Move() string {
// 	return "Running on four legs"
// }

// // Cat implements Animal
// type Cat struct {
// 	name string
// }

// func (c Cat) Speak() string {
// 	return "Meow!"
// }

// func (c Cat) Move() string {
// 	return "Walking gracefully"
// }

// // Bird implements Animal
// type Bird struct {
// 	name string
// }

// func (b Bird) Speak() string {
// 	return "Chirp chirp!"
// }

// func (b Bird) Move() string {
// 	return "Flying in the sky"
// }

// // Fish does NOT fully implement Animal (missing Speak)
// type Fish struct {
// 	name string
// }

// func (f Fish) Move() string {
// 	return "Swimming in water"
// }
// // Missing Speak() ❌

// // Function that works with any Animal
// func describeAnimal(a Animal) {
// 	fmt.Printf("Sound: %s\n", a.Speak())
// 	fmt.Printf("Movement: %s\n\n", a.Move())
// }

// func main() {
// 	dog := Dog{name: "Buddy"}
// 	cat := Cat{name: "Whiskers"}
// 	bird := Bird{name: "Tweety"}
// 	// fish := Fish{name: "Nemo"}

// 	describeAnimal(dog)   // ✅ Works
// 	describeAnimal(cat)   // ✅ Works
// 	describeAnimal(bird)  // ✅ Works
// 	// describeAnimal(fish) // ❌ ERROR! Fish doesn't implement Animal
// }