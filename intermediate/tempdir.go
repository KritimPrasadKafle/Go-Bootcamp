// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     // Create temp file
//     tempFile, err := os.CreateTemp("", "example-*.txt")
//     if err != nil {
//         panic(err)
//     }

//     // Clean up when done
//     defer os.Remove(tempFile.Name())
//     defer tempFile.Close()

//     fmt.Println("Created:", tempFile.Name())

//     // Write to temp file
//     tempFile.WriteString("Hello, temporary world!\n")

//     // Read it back
//     content, _ := os.ReadFile(tempFile.Name())
//     fmt.Println("Content:", string(content))
// }

// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     // Create temp file in current directory instead of /tmp
//     tempFile, err := os.CreateTemp(".", "local-*.tmp")
//     if err != nil {
//         panic(err)
//     }
//     defer os.Remove(tempFile.Name())
//     defer tempFile.Close()

//     fmt.Println("Created in current dir:", tempFile.Name())
//     // Result: ./local-123456789.tmp
// }

// package main

// import (
//     "fmt"
//     "os"
//     "path/filepath"
// )

// func main() {
//     // Create temp directory
//     tempDir, err := os.MkdirTemp("", "workspace-*")
//     if err != nil {
//         panic(err)
//     }
//     defer os.RemoveAll(tempDir)  // Clean up everything inside

//     fmt.Println("Workspace:", tempDir)

//     // Create files inside temp directory
//     file1 := filepath.Join(tempDir, "data.txt")
//     file2 := filepath.Join(tempDir, "config.json")

//     os.WriteFile(file1, []byte("some data"), 0644)
//     os.WriteFile(file2, []byte(`{"key": "value"}`), 0644)

//     // List contents
//     entries, _ := os.ReadDir(tempDir)
//     for _, e := range entries {
//         fmt.Println(" -", e.Name())
//     }
// }

// package main

// import (
//     "fmt"
//     "io"
//     "os"
// )

// func processUpload(data io.Reader) (string, error) {
//     // Create temp file for upload
//     tempFile, err := os.CreateTemp("", "upload-*.tmp")
//     if err != nil {
//         return "", err
//     }
//     defer tempFile.Close()

//     // Copy uploaded data to temp file
//     _, err = io.Copy(tempFile, data)
//     if err != nil {
//         os.Remove(tempFile.Name())
//         return "", err
//     }

//     // Validate, process, etc.
//     // ...

//     return tempFile.Name(), nil
// }

// func main() {
//     // Simulate upload data
//     sourceFile, _ := os.Open("source.txt")
//     defer sourceFile.Close()

//     tempPath, err := processUpload(sourceFile)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     defer os.Remove(tempPath)  // Clean up after processing

//     fmt.Println("Uploaded to:", tempPath)
// }