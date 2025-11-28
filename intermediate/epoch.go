// package main
// func main(){

// }
// // Unix Epoch: 00:00:00 UTC on January 1, 1970
// // It's the "zero point" for Unix time

// now := time.Now()
// unixTime := now.Unix()  // Get Unix timestamp (seconds since epoch)
// fmt.Println("Current Unix Time:", unixTime)  // e.g., 1722345678

// // Convert Unix timestamp back to time.Time
// t := time.Unix(unixTime, 0)  // Parameters: (seconds, nanoseconds)
// fmt.Println(t)  // 2024-07-30 15:04:05 +0530 IST

// fmt.Println("Time:", t.Format("2006-01-02"))  // 2024-07-30

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     now := time.Now()

//     fmt.Println("=== UNIX TIME FORMATS ===")
//     fmt.Println()

//     // Seconds (most common)
//     unixSec := now.Unix()
//     fmt.Printf("Unix (seconds):      %d\n", unixSec)
//     fmt.Printf("                     %d seconds since Jan 1, 1970\n", unixSec)

//     // Milliseconds (JavaScript uses this)
//     unixMilli := now.UnixMilli()
//     fmt.Printf("\nUnix (milliseconds): %d\n", unixMilli)
//     fmt.Printf("                     Used by JavaScript Date.now()\n")

//     // Microseconds
//     unixMicro := now.UnixMicro()
//     fmt.Printf("\nUnix (microseconds): %d\n", unixMicro)

//     // Nanoseconds (highest precision)
//     unixNano := now.UnixNano()
//     fmt.Printf("\nUnix (nanoseconds):  %d\n", unixNano)

//     // Relationship
//     fmt.Println("\n=== RELATIONSHIP ===")
//     fmt.Printf("1 second      = 1,000 milliseconds\n")
//     fmt.Printf("1 second      = 1,000,000 microseconds\n")
//     fmt.Printf("1 second      = 1,000,000,000 nanoseconds\n")
//     fmt.Printf("\nMillis / 1000 = Seconds: %d\n", unixMilli/1000)
// }
// ```

// **Output:**
// ```
// === UNIX TIME FORMATS ===

// Unix (seconds):      1722345678
//                      1722345678 seconds since Jan 1, 1970

// Unix (milliseconds): 1722345678123
//                      Used by JavaScript Date.now()

// Unix (microseconds): 1722345678123456

// Unix (nanoseconds):  1722345678123456789

// === RELATIONSHIP ===
// 1 second      = 1,000 milliseconds
// 1 second      = 1,000,000 microseconds
// 1 second      = 1,000,000,000 nanoseconds

// Millis / 1000 = Seconds: 1722345678

// func main() {
//     fmt.Println("=== CONVERTING TO TIME ===")

//     // From seconds
//     unixSec := int64(1722345678)
//     t1 := time.Unix(unixSec, 0)
//     fmt.Println("From seconds:", t1)

//     // From seconds with nanoseconds
//     t2 := time.Unix(unixSec, 500000000)  // 0.5 seconds
//     fmt.Println("With nanoseconds:", t2)

//     // From milliseconds
//     unixMilli := int64(1722345678123)
//     t3 := time.UnixMilli(unixMilli)
//     fmt.Println("From milliseconds:", t3)

//     // From microseconds
//     unixMicro := int64(1722345678123456)
//     t4 := time.UnixMicro(unixMicro)
//     fmt.Println("From microseconds:", t4)

//     // From nanoseconds (manual conversion)
//     unixNano := int64(1722345678123456789)
//     t5 := time.Unix(0, unixNano)
//     fmt.Println("From nanoseconds:", t5)

//     fmt.Println("\n=== CONVERTING FROM TIME ===")

//     now := time.Now()
//     fmt.Printf("Current time: %s\n", now)
//     fmt.Printf("To seconds:      %d\n", now.Unix())
//     fmt.Printf("To milliseconds: %d\n", now.UnixMilli())
//     fmt.Printf("To microseconds: %d\n", now.UnixMicro())
//     fmt.Printf("To nanoseconds:  %d\n", now.UnixNano())
// }

// func main() {
//     // Unix timestamp is ALWAYS UTC
//     unixTime := int64(1722345678)

//     fmt.Println("=== UNIX TIME IS TIMEZONE-INDEPENDENT ===")
//     fmt.Printf("Unix timestamp: %d\n\n", unixTime)

//     // Convert to different timezones
//     utc := time.Unix(unixTime, 0).UTC()
//     fmt.Println("UTC:          ", utc)

//     // Different locations
//     locations := []string{
//         "America/New_York",
//         "America/Los_Angeles",
//         "Europe/London",
//         "Asia/Tokyo",
//         "Asia/Kolkata",
//     }

//     for _, locName := range locations {
//         loc, _ := time.LoadLocation(locName)
//         localTime := time.Unix(unixTime, 0).In(loc)
//         fmt.Printf("%-20s %s\n", locName+":", localTime.Format("2006-01-02 15:04:05 MST"))
//     }

//     fmt.Println("\n=== IMPORTANT ===")
//     fmt.Println("All these times represent the SAME instant!")
//     fmt.Println("Unix timestamp doesn't change based on timezone.")
// }

// func main() {
//     fmt.Println("=== NOTABLE UNIX TIMESTAMPS ===")
//     fmt.Println()

