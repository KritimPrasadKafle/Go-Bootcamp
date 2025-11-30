package main

import (
	"log"
	"os"
)

func main() {
	// Fatal — use for unrecoverable startup errors
	// Program CANNOT continue without this
	file, err := os.Open("config.yaml")
	if err != nil {
		log.Fatalf("Cannot start without config: %v", err)
		// Program exits with code 1
		// Nothing after this runs
	}
	defer file.Close()

	// Panic — use for programmer errors (bugs)
	// Something that "should never happen"
	var data []int
	if data == nil {
		log.Panic("data slice was not initialized!")
		// This can be recovered with recover()
	}
}


package main

import "log"

func main() {
    log.Println("Default message")
    
    log.SetPrefix("INFO: ")
    log.Println("This is an info message.")
    
    log.SetPrefix("ERROR: ")
    log.Println("This is an error message.")
    
    log.SetPrefix("") // Remove prefix
    log.Println("Back to default")
}
```

**Output:**
```
2025/01/15 14:30:00 Default message
INFO: 2025/01/15 14:30:00 This is an info message.
ERROR: 2025/01/15 14:30:00 This is an error message.
2025/01/15 14:30:00 Back to default


package main

import "log"

func main() {
    // Default flags (Ldate | Ltime)
    log.Println("Default flags")
    
    // Add file and line number
    log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
    log.Println("With short file")
    
    // Full file path
    log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
    log.Println("With long file")
    
    // Microsecond precision
    log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
    log.Println("With microseconds")
    
    // No flags at all
    log.SetFlags(0)
    log.Println("No timestamp, just message")
}
```

**Output:**
```
2025/01/15 14:30:00 Default flags
2025/01/15 14:30:00 main.go:12: With short file
2025/01/15 14:30:00 /home/user/project/main.go:16: With long file
2025/01/15 14:30:00.123456 With microseconds
No timestamp, just message


package main

import (
    "log"
    "os"
)

// Create loggers at package level
var (
    debugLogger = log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
    infoLogger  = log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime)
    warnLogger  = log.New(os.Stdout, "WARN:  ", log.Ldate|log.Ltime)
    errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
)

func main() {
    debugLogger.Println("Starting application...")
    infoLogger.Println("Server listening on port 8080")
    warnLogger.Println("Cache miss for user 123")
    errorLogger.Println("Failed to connect to database")
}
```

**Output:**
```
DEBUG: 2025/01/15 14:30:00 main.go:16: Starting application...
INFO:  2025/01/15 14:30:00 Server listening on port 8080
WARN:  2025/01/15 14:30:00 Cache miss for user 123
ERROR: 2025/01/15 14:30:00 main.go:19: Failed to connect to database


package main

import (
    "log"
    "os"
)

func main() {
    // Open or create log file
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer file.Close()
    
    // Create logger that writes to file
    fileLogger := log.New(file, "", log.Ldate|log.Ltime|log.Lshortfile)
    
    fileLogger.Println("Application started")
    fileLogger.Println("Processing request...")
    fileLogger.Println("Request completed")
}
```

**app.log contents:**
```
2025/01/15 14:30:00 main.go:18: Application started
2025/01/15 14:30:00 main.go:19: Processing request...
2025/01/15 14:30:00 main.go:20: Request completed

package main

import (
    "log"
    "os"
)

func main() {
    file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    defer file.Close()
    
    // Multiple loggers, all writing to same file
    debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
    infoLogger := log.New(file, "INFO:  ", log.Ldate|log.Ltime)
    warnLogger := log.New(file, "WARN:  ", log.Ldate|log.Ltime)
    errorLogger := log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
    
    // Simulate application activity
    infoLogger.Println("Server started")
    debugLogger.Println("Loading configuration")
    infoLogger.Println("Listening on port 8080")
    warnLogger.Println("Slow response time: 2.5s")
    errorLogger.Println("Connection refused to database")
}
```

**app.log contents:**
```
INFO:  2025/01/15 14:30:00 Server started
DEBUG: 2025/01/15 14:30:00 main.go:20: Loading configuration
INFO:  2025/01/15 14:30:00 Listening on port 8080
WARN:  2025/01/15 14:30:00 Slow response time: 2.5s
ERROR: 2025/01/15 14:30:00 main.go:23: Connection refused to database


