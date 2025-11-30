package main

import (
    "fmt"
    "time"
)

func sayHello() {
    time.Sleep(1 * time.Second)
    fmt.Println("Hello from Goroutine!")
}

func main() {
    fmt.Println("1. Program started")
    
    go sayHello()  // Starts in background, doesn't block
    
    fmt.Println("2. After calling sayHello")
    
    time.Sleep(2 * time.Second)  // Wait for goroutine to finish
    
    fmt.Println("3. Program ending")
}
```

**Output:**
```
1. Program started
2. After calling sayHello
Hello from Goroutine!
3. Program ending
```

### Execution Flow
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   TIME    MAIN THREAD              GOROUTINE                    │
│   ════    ═══════════              ═════════                    │
│                                                                 │
│   0ms     Print "1. Program..."                                 │
│           │                                                     │
│   0ms     go sayHello() ──────────► Started                     │
│           │ (doesn't wait)          │                           │
│           │                         │ Sleeping...               │
│   0ms     Print "2. After..."       │                           │
│           │                         │                           │
│           │ Sleeping...             │                           │
│           │                         │                           │
│   1000ms  │                         Print "Hello from..."       │
│           │                         │                           │
│   2000ms  Print "3. Program..."     Done                        │
│           │                                                     │
│           Exit                                                  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘


package main

import (
    "fmt"
    "time"
)

func printNumbers() {
    for i := 1; i <= 5; i++ {
        fmt.Printf("Number: %d\n", i)
        time.Sleep(100 * time.Millisecond)
    }
}

func printLetters() {
    for _, letter := range "ABCDE" {
        fmt.Printf("Letter: %c\n", letter)
        time.Sleep(150 * time.Millisecond)
    }
}

func main() {
    go printNumbers()  // Start goroutine 1
    go printLetters()  // Start goroutine 2
    
    // Both run concurrently!
    time.Sleep(1 * time.Second)
}
```

**Output (interleaved!):**
```
Number: 1
Letter: A
Number: 2
Letter: B
Number: 3
Number: 4
Letter: C
Number: 5
Letter: D
Letter: E
```

### Visual: Concurrent Execution
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   TIME        printNumbers()         printLetters()             │
│   ════        ══════════════         ═════════════              │
│                                                                 │
│   0ms         Number: 1              Letter: A                  │
│   100ms       Number: 2                                         │
│   150ms                              Letter: B                  │
│   200ms       Number: 3                                         │
│   300ms       Number: 4              Letter: C                  │
│   400ms       Number: 5                                         │
│   450ms                              Letter: D                  │
│   600ms                              Letter: E                  │
│                                                                 │
│   They INTERLEAVE because they run at the same time!           │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘


package main

import (
    "fmt"
    "time"
)

func main() {
    // Anonymous goroutine (inline function)
    go func() {
        fmt.Println("Hello from anonymous goroutine!")
    }()
    
    // Anonymous goroutine with parameters
    name := "John"
    go func(n string) {
        fmt.Printf("Hello, %s!\n", n)
    }(name)  // Pass value here
    
    // Capturing variables (closure)
    message := "Captured message"
    go func() {
        fmt.Println(message)  // Uses outer variable
    }()
    
    time.Sleep(100 * time.Millisecond)
}
```

**Output:**
```
Hello from anonymous goroutine!
Hello, John!
Captured message

package main

import (
    "fmt"
    "time"
)

func doWork() error {
    time.Sleep(500 * time.Millisecond)
    return fmt.Errorf("an error occurred in doWork")
}

func main() {
    var err error  // Shared variable
    
    go func() {
        err = doWork()  // Assign result to shared variable
    }()
    
    // Wait for goroutine to complete
    time.Sleep(1 * time.Second)
    
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Success!")
    }
}
```

**Output:**
```
Error: an error occurred in doWork

package main

import "fmt"

func main() {
    go fmt.Println("This might not print!")
    
    // Program exits immediately!
    // Goroutine has no time to run
}
```

**Output:**
```
(nothing — program exits too fast!)
```
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   ⚠️ CRITICAL RULE                                              │
│                                                                 │
│   When main() exits, ALL goroutines are killed immediately!    │
│                                                                 │
│   main() does NOT wait for goroutines to finish.               │
│                                                                 │
│   Solutions:                                                    │
│   1. time.Sleep() — not reliable                               │
│   2. sync.WaitGroup — proper way                               │
│   3. Channels — best for communication                         │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘


package main

import (
    "fmt"
    "sync"
    "time"
)

func worker(id int, wg *sync.WaitGroup) {
    defer wg.Done()  // Signal completion when function returns
    
    fmt.Printf("Worker %d starting\n", id)
    time.Sleep(time.Duration(id) * 100 * time.Millisecond)
    fmt.Printf("Worker %d done\n", id)
}

func main() {
    var wg sync.WaitGroup
    
    for i := 1; i <= 3; i++ {
        wg.Add(1)           // Increment counter
        go worker(i, &wg)   // Start goroutine
    }
    
    wg.Wait()  // Block until all workers call Done()
    fmt.Println("All workers completed!")
}
```

