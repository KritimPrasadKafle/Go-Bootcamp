// package main

// import (
//     "fmt"
//     "net/url"
// )

// func main() {
//     // Complete URL with all components
//     rawURL := "https://john:secret@api.example.com:8080/v1/users/123?sort=name&limit=10#results"

//     parsedURL, err := url.Parse(rawURL)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }

//     fmt.Println("=== URL COMPONENTS ===")
//     fmt.Println()

//     // Scheme (protocol)
//     fmt.Printf("Scheme:      %s\n", parsedURL.Scheme)  // https

//     // User info
//     if parsedURL.User != nil {
//         fmt.Printf("Username:    %s\n", parsedURL.User.Username())
//         password, hasPassword := parsedURL.User.Password()
//         if hasPassword {
//             fmt.Printf("Password:    %s\n", password)
//         }
//     }

//     // Host (includes port if present)
//     fmt.Printf("Host:        %s\n", parsedURL.Host)      // api.example.com:8080
//     fmt.Printf("Hostname:    %s\n", parsedURL.Hostname()) // api.example.com
//     fmt.Printf("Port:        %s\n", parsedURL.Port())     // 8080

//     // Path
//     fmt.Printf("Path:        %s\n", parsedURL.Path)       // /v1/users/123
//     fmt.Printf("RawPath:     %s\n", parsedURL.RawPath)    // (encoded path if different)

//     // Query
//     fmt.Printf("RawQuery:    %s\n", parsedURL.RawQuery)   // sort=name&limit=10
//     fmt.Printf("Query:       %v\n", parsedURL.Query())    // map[limit:[10] sort:[name]]

//     // Fragment
//     fmt.Printf("Fragment:    %s\n", parsedURL.Fragment)   // results

//     // Full URL
//     fmt.Printf("\nFull URL:    %s\n", parsedURL.String())

//     // Request URI (path + query)
//     fmt.Printf("RequestURI:  %s\n", parsedURL.RequestURI())
// }
// ```

// **Output:**
// ```
// === URL COMPONENTS ===

// Scheme:      https
// Username:    john
// Password:    secret
// Host:        api.example.com:8080
// Hostname:    api.example.com
// Port:        8080
// Path:        /v1/users/123
// RawPath:
// RawQuery:    sort=name&limit=10
// Query:       map[limit:[10] sort:[name]]
// Fragment:    results

// Full URL:    https://john:secret@api.example.com:8080/v1/users/123?sort=name&limit=10#results
// RequestURI:  /v1/users/123?sort=name&limit=10

// func main() {
//     fmt.Println("=== QUERY PARAMETER OPERATIONS ===")
//     fmt.Println()

//     // Create url.Values
//     params := url.Values{}

//     // Add - can add multiple values for same key
//     params.Add("color", "red")
//     params.Add("color", "blue")
//     params.Add("size", "large")

//     fmt.Println("After Add:")
//     fmt.Printf("  params: %v\n", params)
//     fmt.Printf("  color (Get): %s\n", params.Get("color"))      // First value only
//     fmt.Printf("  color (all): %v\n", params["color"])          // All values

//     // Set - replaces all values for key
//     params.Set("color", "green")
//     fmt.Println("\nAfter Set(\"color\", \"green\"):")
//     fmt.Printf("  color: %v\n", params["color"])  // [green]

//     // Del - removes key entirely
//     params.Del("color")
//     fmt.Println("\nAfter Del(\"color\"):")
//     fmt.Printf("  params: %v\n", params)

//     // Has - check if key exists (Go 1.17+)
//     fmt.Printf("\nHas \"size\": %v\n", params.Has("size"))
//     fmt.Printf("Has \"color\": %v\n", params.Has("color"))

//     // Encode - convert to query string
//     params.Set("name", "John Doe")
//     params.Set("email", "john@example.com")
//     params.Set("message", "Hello & Goodbye!")

