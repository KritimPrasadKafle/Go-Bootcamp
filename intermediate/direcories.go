package main

import (
	"fmt"
	"os"
)

func main() {
	// Project structure to create
	dirs := []string{
		"myproject/cmd",
		"myproject/internal/handlers",
		"myproject/internal/models",
		"myproject/pkg/utils",
		"myproject/configs",
	}

	for _, dir := range dirs {
		err := os.MkdirAll(dir, 0755)
		if err != nil {
			fmt.Println("Error creating", dir, ":", err)
			continue
		}
		fmt.Println("Created:", dir)
	}

	// Create some files
	files := map[string]string{
		"myproject/main.go":           "package main\n\nfunc main() {}\n",
		"myproject/configs/config.go": "package configs\n",
		"myproject/README.md":         "# My Project\n",
	}

	for path, content := range files {
		os.WriteFile(path, []byte(content), 0644)
		fmt.Println("Created file:", path)
	}
}


package main

import (
    "fmt"
    "os"
)

func main() {
    entries, err := os.ReadDir(".")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("%-30s %-10s %s\n", "NAME", "TYPE", "SIZE")
    fmt.Println("--------------------------------------------------")

    for _, entry := range entries {
        info, _ := entry.Info()
        
        fileType := "FILE"
        if entry.IsDir() {
            fileType = "DIR"
        }

        fmt.Printf("%-30s %-10s %d bytes\n", 
            entry.Name(), 
            fileType, 
            info.Size(),
        )
    }
}

package main

import (
    "fmt"
    "io/fs"
    "path/filepath"
    "strings"
)

func main() {
    targetExt := ".go"
    var goFiles []string

    filepath.WalkDir(".", func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // Skip hidden directories
        if d.IsDir() && strings.HasPrefix(d.Name(), ".") {
            return filepath.SkipDir
        }

        // Collect .go files
        if !d.IsDir() && filepath.Ext(path) == targetExt {
            goFiles = append(goFiles, path)
        }

        return nil
    })

    fmt.Printf("Found %d Go files:\n", len(goFiles))
    for _, f := range goFiles {
        fmt.Println(" ", f)
    }
}

package main

import (
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
)

func copyStructure(src, dst string) error {
    return filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
        if err != nil {
            return err
        }

        // Calculate destination path
        relPath, _ := filepath.Rel(src, path)
        dstPath := filepath.Join(dst, relPath)

        if d.IsDir() {
            return os.MkdirAll(dstPath, 0755)
        }

        // For files, just create empty file (or copy content)
        return os.WriteFile(dstPath, []byte{}, 0644)
    })
}

func main() {
    err := copyStructure("myproject", "myproject_backup")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Structure copied!")
}

package main

import (
    "fmt"
    "os"
)

func main() {
    tempDir := "temp_work"

    // Create temp directory
    os.MkdirAll(tempDir, 0755)

    // Schedule cleanup
    defer func() {
        fmt.Println("Cleaning up...")
        os.RemoveAll(tempDir)
    }()

    // Do work in temp directory
    os.WriteFile(filepath.Join(tempDir, "data.txt"), []byte("temp data"), 0644)
    
    fmt.Println("Working with temp files...")
    // ... more operations

    // Cleanup happens automatically when function exits
}