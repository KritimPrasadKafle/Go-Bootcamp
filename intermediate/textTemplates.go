// package main

// func main() {

// }

// // METHOD 1: Manual error handling
// tmpl := template.New("example")
// tmpl, err := template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n")
// if err != nil {
//     panic(err)
// }

// // METHOD 2: Using template.Must() - cleaner for templates that should never fail
// tmpl := template.Must(template.New("example").Parse("Welcome, {{.name}}! How are you doing?\n"))

// data := map[string]interface{}{
//     "name": "John",
// }
// err := tmpl.Execute(os.Stdout, data)  // Outputs to console

// // 1. READ USER INPUT
// reader := bufio.NewReader(os.Stdin)
// name, _ := reader.ReadString('\n')
// name = strings.TrimSpace(name)  // Remove newline

// // 2. DEFINE MULTIPLE TEMPLATES
// templates := map[string]string{
//     "welcome":      "Welcome, {{.name}}! We're glad you joined.",
//     "notification": "{{.nm}}, you have a new notification: {{.ntf}}",
//     "error":        "Oops! An error occured: {{.em}}",
// }

// // 3. PARSE ALL TEMPLATES
// parsedTemplates := make(map[string]*template.Template)
// for name, tmpl := range templates {
//     parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
// }

// // 4. MENU LOOP - Select and execute templates based on user choice

// package main

// import (
//     "os"
//     "text/template"
// )

// func main() {
//     // Simple field access
//     tmpl := template.Must(template.New("test").Parse("Hello, {{.Name}}!"))

//     data := struct {
//         Name string
//     }{
//         Name: "Alice",
//     }

//     tmpl.Execute(os.Stdout, data)  // Output: Hello, Alice!
// }

// type Address struct {
//     City    string
//     Country string
// }

// type Person struct {
//     Name    string
//     Age     int
//     Address Address
// }

// func main() {
//     tmpl := template.Must(template.New("test").Parse(`
// Name: {{.Name}}
// Age: {{.Age}}
// City: {{.Address.City}}
// Country: {{.Address.Country}}
// `))

//     person := Person{
//         Name: "Bob",
//         Age:  30,
//         Address: Address{
//             City:    "New York",
//             Country: "USA",
//         },
//     }

//     tmpl.Execute(os.Stdout, person)
// }
// // Output:
// // Name: Bob
// // Age: 30
// // City: New York
// // Country: USA

// func main() {
//     tmpl := template.Must(template.New("test").Parse(`
// {{if .IsActive}}
//     Status: ACTIVE
// {{else}}
//     Status: INACTIVE
// {{end}}

// {{if gt .Age 18}}
//     You are an adult.
// {{else}}
//     You are a minor.
// {{end}}
// `))

//     data := map[string]interface{}{
//         "IsActive": true,
//         "Age":      25,
//     }

//     tmpl.Execute(os.Stdout, data)
// }
// // Output:
// // Status: ACTIVE
// // You are an adult.

// func main() {
//     tmpl := template.Must(template.New("test").Parse(`
// Users:
// {{range .Users}}
// - Name: {{.Name}}, Age: {{.Age}}
// {{end}}

// Fruits:
// {{range .Fruits}}
// - {{.}}
// {{end}}

// Scores:
// {{range $key, $value := .Scores}}
// - {{$key}}: {{$value}}
// {{end}}
// `))

//     data := map[string]interface{}{
//         "Users": []struct {
//             Name string
//             Age  int
//         }{
//             {"Alice", 25},
//             {"Bob", 30},
//             {"Charlie", 35},
//         },
//         "Fruits": []string{"Apple", "Banana", "Cherry"},
//         "Scores": map[string]int{
//             "Math":    90,
//             "English": 85,
//             "Science": 92,
//         },
//     }

