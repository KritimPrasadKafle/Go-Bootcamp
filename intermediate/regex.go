// package main

// import "fmt"

// func main() {

// }

// // REGULAR STRING - need to escape quotes
// fmt.Println("He said, \"I am great\"")  // Output: He said, "I am great"

// // RAW STRING - no escaping needed
// fmt.Println(`He said, "I am great"`)   // Output: He said, "I am great"

// // Compile regex pattern for email
// re := regexp.MustCompile(`[a-zA-Z0-9._+%-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)

// email1 := "user@email.com"
// email2 := "invalid_email"

// fmt.Println("Email1:", re.MatchString(email1))  // true
// fmt.Println("Email2:", re.MatchString(email2))  // false

// // Pattern with capturing groups (parentheses)
// re = regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

// date := "2024-07-30"

// // Extract all matches including full match and groups
// submatches := re.FindStringSubmatch(date)

// fmt.Println(submatches)     // [2024-07-30 2024 07 30]
// fmt.Println(submatches[0])  // 2024-07-30 (full match)
// fmt.Println(submatches[1])  // 2024 (first group - year)
// fmt.Println(submatches[2])  // 07 (second group - month)
// fmt.Println(submatches[3])  // 30 (third group - day)

// str := "Hello World"
// re = regexp.MustCompile(`[aeiou]`)  // Match vowels

// result := re.ReplaceAllString(str, "*")
// fmt.Println(result)  // Output: H*ll* W*rld

// package main

// import (
//     "fmt"
//     "regexp"
// )

// func main() {
//     // Match specific word
//     re := regexp.MustCompile(`hello`)
//     fmt.Println(re.MatchString("hello world"))     // true
//     fmt.Println(re.MatchString("Hello world"))     // false (case-sensitive)

//     // Case-insensitive matching
//     re = regexp.MustCompile(`(?i)hello`)
//     fmt.Println(re.MatchString("Hello world"))     // true
//     fmt.Println(re.MatchString("HELLO world"))     // true

//     // Match any digit
//     re = regexp.MustCompile(`\d`)
//     fmt.Println(re.MatchString("abc123"))          // true
//     fmt.Println(re.MatchString("abc"))             // false

//     // Match multiple digits
//     re = regexp.MustCompile(`\d+`)
//     fmt.Println(re.FindString("Price: $123"))      // "123"

//     // Match word characters (letters, digits, underscore)
//     re = regexp.MustCompile(`\w+`)
//     fmt.Println(re.FindString("hello_world123"))   // "hello_world123"

//     // Match whitespace
//     re = regexp.MustCompile(`\s+`)
//     fmt.Println(re.FindString("hello   world"))    // "   "
// }

// func main() {
//     text := "The price is $50 and quantity is 10"

//     // Match specific characters
//     re := regexp.MustCompile(`[aeiou]`)
//     fmt.Println(re.FindAllString(text, -1))
//     // Output: [e i e i a u a i i i]

//     // Match digits
//     re = regexp.MustCompile(`[0-9]+`)
//     fmt.Println(re.FindAllString(text, -1))
//     // Output: [50 10]

//     // Match letters (case insensitive range)
//     re = regexp.MustCompile(`[a-zA-Z]+`)
//     fmt.Println(re.FindAllString(text, -1))
//     // Output: [The price is and quantity is]

//     // Negated character class (NOT digits)
//     re = regexp.MustCompile(`[^0-9]+`)
//     fmt.Println(re.FindString("abc123def"))
//     // Output: abc

//     // Match special characters
//     re = regexp.MustCompile(`[$#@!]+`)
//     fmt.Println(re.FindString("Price: $50 #discount"))
//     // Output: $
// }

// func main() {
//     // * - 0 or more
//     re := regexp.MustCompile(`ab*c`)
//     fmt.Println(re.FindString("ac"))       // ac (0 b's)
//     fmt.Println(re.FindString("abc"))      // abc (1 b)
//     fmt.Println(re.FindString("abbc"))     // abbc (2 b's)

//     // + - 1 or more
//     re = regexp.MustCompile(`ab+c`)
//     fmt.Println(re.FindString("ac"))       // "" (no match, needs at least 1 b)
//     fmt.Println(re.FindString("abc"))      // abc
//     fmt.Println(re.FindString("abbbc"))    // abbbc

//     // ? - 0 or 1
//     re = regexp.MustCompile(`colou?r`)
//     fmt.Println(re.FindString("color"))    // color
//     fmt.Println(re.FindString("colour"))   // colour

