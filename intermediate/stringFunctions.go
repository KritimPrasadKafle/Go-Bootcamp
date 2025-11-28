// package main

// import (
// 	"fmt"
// 	"strings"
// )

// var builder strings.Builder

// func main() {
// 	builder.WriteString("Hello")
// 	builder.WriteString(", ")
// 	builder.WriteString("world!")

// 	result := builder.String()
// 	fmt.Println(result)

// 	builder.WriteRune(' ')
// 	builder.WriteString("How are you")
// 	result = builder.String()
// 	fmt.Println(result)

// 	builder.Reset()

// 	builder.WriteString("Starting fresh!")
// 	result = builder.String()
// 	fmt.Println(result)
// }

// func buildUserInfo(name string, age int, city string) string {
//     var builder strings.Builder

//     builder.WriteString("Name: ")
//     builder.WriteString(name)
//     builder.WriteString("\nAge: ")
//     builder.WriteString(strconv.Itoa(age))
//     builder.WriteString("\nCity: ")
//     builder.WriteString(city)

//     return builder.String()
// }

// // Usage
// info := buildUserInfo("Alice", 25, "New York")
// fmt.Println(info)
// // Output:
// // Name: Alice
// // Age: 25
// // City: New York

// func buildHTML(title string, items []string) string {
//     var builder strings.Builder

//     builder.WriteString("<html>\n")
//     builder.WriteString("  <head><title>")
//     builder.WriteString(title)
//     builder.WriteString("</title></head>\n")
//     builder.WriteString("  <body>\n")
//     builder.WriteString("    <ul>\n")

//     for _, item := range items {
//         builder.WriteString("      <li>")
//         builder.WriteString(item)
//         builder.WriteString("</li>\n")
//     }

//     builder.WriteString("    </ul>\n")
//     builder.WriteString("  </body>\n")
//     builder.WriteString("</html>")

//     return builder.String()
// }

// // Usage
// items := []string{"Apple", "Banana", "Cherry"}
// html := buildHTML("My List", items)
// fmt.Println(html)

// // BAD: Creates many temporary strings
// func badConcat(words []string) string {
//     result := ""
//     for _, word := range words {
//         result += word + " "  // Creates new string each iteration!
//     }
//     return result
// }

// // GOOD: Uses Builder efficiently
// func goodConcat(words []string) string {
//     var builder strings.Builder
//     for _, word := range words {
//         builder.WriteString(word)
//         builder.WriteRune(' ')
//     }
//     return builder.String()
// }

// // Usage
// words := []string{"Go", "is", "awesome"}
// fmt.Println(goodConcat(words))  // Output: Go is awesome

// func buildLargeString(items []string) string {
//     var builder strings.Builder

//     // Pre-allocate space if you know approximate size
//     estimatedSize := 0
//     for _, item := range items {
//         estimatedSize += len(item) + 2  // +2 for ", "
//     }
//     builder.Grow(estimatedSize)

//     for i, item := range items {
//         builder.WriteString(item)
//         if i < len(items)-1 {
//             builder.WriteString(", ")
//         }
//     }

//     return builder.String()
// }

// // Usage
// items := []string{"apple", "banana", "cherry", "date"}
// fmt.Println(buildLargeString(items))
// // Output: apple, banana, cherry, date

// func buildJSON(data map[string]string) string {
//     var builder strings.Builder

//     builder.WriteRune('{')
//     builder.WriteRune('\n')

//     count := 0
//     for key, value := range data {
//         builder.WriteString("  \"")
//         builder.WriteString(key)
//         builder.WriteString("\": \"")
//         builder.WriteString(value)
//         builder.WriteRune('"')

//         if count < len(data)-1 {
//             builder.WriteRune(',')
//         }
//         builder.WriteRune('\n')
//         count++
//     }

//     builder.WriteRune('}')
//     return builder.String()
// }

// // Usage
// data := map[string]string{
//     "name": "John",
//     "city": "Paris",
// }
// fmt.Println(buildJSON(data))