**Output:**
```
Worker 1 starting
Worker 2 starting
Worker 3 starting
Worker 1 done
Worker 2 done
Worker 3 done
All workers completed!
```

### How WaitGroup Works
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   sync.WaitGroup                                                │
│                                                                 │
│   ┌────────────────────────────────────────────────────┐       │
│   │                                                    │       │
│   │   Counter: [  0  ]                                 │       │
│   │                                                    │       │
│   └────────────────────────────────────────────────────┘       │
│                                                                 │
│   wg.Add(1)   →   Counter: [  1  ]   (one more to wait for)    │
│   wg.Add(1)   →   Counter: [  2  ]                              │
│   wg.Add(1)   →   Counter: [  3  ]                              │
│                                                                 │
│   wg.Done()   →   Counter: [  2  ]   (one finished)            │
│   wg.Done()   →   Counter: [  1  ]                              │
│   wg.Done()   →   Counter: [  0  ]                              │
│                                                                 │
│   wg.Wait()   →   Blocks until counter reaches 0               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## 7. Goroutines vs Threads
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   GOROUTINES                       OS THREADS                   │
│   ══════════                       ══════════                   │
│                                                                 │
│   • Managed by Go runtime          • Managed by OS              │
│   • ~2KB stack (grows as needed)   • ~1-8MB fixed stack         │
│   • Can run millions               • Limited (thousands)        │
│   • Fast to create/destroy         • Expensive to create        │
│   • Multiplexed onto threads       • 1:1 with CPU               │
│                                                                 │
│   ┌────────────────────────────────────────────────────────┐   │
│   │                                                        │   │
│   │   100,000 Goroutines  ─────►  Few OS Threads           │   │
│   │                               (Go runtime manages)     │   │
│   │                                                        │   │
│   └────────────────────────────────────────────────────────┘   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

package main

import (
    "fmt"
    "sync"
)

func main() {
    var wg sync.WaitGroup
    
    // Create 10,000 goroutines — no problem!
    for i := 0; i < 10000; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            // Do some work
            _ = id * id
        }(i)
    }
    
    wg.Wait()
    fmt.Println("10,000 goroutines completed!")
}


package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 3; i++ {
        go func() {
            fmt.Println(i)  // ❌ Captures variable, not value!
        }()
    }
    time.Sleep(100 * time.Millisecond)
}
```

**Output (unexpected!):**
```
3
3
3


package main

import (
    "fmt"
    "time"
)

func main() {
    for i := 0; i < 3; i++ {
        go func(n int) {
            fmt.Println(n)  // ✅ Each goroutine gets its own copy
        }(i)  // Pass value here
    }
    time.Sleep(100 * time.Millisecond)
}
```

**Output (correct!):**
```
0
1
2
```
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   WRONG: go func() { fmt.Println(i) }()                        │
│                                                                 │
│   Loop:  i=0 → i=1 → i=2 → i=3 (exits)                         │
│                                                                 │
│   Goroutine 1 reads i ──────────────────────► 3                │
│   Goroutine 2 reads i ──────────────────────► 3                │
│   Goroutine 3 reads i ──────────────────────► 3                │
│                                                                 │
│   All read the FINAL value of i!                               │
│                                                                 │
│   ─────────────────────────────────────────────────────────    │
│                                                                 │
│   CORRECT: go func(n int) { fmt.Println(n) }(i)                │
│                                                                 │
│   Loop:  i=0 → i=1 → i=2                                       │
│                                                                 │
│   Goroutine 1 receives copy: n=0 ──► prints 0                  │
│   Goroutine 2 receives copy: n=1 ──► prints 1                  │
│   Goroutine 3 receives copy: n=2 ──► prints 2                  │
│                                                                 │
│   Each goroutine has its OWN copy!                             │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘


package main

import (
    "fmt"
    "sync"
    "time"
)

func downloadFile(filename string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    fmt.Printf("Starting download: %s\n", filename)
    
    // Simulate download time
    time.Sleep(time.Duration(len(filename)*100) * time.Millisecond)
    
    fmt.Printf("Finished download: %s\n", filename)
}

func main() {
    files := []string{"image.jpg", "video.mp4", "document.pdf", "data.json"}
    
    var wg sync.WaitGroup
    
    start := time.Now()
    
    for _, file := range files {
        wg.Add(1)
        go downloadFile(file, &wg)
    }
    
    wg.Wait()
    
    fmt.Printf("\nAll downloads completed in %v\n", time.Since(start))
}
```

**Output:**
```
Starting download: image.jpg
Starting download: video.mp4
Starting download: document.pdf
Starting download: data.json
Finished download: data.json
Finished download: image.jpg
Finished download: document.pdf
Finished download: video.mp4

All downloads completed in 900ms
```

Without goroutines, this would take ~3.5 seconds!