//     // {n} - exactly n times
//     re = regexp.MustCompile(`\d{3}`)
//     fmt.Println(re.FindString("12345"))    // 123

//     // {n,} - n or more times
//     re = regexp.MustCompile(`\d{3,}`)
//     fmt.Println(re.FindString("12"))       // "" (needs at least 3)
//     fmt.Println(re.FindString("12345"))    // 12345

//     // {n,m} - between n and m times
//     re = regexp.MustCompile(`\d{2,4}`)
//     fmt.Println(re.FindString("1"))        // "" (too short)
//     fmt.Println(re.FindString("12"))       // 12
//     fmt.Println(re.FindString("1234"))     // 1234
//     fmt.Println(re.FindString("12345"))    // 1234 (only matches first 4)
// }

// func main() {
//     // ^ - Start of string
//     re := regexp.MustCompile(`^Hello`)
//     fmt.Println(re.MatchString("Hello World"))     // true
//     fmt.Println(re.MatchString("Say Hello"))       // false

//     // $ - End of string
//     re = regexp.MustCompile(`World$`)
//     fmt.Println(re.MatchString("Hello World"))     // true
//     fmt.Println(re.MatchString("World is big"))    // false

//     // Both start and end
//     re = regexp.MustCompile(`^Hello World$`)
//     fmt.Println(re.MatchString("Hello World"))     // true
//     fmt.Println(re.MatchString("Hello World!"))    // false

//     // \b - Word boundary
//     re = regexp.MustCompile(`\bcat\b`)
//     fmt.Println(re.MatchString("cat"))             // true
//     fmt.Println(re.MatchString("the cat sat"))     // true
//     fmt.Println(re.MatchString("category"))        // false
// }

// func main() {
//     // Basic capturing group
//     re := regexp.MustCompile(`(\w+)@(\w+)\.(\w+)`)
//     email := "user@example.com"

//     matches := re.FindStringSubmatch(email)
//     fmt.Println(matches)
//     // Output: [user@example.com user example com]
//     // Index 0: full match
//     // Index 1: first group (user)
//     // Index 2: second group (example)
//     // Index 3: third group (com)

//     // Named groups
//     re = regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+)\.(?P<tld>\w+)`)
//     matches = re.FindStringSubmatch(email)

//     // Get group names
//     names := re.SubexpNames()
//     result := make(map[string]string)
//     for i, name := range names {
//         if i > 0 && name != "" {
//             result[name] = matches[i]
//         }
//     }
//     fmt.Println(result)
//     // Output: map[domain:example tld:com user:user]

//     // Non-capturing group (?:...)
//     re = regexp.MustCompile(`(?:Mr|Mrs|Ms)\. (\w+)`)
//     name := "Mr. Smith"
//     matches = re.FindStringSubmatch(name)
//     fmt.Println(matches)
//     // Output: [Mr. Smith Smith]
//     // Only captures the name, not the title
// }

// func main() {
//     // US Phone number validation
//     patterns := []string{
//         `^\d{3}-\d{3}-\d{4}$`,           // 123-456-7890
//         `^\(\d{3}\)\s\d{3}-\d{4}$`,      // (123) 456-7890
//         `^\d{10}$`,                       // 1234567890
//     }

//     phones := []string{
//         "123-456-7890",
//         "(123) 456-7890",
//         "1234567890",
//         "12-345-6789",  // Invalid
//     }

//     for _, pattern := range patterns {
//         re := regexp.MustCompile(pattern)
//         fmt.Printf("Pattern: %s\n", pattern)
//         for _, phone := range phones {
//             fmt.Printf("  %s: %v\n", phone, re.MatchString(phone))
//         }
//         fmt.Println()
//     }

//     // Extract phone number parts
//     re := regexp.MustCompile(`\(?(\d{3})\)?[-.\s]?(\d{3})[-.\s]?(\d{4})`)

//     testPhones := []string{
//         "123-456-7890",
//         "(123) 456-7890",
//         "123.456.7890",
//         "1234567890",
//     }

//     for _, phone := range testPhones {
//         matches := re.FindStringSubmatch(phone)
//         if matches != nil {
//             fmt.Printf("Phone: %s -> Area: %s, Prefix: %s, Line: %s\n",
//                 matches[0], matches[1], matches[2], matches[3])
//         }
//     }
// }