//     tmpl.Execute(os.Stdout, data)
// }
// // Output:
// // Users:
// // - Name: Alice, Age: 25
// // - Name: Bob, Age: 30
// // - Name: Charlie, Age: 35
// //
// // Fruits:
// // - Apple
// // - Banana
// // - Cherry
// //
// // Scores:
// // - English: 85
// // - Math: 90
// // - Science: 92

// func main() {
//     tmpl := template.Must(template.New("test").Parse(`
// {{$name := .Name}}
// {{$age := .Age}}

// Hello, {{$name}}!
// You are {{$age}} years old.
// {{if gt $age 18}}
//     Welcome, adult {{$name}}!
// {{end}}
// `))

//     data := map[string]interface{}{
//         "Name": "Alice",
//         "Age":  25,
//     }

//     tmpl.Execute(os.Stdout, data)
// }

// func main() {
//     tmpl := template.Must(template.New("test").Parse(`
// {{with .User}}
//     Name: {{.Name}}
//     Email: {{.Email}}
// {{end}}

// {{with .Address}}
//     City: {{.City}}
//     Zip: {{.Zip}}
// {{else}}
//     No address provided
// {{end}}
// `))

//     data := map[string]interface{}{
//         "User": map[string]string{
//             "Name":  "John",
//             "Email": "john@example.com",
//         },
//         // Address is nil
//     }

//     tmpl.Execute(os.Stdout, data)
// }
// // Output:
// // Name: John
// // Email: john@example.com
// // No address provided

// func main() {
//     // Built-in functions
//     tmpl := template.Must(template.New("test").Parse(`
// Length of name: {{len .Name}}
// Print with printf: {{printf "Hello, %s!" .Name}}
// {{if and .IsActive .IsVerified}}
//     User is active and verified
// {{end}}
// {{if or .IsAdmin .IsModerator}}
//     User has elevated privileges
// {{end}}
// `))

//     data := map[string]interface{}{
//         "Name":        "Alice",
//         "IsActive":    true,
//         "IsVerified":  true,
//         "IsAdmin":     false,
//         "IsModerator": true,
//     }

//     tmpl.Execute(os.Stdout, data)
// }

// import (
//     "strings"
//     "text/template"
// )

// func main() {
//     // Define custom functions
//     funcMap := template.FuncMap{
//         "toUpper": strings.ToUpper,
//         "toLower": strings.ToLower,
//         "repeat": func(s string, n int) string {
//             return strings.Repeat(s, n)
//         },
//         "multiply": func(a, b int) int {
//             return a * b
//         },
//     }

//     tmpl := template.Must(template.New("test").Funcs(funcMap).Parse(`
// Original: {{.Name}}
// Upper: {{toUpper .Name}}
// Lower: {{toLower .Name}}
// Repeated: {{repeat "Go! " 3}}
// Calculation: 5 x 10 = {{multiply 5 10}}
// `))

//     data := map[string]string{
//         "Name": "Alice",
//     }

//     tmpl.Execute(os.Stdout, data)
// }
// // Output:
// // Original: Alice
// // Upper: ALICE
// // Lower: alice
// // Repeated: Go! Go! Go!
// // Calculation: 5 x 10 = 50

// package main

// import (
//     "fmt"
//     "os"
//     "text/template"
//     "time"
// )

// func main() {
//     emailTemplates := map[string]string{
//         "welcome": `
// Subject: Welcome to {{.CompanyName}}!

// Dear {{.Name}},

// Welcome aboard! We're thrilled to have you join {{.CompanyName}}.

// Your account has been created with the following details:
// - Username: {{.Username}}
// - Email: {{.Email}}
// - Registration Date: {{.Date}}

// Get started by logging in at: {{.LoginURL}}

// Best regards,
// The {{.CompanyName}} Team
// `,
//         "resetPassword": `
// Subject: Password Reset Request

// Hi {{.Name}},

// We received a request to reset your password.

// Click the link below to reset your password:
// {{.ResetURL}}

// This link will expire in {{.ExpiryHours}} hours.

// If you didn't request this, please ignore this email.

