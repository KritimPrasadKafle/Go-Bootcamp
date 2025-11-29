// package main

// import (
// 	_ "embed"
// 	"fmt"
// )

// //go:embed config.json
// var configData string

// //go:embed logo.png
// var logoBytes []byte

// func main() {
// 	fmt.Println("Config:", configData)
// 	fmt.Printf("Logo size: %d bytes\n", len(logoBytes))
// }

// package main

// import (
//     "embed"
//     "fmt"
// )

// //go:embed file1.txt file2.txt file3.txt
// var files embed.FS

// func main() {
//     // Read each file
//     for _, name := range []string{"file1.txt", "file2.txt", "file3.txt"} {
//         data, err := files.ReadFile(name)
//         if err != nil {
//             fmt.Println("Error:", err)
//             continue
//         }
//         fmt.Printf("%s: %s\n", name, string(data))
//     }
// }

// package main

// import (
//     "embed"
//     "fmt"
// )

// //go:embed templates/*.html
// var templates embed.FS

// //go:embed static/*.css static/*.js
// var staticFiles embed.FS

// //go:embed assets/*
// var allAssets embed.FS

// func main() {
//     // Read a template
//     html, _ := templates.ReadFile("templates/index.html")
//     fmt.Println(string(html))

//     // List all entries
//     entries, _ := templates.ReadDir("templates")
//     for _, e := range entries {
//         fmt.Println("Template:", e.Name())
//     }
// }

// package main

// import (
//     "embed"
//     "fmt"
//     "io/fs"
// )

// //go:embed static
// var staticFS embed.FS

// func main() {
//     // Walk entire embedded tree
//     fs.WalkDir(staticFS, ".", func(path string, d fs.DirEntry, err error) error {
//         if err != nil {
//             return err
//         }

//         if d.IsDir() {
//             fmt.Printf("[DIR]  %s\n", path)
//         } else {
//             info, _ := d.Info()
//             fmt.Printf("[FILE] %s (%d bytes)\n", path, info.Size())
//         }
//         return nil
//     })
// }

// package main

// import (
//     "embed"
//     "io/fs"
//     "log"
//     "net/http"
// )

// //go:embed static
// var staticFiles embed.FS

// func main() {
//     // Strip "static" prefix for cleaner URLs
//     staticContent, err := fs.Sub(staticFiles, "static")
//     if err != nil {
//         log.Fatal(err)
//     }

//     // Serve embedded files
//     http.Handle("/", http.FileServer(http.FS(staticContent)))

//     log.Println("Server starting on :8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }
// ```

// **Directory structure:**
// ```
// project/
// ├── main.go
// └── static/
//     ├── index.html
//     ├── css/
//     │   └── style.css
//     └── js/
//         └── app.js