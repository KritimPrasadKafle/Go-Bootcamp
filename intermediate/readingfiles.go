// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	// METHOD 1: Read entire file at once (simplest)
// 	data, err := os.ReadFile("output.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println("Full content:", string(data))

// 	// METHOD 2: Read with fixed buffer
// 	file, _ := os.Open("output.txt")
// 	defer file.Close()

// 	buffer := make([]byte, 10) // Read 10 bytes at a time
// 	for {
// 		n, err := file.Read(buffer)
// 		if err == io.EOF {
// 			break // End of file
// 		}
// 		fmt.Print(string(buffer[:n]))
// 	}

// 	// METHOD 3: Scanner - line by line
// 	file2, _ := os.Open("output.txt")
// 	defer file2.Close()

// 	scanner := bufio.NewScanner(file2)
// 	lineNum := 1
// 	for scanner.Scan() {
// 		fmt.Printf("%d: %s\n", lineNum, scanner.Text())
// 		lineNum++
// 	}

// 	// METHOD 4: Buffered reader - more control
// 	file3, _ := os.Open("output.txt")
// 	defer file3.Close()

// 	reader := bufio.NewReader(file3)
// 	for {
// 		line, err := reader.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Print(line)
// 	}
// }

// package main

// import (
//     "bufio"
//     "fmt"
//     "os"
//     "strings"
// )

// func main() {
//     config := make(map[string]string)

//     file, err := os.Open("config.txt")
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     defer file.Close()

//     scanner := bufio.NewScanner(file)
//     for scanner.Scan() {
//         line := scanner.Text()

//         // Skip empty lines and comments
//         if line == "" || strings.HasPrefix(line, "#") {
//             continue
//         }

//         // Parse key=value
//         parts := strings.SplitN(line, "=", 2)
//         if len(parts) == 2 {
//             config[parts[0]] = parts[1]
//         }
//     }

//     fmt.Println("Config loaded:", config)
// }
// ```

// **config.txt:**
// ```
// # Database settings
// host=localhost
// port=5432
// user=admin
// ```

// **Output:**
// ```
// Config loaded: map[host:localhost port:5432 user:admin]