//     encoded := params.Encode()
//     fmt.Println("\nEncoded:")
//     fmt.Printf("  %s\n", encoded)

//     // Parse query string back
//     fmt.Println("\n=== PARSE QUERY STRING ===")
//     queryString := "name=Jane&age=25&city=New+York"
//     parsed, err := url.ParseQuery(queryString)
//     if err != nil {
//         fmt.Println("Error:", err)
//         return
//     }
//     fmt.Printf("Parsed: %v\n", parsed)
//     fmt.Printf("Name: %s\n", parsed.Get("name"))
//     fmt.Printf("City: %s\n", parsed.Get("city"))
// }

// func main() {
//     fmt.Println("=== URL ENCODING/DECODING ===")
//     fmt.Println()

//     // Special characters that need encoding
//     special := "Hello World! & More? #Fragment"

//     // QueryEscape - for query parameter values
//     queryEncoded := url.QueryEscape(special)
//     fmt.Println("QueryEscape:")
//     fmt.Printf("  Original: %s\n", special)
//     fmt.Printf("  Encoded:  %s\n", queryEncoded)

//     // QueryUnescape - decode query parameter values
//     decoded, _ := url.QueryUnescape(queryEncoded)
//     fmt.Printf("  Decoded:  %s\n", decoded)

//     // PathEscape - for URL path segments
//     path := "/users/john doe/profile"
//     pathEncoded := url.PathEscape(path)
//     fmt.Println("\nPathEscape:")
//     fmt.Printf("  Original: %s\n", path)
//     fmt.Printf("  Encoded:  %s\n", pathEncoded)

//     // PathUnescape - decode path segments
//     pathDecoded, _ := url.PathUnescape(pathEncoded)
//     fmt.Printf("  Decoded:  %s\n", pathDecoded)

//     // Difference between QueryEscape and PathEscape
//     fmt.Println("\n=== ESCAPE DIFFERENCES ===")
//     text := "hello/world test"
//     fmt.Printf("Original:    %s\n", text)
//     fmt.Printf("QueryEscape: %s\n", url.QueryEscape(text))  // hello%2Fworld+test
//     fmt.Printf("PathEscape:  %s\n", url.PathEscape(text))   // hello%2Fworld%20test

//     // Characters that need encoding
//     fmt.Println("\n=== SPECIAL CHARACTERS ===")
//     chars := []string{" ", "&", "=", "?", "#", "/", "+", "%"}
//     for _, c := range chars {
//         fmt.Printf("  '%s' → QueryEscape: '%s', PathEscape: '%s'\n",
//             c, url.QueryEscape(c), url.PathEscape(c))
//     }
// }
// ```

// **Output:**
// ```
// === URL ENCODING/DECODING ===

// QueryEscape:
//   Original: Hello World! & More? #Fragment
//   Encoded:  Hello+World%21+%26+More%3F+%23Fragment
//   Decoded:  Hello World! & More? #Fragment

// PathEscape:
//   Original: /users/john doe/profile
//   Encoded:  %2Fusers%2Fjohn%20doe%2Fprofile
//   Decoded:  /users/john doe/profile

// === ESCAPE DIFFERENCES ===
// Original:    hello/world test
// QueryEscape: hello%2Fworld+test
// PathEscape:  hello%2Fworld%20test

// === SPECIAL CHARACTERS ===
//   ' ' → QueryEscape: '+', PathEscape: '%20'
//   '&' → QueryEscape: '%26', PathEscape: '%26'
//   '=' → QueryEscape: '%3D', PathEscape: '%3D'
//   '?' → QueryEscape: '%3F', PathEscape: '%3F'
//   '#' → QueryEscape: '%23', PathEscape: '%23'
//   '/' → QueryEscape: '%2F', PathEscape: '%2F'
//   '+' → QueryEscape: '%2B', PathEscape: '%2B'
//   '%' → QueryEscape: '%25', PathEscape: '%25'