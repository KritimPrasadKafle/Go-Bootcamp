┌─────────────────────────────────────────────────────────────────────────────┐
│                        WHAT IS BASE64?                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Binary Data (bytes)        Base64 Text (ASCII)                             │
│  ─────────────────────      ─────────────────────                           │
│  [0x48 0x65 0x6C 0x6C]  →   "SGVsbG8="                                      │
│  (Hello)                    (safe text)                                     │
│                                                                             │
│  Base64 Alphabet (64 characters):                                           │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ A B C D E F G H I J K L M N O P Q R S T U V W X Y Z                 │   │
│  │ a b c d e f g h i j k l m n o p q r s t u v w x y z                 │   │
│  │ 0 1 2 3 4 5 6 7 8 9                                                 │   │
│  │ + /                                                                  │   │
│  │ = (padding)                                                          │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
│  URL-Safe Alphabet uses: - _ instead of + /                                 │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│                     WHY BASE64 IS NEEDED                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ❌ Binary data can contain ANY byte (0-255)                                │
│     • Control characters (break systems)                                    │
│     • Null bytes (terminate strings)                                        │
│     • Non-printable characters                                              │
│     • Characters with special meaning                                       │
│                                                                             │
│  ❌ Many systems only handle text safely:                                   │
│     • Email (SMTP protocol)                                                 │
│     • JSON (text-based format)                                              │
│     • XML (text-based format)                                               │
│     • URLs (limited character set)                                          │
│     • HTML attributes                                                       │
│     • Configuration files                                                   │
│                                                                             │
│  ✅ Base64 converts binary → safe ASCII text                                │
│     • Only uses 64 printable characters                                     │
│     • Works everywhere text works                                           │
│     • Reversible (can decode back)                                          │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────────────────────┐
│                     BASE64 USE CASES                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  📧 Email Attachments                                                       │
│     Images, PDFs, documents sent via SMTP                                   │
│                                                                             │
│  🔐 Authentication                                                          │
│     HTTP Basic Auth: "username:password" → Base64                           │
│     JWT tokens contain Base64 encoded parts                                 │
│                                                                             │
│  🖼️ Data URLs                                                               │
│     Embed images in HTML/CSS: data:image/png;base64,iVBORw0K...             │
│                                                                             │
│  📡 API Communication                                                       │
│     Send binary data in JSON payloads                                       │
│                                                                             │
│  🔑 Cryptographic Keys                                                      │
│     Store/transmit encryption keys as text                                  │
│                                                                             │
│  📝 Configuration                                                           │
│     Store binary config values in text files                                │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────────────────────┐
│                     BASE64 ENCODING PROCESS                                 │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Step 1: Take 3 bytes (24 bits)                                             │
│  ────────────────────────────────                                           │
│  "Man" = M(77) a(97) n(110)                                                 │
│        = 01001101 01100001 01101110                                         │
│                                                                             │
│  Step 2: Split into 4 groups of 6 bits                                      │
│  ────────────────────────────────────                                       │
│  010011 | 010110 | 000101 | 101110                                          │
│    19       22       5       46                                             │
│                                                                             │
│  Step 3: Map each 6-bit value to Base64 character                           │
│  ─────────────────────────────────────────────                              │
│    19 = T                                                                   │
│    22 = W                                                                   │
│     5 = F                                                                   │
│    46 = u                                                                   │
│                                                                             │
│  Result: "Man" → "TWFu"                                                     │
│                                                                             │
│  ═══════════════════════════════════════════════════════════════════════   │
│                                                                             │
│  PADDING (when input isn't multiple of 3 bytes):                            │
│  ─────────────────────────────────────────────                              │
│  "Ma" (2 bytes) → "TWE=" (one = padding)                                    │
│  "M"  (1 byte)  → "TQ==" (two = padding)                                    │
│                                                                             │
│  Size increase: Base64 output is ~33% larger than input                     │
│  (3 bytes → 4 characters)                                                   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘


package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    // Original data as bytes
    data := []byte("He~lo, Base64 Encoding")
    
    // ENCODE: Convert bytes to Base64 string
    // StdEncoding uses standard Base64 alphabet (A-Z, a-z, 0-9, +, /)
    encoded := base64.StdEncoding.EncodeToString(data)
    fmt.Println(encoded)  // SGV+bG8sIEJhc2U2NCBFbmNvZGluZw==
    
    // DECODE: Convert Base64 string back to bytes
    decoded, err := base64.StdEncoding.DecodeString(encoded)
    if err != nil {
        fmt.Println("Error in decoding:", err)
        return
    }
    fmt.Println("Decoded:", string(decoded))  // He~lo, Base64 Encoding
    
    // URL-SAFE ENCODING: Uses - and _ instead of + and /
    // This is safe to use in URLs without additional escaping
    urlSafeEncoded := base64.URLEncoding.EncodeToString(data)
    fmt.Println("URL Safe encoded:", urlSafeEncoded)
}