//     timestamps := []struct {
//         name  string
//         unix  int64
//     }{
//         {"Unix Epoch (start)", 0},
//         {"First Moon Landing (Apollo 11)", -14182940},  // July 20, 1969
//         {"Y2K (Year 2000)", 946684800},
//         {"Unix Billennium (1 billion)", 1000000000},
//         {"iPhone Launch (2007)", 1183248000},
//         {"Bitcoin Genesis Block", 1231006505},
//         {"Year 2038 Problem", 2147483647},  // Max 32-bit signed integer
//     }

//     for _, ts := range timestamps {
//         t := time.Unix(ts.unix, 0).UTC()
//         fmt.Printf("%-35s Unix: %-12d  Date: %s\n",
//             ts.name, ts.unix, t.Format("2006-01-02 15:04:05"))
//     }

//     fmt.Println("\n=== THE YEAR 2038 PROBLEM ===")
//     fmt.Println("32-bit systems store Unix time as signed 32-bit integer")
//     fmt.Println("Maximum value: 2,147,483,647 (2^31 - 1)")
//     fmt.Println("This overflows on: January 19, 2038, 03:14:07 UTC")
//     fmt.Println("Go uses 64-bit integers, so it's not affected!")

//     // Demonstrate the limit
//     max32bit := time.Unix(2147483647, 0).UTC()
//     fmt.Printf("\nMax 32-bit timestamp: %s\n", max32bit)

//     // Go can handle much larger timestamps
//     farFuture := time.Unix(4102444800, 0).UTC()  // Year 2100
//     fmt.Printf("Year 2100: %s\n", farFuture)
// }

// func main() {
//     fmt.Println("=== TIME DIFFERENCES WITH UNIX ===")
//     fmt.Println()

//     // Two events
//     event1Unix := int64(1722300000)  // Earlier
//     event2Unix := int64(1722345678)  // Later

//     event1 := time.Unix(event1Unix, 0)
//     event2 := time.Unix(event2Unix, 0)

//     fmt.Println("Event 1:", event1.Format("2006-01-02 15:04:05"))
//     fmt.Println("Event 2:", event2.Format("2006-01-02 15:04:05"))

//     // Calculate difference
//     diffSeconds := event2Unix - event1Unix
//     diffDuration := event2.Sub(event1)

//     fmt.Printf("\nDifference in seconds: %d\n", diffSeconds)
//     fmt.Printf("Difference as duration: %v\n", diffDuration)
//     fmt.Printf("Hours: %.2f\n", float64(diffSeconds)/3600)

//     // Age calculation using Unix
//     fmt.Println("\n=== AGE FROM UNIX TIMESTAMP ===")
//     birthdayUnix := int64(631152000)  // Jan 1, 1990
//     nowUnix := time.Now().Unix()

//     ageSeconds := nowUnix - birthdayUnix
//     ageYears := float64(ageSeconds) / (365.25 * 24 * 60 * 60)

//     fmt.Printf("Birth timestamp: %d\n", birthdayUnix)
//     fmt.Printf("Current timestamp: %d\n", nowUnix)
//     fmt.Printf("Age in seconds: %d\n", ageSeconds)
//     fmt.Printf("Age in years: %.2f\n", ageYears)
// }

// package main

// import (
//     "encoding/json"
//     "fmt"
//     "time"
// )

// // Common patterns for APIs

// // Pattern 1: Unix seconds
// type EventV1 struct {
//     Name      string `json:"name"`
//     Timestamp int64  `json:"timestamp"`  // Unix seconds
// }

// // Pattern 2: Unix milliseconds (JavaScript friendly)
// type EventV2 struct {
//     Name      string `json:"name"`
//     Timestamp int64  `json:"timestamp_ms"`  // Unix milliseconds
// }

// // Pattern 3: ISO 8601 string
// type EventV3 struct {
//     Name      string `json:"name"`
//     Timestamp string `json:"timestamp"`  // ISO 8601
// }

// func main() {
//     now := time.Now()

//     fmt.Println("=== API TIMESTAMP FORMATS ===")
//     fmt.Println()

//     // Unix seconds (common in many APIs)
//     eventV1 := EventV1{
//         Name:      "user_login",
//         Timestamp: now.Unix(),
//     }
//     jsonV1, _ := json.MarshalIndent(eventV1, "", "  ")
//     fmt.Println("Unix Seconds:")
//     fmt.Println(string(jsonV1))

//     // Unix milliseconds (JavaScript, Java)
//     eventV2 := EventV2{
//         Name:      "user_login",
//         Timestamp: now.UnixMilli(),
//     }
//     jsonV2, _ := json.MarshalIndent(eventV2, "", "  ")
//     fmt.Println("\nUnix Milliseconds:")
//     fmt.Println(string(jsonV2))

//     // ISO 8601 (most readable)
//     eventV3 := EventV3{
//         Name:      "user_login",
//         Timestamp: now.Format(time.RFC3339),
//     }
//     jsonV3, _ := json.MarshalIndent(eventV3, "", "  ")
//     fmt.Println("\nISO 8601:")
//     fmt.Println(string(jsonV3))

//     // Parsing back
//     fmt.Println("\n=== PARSING API RESPONSES ===")

//     // From Unix seconds
//     t1 := time.Unix(eventV1.Timestamp, 0)
//     fmt.Println("From seconds:", t1)

//     // From Unix milliseconds
//     t2 := time.UnixMilli(eventV2.Timestamp)
//     fmt.Println("From milliseconds:", t2)

//     // From ISO 8601
//     t3, _ := time.Parse(time.RFC3339, eventV3.Timestamp)
//     fmt.Println("From ISO 8601:", t3)
// }