// func buildQuery(table string, conditions map[string]string) string {
//     var builder strings.Builder

//     builder.WriteString("SELECT * FROM ")
//     builder.WriteString(table)

//     if len(conditions) > 0 {
//         builder.WriteString(" WHERE ")

//         count := 0
//         for key, value := range conditions {
//             builder.WriteString(key)
//             builder.WriteString(" = '")
//             builder.WriteString(value)
//             builder.WriteRune('\'')

//             if count < len(conditions)-1 {
//                 builder.WriteString(" AND ")
//             }
//             count++
//         }
//     }

//     builder.WriteRune(';')
//     return builder.String()
// }

// // Usage
// query := buildQuery("users", map[string]string{
//     "status": "active",
//     "role":   "admin",
// })
// fmt.Println(query)
// // Output: SELECT * FROM users WHERE status = 'active' AND role = 'admin';

// str := "Hello Go!"
// fmt.Println(len(str))  // Output: 9 (length of string)

// str1 := "Hello"
// str2 := "World"
// result := str1 + " " + str2  // String concatenation
// fmt.Println(result)  // Output: Hello World

// fmt.Println(str[0])      // Output: 72 (ASCII value of 'H')
// fmt.Println(str[1:7])    // Output: ello G (substring from index 1 to 6)

// num := 18
// str3 := strconv.Itoa(num)  // Integer to ASCII (string)
// fmt.Println(len(str3))     // Output: 2 (length of "18")

// // String to int
// str := "123"
// num, err := strconv.Atoi(str)
// if err == nil {
//     fmt.Println(num)  // 123
// }

// // Int to string
// intVal := 456
// strVal := strconv.Itoa(intVal)  // "456"

// // Float to string
// floatVal := 3.14
// strFloat := strconv.FormatFloat(floatVal, 'f', 2, 64)  // "3.14"

// // String to float
// strNum := "3.14"
// floatNum, _ := strconv.ParseFloat(strNum, 64)  // 3.14

// fruits := "apple, orange, banana"
// fruits1 := "apple-orange-banana"

// parts := strings.Split(fruits, ",")    // ["apple", " orange", " banana"]
// parts1 := strings.Split(fruits1, "-") // ["apple", "orange", "banana"]

// fmt.Println(parts)
// fmt.Println(parts1)

// countries := []string{"Germany", "France", "Italy"}
// joined := strings.Join(countries, ", ")  // "Germany, France, Italy"
// fmt.Println(joined)

// // Split by multiple separators
// text := "one,two;three|four"
// // Method 1: Split multiple times
// parts := strings.FieldsFunc(text, func(r rune) bool {
//     return r == ',' || r == ';' || r == '|'
// })
// fmt.Println(parts)  // [one two three four]

// // Split with limit
// data := "a:b:c:d:e"
// limited := strings.SplitN(data, ":", 3)
// fmt.Println(limited)  // [a b c:d:e]

// // Join with custom separator
// words := []string{"Go", "is", "fun"}
// sentence := strings.Join(words, " ")
// fmt.Println(sentence)  // "Go is fun"

// fmt.Println(strings.Contains(str, "Go?"))  // false (exact match needed)

// replaced := strings.Replace(str, "Go", "Universe", 1)  // Replace first occurrence
// fmt.Println(replaced)  // "Hello Universe!"

// text := "Hello Go! Go is great! Go rocks!"

// // Replace all occurrences (use -1)
// allReplaced := strings.Replace(text, "Go", "Golang", -1)
// fmt.Println(allReplaced)  // "Hello Golang! Golang is great! Golang rocks!"

// // Replace only first 2
// twoReplaced := strings.Replace(text, "Go", "Golang", 2)
// fmt.Println(twoReplaced)  // "Hello Golang! Golang is great! Go rocks!"

// // ReplaceAll (Go 1.12+) - simpler syntax
// replaced := strings.ReplaceAll(text, "Go", "Golang")

