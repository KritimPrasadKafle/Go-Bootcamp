// What is Password Hashing?
// Password hashing is a one-way transformation that converts a password into a fixed-length string of characters. Unlike encryption, hashing cannot be reversed.

┌─────────────────────────────────────────────────────────────────────────────┐
│                     HASHING vs ENCRYPTION                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ENCRYPTION (Two-way):                                                      │
│  ─────────────────────                                                      │
│  "password" ──encrypt──► "X8f#kL9..." ──decrypt──► "password"               │
│              (with key)                 (with key)                          │
│                                                                             │
│  HASHING (One-way):                                                         │
│  ──────────────────                                                         │
│  "password" ──hash──► "5e884898da28..." ──???──► IMPOSSIBLE!                │
│                       (always same output)                                  │
│                                                                             │
│  ✓ Same input ALWAYS produces same hash                                     │
│  ✓ Cannot reverse hash to get original                                      │
│  ✓ Small change in input = completely different hash                        │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│                     THE PROBLEM WITHOUT SALT                                │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  User A: password = "password123" → hash = "ef92b778..."                    │
│  User B: password = "password123" → hash = "ef92b778..."                    │
│  User C: password = "password123" → hash = "ef92b778..."                    │
│                                                                             │
│  ❌ PROBLEM: Same password = Same hash                                      │
│                                                                             │
│  Attacker can:                                                              │
│  • Pre-compute hashes for common passwords (Rainbow Tables)                 │
│  • If they crack one, they crack ALL users with same password               │
│  • Easily identify users with same passwords                                │
│                                                                             │
├─────────────────────────────────────────────────────────────────────────────┤
│                     THE SOLUTION WITH SALT                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Salt = Random bytes added to password before hashing                       │
│                                                                             │
│  User A: password + salt_A → "password123" + "x9Kj2..." → hash = "a1b2c3..."│
│  User B: password + salt_B → "password123" + "m4Np7..." → hash = "d4e5f6..."│
│  User C: password + salt_C → "password123" + "qR8sT..." → hash = "g7h8i9..."│
│                                                                             │
│  ✓ Same password = Different hashes (because different salts)               │
│  ✓ Rainbow tables become useless                                            │
│  ✓ Each password must be cracked individually                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘


┌─────────────────────────────────────────────────────────────────────────────┐
│                     PASSWORD HASHING FLOW                                   │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  SIGN UP (Registration):                                                    │
│  ────────────────────────                                                   │
│                                                                             │
│  1. User enters password: "password123"                                     │
│  2. Generate random salt: [16 random bytes]                                 │
│  3. Combine: salt + password                                                │
│  4. Hash the combination: SHA-256(salt + password)                          │
│  5. Store in database: salt (base64) + hash (base64)                        │
│                                                                             │
│  LOGIN (Verification):                                                      │
│  ─────────────────────                                                      │
│                                                                             │
│  1. User enters password: "password123"                                     │
│  2. Retrieve stored salt from database                                      │
│  3. Decode salt from base64                                                 │
│  4. Hash: SHA-256(retrieved_salt + entered_password)                        │
│  5. Compare with stored hash                                                │
│  6. If match → Login success!                                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

package main

import (
    "crypto/rand"      // Cryptographically secure random numbers
    "crypto/sha256"    // SHA-256 hashing algorithm
    "encoding/base64"  // For storing binary as text
    "fmt"
    "io"
)

func main() {
    password := "password123"
    
    // ═══════════════════════════════════════════════════════════════════
    // STEP 1: SIGN UP - Generate Salt
    // ═══════════════════════════════════════════════════════════════════
    
    salt, err := generateSalt()  // Generate 16 random bytes
    if err != nil {
        fmt.Println("Error generating salt:", err)
        return
    }
    
    fmt.Println("Original Salt:", salt)        // Raw bytes: [72 156 43 ...]
    fmt.Printf("Original Salt: %x\n", salt)    // Hex: 489c2b...
    
    // ═══════════════════════════════════════════════════════════════════
    // STEP 2: SIGN UP - Hash Password with Salt
    // ═══════════════════════════════════════════════════════════════════
    
    signUpHash := hashPassword(password, salt)
    
    // ═══════════════════════════════════════════════════════════════════
    // STEP 3: SIGN UP - Store in Database (simulated)
    // ═══════════════════════════════════════════════════════════════════
    
    // Convert salt to Base64 for storage (binary → text)
    saltStr := base64.StdEncoding.EncodeToString(salt)
    
    fmt.Println("Salt:", saltStr)      // Store this in database
    fmt.Println("Hash:", signUpHash)   // Store this in database
    
    // Show difference: hash without salt vs with salt
    hashOriginalPassword := sha256.Sum256([]byte(password))
    fmt.Println("Hash WITHOUT salt:", base64.StdEncoding.EncodeToString(hashOriginalPassword[:]))
    // Note: signUpHash is different because it includes salt!
    
    // ═══════════════════════════════════════════════════════════════════
    // STEP 4: LOGIN - Retrieve Salt from Database
    // ═══════════════════════════════════════════════════════════════════
    
    decodedSalt, err := base64.StdEncoding.DecodeString(saltStr)
    if err != nil {
        fmt.Println("Unable to decode salt:", err)
        return
    }
    
    // ═══════════════════════════════════════════════════════════════════
    // STEP 5: LOGIN - Hash Entered Password with Retrieved Salt
    // ═══════════════════════════════════════════════════════════════════
    
    // User entered "password124" (wrong password - note the 4 at the end)
    loginHash := hashPassword("password124", decodedSalt)
    
    // ═══════════════════════════════════════════════════════════════════
    // STEP 6: LOGIN - Compare Hashes
    // ═══════════════════════════════════════════════════════════════════
    
    if signUpHash == loginHash {
        fmt.Println("Password is correct. You are logged in.")
    } else {
        fmt.Println("Login failed. Please check user credentials.")
    }
    // Output: "Login failed" because "password124" ≠ "password123"
}