package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    data := []byte("Hello+World/Test")
    
    fmt.Println("=== BASE64 ENCODING VARIANTS ===")
    fmt.Println()
    fmt.Printf("Original: %s\n\n", data)
    
    // 1. Standard Encoding (RFC 4648)
    // Alphabet: A-Z, a-z, 0-9, +, /
    // Padding: = (required)
    std := base64.StdEncoding.EncodeToString(data)
    fmt.Println("1. StdEncoding:")
    fmt.Printf("   Encoded: %s\n", std)
    fmt.Println("   Uses: + and /")
    fmt.Println("   Padding: Yes (=)")
    
    // 2. URL Encoding (RFC 4648)
    // Alphabet: A-Z, a-z, 0-9, -, _
    // Padding: = (required)
    urlEnc := base64.URLEncoding.EncodeToString(data)
    fmt.Println("\n2. URLEncoding:")
    fmt.Printf("   Encoded: %s\n", urlEnc)
    fmt.Println("   Uses: - and _ (URL safe)")
    fmt.Println("   Padding: Yes (=)")
    
    // 3. Raw Standard Encoding
    // Same as StdEncoding but WITHOUT padding
    rawStd := base64.RawStdEncoding.EncodeToString(data)
    fmt.Println("\n3. RawStdEncoding:")
    fmt.Printf("   Encoded: %s\n", rawStd)
    fmt.Println("   Uses: + and /")
    fmt.Println("   Padding: No")
    
    // 4. Raw URL Encoding
    // Same as URLEncoding but WITHOUT padding
    rawURL := base64.RawURLEncoding.EncodeToString(data)
    fmt.Println("\n4. RawURLEncoding:")
    fmt.Printf("   Encoded: %s\n", rawURL)
    fmt.Println("   Uses: - and _ (URL safe)")
    fmt.Println("   Padding: No")
    
    // Comparison
    fmt.Println("\n=== COMPARISON ===")
    fmt.Printf("StdEncoding:    %s\n", std)
    fmt.Printf("URLEncoding:    %s\n", urlEnc)
    fmt.Printf("RawStdEncoding: %s\n", rawStd)
    fmt.Printf("RawURLEncoding: %s\n", rawURL)
}
```

**Output:**
```
=== BASE64 ENCODING VARIANTS ===

Original: Hello+World/Test

1. StdEncoding:
   Encoded: SGVsbG8rV29ybGQvVGVzdA==
   Uses: + and /
   Padding: Yes (=)

2. URLEncoding:
   Encoded: SGVsbG8rV29ybGQvVGVzdA==
   Uses: - and _ (URL safe)
   Padding: Yes (=)

3. RawStdEncoding:
   Encoded: SGVsbG8rV29ybGQvVGVzdA
   Uses: + and /
   Padding: No

4. RawURLEncoding:
   Encoded: SGVsbG8rV29ybGQvVGVzdA
   Uses: - and _ (URL safe)
   Padding: No

=== COMPARISON ===
StdEncoding:    SGVsbG8rV29ybGQvVGVzdA==
URLEncoding:    SGVsbG8rV29ybGQvVGVzdA==
RawStdEncoding: SGVsbG8rV29ybGQvVGVzdA
RawURLEncoding: SGVsbG8rV29ybGQvVGVzdA




package main

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "io"
)