// func main() {
//     // URL pattern
//     urlPattern := `^(https?://)?([\w\-]+\.)+[\w\-]+(:\d+)?(/[\w\-./]*)?(\?[\w=&]*)?$`
//     re := regexp.MustCompile(urlPattern)

//     urls := []string{
//         "https://example.com",
//         "http://sub.example.com:8080/path/to/page",
//         "example.com/page?id=123&name=test",
//         "invalid url with spaces",
//     }

//     for _, url := range urls {
//         fmt.Printf("%s: %v\n", url, re.MatchString(url))
//     }

//     // Extract URL components
//     re = regexp.MustCompile(`^(https?://)?([\w\-]+\.[\w\-]+)(:\d+)?(/.*)?$`)

//     url := "https://www.example.com:8080/path/to/page"
//     matches := re.FindStringSubmatch(url)

//     if matches != nil {
//         fmt.Printf("Full URL: %s\n", matches[0])
//         fmt.Printf("Protocol: %s\n", matches[1])
//         fmt.Printf("Domain: %s\n", matches[2])
//         fmt.Printf("Port: %s\n", matches[3])
//         fmt.Printf("Path: %s\n", matches[4])
//     }
// }

// func main() {
//     text := `
// Contact Information:
// Email: john@example.com
// Phone: 555-123-4567
// Website: https://example.com
// Price: $1,234.56
// Date: 2024-07-30
// `

//     // Extract emails
//     emailRe := regexp.MustCompile(`[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
//     emails := emailRe.FindAllString(text, -1)
//     fmt.Println("Emails:", emails)

//     // Extract phone numbers
//     phoneRe := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)
//     phones := phoneRe.FindAllString(text, -1)
//     fmt.Println("Phones:", phones)

//     // Extract URLs
//     urlRe := regexp.MustCompile(`https?://[^\s]+`)
//     urls := urlRe.FindAllString(text, -1)
//     fmt.Println("URLs:", urls)

//     // Extract prices
//     priceRe := regexp.MustCompile(`\$[\d,]+\.?\d*`)
// 	prices := priceRe.FindAllString(text, -1)
//     fmt.Println("Prices:", prices)

//     // Extract dates (YYYY-MM-DD)
//     dateRe := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
//     dates := dateRe.FindAllString(text, -1)
//     fmt.Println("Dates:", dates)
// }

// fmt.Println(time.Now())  // Output: 2024-07-30 15:04:05.123456789 +0530 IST

// specificTime := time.Date(2024, time.July, 30, 12, 0, 0, 0, time.UTC)
// fmt.Println("Specific time:", specificTime)
// // Output: 2024-07-30 12:00:00 +0000 UTC

// // Go uses a REFERENCE TIME for parsing: Mon Jan 2 15:04:05 MST 2006
// // Remember: 01/02 03:04:05 PM '06 -0700

// parsedTime, _ := time.Parse("2006-01-02", "2020-05-01")
// fmt.Println(parsedTime)  // 2020-05-01 00:00:00 +0000 UTC

// parsedTime1, _ := time.Parse("06-01-02", "20-05-01")
// fmt.Println(parsedTime1)  // 2020-05-01 00:00:00 +0000 UTC

// parsedTime2, _ := time.Parse("06-1-2", "20-5-1")
// fmt.Println(parsedTime2)  // 2020-05-01 00:00:00 +0000 UTC

// parsedTime3, _ := time.Parse("06-1-2 15-04", "20-5-1 18-03")
// fmt.Println(parsedTime3)  // 2020-05-01 18:03:00 +0000 UTC

// t := time.Now()
// fmt.Println("Formatted time:", t.Format("Monday 06-01-02 15-04-05"))
// // Output: Tuesday 24-07-30 15-04-05

// oneDayLater := t.Add(time.Hour * 24)
// fmt.Println(oneDayLater)           // Tomorrow's date
// fmt.Println(oneDayLater.Weekday()) // Day of week

// fmt.Println("Rounded Time:", t.Round(time.Hour))      // Rounds to nearest hour
// fmt.Println("Truncated Time:", t.Truncate(time.Hour)) // Removes minutes/seconds

// loc, _ := time.LoadLocation("America/New_York")
// tInNY := time.Now().In(loc)
// fmt.Println("New York Time:", tInNY)

// t1 := time.Date(2024, time.July, 4, 12, 0, 0, 0, time.UTC)
// t2 := time.Date(2024, time.July, 4, 18, 0, 0, 0, time.UTC)

// duration := t2.Sub(t1)
// fmt.Println("Duration:", duration)  // 6h0m0s