// // Contains examples
// fmt.Println(strings.Contains("Hello World", "World"))  // true
// fmt.Println(strings.Contains("Hello World", "world"))  // false (case-sensitive)
// fmt.Println(strings.ContainsAny("Hello", "aeiou"))     // true (contains 'e' or 'o')

// strwspace := " Hello Everyone! "

// fmt.Println(strwspace)                      // " Hello Everyone! "
// fmt.Println(strings.TrimSpace(strwspace))   // "Hello Everyone!"
// fmt.Println(strings.ToLower(strwspace))     // " hello everyone! "
// fmt.Println(strings.ToUpper(strwspace))     // " HELLO EVERYONE! "

// // Trim specific characters from both ends
// str := "!!!Hello World!!!"
// trimmed := strings.Trim(str, "!")
// fmt.Println(trimmed)  // "Hello World"

// // Trim from left only
// leftTrim := strings.TrimLeft(str, "!")
// fmt.Println(leftTrim)  // "Hello World!!!"

// // Trim from right only
// rightTrim := strings.TrimRight(str, "!")
// fmt.Println(rightTrim)  // "!!!Hello World"

// // Trim prefix
// url := "https://example.com"
// withoutProtocol := strings.TrimPrefix(url, "https://")
// fmt.Println(withoutProtocol)  // "example.com"

// // Trim suffix
// file := "document.txt"
// withoutExt := strings.TrimSuffix(file, ".txt")
// fmt.Println(withoutExt)  // "document"

// // Case conversion
// mixed := "HeLLo WoRLd"
// fmt.Println(strings.ToLower(mixed))  // "hello world"
// fmt.Println(strings.ToUpper(mixed))  // "HELLO WORLD"
// fmt.Println(strings.Title(mixed))    // "Hello World" (deprecated, use cases.Title)

// fmt.Println(strings.Repeat("foo ", 3))      // "foo foo foo "
// fmt.Println(strings.Count("Hello", "l"))    // 2 (counts occurrences)

// fmt.Println(strings.HasPrefix("Hello", "he"))  // false (case-sensitive)
// fmt.Println(strings.HasSuffix("Hello", "lo"))  // true
// fmt.Println(strings.HasSuffix("Hello", "la"))  // false

// // Repeat
// stars := strings.Repeat("*", 10)
// fmt.Println(stars)  // "**********"

// divider := strings.Repeat("-", 50)
// fmt.Println(divider)  // "--------------------------------------------------"

// // Count
// text := "Mississippi"
// fmt.Println(strings.Count(text, "s"))   // 4
// fmt.Println(strings.Count(text, "ss"))  // 2
// fmt.Println(strings.Count(text, "i"))   // 4

// // Prefix/Suffix checks
// filename := "document.pdf"
// fmt.Println(strings.HasSuffix(filename, ".pdf"))   // true
// fmt.Println(strings.HasSuffix(filename, ".docx"))  // false

// url := "https://example.com"
// fmt.Println(strings.HasPrefix(url, "https://"))  // true
// fmt.Println(strings.HasPrefix(url, "http://"))   // false

// // Practical example: file type check
// func isImageFile(filename string) bool {
//     return strings.HasSuffix(filename, ".jpg") ||
//            strings.HasSuffix(filename, ".png") ||
//            strings.HasSuffix(filename, ".gif")
// }

// // CORRECT VERSION (your code was missing parentheses)
// str5 := "Hel1lo, 123 Go 11!"
// re := regexp.MustCompile(`\d+`)  // Match one or more digits
// matches := re.FindAllString(str5, -1)
// fmt.Println(matches)  // [1 123 11]

// import "regexp"

// // Find email
// text := "Contact us at info@example.com or support@test.com"
// emailRe := regexp.MustCompile(`\w+@\w+\.\w+`)
// emails := emailRe.FindAllString(text, -1)
// fmt.Println(emails)  // [info@example.com support@test.com]

