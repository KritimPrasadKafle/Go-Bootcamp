package main



numStr := "12345"
num, err := strconv.Atoi(numStr)
if err != nil {
    fmt.Println("Error parsing the value:", err)
}
fmt.Println("Parsed Integer:", num)    // 12345
fmt.Println("Parsed Integer:", num+1)  // 12346

// 2. ParseInt with base and bit size
// ParseInt(string, base, bitSize)
// base 10 = decimal, bitSize 32 = int32
numistr, err := strconv.ParseInt(numStr, 10, 32)
fmt.Println("Parsed Integer:", numistr)  // 12345

// 3. ParseFloat
floatstr := "3.14"
floatval, err := strconv.ParseFloat(floatstr, 64)  // 64-bit float
fmt.Printf("Parsed float: %.2f\n", floatval)  // 3.14

// 4. Binary to Decimal (base 2)
binaryStr := "1010"  // 0 + 2 + 0 + 8 = 10
decimal, err := strconv.ParseInt(binaryStr, 2, 64)
fmt.Println("Parsed binary to decimal:", decimal)  // 10

// 5. Hexadecimal to Decimal (base 16)
hexStr := "FF"
hex, err := strconv.ParseInt(hexStr, 16, 64)
fmt.Println("Parsed hex to decimal:", hex)  // 255

// 6. Error Handling for Invalid Input
invalidNum := "456abc"
invalidParse, err := strconv.Atoi(invalidNum)
if err != nil {
    fmt.Println("Error parsing value:", err)  // Error!
    return
}





import (
    "fmt"
    "strconv"
)

func main() {
    fmt.Println("=== STRING TO NUMBER ===")
    fmt.Println()
    
    // Atoi (ASCII to Integer) - simplest way
    s1 := "42"
    n1, err := strconv.Atoi(s1)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("Atoi(\"%s\") = %d (type: %T)\n", s1, n1, n1)
    }
    
    // ParseInt - more control
    // ParseInt(s string, base int, bitSize int) (int64, error)
    s2 := "12345"
    n2, err := strconv.ParseInt(s2, 10, 64)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("ParseInt(\"%s\", 10, 64) = %d (type: %T)\n", s2, n2, n2)
    }
    
    // ParseUint - unsigned integers
    s3 := "65535"
    n3, err := strconv.ParseUint(s3, 10, 16)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("ParseUint(\"%s\", 10, 16) = %d (type: %T)\n", s3, n3, n3)
    }
    
    // ParseFloat
    s4 := "3.14159"
    n4, err := strconv.ParseFloat(s4, 64)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Printf("ParseFloat(\"%s\", 64) = %f (type: %T)\n", s4, n4, n4)
    }
    
    // ParseBool
    boolStrings := []string{"true", "false", "1", "0", "T", "F", "TRUE", "FALSE"}
    fmt.Println("\nParseBool examples:")
    for _, s := range boolStrings {
        b, err := strconv.ParseBool(s)
        if err != nil {
            fmt.Printf("  ParseBool(\"%s\") = ERROR\n", s)
        } else {
            fmt.Printf("  ParseBool(\"%s\") = %v\n", s, b)
        }
    }
}
```

**Output:**
```
=== STRING TO NUMBER ===

Atoi("42") = 42 (type: int)
ParseInt("12345", 10, 64) = 12345 (type: int64)
ParseUint("65535", 10, 16) = 65535 (type: uint64)
ParseFloat("3.14159", 64) = 3.141590 (type: float64)

