// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// Read server configuration from environment
// 	port := os.Getenv("PORT")
// 	host := os.Getenv("HOST")

// 	// What if they're not set? Empty string!
// 	fmt.Printf("Port: '%s'\n", port) // Port: ''
// 	fmt.Printf("Host: '%s'\n", host) // Host: ''

// 	// Always provide defaults!
// 	if port == "" {
// 		port = "8080"
// 	}
// 	if host == "" {
// 		host = "localhost"
// 	}

// 	fmt.Printf("Server starting at %s:%s\n", host, port)
// }

// package main

// import (
//     "fmt"
//     "os"
// )

// type DBConfig struct {
//     Host     string
//     Port     string
//     User     string
//     Password string
//     Database string
// }

// func loadDBConfig() DBConfig {
//     return DBConfig{
//         Host:     getEnvOrDefault("DB_HOST", "localhost"),
//         Port:     getEnvOrDefault("DB_PORT", "5432"),
//         User:     getEnvOrDefault("DB_USER", "postgres"),
//         Password: os.Getenv("DB_PASSWORD"), // No default for secrets!
//         Database: getEnvOrDefault("DB_NAME", "myapp"),
//     }
// }

// func getEnvOrDefault(key, defaultValue string) string {
//     value := os.Getenv(key)
//     if value == "" {
//         return defaultValue
//     }
//     return value
// }

// func main() {
//     config := loadDBConfig()

//     if config.Password == "" {
//         fmt.Println("ERROR: DB_PASSWORD is required!")
//         os.Exit(1)
//     }

//     fmt.Printf("Connecting to %s@%s:%s/%s\n",
//         config.User, config.Host, config.Port, config.Database)
// }

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func isFeatureEnabled(featureName string) bool {
// 	value, exists := os.LookupEnv(featureName)

// 	if !exists {
// 		// Not configured — use default behavior
// 		return false
// 	}

// 	// Explicitly set to empty or "false" means disabled
// 	// Any other value means enabled
// 	return value != "" && value != "false" && value != "0"
// }

// func main() {
// 	os.Setenv("FEATURE_DARK_MODE", "true")
// 	os.Setenv("FEATURE_BETA", "") // Explicitly disabled
// 	// FEATURE_NEW_UI is not set at all

// 	fmt.Println("Dark Mode:", isFeatureEnabled("FEATURE_DARK_MODE")) // true
// 	fmt.Println("Beta:", isFeatureEnabled("FEATURE_BETA"))           // false
// 	fmt.Println("New UI:", isFeatureEnabled("FEATURE_NEW_UI"))       // false
// }

// package main

// import (
//     "fmt"
//     "os"
// )

// func isFeatureEnabled(featureName string) bool {
//     value, exists := os.LookupEnv(featureName)

//     if !exists {
//         // Not configured — use default behavior
//         return false
//     }

//     // Explicitly set to empty or "false" means disabled
//     // Any other value means enabled
//     return value != "" && value != "false" && value != "0"
// }

// func main() {
//     os.Setenv("FEATURE_DARK_MODE", "true")
//     os.Setenv("FEATURE_BETA", "")        // Explicitly disabled
//     // FEATURE_NEW_UI is not set at all

//     fmt.Println("Dark Mode:", isFeatureEnabled("FEATURE_DARK_MODE"))  // true
//     fmt.Println("Beta:", isFeatureEnabled("FEATURE_BETA"))            // false
//     fmt.Println("New UI:", isFeatureEnabled("FEATURE_NEW_UI"))        // false
// }

// package main

// import (
//     "fmt"
//     "os"
// )

// func main() {
//     // This only affects THIS Go program
//     os.Setenv("MY_VAR", "hello")

//     fmt.Println("Inside Go:", os.Getenv("MY_VAR"))  // hello

//     // When program exits, MY_VAR is gone!
//     // Your shell won't see it
// }