// generateSalt creates 16 cryptographically secure random bytes
func generateSalt() ([]byte, error) {
    salt := make([]byte, 16)                    // Create 16-byte slice
    _, err := io.ReadFull(rand.Reader, salt)    // Fill with random bytes
    if err != nil {
        return nil, err
    }
    return salt, nil
}

// hashPassword combines salt + password and hashes with SHA-256
func hashPassword(password string, salt []byte) string {
    // Combine salt and password
    saltedPassword := append(salt, []byte(password)...)
    
    // Hash the combination
    hash := sha256.Sum256(saltedPassword)
    
    // Return as Base64 string (for storage)
    return base64.StdEncoding.EncodeToString(hash[:])
}
```

---

## Visual Explanation
```
┌─────────────────────────────────────────────────────────────────────────────┐
│                     SIGN UP PROCESS                                         │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  Password: "password123"                                                    │
│                │                                                            │
│                ▼                                                            │
│  ┌────────────────────────┐                                                │
│  │   Generate Salt        │                                                │
│  │   (16 random bytes)    │                                                │
│  └────────────────────────┘                                                │
│                │                                                            │
│                ▼                                                            │
│  Salt: [0x48, 0x9c, 0x2b, 0x7f, 0x3a, 0xe1, ...]                           │
│                │                                                            │
│                ▼                                                            │
│  ┌────────────────────────┐                                                │
│  │  Combine: salt +       │                                                │
│  │  password bytes        │                                                │
│  └────────────────────────┘                                                │
│                │                                                            │
│                ▼                                                            │
│  Combined: [0x48, 0x9c, ..., 0x70, 0x61, 0x73, 0x73, ...]                  │
│            ├─── salt ────┤  ├──── "password123" ────────┤                  │
│                │                                                            │
│                ▼                                                            │
│  ┌────────────────────────┐                                                │
│  │   SHA-256 Hash         │                                                │
│  └────────────────────────┘                                                │
│                │                                                            │
│                ▼                                                            │
│  Hash: [0xa1, 0xb2, 0xc3, ...] (32 bytes)                                  │
│                │                                                            │
│                ▼                                                            │
│  ┌────────────────────────┐                                                │
│  │   Base64 Encode        │                                                │
│  └────────────────────────┘                                                │
│                │                                                            │
│                ▼                                                            │
│  ┌─────────────────────────────────────────────┐                           │
│  │  DATABASE                                    │                           │
│  │  ─────────                                   │                           │
│  │  Salt: "SJwrfzrh..." (Base64)               │                           │
│  │  Hash: "obLD8mK..." (Base64)                │                           │
│  └─────────────────────────────────────────────┘                           │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘

┌─────────────────────────────────────────────────────────────────────────────┐
│                     LOGIN VERIFICATION                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  User enters: "password123"                                                 │
│                │                                                            │
│                ▼                                                            │
│  ┌─────────────────────────────────────────────┐                           │
│  │  DATABASE                                    │                           │
│  │  ─────────                                   │                           │
│  │  Retrieve Salt: "SJwrfzrh..."               │                           │
│  │  Retrieve Hash: "obLD8mK..." (stored)       │                           │
│  └─────────────────────────────────────────────┘                           │
│                │                                                            │
│                ▼                                                            │
│  ┌────────────────────────┐                                                │
│  │   Decode Salt          │                                                │
│  │   (Base64 → bytes)     │                                                │
│  └────────────────────────┘                                                │
│                │                                                            │
│                ▼                                                            │
│  ┌────────────────────────┐                                                │
│  │  Hash entered password │                                                │
│  │  with retrieved salt   │                                                │
│  └────────────────────────┘                                                │
│                │                                                            │
│                ▼                                                            │
│  New Hash: "obLD8mK..." (computed)                                         │
│                │                                                            │
│                ▼                                                            │
│  ┌────────────────────────┐                                                │
│  │  Compare:              │                                                │
│  │  stored == computed?   │                                                │
│  └────────────────────────┘                                                │
│                │                                                            │
│       ┌────────┴────────┐                                                  │
│       ▼                 ▼                                                  │
│    MATCH             NO MATCH                                              │
│  ✓ Login OK        ✗ Login Failed                                         │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘


package main

import (
    "crypto/rand"
    "crypto/sha256"
    "crypto/subtle"
    "encoding/base64"
    "errors"
    "fmt"
    "io"
)

// User represents a user in the database
type User struct {
    Username     string
    PasswordHash string  // Base64 encoded hash
    Salt         string  // Base64 encoded salt
}

// Simulated database
var userDatabase = make(map[string]*User)

func main() {
    fmt.Println("=== PASSWORD AUTHENTICATION SYSTEM ===")
    fmt.Println()
    
    // ═══════════════════════════════════════════════════════════════════
    // REGISTRATION
    // ═══════════════════════════════════════════════════════════════════
    
    fmt.Println("1. REGISTRATION")
    fmt.Println("───────────────")
    
    err := registerUser("john_doe", "MySecurePassword123!")
    if err != nil {
        fmt.Println("Registration failed:", err)
        return
    }
    fmt.Println("User 'john_doe' registered successfully!")
    
    err = registerUser("jane_doe", "MySecurePassword123!")  // Same password!
    if err != nil {
        fmt.Println("Registration failed:", err)
        return
    }
    fmt.Println("User 'jane_doe' registered successfully!")
    
    // Show that same password = different hash
    fmt.Println("\n2. DIFFERENT HASHES FOR SAME PASSWORD")
    fmt.Println("─────────────────────────────────────")
    fmt.Printf("john_doe hash: %s...\n", userDatabase["john_doe"].PasswordHash[:20])
    fmt.Printf("jane_doe hash: %s...\n", userDatabase["jane_doe"].PasswordHash[:20])
    fmt.Println("(Different hashes because of different salts!)")
    
    // ═══════════════════════════════════════════════════════════════════
    // LOGIN ATTEMPTS
    // ═══════════════════════════════════════════════════════════════════
    
    fmt.Println("\n3. LOGIN ATTEMPTS")
    fmt.Println("─────────────────")
    
    // Correct password
    success, err := loginUser("john_doe", "MySecurePassword123!")
    printLoginResult("john_doe", "MySecurePassword123!", success, err)
    
    // Wrong password
    success, err = loginUser("john_doe", "WrongPassword!")
    printLoginResult("john_doe", "WrongPassword!", success, err)
    
    // Wrong username
    success, err = loginUser("unknown_user", "MySecurePassword123!")
    printLoginResult("unknown_user", "MySecurePassword123!", success, err)
    
    // Case-sensitive password
    success, err = loginUser("john_doe", "mysecurepassword123!")
    printLoginResult("john_doe", "mysecurepassword123!", success, err)
}

func registerUser(username, password string) error {
    // Check if user exists
    if _, exists := userDatabase[username]; exists {
        return errors.New("username already exists")
    }
    
    // Validate password
    if len(password) < 8 {
        return errors.New("password must be at least 8 characters")
    }
    
    // Generate salt
    salt, err := generateSalt(16)
    if err != nil {
        return fmt.Errorf("failed to generate salt: %v", err)
    }
    
    // Hash password
    hash := hashPasswordWithSalt(password, salt)
    
    // Store user
    userDatabase[username] = &User{
        Username:     username,
        PasswordHash: base64.StdEncoding.EncodeToString(hash),
        Salt:         base64.StdEncoding.EncodeToString(salt),
    }
    
    return nil
}

func loginUser(username, password string) (bool, error) {
    // Find user
    user, exists := userDatabase[username]
    if !exists {
        return false, errors.New("user not found")
    }
    
    // Decode stored salt
    salt, err := base64.StdEncoding.DecodeString(user.Salt)
    if err != nil {
        return false, fmt.Errorf("invalid salt: %v", err)
    }
    
    // Hash provided password with stored salt
    computedHash := hashPasswordWithSalt(password, salt)
    
    // Decode stored hash
    storedHash, err := base64.StdEncoding.DecodeString(user.PasswordHash)
    if err != nil {
        return false, fmt.Errorf("invalid stored hash: %v", err)
    }
    
    // Compare hashes (constant-time to prevent timing attacks)
    if subtle.ConstantTimeCompare(computedHash, storedHash) == 1 {
        return true, nil
    }
    
    return false, nil
}

func generateSalt(length int) ([]byte, error) {
    salt := make([]byte, length)
    _, err := io.ReadFull(rand.Reader, salt)
    return salt, err
}

func hashPasswordWithSalt(password string, salt []byte) []byte {
    combined := append(salt, []byte(password)...)
    hash := sha256.Sum256(combined)
    return hash[:]
}

func printLoginResult(username, password string, success bool, err error) {
    maskedPassword := password[:3] + "***"
    if err != nil {
        fmt.Printf("Login(%s, %s): ❌ Error - %v\n", username, maskedPassword, err)
    } else if success {
        fmt.Printf("Login(%s, %s): ✓ Success!\n", username, maskedPassword)
    } else {
        fmt.Printf("Login(%s, %s): ❌ Invalid password\n", username, maskedPassword)
    }
}