// Best regards,
// Security Team
// `,
//         "orderConfirmation": `
// Subject: Order Confirmation #{{.OrderID}}

// Dear {{.Customer.Name}},

// Thank you for your order!

// Order Details:
// - Order ID: {{.OrderID}}
// - Date: {{.Date}}
// - Total: ${{.Total}}

// Items:
// {{range .Items}}
// - {{.Name}} (x{{.Quantity}}) - ${{.Price}}
// {{end}}

// Shipping Address:
// {{.Customer.Address}}

// Estimated delivery: {{.DeliveryDate}}

// Thank you for shopping with us!
// `,
//     }

//     // Parse templates
//     tmplMap := make(map[string]*template.Template)
//     for name, content := range emailTemplates {
//         tmplMap[name] = template.Must(template.New(name).Parse(content))
//     }

//     // Example 1: Welcome email
//     fmt.Println("=== WELCOME EMAIL ===")
//     welcomeData := map[string]interface{}{
//         "CompanyName": "TechCorp",
//         "Name":        "John Doe",
//         "Username":    "johndoe",
//         "Email":       "john@example.com",
//         "Date":        time.Now().Format("2006-01-02"),
//         "LoginURL":    "https://example.com/login",
//     }
//     tmplMap["welcome"].Execute(os.Stdout, welcomeData)

//     // Example 2: Order confirmation
//     fmt.Println("\n\n=== ORDER CONFIRMATION ===")
//     orderData := map[string]interface{}{
//         "OrderID": "ORD-12345",
//         "Date":    time.Now().Format("2006-01-02 15:04:05"),
//         "Total":   149.97,
//         "Customer": map[string]string{
//             "Name":    "Jane Smith",
//             "Address": "123 Main St, New York, NY 10001",
//         },
//         "Items": []map[string]interface{}{
//             {"Name": "Laptop", "Quantity": 1, "Price": 99.99},
//             {"Name": "Mouse", "Quantity": 2, "Price": 24.99},
//         },
//         "DeliveryDate": "2024-01-20",
//     }
//     tmplMap["orderConfirmation"].Execute(os.Stdout, orderData)
// }

// package main

// import (
//     "html/template"
//     "os"
// )

// func main() {
//     htmlTemplate := `
// <!DOCTYPE html>
// <html>
// <head>
//     <title>{{.Title}}</title>
//     <style>
//         body { font-family: Arial, sans-serif; margin: 20px; }
//         table { border-collapse: collapse; width: 100%; }
//         th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
//         th { background-color: #4CAF50; color: white; }
//         .summary { background-color: #f2f2f2; padding: 15px; margin-bottom: 20px; }
//         .total { font-weight: bold; }
//     </style>
// </head>
// <body>
//     <h1>{{.Title}}</h1>

//     <div class="summary">
//         <h2>Summary</h2>
//         <p>Report Date: {{.Date}}</p>
//         <p>Total Records: {{len .Records}}</p>
//         <p class="total">Total Revenue: ${{.TotalRevenue}}</p>
//     </div>

//     <h2>Sales Records</h2>
//     <table>
//         <thead>
//             <tr>
//                 <th>Product</th>
//                 <th>Quantity</th>
//                 <th>Price</th>
//                 <th>Total</th>
//             </tr>
//         </thead>
//         <tbody>
//             {{range .Records}}
//             <tr>
//                 <td>{{.Product}}</td>
//                 <td>{{.Quantity}}</td>
//                 <td>${{.Price}}</td>
//                 <td>${{multiply .Quantity .Price}}</td>
//             </tr>
//             {{end}}
//         </tbody>
//     </table>

//     {{if .Notes}}
//     <h3>Notes</h3>
//     <ul>
//         {{range .Notes}}
//         <li>{{.}}</li>
//         {{end}}
//     </ul>
//     {{end}}
// </body>
// </html>
// `

//     // Custom function for multiplication
//     funcMap := template.FuncMap{
//         "multiply": func(a, b float64) float64 {
//             return a * b
//         },
//     }