// fmt.Println("t2 is after t1?", t2.After(t1))  // true

// package main

// import (
//     "fmt"
//     "time"
// )

// func main() {
//     t := time.Now()

//     // Go's reference time: Mon Jan 2 15:04:05 MST 2006
//     // Easy to remember: 1/2 3:4:5 PM 2006 (01-02-03-04-05-06)

//     formats := map[string]string{
//         // Date formats
//         "2006-01-02":          "YYYY-MM-DD",
//         "02-01-2006":          "DD-MM-YYYY",
//         "01/02/2006":          "MM/DD/YYYY",
//         "January 2, 2006":     "Month Day, Year",
//         "Jan 2, 2006":         "Mon Day, Year",
//         "2 Jan 2006":          "Day Mon Year",
//         "Monday, January 2, 2006": "Full date",

//         // Time formats
//         "15:04":               "24-hour HH:MM",
//         "15:04:05":            "24-hour HH:MM:SS",
//         "3:04 PM":             "12-hour H:MM AM/PM",
//         "03:04:05 PM":         "12-hour HH:MM:SS AM/PM",

//         // Combined
//         "2006-01-02 15:04:05": "YYYY-MM-DD HH:MM:SS",
//         "Mon, 02 Jan 2006 15:04:05 MST": "RFC1123",
//         "2006-01-02T15:04:05Z07:00":     "ISO 8601 / RFC3339",
//     }

//     fmt.Println("Current time:", t)
//     fmt.Println("\nFormatting Examples:")
//     fmt.Println(strings.Repeat("-", 60))

//     for format, description := range formats {
//         fmt.Printf("%-35s → %s\n", description, t.Format(format))
//     }
// }
// ```

// **Output:**
// ```
// Current time: 2024-07-30 15:30:45.123456789 +0530 IST

// Formatting Examples:
// ------------------------------------------------------------
// YYYY-MM-DD                          → 2024-07-30
// DD-MM-YYYY                          → 30-07-2024
// MM/DD/YYYY                          → 07/30/2024
// Month Day, Year                     → July 30, 2024
// 24-hour HH:MM:SS                    → 15:30:45
// 12-hour H:MM AM/PM                  → 3:30 PM
// ISO 8601 / RFC3339                  → 2024-07-30T15:30:45+05:30

// func main() {
//     dateStrings := []struct {
//         layout string
//         value  string
//     }{
//         {"2006-01-02", "2024-07-30"},
//         {"02/01/2006", "30/07/2024"},
//         {"January 2, 2006", "July 30, 2024"},
//         {"2006-01-02 15:04:05", "2024-07-30 14:30:00"},
//         {"Mon, 02 Jan 2006", "Tue, 30 Jul 2024"},
//         {time.RFC3339, "2024-07-30T14:30:00Z"},
//         {time.RFC1123, "Tue, 30 Jul 2024 14:30:00 UTC"},
//     }

//     for _, ds := range dateStrings {
//         t, err := time.Parse(ds.layout, ds.value)
//         if err != nil {
//             fmt.Printf("Error parsing %s: %v\n", ds.value, err)
//             continue
//         }
//         fmt.Printf("Parsed %-30s → %v\n", ds.value, t)
//     }

//     // Parsing with location
//     loc, _ := time.LoadLocation("America/New_York")
//     t, _ := time.ParseInLocation("2006-01-02 15:04", "2024-07-30 14:30", loc)
//     fmt.Println("\nParsed with location:", t)
// }

// func main() {
//     t := time.Now()

//     fmt.Println("=== TIME COMPONENTS ===")
//     fmt.Println()

//     // Date components
//     fmt.Printf("Year:       %d\n", t.Year())
//     fmt.Printf("Month:      %s (%d)\n", t.Month(), t.Month())
//     fmt.Printf("Day:        %d\n", t.Day())
//     fmt.Printf("Weekday:    %s (%d)\n", t.Weekday(), t.Weekday())
//     fmt.Printf("Year Day:   %d\n", t.YearDay())

//     // Time components
//     fmt.Printf("Hour:       %d\n", t.Hour())
//     fmt.Printf("Minute:     %d\n", t.Minute())
//     fmt.Printf("Second:     %d\n", t.Second())
//     fmt.Printf("Nanosecond: %d\n", t.Nanosecond())

//     // Other info
//     fmt.Printf("Unix:       %d\n", t.Unix())
//     fmt.Printf("UnixMilli:  %d\n", t.UnixMilli())
//     fmt.Printf("UnixNano:   %d\n", t.UnixNano())

