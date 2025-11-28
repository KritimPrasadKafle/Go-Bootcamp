// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     file, _ := os.Open("app.log")
//     defer file.Close()

//     scanner := bufio.NewScanner(file)

//     fmt.Println("=== ERRORS FOUND ===")
//     for scanner.Scan() {
//         line := scanner.Text()
//         if strings.Contains(line, "ERROR") {
//             fmt.Println(line)
//         }
//     }
// }
// ```

// **app.log:**
// ```
// 2024-01-15 INFO: Server started
// 2024-01-15 ERROR: Database connection failed
// 2024-01-15 INFO: Retrying...
// 2024-01-15 ERROR: Timeout exceeded
// ```

// **Output:**
// ```
// === ERRORS FOUND ===
// 2024-01-15 ERROR: Database connection failed
// 2024-01-15 ERROR: Timeout exceeded

// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     file, _ := os.Open("app.log")
//     defer file.Close()

//     scanner := bufio.NewScanner(file)

//     stats := map[string]int{
//         "INFO":    0,
//         "ERROR":   0,
//         "WARNING": 0,
//     }

//     for scanner.Scan() {
//         line := scanner.Text()
//         for level := range stats {
//             if strings.Contains(line, level) {
//                 stats[level]++
//             }
//         }
//     }

//     fmt.Println("Log Statistics:")
//     for level, count := range stats {
//         fmt.Printf("  %s: %d\n", level, count)
//     }
// }

// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     // Open source file
//     input, _ := os.Open("input.txt")
//     defer input.Close()

//     // Create output file
//     output, _ := os.Create("output.txt")
//     defer output.Close()

//     scanner := bufio.NewScanner(input)
//     writer := bufio.NewWriter(output)

//     for scanner.Scan() {
//         line := scanner.Text()

//         // Transform: uppercase and add prefix
//         transformed := ">> " + strings.ToUpper(line)

//         writer.WriteString(transformed + "\n")
//     }

//     writer.Flush()  // Important: write buffered data to file
//     fmt.Println("Transformation complete!")
// }

// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     file, _ := os.Open("users.txt")
//     defer file.Close()

//     scanner := bufio.NewScanner(file)

//     fmt.Println("Active Users:")
//     for scanner.Scan() {
//         line := scanner.Text()

//         // Format: name,email,status
//         parts := strings.Split(line, ",")
//         if len(parts) == 3 {
//             name := parts[0]
//             email := parts[1]
//             status := parts[2]

//             if status == "active" {
//                 fmt.Printf("  %s (%s)\n", name, email)
//             }
//         }
//     }
// }
// ```

// **users.txt:**
// ```
// John,john@email.com,active
// Jane,jane@email.com,inactive
// Bob,bob@email.com,active
// ```

// **Output:**
// ```
// Active Users:
//   John (john@email.com)
//   Bob (bob@email.com)

//   package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     scanner := bufio.NewScanner(os.Stdin)  // Read from terminal input

//     for scanner.Scan() {
//         line := scanner.Text()

//         // Only print lines with "go"
//         if strings.Contains(strings.ToLower(line), "go") {
//             fmt.Println(line)
//         }
//     }
// }