//     tmpl := template.Must(template.New("report").Funcs(funcMap).Parse(htmlTemplate))

//     data := map[string]interface{}{
//         "Title": "Monthly Sales Report",
//         "Date":  "2024-01-15",
//         "Records": []map[string]interface{}{
//             {"Product": "Laptop", "Quantity": 5.0, "Price": 999.99},
//             {"Product": "Mouse", "Quantity": 20.0, "Price": 29.99},
//             {"Product": "Keyboard", "Quantity": 15.0, "Price": 79.99},
//         },
//         "TotalRevenue": 6798.75,
//         "Notes": []string{
//             "Strong performance in laptop sales",
//             "Mouse inventory running low",
//             "Consider promotion for keyboards",
//         },
//     }

//     // Output to file
//     file, _ := os.Create("report.html")
//     defer file.Close()
//     tmpl.Execute(file, data)

//     fmt.Println("HTML report generated: report.html")
// }

// func main() {
//     configTemplate := `
// # Application Configuration
// # Generated on: {{.GeneratedDate}}

// [server]
// host = {{.Server.Host}}
// port = {{.Server.Port}}
// {{if .Server.SSL}}
// ssl_enabled = true
// ssl_cert = {{.Server.SSLCert}}
// ssl_key = {{.Server.SSLKey}}
// {{end}}

// [database]
// type = {{.Database.Type}}
// host = {{.Database.Host}}
// port = {{.Database.Port}}
// name = {{.Database.Name}}
// {{if .Database.MaxConnections}}
// max_connections = {{.Database.MaxConnections}}
// {{end}}

// [features]
// {{range $key, $value := .Features}}
// {{$key}} = {{$value}}
// {{end}}

// [logging]
// level = {{.Logging.Level}}
// file = {{.Logging.File}}
// {{if .Logging.Rotate}}
// rotate = true
// max_size = {{.Logging.MaxSize}}MB
// {{end}}
// `

//     tmpl := template.Must(template.New("config").Parse(configTemplate))

//     data := map[string]interface{}{
//         "GeneratedDate": time.Now().Format("2006-01-02 15:04:05"),
//         "Server": map[string]interface{}{
//             "Host":    "localhost",
//             "Port":    8080,
//             "SSL":     true,
//             "SSLCert": "/path/to/cert.pem",
//             "SSLKey":  "/path/to/key.pem",
//         },
//         "Database": map[string]interface{}{
//             "Type":           "postgresql",
//             "Host":           "db.example.com",
//             "Port":           5432,
//             "Name":           "myapp_db",
//             "MaxConnections": 100,
//         },
//         "Features": map[string]bool{
//             "enable_api":      true,
//             "enable_websocket": true,
//             "enable_cache":    false,
//         },
//         "Logging": map[string]interface{}{
//             "Level":   "info",
//             "File":    "/var/log/app.log",
//             "Rotate":  true,
//             "MaxSize": 100,
//         },
//     }

//     file, _ := os.Create("config.ini")
//     defer file.Close()
//     tmpl.Execute(file, data)

//     fmt.Println("Configuration file generated: config.ini")
// }

// func main() {
//     structTemplate := `
// // Code generated by template. DO NOT EDIT.
// package {{.Package}}

// import (
// {{range .Imports}}
//     "{{.}}"
// {{end}}
// )

// // {{.StructName}} represents {{.Description}}
// type {{.StructName}} struct {
// {{range .Fields}}
//     {{.Name}} {{.Type}} ` + "`json:\"{{.JSONTag}}\"`" + `
// {{end}}
// }

// // New{{.StructName}} creates a new instance
// func New{{.StructName}}() *{{.StructName}} {
//     return &{{.StructName}}{}
// }