//     // Get year, month, day at once
//     year, month, day := t.Date()
//     fmt.Printf("\nDate(): %d-%s-%d\n", year, month, day)

//     // Get hour, minute, second at once
//     hour, min, sec := t.Clock()
//     fmt.Printf("Clock(): %02d:%02d:%02d\n", hour, min, sec)

//     // Week info
//     year, week := t.ISOWeek()
//     fmt.Printf("ISO Week: Year %d, Week %d\n", year, week)
// }

// func main() {
//     t := time.Now()
//     fmt.Println("Current Time:", t.Format("2006-01-02 15:04:05"))
//     fmt.Println()

//     // Add duration
//     fmt.Println("=== ADDING TIME ===")
//     fmt.Println("+ 1 hour:     ", t.Add(time.Hour).Format("2006-01-02 15:04:05"))
//     fmt.Println("+ 30 minutes: ", t.Add(30*time.Minute).Format("2006-01-02 15:04:05"))
//     fmt.Println("+ 24 hours:   ", t.Add(24*time.Hour).Format("2006-01-02 15:04:05"))
//     fmt.Println("+ 7 days:     ", t.Add(7*24*time.Hour).Format("2006-01-02 15:04:05"))

//     // AddDate for calendar arithmetic
//     fmt.Println("\n=== ADDING DATE ===")
//     fmt.Println("+ 1 year:     ", t.AddDate(1, 0, 0).Format("2006-01-02"))
//     fmt.Println("+ 1 month:    ", t.AddDate(0, 1, 0).Format("2006-01-02"))
//     fmt.Println("+ 1 week:     ", t.AddDate(0, 0, 7).Format("2006-01-02"))
//     fmt.Println("+ 1 day:      ", t.AddDate(0, 0, 1).Format("2006-01-02"))

//     // Subtract time
//     fmt.Println("\n=== SUBTRACTING TIME ===")
//     fmt.Println("- 1 hour:     ", t.Add(-time.Hour).Format("2006-01-02 15:04:05"))
//     fmt.Println("- 1 day:      ", t.AddDate(0, 0, -1).Format("2006-01-02"))
//     fmt.Println("- 1 month:    ", t.AddDate(0, -1, 0).Format("2006-01-02"))
//     fmt.Println("- 1 year:     ", t.AddDate(-1, 0, 0).Format("2006-01-02"))

//     // Calculate duration between times
//     fmt.Println("\n=== DURATION BETWEEN TIMES ===")
//     birthday := time.Date(2000, time.January, 15, 0, 0, 0, 0, time.UTC)
//     age := t.Sub(birthday)
//     fmt.Printf("Age: %.2f years\n", age.Hours()/24/365.25)

//     // Until and Since
//     newYear := time.Date(t.Year()+1, time.January, 1, 0, 0, 0, 0, time.Local)
//     fmt.Printf("Days until New Year: %.0f\n", time.Until(newYear).Hours()/24)

//     startOfYear := time.Date(t.Year(), time.January, 1, 0, 0, 0, 0, time.Local)
//     fmt.Printf("Days since start of year: %.0f\n", time.Since(startOfYear).Hours()/24)
// }

// func main() {
//     // Creating durations
//     fmt.Println("=== CREATING DURATIONS ===")

//     d1 := 2 * time.Hour
//     d2 := 30 * time.Minute
//     d3 := 45 * time.Second
//     d4 := 500 * time.Millisecond
//     d5 := 1000 * time.Microsecond
//     d6 := 1000000 * time.Nanosecond

//     fmt.Println("2 hours:", d1)
//     fmt.Println("30 minutes:", d2)
//     fmt.Println("45 seconds:", d3)
//     fmt.Println("500 milliseconds:", d4)
//     fmt.Println("1000 microseconds:", d5)
//     fmt.Println("1000000 nanoseconds:", d6)

//     // Parse duration from string
//     parsed, _ := time.ParseDuration("2h30m45s")
//     fmt.Println("\nParsed '2h30m45s':", parsed)

//     parsed2, _ := time.ParseDuration("1.5h")
//     fmt.Println("Parsed '1.5h':", parsed2)

