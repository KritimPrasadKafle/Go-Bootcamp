// // package main

// // import (
// //     "fmt"
// //     "path/filepath"
// //     "strings"
// // )

// // func main() {
// //     path := "/home/user/projects/app/main.go"

// //     fmt.Println("=== Path Analysis ===")
// //     fmt.Printf("Full path:  %s\n", path)
// //     fmt.Printf("Directory:  %s\n", filepath.Dir(path))
// //     fmt.Printf("Filename:   %s\n", filepath.Base(path))
// //     fmt.Printf("Extension:  %s\n", filepath.Ext(path))
// //     fmt.Printf("Name only:  %s\n", strings.TrimSuffix(filepath.Base(path), filepath.Ext(path)))
// //     fmt.Printf("Is absolute: %t\n", filepath.IsAbs(path))
// // }
// // ```

// // **Output:**
// // ```
// // === Path Analysis ===
// // Full path:  /home/user/projects/app/main.go
// // Directory:  /home/user/projects/app
// // Filename:   main.go
// // Extension:  .go
// // Name only:  main
// // Is absolute: true

// // package main

// // import (
// //     "fmt"
// //     "path/filepath"
// // )

// // func main() {
// //     // Find all .txt files
// //     matches, err := filepath.Glob("data/*.txt")
// //     if err != nil {
// //         fmt.Println("Error:", err)
// //         return
// //     }

// //     for _, match := range matches {
// //         fmt.Println("Found:", match)
// //     }
// // }

// package main

// import (
//     "fmt"
//     "os"
//     "path/filepath"
// )

// func main() {
//     root := "."

//     filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
//         if err != nil {
//             return err
//         }

//         if !info.IsDir() && filepath.Ext(path) == ".go" {
//             fmt.Println("Go file:", path)
//         }
//         return nil
//     })
// }

// package main

// import (
//     "fmt"
//     "path/filepath"
// )

// func main() {
//     baseDir := "/var/app/uploads"
//     userInput := "../../../etc/passwd"  // Malicious input!

//     // Clean and join
//     safePath := filepath.Join(baseDir, filepath.Clean(userInput))
//     fmt.Println("Combined:", safePath)

//     // Verify it's still under baseDir
//     if !strings.HasPrefix(safePath, baseDir) {
//         fmt.Println("Security warning: Path escape detected!")
//     }
// }

// package main

// import (
//     "fmt"
//     "path/filepath"
//     "strings"
// )

// func changeExtension(path, newExt string) string {
//     return strings.TrimSuffix(path, filepath.Ext(path)) + newExt
// }

// func main() {
//     original := "/home/user/photo.jpg"
//     converted := changeExtension(original, ".png")
//     fmt.Println(converted)  // /home/user/photo.png
// }