// // Validate validates the {{.StructName}}
// func (s *{{.StructName}}) Validate() error {
//     {{range .Fields}}
//     {{if .Required}}
//     if s.{{.Name}} == {{.ZeroValue}} {
//         return fmt.Errorf("{{.Name}} is required")
//     }
//     {{end}}
//     {{end}}
//     return nil
// }
// `

//     tmpl := template.Must(template.New("struct").Parse(structTemplate))

//     data := map[string]interface{}{
//         "Package":    "models",
//         "StructName": "User",
//         "Description": "a user in the system",
//         "Imports":    []string{"fmt", "time"},
//         "Fields": []map[string]interface{}{
//             {
//                 "Name":      "ID",
//                 "Type":      "int64",
//                 "JSONTag":   "id",
//                 "Required":  true,
//                 "ZeroValue": "0",
//             },
//             {
//                 "Name":      "Name",
//                 "Type":      "string",
//                 "JSONTag":   "name",
//                 "Required":  true,
//                 "ZeroValue": `""`,
//             },
//             {
//                 "Name":      "Email",
//                 "Type":      "string",
//                 "JSONTag":   "email",
//                 "Required":  true,
//                 "ZeroValue": `""`,
//             },
//             {
//                 "Name":      "CreatedAt",
//                 "Type":      "time.Time",
//                 "JSONTag":   "created_at",
//                 "Required":  false,
//                 "ZeroValue": "time.Time{}",
//             },
//         },
//     }

//     file, _ := os.Create("user_generated.go")
//     defer file.Close()
//     tmpl.Execute(file, data)

//     fmt.Println("Code generated: user_generated.go")
// }

// func main() {
//     // Define multiple templates
//     templates := `
// {{define "header"}}
// =================================
// {{.Title}}
// =================================
// {{end}}

// {{define "footer"}}
// ---------------------------------
// Generated: {{.Date}}
// ---------------------------------
// {{end}}

// {{define "userList"}}
// {{template "header" .}}

// Users:
// {{range .Users}}
// - {{.Name}} ({{.Email}})
// {{end}}

// {{template "footer" .}}
// {{end}}
// `

//     tmpl := template.Must(template.New("main").Parse(templates))

//     data := map[string]interface{}{
//         "Title": "User Directory",
//         "Date":  time.Now().Format("2006-01-02"),
//         "Users": []map[string]string{
//             {"Name": "Alice", "Email": "alice@example.com"},
//             {"Name": "Bob", "Email": "bob@example.com"},
//             {"Name": "Charlie", "Email": "charlie@example.com"},
//         },
//     }

//     tmpl.ExecuteTemplate(os.Stdout, "userList", data)
// }

// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strconv"
//     "strings"
//     "text/template"
//     "time"
// )

// func main() {
//     reader := bufio.NewReader(os.Stdin)

//     // Define templates with custom functions
//     funcMap := template.FuncMap{
//         "upper": strings.ToUpper,
//         "formatDate": func(t time.Time) string {
//             return t.Format("2006-01-02")
//         },
//     }

//     templates := map[string]string{
//         "invoice": `
// INVOICE #{{.InvoiceID}}
// Date: {{formatDate .Date}}

// Bill To:
// {{.CustomerName}}
// {{.CustomerAddress}}

// Items:
// {{range $i, $item := .Items}}
// {{add $i 1}}. {{$item.Name}} - ${{$item.Price}} x {{$item.Quantity}} = ${{multiply $item.Price $item.Quantity}}
// {{end}}

// Subtotal: ${{.Subtotal}}
// Tax ({{.TaxRate}}%): ${{.Tax}}
// -------------------
// TOTAL: ${{.Total}}

// Thank you for your business!
// `,
//         "certificate": `
// ╔══════════════════════════════════════════╗
// ║       CERTIFICATE OF COMPLETION          ║
// ╚══════════════════════════════════════════╝

// This is to certify that

//         {{upper .Name}}

// has successfully completed the course

//         "{{.CourseName}}"

// on {{formatDate .Date}}

//         Score: {{.Score}}%

// Instructor: {{.Instructor}}
// `,
//     }