//     // Duration components
//     fmt.Println("\n=== DURATION COMPONENTS ===")
//     d := 90*time.Minute + 30*time.Second
//     fmt.Println("Duration:", d)
//     fmt.Printf("Hours: %.2f\n", d.Hours())
//     fmt.Printf("Minutes: %.2f\n", d.Minutes())
//     fmt.Printf("Seconds: %.2f\n", d.Seconds())
//     fmt.Printf("Milliseconds: %d\n", d.Milliseconds())
//     fmt.Printf("Nanoseconds: %d\n", d.Nanoseconds())

//     // Custom duration formatting
//     fmt.Println("\n=== CUSTOM DURATION FORMAT ===")
//     totalDuration := 3*time.Hour + 25*time.Minute + 45*time.Second
//     formatDuration(totalDuration)
// }

// func formatDuration(d time.Duration) {
//     hours := int(d.Hours())
//     minutes := int(d.Minutes()) % 60
//     seconds := int(d.Seconds()) % 60

//     fmt.Printf("Duration: %dh %dm %ds\n", hours, minutes, seconds)
//     fmt.Printf("In words: %d hours, %d minutes, %d seconds\n", hours, minutes, seconds)
// }

// func main() {
//     t1 := time.Date(2024, 7, 15, 10, 0, 0, 0, time.UTC)
//     t2 := time.Date(2024, 7, 15, 14, 0, 0, 0, time.UTC)
//     t3 := time.Date(2024, 7, 15, 10, 0, 0, 0, time.UTC)

//     fmt.Println("t1:", t1)
//     fmt.Println("t2:", t2)
//     fmt.Println("t3:", t3)
//     fmt.Println()

//     // Comparison methods
//     fmt.Println("=== COMPARISONS ===")
//     fmt.Printf("t1.Before(t2): %v\n", t1.Before(t2))  // true
//     fmt.Printf("t1.After(t2):  %v\n", t1.After(t2))   // false
//     fmt.Printf("t1.Equal(t3):  %v\n", t1.Equal(t3))   // true

//     // Compare method returns -1, 0, or 1
//     fmt.Println("\n=== COMPARE METHOD ===")
//     fmt.Printf("t1.Compare(t2): %d (t1 < t2)\n", t1.Compare(t2))
//     fmt.Printf("t2.Compare(t1): %d (t2 > t1)\n", t2.Compare(t1))
//     fmt.Printf("t1.Compare(t3): %d (t1 = t3)\n", t1.Compare(t3))

//     // Check if time is zero
//     fmt.Println("\n=== ZERO TIME ===")
//     var zeroTime time.Time
//     fmt.Printf("zeroTime.IsZero(): %v\n", zeroTime.IsZero())
//     fmt.Printf("t1.IsZero(): %v\n", t1.IsZero())

//     // Practical example: Check if time is within range
//     fmt.Println("\n=== TIME RANGE CHECK ===")
//     start := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
//     end := time.Date(2024, 7, 31, 23, 59, 59, 0, time.UTC)
//     checkTime := time.Date(2024, 7, 15, 12, 0, 0, 0, time.UTC)

//     isInRange := checkTime.After(start) && checkTime.Before(end)
//     fmt.Printf("Is %s within July 2024? %v\n",
//         checkTime.Format("2006-01-02"), isInRange)
// }

// func main() {
//     t := time.Now()

//     fmt.Println("=== TIME ZONES ===")
//     fmt.Println("Local time:", t)

//     // UTC
//     fmt.Println("UTC time:", t.UTC())

//     // Load different locations
//     locations := []string{
//         "America/New_York",
//         "America/Los_Angeles",
//         "Europe/London",
//         "Europe/Paris",
//         "Asia/Tokyo",
//         "Asia/Kolkata",
//         "Australia/Sydney",
//     }

//     fmt.Println("\n=== WORLD CLOCK ===")
//     for _, locName := range locations {
//         loc, err := time.LoadLocation(locName)
//         if err != nil {
//             fmt.Printf("Error loading %s: %v\n", locName, err)
//             continue
//         }
//         localTime := t.In(loc)
//         fmt.Printf("%-20s: %s\n", locName, localTime.Format("2006-01-02 15:04:05 MST"))
//     }

//     // Get timezone info
//     fmt.Println("\n=== TIMEZONE INFO ===")
//     name, offset := t.Zone()
//     fmt.Printf("Current zone: %s (UTC%+d)\n", name, offset/3600)

//     // Fixed offset timezone
//     fixedZone := time.FixedZone("IST", 5*60*60+30*60) // +5:30
//     tInIST := t.In(fixedZone)
//     fmt.Println("Fixed zone (IST):", tInIST)