ParseBool examples:
  ParseBool("true") = true
  ParseBool("false") = false
  ParseBool("1") = true
  ParseBool("0") = false
  ParseBool("T") = true
  ParseBool("F") = false
  ParseBool("TRUE") = true
  ParseBool("FALSE") = false



  func main() {
    fmt.Println("=== ParseInt BASES ===")
    fmt.Println()
    
    // Base 2 (Binary)
    binary := "11111111"
    n, _ := strconv.ParseInt(binary, 2, 64)
    fmt.Printf("Binary  \"%s\" = %d\n", binary, n)
    
    // Base 8 (Octal)
    octal := "377"
    n, _ = strconv.ParseInt(octal, 8, 64)
    fmt.Printf("Octal   \"%s\" = %d\n", octal, n)
    
    // Base 10 (Decimal)
    decimal := "255"
    n, _ = strconv.ParseInt(decimal, 10, 64)
    fmt.Printf("Decimal \"%s\" = %d\n", decimal, n)
    
    // Base 16 (Hexadecimal)
    hex := "FF"
    n, _ = strconv.ParseInt(hex, 16, 64)
    fmt.Printf("Hex     \"%s\" = %d\n", hex, n)
    
    // Base 0 (Auto-detect from prefix)
    fmt.Println("\n=== AUTO-DETECT BASE (base=0) ===")
    autoDetect := []string{"42", "0b101010", "052", "0o52", "0x2A", "0X2a"}
    
    for _, s := range autoDetect {
        n, _ := strconv.ParseInt(s, 0, 64)
        fmt.Printf("ParseInt(\"%s\", 0, 64) = %d\n", s, n)
    }
    
    // BitSize affects range checking
    fmt.Println("\n=== BIT SIZE ===")
    
    value := "200"
    
    // 8-bit signed: -128 to 127
    n8, err := strconv.ParseInt(value, 10, 8)
    if err != nil {
        fmt.Printf("ParseInt(\"%s\", 10, 8) = ERROR (out of int8 range)\n", value)
    } else {
        fmt.Printf("ParseInt(\"%s\", 10, 8) = %d\n", value, n8)
    }
    
    // 16-bit signed: -32768 to 32767
    n16, _ := strconv.ParseInt(value, 10, 16)
    fmt.Printf("ParseInt(\"%s\", 10, 16) = %d\n", value, n16)
    
    // 32-bit signed: -2147483648 to 2147483647
    n32, _ := strconv.ParseInt(value, 10, 32)
    fmt.Printf("ParseInt(\"%s\", 10, 32) = %d\n", value, n32)
    
    // 64-bit signed: full range
    n64, _ := strconv.ParseInt(value, 10, 64)
    fmt.Printf("ParseInt(\"%s\", 10, 64) = %d\n", value, n64)
}
```

**Output:**
```
=== ParseInt BASES ===

Binary  "11111111" = 255
Octal   "377" = 255
Decimal "255" = 255
Hex     "FF" = 255

=== AUTO-DETECT BASE (base=0) ===
ParseInt("42", 0, 64) = 42
ParseInt("0b101010", 0, 64) = 42
ParseInt("052", 0, 64) = 42
ParseInt("0o52", 0, 64) = 42
ParseInt("0x2A", 0, 64) = 42
ParseInt("0X2a", 0, 64) = 42

=== BIT SIZE ===
ParseInt("200", 10, 8) = ERROR (out of int8 range)
ParseInt("200", 10, 16) = 200
ParseInt("200", 10, 32) = 200
ParseInt("200", 10, 64) = 200



func main() {
    fmt.Println("=== ParseFloat ===")
    fmt.Println()
    
    floatStrings := []string{
        "3.14159",
        "-273.15",
        "1.23e10",       // Scientific notation
        "1.23E-5",       // Scientific notation
        "+Inf",          // Positive infinity
        "-Inf",          // Negative infinity
        "NaN",           // Not a Number
        ".5",            // Leading decimal
        "5.",            // Trailing decimal
    }
    
    for _, s := range floatStrings {
        f, err := strconv.ParseFloat(s, 64)
        if err != nil {
            fmt.Printf("ParseFloat(\"%s\") = ERROR: %v\n", s, err)
        } else {
            fmt.Printf("ParseFloat(\"%s\") = %v\n", s, f)
        }
    }
    
    // 32-bit vs 64-bit precision
    fmt.Println("\n=== FLOAT PRECISION ===")
    
    preciseStr := "3.141592653589793238"
    
    f32, _ := strconv.ParseFloat(preciseStr, 32)
    f64, _ := strconv.ParseFloat(preciseStr, 64)
    
    fmt.Printf("Original:  %s\n", preciseStr)
    fmt.Printf("float32:   %.20f\n", f32)
    fmt.Printf("float64:   %.20f\n", f64)
}