//     // Add arithmetic functions
//     funcMap["add"] = func(a, b int) int { return a + b }
//     funcMap["multiply"] = func(a, b float64) float64 { return a * b }

//     // Parse templates
//     tmplMap := make(map[string]*template.Template)
//     for name, content := range templates {
//         tmplMap[name] = template.Must(template.New(name).Funcs(funcMap).Parse(content))
//     }

//     for {
//         fmt.Println("\n=== DOCUMENT GENERATOR ===")
//         fmt.Println("1. Generate Invoice")
//         fmt.Println("2. Generate Certificate")
//         fmt.Println("3. Exit")
//         fmt.Print("Choose option: ")

//         choice, _ := reader.ReadString('\n')
//         choice = strings.TrimSpace(choice)

//         switch choice {
//         case "1":
//             generateInvoice(reader, tmplMap["invoice"])
//         case "2":
//             generateCertificate(reader, tmplMap["certificate"])
//         case "3":
//             fmt.Println("Goodbye!")
//             return
//         default:
//             fmt.Println("Invalid choice")
//         }
//     }
// }

// func generateInvoice(reader *bufio.Reader, tmpl *template.Template) {
//     fmt.Print("Customer Name: ")
//     name, _ := reader.ReadString('\n')

//     fmt.Print("Customer Address: ")
//     address, _ := reader.ReadString('\n')

//     items := []map[string]interface{}{}
//     for {
//         fmt.Print("Item name (or 'done'): ")
//         itemName, _ := reader.ReadString('\n')
//         itemName = strings.TrimSpace(itemName)

//         if itemName == "done" {
//             break
//         }

//         fmt.Print("Price: ")
//         priceStr, _ := reader.ReadString('\n')
//         price, _ := strconv.ParseFloat(strings.TrimSpace(priceStr), 64)

//         fmt.Print("Quantity: ")
//         qtyStr, _ := reader.ReadString('\n')
//         qty, _ := strconv.Atoi(strings.TrimSpace(qtyStr))

//         items = append(items, map[string]interface{}{
//             "Name":     itemName,
//             "Price":    price,
//             "Quantity": qty,
//         })
//     }

//     // Calculate totals
//     subtotal := 0.0
//     for _, item := range items {
//         subtotal += item["Price"].(float64) * float64(item["Quantity"].(int))
//     }
//     tax := subtotal * 0.1
//     total := subtotal + tax

//     data := map[string]interface{}{
//         "InvoiceID":       fmt.Sprintf("INV-%d", time.Now().Unix()),
//         "Date":            time.Now(),
//         "CustomerName":    strings.TrimSpace(name),
//         "CustomerAddress": strings.TrimSpace(address),
//         "Items":           items,
//         "Subtotal":        subtotal,
//         "TaxRate":         10,
//         "Tax":             tax,
//         "Total":           total,
//     }

//     fmt.Println("\n" + strings.Repeat("=", 50))
//     tmpl.Execute(os.Stdout, data)
//     fmt.Println(strings.Repeat("=", 50))
// }

// func generateCertificate(reader *bufio.Reader, tmpl *template.Template) {
//     fmt.Print("Student Name: ")
//     name, _ := reader.ReadString('\n')

//     fmt.Print("Course Name: ")
//     course, _ := reader.ReadString('\n')

//     fmt.Print("Score: ")
//     scoreStr, _ := reader.ReadString('\n')
//     score, _ := strconv.Atoi(strings.TrimSpace(scoreStr))

//     fmt.Print("Instructor: ")
//     instructor, _ := reader.ReadString('\n')

//     data := map[string]interface{}{
//         "Name":       strings.TrimSpace(name),
//         "CourseName": strings.TrimSpace(course),
//         "Date":       time.Now(),
//         "Score":      score,
//         "Instructor": strings.TrimSpace(instructor),
//     }

//     fmt.Println()
//     tmpl.Execute(os.Stdout, data)
// }