//     // Convert between zones
//     fmt.Println("\n=== TIMEZONE CONVERSION ===")
//     nyLoc, _ := time.LoadLocation("America/New_York")
//     tokyoLoc, _ := time.LoadLocation("Asia/Tokyo")

//     meeting := time.Date(2024, 7, 30, 14, 0, 0, 0, nyLoc)
//     fmt.Println("Meeting in New York:", meeting.Format("2006-01-02 15:04 MST"))
//     fmt.Println("Meeting in Tokyo:", meeting.In(tokyoLoc).Format("2006-01-02 15:04 MST"))
//     fmt.Println("Meeting in UTC:", meeting.UTC().Format("2006-01-02 15:04 MST"))
// }

// func main() {
//     t := time.Date(2024, 7, 30, 14, 37, 45, 123456789, time.UTC)

//     fmt.Println("Original time:", t)
//     fmt.Println()

//     // Truncate (always rounds down)
//     fmt.Println("=== TRUNCATE (Floor) ===")
//     fmt.Println("To hour:        ", t.Truncate(time.Hour))
//     fmt.Println("To 30 minutes:  ", t.Truncate(30*time.Minute))
//     fmt.Println("To 15 minutes:  ", t.Truncate(15*time.Minute))
//     fmt.Println("To minute:      ", t.Truncate(time.Minute))
//     fmt.Println("To second:      ", t.Truncate(time.Second))

//     // Round (rounds to nearest)
//     fmt.Println("\n=== ROUND (Nearest) ===")
//     fmt.Println("To hour:        ", t.Round(time.Hour))
//     fmt.Println("To 30 minutes:  ", t.Round(30*time.Minute))
//     fmt.Println("To 15 minutes:  ", t.Round(15*time.Minute))
//     fmt.Println("To minute:      ", t.Round(time.Minute))
//     fmt.Println("To second:      ", t.Round(time.Second))

//     // Practical example: Round to 5-minute intervals
//     fmt.Println("\n=== PRACTICAL: 5-MINUTE INTERVALS ===")
//     times := []time.Time{
//         time.Date(2024, 7, 30, 14, 32, 0, 0, time.UTC),
//         time.Date(2024, 7, 30, 14, 33, 0, 0, time.UTC),
//         time.Date(2024, 7, 30, 14, 37, 0, 0, time.UTC),
//         time.Date(2024, 7, 30, 14, 38, 0, 0, time.UTC),
//     }

//     for _, tm := range times {
//         rounded := tm.Round(5 * time.Minute)
//         truncated := tm.Truncate(5 * time.Minute)
//         fmt.Printf("%s → Rounded: %s, Truncated: %s\n",
//             tm.Format("15:04"),
//             rounded.Format("15:04"),
//             truncated.Format("15:04"))
//     }
// }

// func main() {
//     // Timer - fires once after duration
//     fmt.Println("=== TIMER ===")
//     timer := time.NewTimer(2 * time.Second)
//     fmt.Println("Timer started at", time.Now().Format("15:04:05"))

//     <-timer.C // Wait for timer
//     fmt.Println("Timer fired at", time.Now().Format("15:04:05"))

//     // After - simpler way to wait
//     fmt.Println("\n=== AFTER ===")
//     fmt.Println("Waiting 1 second...")
//     <-time.After(1 * time.Second)
//     fmt.Println("Done!")

//     // AfterFunc - execute function after duration
//     fmt.Println("\n=== AFTER FUNC ===")
//     done := make(chan bool)
//     time.AfterFunc(1*time.Second, func() {
//         fmt.Println("AfterFunc executed!")
//         done <- true
//     })
//     fmt.Println("AfterFunc scheduled...")
//     <-done

//     // Ticker - fires repeatedly
//     fmt.Println("\n=== TICKER ===")
//     ticker := time.NewTicker(500 * time.Millisecond)
//     count := 0

//     for t := range ticker.C {
//         count++
//         fmt.Printf("Tick %d at %s\n", count, t.Format("15:04:05.000"))
//         if count >= 5 {
//             ticker.Stop()
//             break
//         }
//     }
//     fmt.Println("Ticker stopped")

//     // Timeout pattern
//     fmt.Println("\n=== TIMEOUT PATTERN ===")
//     select {
//     case <-time.After(1 * time.Second):
//         fmt.Println("Operation timed out")
//     case result := <-doSomething():
//         fmt.Println("Result:", result)
//     }
// }