// // Find phone numbers
// phones := "Call 123-456-7890 or 098-765-4321"
// phoneRe := regexp.MustCompile(`\d{3}-\d{3}-\d{4}`)
// numbers := phoneRe.FindAllString(phones, -1)
// fmt.Println(numbers)  // [123-456-7890 098-765-4321]

// // Match and replace
// text2 := "Price is $100 and $200"
// priceRe := regexp.MustCompile(`\$\d+`)
// replaced := priceRe.ReplaceAllString(text2, "€XX")
// fmt.Println(replaced)  // "Price is €XX and €XX"

// // Check if matches pattern
// re := regexp.MustCompile(`^[a-zA-Z]+$`)  // Only letters
// fmt.Println(re.MatchString("Hello"))   // true
// fmt.Println(re.MatchString("Hello123")) // false

// // Extract parts with groups
// dateText := "Date: 2024-12-25"
// dateRe := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)
// matches := dateRe.FindStringSubmatch(dateText)
// if len(matches) == 4 {
//     year, month, day := matches[1], matches[2], matches[3]
//     fmt.Printf("Year: %s, Month: %s, Day: %s\n", year, month, day)
// }

// str6 := "Hello, 世界"
// fmt.Println(utf8.RuneCountInString(str6))  // 9 (counts characters, not bytes)

// package main

// import (
// 	"fmt"
// 	"regexp"
// 	"strconv"
// 	"strings"
// 	"unicode/utf8"
// )

// func main() {
// 	fmt.Println("=== 1. BASIC STRING OPS ===")
// 	str := "Hello Go!"
// 	fmt.Println("Length:", len(str))
// 	fmt.Println("Concatenation:", "Hello" + " " + "World")
// 	fmt.Println("First char (byte):", str[0])
// 	fmt.Println("Substring:", str[1:7])

// 	fmt.Println("\n=== 2. CONVERSIONS ===")
// 	num := 18
// 	str3 := strconv.Itoa(num)
// 	fmt.Println("Int to string:", str3)

// 	fmt.Println("\n=== 3. SPLIT & JOIN ===")
// 	fruits := "apple,orange,banana"
// 	parts := strings.Split(fruits, ",")
// 	fmt.Println("Split:", parts)
// 	joined := strings.Join(parts, " | ")
// 	fmt.Println("Join:", joined)

// 	fmt.Println("\n=== 4. CONTAINS & REPLACE ===")
// 	fmt.Println("Contains 'Go':", strings.Contains(str, "Go"))
// 	fmt.Println("Replace:", strings.Replace(str, "Go", "World", 1))

// 	fmt.Println("\n=== 5. TRIM & CASE ===")
// 	messy := "  HELLO world  "
// 	fmt.Println("Trimmed:", strings.TrimSpace(messy))
// 	fmt.Println("Lower:", strings.ToLower(messy))
// 	fmt.Println("Upper:", strings.ToUpper(messy))

// 	fmt.Println("\n=== 6. REPEAT & COUNT ===")
// 	fmt.Println("Repeat:", strings.Repeat("Go! ", 3))
// 	fmt.Println("Count 'l':", strings.Count("Hello", "l"))
// 	fmt.Println("Has prefix 'He':", strings.HasPrefix("Hello", "He"))

// 	fmt.Println("\n=== 7. REGEX ===")
// 	text := "Order #123 costs $45 and #789 costs $90"
// 	re := regexp.MustCompile(`\d+`)
// 	numbers := re.FindAllString(text, -1)
// 	fmt.Println("Numbers found:", numbers)

// 	fmt.Println("\n=== 8. UTF-8 ===")
// 	str6 := "Hello, 世界 👋"
// 	fmt.Println("Bytes:", len(str6))
// 	fmt.Println("Characters:", utf8.RuneCountInString(str6))

// 	fmt.Println("\n=== 9. STRING BUILDER ===")
// 	var builder strings.Builder
// 	builder.WriteString("Built ")
// 	builder.WriteString("with ")
// 	builder.WriteString("efficiency!")
// 	fmt.Println(builder.String())
// }