func main() {
    data := []byte("Hello, Base64 World!")
    
    fmt.Println("=== ENCODING METHODS ===")
    fmt.Println()
    
    // Method 1: EncodeToString (most common)
    fmt.Println("1. EncodeToString:")
    encoded := base64.StdEncoding.EncodeToString(data)
    fmt.Printf("   Result: %s\n", encoded)
    
    // Method 2: Encode to byte slice
    fmt.Println("\n2. Encode to []byte:")
    dst := make([]byte, base64.StdEncoding.EncodedLen(len(data)))
    base64.StdEncoding.Encode(dst, data)
    fmt.Printf("   Result: %s\n", dst)
    
    // Method 3: Using Encoder (streaming)
    fmt.Println("\n3. NewEncoder (streaming):")
    var buf bytes.Buffer
    encoder := base64.NewEncoder(base64.StdEncoding, &buf)
    encoder.Write(data)
    encoder.Close()  // Important: must close to flush
    fmt.Printf("   Result: %s\n", buf.String())
    
    fmt.Println("\n=== DECODING METHODS ===")
    fmt.Println()
    
    encodedStr := "SGVsbG8sIEJhc2U2NCBXb3JsZCE="
    
    // Method 1: DecodeString (most common)
    fmt.Println("1. DecodeString:")
    decoded, err := base64.StdEncoding.DecodeString(encodedStr)
    if err != nil {
        fmt.Println("   Error:", err)
    } else {
        fmt.Printf("   Result: %s\n", decoded)
    }
    
    // Method 2: Decode to byte slice
    fmt.Println("\n2. Decode to []byte:")
    dst2 := make([]byte, base64.StdEncoding.DecodedLen(len(encodedStr)))
    n, err := base64.StdEncoding.Decode(dst2, []byte(encodedStr))
    if err != nil {
        fmt.Println("   Error:", err)
    } else {
        fmt.Printf("   Result: %s\n", dst2[:n])
    }
    
    // Method 3: Using Decoder (streaming)
    fmt.Println("\n3. NewDecoder (streaming):")
    decoder := base64.NewDecoder(base64.StdEncoding, bytes.NewReader([]byte(encodedStr)))
    decodedData, err := io.ReadAll(decoder)
    if err != nil {
        fmt.Println("   Error:", err)
    } else {
        fmt.Printf("   Result: %s\n", decodedData)
    }
}

package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    fmt.Println("=== ENCODING BINARY DATA ===")
    fmt.Println()
    
    // Binary data (not valid text)
    binaryData := []byte{0x00, 0x01, 0xFF, 0xFE, 0x89, 0x50, 0x4E, 0x47}
    
    fmt.Println("1. Raw binary bytes:")
    fmt.Printf("   Hex: %X\n", binaryData)
    fmt.Printf("   As string (corrupted): %q\n", string(binaryData))
    
    // Encode to Base64
    encoded := base64.StdEncoding.EncodeToString(binaryData)
    fmt.Println("\n2. Base64 encoded:")
    fmt.Printf("   %s\n", encoded)
    fmt.Println("   (Safe to transmit as text)")
    
    // Decode back
    decoded, _ := base64.StdEncoding.DecodeString(encoded)
    fmt.Println("\n3. Decoded back:")
    fmt.Printf("   Hex: %X\n", decoded)
    fmt.Printf("   Match: %v\n", bytesEqual(binaryData, decoded))
    
    // Simulating image data (PNG header)
    fmt.Println("\n=== SIMULATED IMAGE DATA ===")
    
    // PNG file signature
    pngHeader := []byte{0x89, 0x50, 0x4E, 0x47, 0x0D, 0x0A, 0x1A, 0x0A}
    
    fmt.Printf("PNG Header (hex): %X\n", pngHeader)
    encoded = base64.StdEncoding.EncodeToString(pngHeader)
    fmt.Printf("PNG Header (base64): %s\n", encoded)
}

func bytesEqual(a, b []byte) bool {
    if len(a) != len(b) {
        return false
    }
    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }
    return true
}


package main

import (
    "encoding/base64"
    "fmt"
    "strings"
)

func main() {
    fmt.Println("=== HTTP BASIC AUTHENTICATION ===")
    fmt.Println()
    
    // Basic Auth: username:password → Base64
    username := "admin"
    password := "secret123"
    
    // Create credentials string
    credentials := username + ":" + password
    fmt.Printf("Credentials: %s\n", credentials)
    
    // Encode to Base64
    encoded := base64.StdEncoding.EncodeToString([]byte(credentials))
    fmt.Printf("Base64: %s\n", encoded)
    
    // Create Authorization header
    authHeader := "Basic " + encoded
    fmt.Printf("Header: Authorization: %s\n", authHeader)
    
    // Decoding (server side)
    fmt.Println("\n=== DECODING ON SERVER ===")
    
    // Parse header
    parts := strings.SplitN(authHeader, " ", 2)
    if len(parts) != 2 || parts[0] != "Basic" {
        fmt.Println("Invalid header format")
        return
    }
    
    // Decode Base64
    decoded, err := base64.StdEncoding.DecodeString(parts[1])
    if err != nil {
        fmt.Println("Decode error:", err)
        return
    }
    
    // Split credentials
    credParts := strings.SplitN(string(decoded), ":", 2)
    if len(credParts) != 2 {
        fmt.Println("Invalid credentials format")
        return
    }
    
    fmt.Printf("Username: %s\n", credParts[0])
    fmt.Printf("Password: %s\n", credParts[1])
}
```

**Output:**
```
=== HTTP BASIC AUTHENTICATION ===

Credentials: admin:secret123
Base64: YWRtaW46c2VjcmV0MTIz
Header: Authorization: Basic YWRtaW46c2VjcmV0MTIz

=== DECODING ON SERVER ===
Username: admin
Password: secret123