// func doSomething() chan string {
//     ch := make(chan string)
//     go func() {
//         time.Sleep(2 * time.Second) // Simulates slow operation
//         ch <- "done"
//     }()
//     return ch
// }

// func calculateAge(birthday time.Time) (years, months, days int) {
//     now := time.Now()

//     years = now.Year() - birthday.Year()
//     months = int(now.Month()) - int(birthday.Month())
//     days = now.Day() - birthday.Day()

//     if days < 0 {
//         months--
//         // Get days in previous month
//         prevMonth := now.AddDate(0, -1, 0)
//         days += time.Date(prevMonth.Year(), prevMonth.Month()+1, 0, 0, 0, 0, 0, time.UTC).Day()
//     }

//     if months < 0 {
//         years--
//         months += 12
//     }

//     return
// }

// func main() {
//     birthday := time.Date(1990, time.March, 15, 0, 0, 0, 0, time.UTC)
//     years, months, days := calculateAge(birthday)
//     fmt.Printf("Age: %d years, %d months, %d days\n", years, months, days)
// }

// func isWeekend(t time.Time) bool {
//     day := t.Weekday()
//     return day == time.Saturday || day == time.Sunday
// }

// func addBusinessDays(start time.Time, days int) time.Time {
//     result := start
//     added := 0

//     for added < days {
//         result = result.AddDate(0, 0, 1)
//         if !isWeekend(result) {
//             added++
//         }
//     }

//     return result
// }

// func countBusinessDays(start, end time.Time) int {
//     count := 0
//     current := start

//     for current.Before(end) || current.Equal(end) {
//         if !isWeekend(current) {
//             count++
//         }
//         current = current.AddDate(0, 0, 1)
//     }

//     return count
// }

// func main() {
//     start := time.Date(2024, 7, 29, 0, 0, 0, 0, time.UTC) // Monday

//     // Add 5 business days
//     result := addBusinessDays(start, 5)
//     fmt.Printf("5 business days from %s: %s\n",
//         start.Format("Mon 2006-01-02"),
//         result.Format("Mon 2006-01-02"))

//     // Count business days in a month
//     monthStart := time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)
//     monthEnd := time.Date(2024, 7, 31, 0, 0, 0, 0, time.UTC)
//     businessDays := countBusinessDays(monthStart, monthEnd)
//     fmt.Printf("Business days in July 2024: %d\n", businessDays)
// }

// type Event struct {
//     Name      string
//     StartTime time.Time
//     Duration  time.Duration
// }

// func (e Event) EndTime() time.Time {
//     return e.StartTime.Add(e.Duration)
// }

// func (e Event) IsOngoing() bool {
//     now := time.Now()
//     return now.After(e.StartTime) && now.Before(e.EndTime())
// }

// func (e Event) TimeUntil() time.Duration {
//     return time.Until(e.StartTime)
// }

// func main() {
//     events := []Event{
//         {"Meeting", time.Now().Add(1 * time.Hour), 30 * time.Minute},
//         {"Lunch", time.Now().Add(3 * time.Hour), 1 * time.Hour},
//         {"Review", time.Now().Add(5 * time.Hour), 45 * time.Minute},
//     }

//     fmt.Println("=== TODAY'S EVENTS ===")
//     for _, event := range events {
//         fmt.Printf("\n%s:\n", event.Name)
//         fmt.Printf("  Start: %s\n", event.StartTime.Format("15:04"))
//         fmt.Printf("  End: %s\n", event.EndTime().Format("15:04"))
//         fmt.Printf("  Duration: %v\n", event.Duration)
//         fmt.Printf("  Starts in: %v\n", event.TimeUntil().Round(time.Minute))
//     }
// }

// func countdown(target time.Time) {
//     for {
//         remaining := time.Until(target)

//         if remaining <= 0 {
//             fmt.Println("\n🎉 TIME'S UP!")
//             return
//         }

//         days := int(remaining.Hours()) / 24
//         hours := int(remaining.Hours()) % 24
//         minutes := int(remaining.Minutes()) % 60
//         seconds := int(remaining.Seconds()) % 60

//         fmt.Printf("\r⏰ %dd %02dh %02dm %02ds remaining", days, hours, minutes, seconds)

//         time.Sleep(1 * time.Second)
//     }
// }

// func main() {
//     // Countdown to New Year
//     newYear := time.Date(time.Now().Year()+1, 1, 1, 0, 0, 0, 0, time.Local)
//     fmt.Printf("Counting down to %s\n\n", newYear.Format("January 2, 2006"))
//     countdown(newYear)
// }




