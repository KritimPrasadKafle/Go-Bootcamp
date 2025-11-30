package main

import "fmt"

func main() {
    var a int = 32
    
    // int → int32
    b := int32(a)
    
    // int → int64
    c := int64(a)
    
    // int → int8 (⚠️ be careful with overflow!)
    d := int8(a)
    
    fmt.Printf("int:   %d (type: %T)\n", a, a)
    fmt.Printf("int32: %d (type: %T)\n", b, b)
    fmt.Printf("int64: %d (type: %T)\n", c, c)
    fmt.Printf("int8:  %d (type: %T)\n", d, d)
}
```

**Output:**
```
int:   32 (type: int)
int32: 32 (type: int32)
int64: 32 (type: int64)
int8:  32 (type: int8)


package main

import "fmt"

func main() {
    var big int = 300
    
    // int8 can only hold -128 to 127
    small := int8(big)  // OVERFLOW!
    
    fmt.Printf("Original: %d\n", big)
    fmt.Printf("As int8:  %d (WRONG!)\n", small)
    
    // What happened:
    // 300 in binary: 100101100
    // int8 takes only last 8 bits: 00101100 = 44
}
```

**Output:**
```
Original: 300
As int8:  44 (WRONG!)
```
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   INTEGER TYPE RANGES                                           │
│                                                                 │
│   Type     │ Min                  │ Max                         │
│   ─────────┼──────────────────────┼─────────────────────────    │
│   int8     │ -128                 │ 127                         │
│   int16    │ -32,768              │ 32,767                      │
│   int32    │ -2,147,483,648       │ 2,147,483,647               │
│   int64    │ -9 quintillion       │ 9 quintillion               │
│   uint8    │ 0                    │ 255                         │
│   uint16   │ 0                    │ 65,535                      │
│   uint32   │ 0                    │ 4,294,967,295               │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘



package main

import "fmt"

func main() {
    var a int = 42
    
    // int → float64 (safe, no data loss)
    b := float64(a)
    
    // int → float32
    c := float32(a)
    
    fmt.Printf("int:     %d\n", a)
    fmt.Printf("float64: %f\n", b)
    fmt.Printf("float32: %f\n", c)
}
```

**Output:**
```
int:     42
float64: 42.000000
float32: 42.000000