---

## 10. Goroutine Lifecycle
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   GOROUTINE LIFECYCLE                                           │
│                                                                 │
│   ┌─────────┐                                                   │
│   │ Created │ ──► go func()                                     │
│   └────┬────┘                                                   │
│        │                                                        │
│        ▼                                                        │
│   ┌─────────┐                                                   │
│   │ Runnable│ ──► Ready to run, waiting for CPU                 │
│   └────┬────┘                                                   │
│        │                                                        │
│        ▼                                                        │
│   ┌─────────┐                                                   │
│   │ Running │ ──► Actually executing on CPU                     │
│   └────┬────┘                                                   │
│        │                                                        │
│        ├──────────────────────────────────┐                     │
│        ▼                                  ▼                     │
│   ┌─────────┐                       ┌───────────┐               │
│   │ Blocked │ ◄─────────────────────│ Completed │               │
│   │ (I/O,   │                       └───────────┘               │
│   │ channel,│                                                   │
│   │ mutex)  │                                                   │
│   └────┬────┘                                                   │
│        │                                                        │
│        └──► Back to Runnable when unblocked                     │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Summary: Goroutines Key Points
```
┌─────────────────────────────────────────────────────────────────┐
│                    GOROUTINES SUMMARY                           │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│   SYNTAX                                                        │
│   ══════                                                        │
│   go functionName()                  // Named function          │
│   go functionName(arg1, arg2)        // With arguments          │
│   go func() { ... }()                // Anonymous function      │
│   go func(x int) { ... }(value)      // Anonymous with params   │
│                                                                 │
│   KEY RULES                                                     │
│   ═════════                                                     │
│   • Goroutines are NON-BLOCKING                                │
│   • main() does NOT wait for goroutines                        │
│   • Cannot directly get return values                          │
│   • Pass loop variables as parameters                          │
│   • Use sync.WaitGroup to wait for completion                  │
│                                                                 │
│   WAITING METHODS                                               │
│   ═══════════════                                               │
│   • time.Sleep()      — unreliable (avoid in production)       │
│   • sync.WaitGroup    — proper way to wait                     │
│   • Channels          — best for communication (next topic)    │
│                                                                 │
│   BENEFITS                                                      │
│   ════════                                                      │
│   • Lightweight (~2KB vs ~1MB for OS threads)                  │
│   • Can run millions concurrently                              │
│   • Automatically scheduled by Go runtime                       │
│   • Great for I/O-bound tasks                                  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘


// === BASIC GOROUTINE ===
go myFunction()

// === WITH PARAMETERS ===
go myFunction(arg1, arg2)

// === ANONYMOUS ===
go func() {
    // code
}()

// === ANONYMOUS WITH PARAMS ===
go func(n int) {
    fmt.Println(n)
}(value)

// === WAITGROUP ===
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // work
}()
wg.Wait()

// === LOOP PATTERN ===
for i := 0; i < 10; i++ {
    wg.Add(1)
    go func(n int) {  // Pass i as parameter!
        defer wg.Done()
        fmt.Println(n)
    }(i)
}
wg.Wait()


┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   defer func()                    go func()                     │
│   ════════════                    ═════════                     │
│                                                                 │
│   "Run this LATER,                "Run this NOW,                │
│    when function exits"            in the BACKGROUND"           │
│                                                                 │
│   • Delays execution              • Starts immediately          │
│   • Runs in SAME goroutine        • Runs in NEW goroutine       │
│   • BLOCKS at function end        • NON-BLOCKING                │
│   • Guaranteed to run             • May not complete            │
│   • For cleanup                   • For concurrency             │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   DEFER — "Do this when I'm leaving"                           │
│   ══════════════════════════════════                            │
│                                                                 │
│   func example() {                                              │
│       defer cleanup()  ─────────────────────┐                   │
│       │                                     │ (scheduled)       │
│       step1()                               │                   │
│       │                                     │                   │
│       step2()                               │                   │
│       │                                     │                   │
│       step3()                               │                   │
│       │                                     │                   │
│       return ◄──────────────────────────────┘                   │
│       │        cleanup() runs HERE                              │
│   }                                                             │
│                                                                 │
│   Order: step1 → step2 → step3 → cleanup                       │
│   Everything in SAME thread, SEQUENTIAL                         │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   GOROUTINE — "Do this in parallel"                            │
│   ═════════════════════════════════                             │
│                                                                 │
│   func example() {                                              │
│       go backgroundTask() ──────► Starts in NEW thread          │
│       │                           │                             │
│       step1()                     │ (running simultaneously)    │
│       │                           │                             │
│       step2()                     │                             │
│       │                           │                             │
│       step3()                     │                             │
│       │                           │                             │
│       return                      │ (might still be running!)   │
│   }                               ▼                             │
│                                                                 │
│   Order: backgroundTask + step1 + step2 + step3 (CONCURRENT)   │
│   DIFFERENT threads, PARALLEL                                   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘


