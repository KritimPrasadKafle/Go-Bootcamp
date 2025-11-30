package main

import (
    "fmt"
    "time"
)

// Simulated API calls
func fetchUserAPI(ch chan string) {
    time.Sleep(1 * time.Second)
    ch <- "User: John Doe"
}

func fetchOrdersAPI(ch chan string) {
    time.Sleep(1 * time.Second)
    ch <- "Orders: 5 items"
}

func fetchRecommendationsAPI(ch chan string) {
    time.Sleep(1 * time.Second)
    ch <- "Recommendations: 10 products"
}

func main() {
    start := time.Now()
    
    // Create channels
    userCh := make(chan string)
    ordersCh := make(chan string)
    recoCh := make(chan string)
    
    // Start all API calls in PARALLEL (goroutines)
    go fetchUserAPI(userCh)
    go fetchOrdersAPI(ordersCh)
    go fetchRecommendationsAPI(recoCh)
    
    // WAIT for all results (channels)
    user := <-userCh
    orders := <-ordersCh
    recommendations := <-recoCh
    
    fmt.Println(user)
    fmt.Println(orders)
    fmt.Println(recommendations)
    fmt.Printf("\nTotal time: %v\n", time.Since(start))
}
```

**Output:**
```
User: John Doe
Orders: 5 items
Recommendations: 10 products

Total time: 1.001234s
```

**Without goroutines, it would take 3 seconds!**
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   SEQUENTIAL (No goroutines)       PARALLEL (With goroutines)   │
│   ══════════════════════════       ═════════════════════════    │
│                                                                 │
│   fetchUser()    1 sec             go fetchUser()    ─┐         │
│        ↓                           go fetchOrders()   ├─ 1 sec  │
│   fetchOrders()  1 sec             go fetchReco()    ─┘         │
│        ↓                                                        │
│   fetchReco()    1 sec             <-userCh     ─┐              │
│        ↓                           <-ordersCh    ├─ ~0 sec      │
│   Done           3 sec             <-recoCh     ─┘  (all ready) │
│                                    Done          1 sec          │
│                                                                 │
│   3x SLOWER!                       3x FASTER!                   │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Summary
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   YOUR UNDERSTANDING IS CORRECT! ✅                             │
│                                                                 │
│   GOROUTINES (go func)                                         │
│   ════════════════════                                          │
│   • "Run this in the background"                               │
│   • Parallel execution                                          │
│   • Non-blocking                                                │
│   • Like: async, Promise, Thread.start()                       │
│                                                                 │
│   CHANNELS (<-ch)                                              │
│   ═══════════════                                               │
│   • "Wait here until data arrives"                             │
│   • Synchronization                                             │
│   • Blocking (like await)                                      │
│   • Safe communication between goroutines                      │
│                                                                 │
│   TOGETHER THEY SOLVE                                          │
│   ═══════════════════                                           │
│   • Parallel execution (goroutines)                            │
│   • Waiting for results (channels)                             │
│   • Getting return values (channels)                           │
│   • Knowing when done (channels)                               │
│   • Thread-safe communication (channels)                       │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘



package main

import (
    "fmt"
    "time"
)

func fetchUser(ch chan map[string]string) {
    time.Sleep(1 * time.Second)
    ch <- map[string]string{"name": "John"}
}

func fetchPosts(ch chan []string) {
    time.Sleep(1 * time.Second)
    ch <- []string{"Post 1", "Post 2"}
}

func main() {
    userCh := make(chan map[string]string)
    postsCh := make(chan []string)
    
    // Run in parallel (like Promise.all)
    go fetchUser(userCh)
    go fetchPosts(postsCh)
    
    // Wait for both (like await)
    user := <-userCh
    posts := <-postsCh
    
    fmt.Println(user, posts)
}
```

---

## Complete Mental Model
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   GOROUTINES + CHANNELS = COMPLETE ASYNC SOLUTION               │
│                                                                 │
│   ┌─────────────────────────────────────────────────────────┐  │
│   │                                                         │  │
│   │   go func()        "Start this task in background"      │  │
│   │   ═════════        (non-blocking, parallel)             │  │
│   │                                                         │  │
│   │   Similar to:                                           │  │
│   │   • JavaScript: async function, Promise                 │  │
│   │   • Python: asyncio.create_task()                       │  │
│   │   • Java: new Thread().start()                          │  │
│   │                                                         │  │
│   └─────────────────────────────────────────────────────────┘  │
│                              +                                  │
│   ┌─────────────────────────────────────────────────────────┐  │
│   │                                                         │  │
│   │   <-channel        "Wait here until data arrives"       │  │
│   │   ═════════        (blocking, synchronization)          │  │
│   │                                                         │  │
│   │   Similar to:                                           │  │
│   │   • JavaScript: await                                   │  │
│   │   • Python: await                                       │  │
│   │   • Java: future.get()                                  │  │
│   │                                                         │  │
│   └─────────────────────────────────────────────────────────┘  │
│                              =                                  │
│   ┌─────────────────────────────────────────────────────────┐  │
│   │                                                         │  │
│   │   SAFE, EFFICIENT CONCURRENT PROGRAMMING                │  │
│   │                                                         │  │
│   │   • Run tasks in parallel (goroutines)                 │  │
│   │   • Wait for results when needed (channels)            │  │
│   │   • No race conditions (channels are thread-safe)      │  │
│   │   • No callbacks hell (sequential-looking code)        │  │
│   │                                                         │  │
│   └─────────────────────────────────────────────────────────┘  │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘



package main

import "fmt"

func calculate(ch chan int) {
    result := 42 * 2
    ch <- result  // Send result through channel
}

func main() {
    ch := make(chan int)
    
    go calculate(ch)
    
    result := <-ch  // Receive result
    
    fmt.Println("Result:", result)
}
```

**Output:**
```
Result: 84
```

---

## JavaScript Comparison (async/await)

You mentioned `await` — let's compare:
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   JAVASCRIPT                       GO                           │
│   ══════════                       ══                           │
│                                                                 │
│   async function fetchData() {     func fetchData(ch chan str) {│
│       // async work                    // work                  │
│       return "data"                    ch <- "data"             │
│   }                                }                            │
│                                                                 │
│   async function main() {          func main() {                │
│       const result = await fetch()     ch := make(chan string)  │
│       //             ═════             go fetchData(ch)         │
│       //             waits here        result := <-ch           │
│       console.log(result)              //         ═══           │
│   }                                    //         waits here    │
│                                        fmt.Println(result)      │
│                                    }                            │
│                                                                 │
│   await = <-ch (both wait for result)                          │
│   async = go   (both run concurrently)                         │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘




