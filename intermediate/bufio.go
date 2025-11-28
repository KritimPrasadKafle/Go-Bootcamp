// package intermediate
// // Without buffering - SLOW (many system calls)
// ┌─────────────────────────────────────────────────────────────────────────────┐
// │                        WHY BUFFERED I/O?                                    │
// ├─────────────────────────────────────────────────────────────────────────────┤
// │                                                                             │
// │  WITHOUT BUFFER (Direct I/O):                                               │
// │  ─────────────────────────────                                              │
// │                                                                             │
// │  Program ──read 1 byte──► Disk  (slow - disk access every time)             │
// │  Program ──read 1 byte──► Disk                                              │
// │  Program ──read 1 byte──► Disk                                              │
// │  Program ──read 1 byte──► Disk                                              │
// │  ... 1000 times = 1000 disk accesses! 🐌                                    │
// │                                                                             │
// │  WITH BUFFER (bufio):                                                       │
// │  ────────────────────                                                       │
// │                                                                             │
// │  Program ──read──► [Buffer 4KB] ◄──single read──► Disk                      │
// │  Program ──read──► [Buffer 4KB] (from memory - fast!)                       │
// │  Program ──read──► [Buffer 4KB] (from memory - fast!)                       │
// │  ... 1000 times = only few disk accesses! 🚀                                │
// │                                                                             │
// │  Buffer acts as a middleman, reducing expensive I/O operations              │
// │                                                                             │
// └─────────────────────────────────────────────────────────────────────────────┘

// file, _ := os.Open("large_file.txt")
// for {
//     b := make([]byte, 1)
//     file.Read(b)  // System call for EVERY byte!
// }

// // With buffering - FAST (fewer system calls)
// file, _ := os.Open("large_file.txt")
// reader := bufio.NewReader(file)
// for {
//     b, _ := reader.ReadByte()  // Reads from buffer, refills when empty
// }

// // Raw io.Reader only has Read([]byte) method
// // bufio.Reader adds many useful methods:

// reader.ReadByte()      // Read single byte
// reader.ReadRune()      // Read single UTF-8 character
// reader.ReadString('\n') // Read until delimiter
// reader.ReadLine()      // Read a line
// reader.Peek(n)         // Look ahead without consuming

// // Create a reader from a string
// reader := bufio.NewReader(strings.NewReader("Hello, bufio package!\nHow are you doing?"))

// // Read into byte slice
// data := make([]byte, 20)
// n, err := reader.Read(data)
// if err != nil {
//     fmt.Println("Error reading:", err)
//     return
// }
// fmt.Printf("Read %d bytes: %s\n", n, data[:n])  // Read 20 bytes: Hello, bufio package

// // Read until newline delimiter
// line, err := reader.ReadString('\n')
// if err != nil {
//     fmt.Println("Error reading string:", err)
//     return
// }
// fmt.Println("Read string:", line)  // Read string: !\n

// // Create buffered writer wrapping os.Stdout
// writer := bufio.NewWriter(os.Stdout)

// // Write byte slice - data goes to BUFFER, not stdout yet!
// data := []byte("Hello, bufio package!\n")
// n, err := writer.Write(data)
// if err != nil {
//     fmt.Println("Error writing:", err)
//     return
// }
// fmt.Printf("Wrote %d bytes\n", n)

// // CRITICAL: Flush sends buffer contents to actual output
// err = writer.Flush()
// if err != nil {
//     fmt.Println("Error flushing writer:", err)
//     return
// }

// // Write string directly
// str := "This is a string.\n"
// n, err = writer.WriteString(str)
// if err != nil {
//     fmt.Println("Error writing string:", err)
//     return
// }
// fmt.Printf("Wrote %d bytes.\n", n)

// // Don't forget to flush!
// err = writer.Flush()

// package main

// import (
//     "bufio"
//     "fmt"
//     "io"
//     "strings"
// )

// func main() {
//     fmt.Println("=== bufio.Reader METHODS ===")
//     fmt.Println()

//     // Create reader
//     text := "Hello, World!\nLine 2\nLine 3\n日本語"
//     reader := bufio.NewReader(strings.NewReader(text))

//     // 1. Read - read into byte slice
//     fmt.Println("1. Read([]byte):")
//     data := make([]byte, 5)
//     n, _ := reader.Read(data)
//     fmt.Printf("   Read %d bytes: %q\n", n, data[:n])

//     // Reset reader for next examples
//     reader = bufio.NewReader(strings.NewReader(text))

//     // 2. ReadByte - read single byte
//     fmt.Println("\n2. ReadByte():")
//     b, _ := reader.ReadByte()
//     fmt.Printf("   Read byte: %c (value: %d)\n", b, b)

//     // 3. UnreadByte - put byte back
//     reader.UnreadByte()
//     b, _ = reader.ReadByte()
//     fmt.Printf("   After UnreadByte, read again: %c\n", b)

//     // Reset
//     reader = bufio.NewReader(strings.NewReader("日本語 Hello"))

//     // 4. ReadRune - read UTF-8 character
//     fmt.Println("\n3. ReadRune():")
//     r, size, _ := reader.ReadRune()
//     fmt.Printf("   Read rune: %c (size: %d bytes)\n", r, size)

//     // Reset
//     reader = bufio.NewReader(strings.NewReader("Hello\nWorld\nGo"))

//     // 5. ReadString - read until delimiter
//     fmt.Println("\n4. ReadString(delimiter):")
//     line, _ := reader.ReadString('\n')
//     fmt.Printf("   Read until '\\n': %q\n", line)

//     // 6. ReadBytes - similar but returns []byte
//     fmt.Println("\n5. ReadBytes(delimiter):")
//     lineBytes, _ := reader.ReadBytes('\n')
//     fmt.Printf("   Read bytes until '\\n': %q\n", lineBytes)

//     // Reset
//     reader = bufio.NewReader(strings.NewReader("Hello World"))

//     // 7. Peek - look ahead without consuming
//     fmt.Println("\n6. Peek(n):")
//     peeked, _ := reader.Peek(5)
//     fmt.Printf("   Peeked: %q\n", peeked)

//     // Read again to show data wasn't consumed
//     data = make([]byte, 5)
//     reader.Read(data)
//     fmt.Printf("   Read after peek: %q (same data!)\n", data)

//     // 8. Buffered - bytes available in buffer
//     fmt.Println("\n7. Buffered():")
//     reader = bufio.NewReader(strings.NewReader("Hello World"))
//     reader.Peek(5)  // Force some data into buffer
//     fmt.Printf("   Bytes buffered: %d\n", reader.Buffered())

//     // 9. ReadSlice - returns slice of internal buffer
//     fmt.Println("\n8. ReadSlice(delimiter):")
//     reader = bufio.NewReader(strings.NewReader("Hello,World"))
//     slice, _ := reader.ReadSlice(',')
//     fmt.Printf("   ReadSlice until ',': %q\n", slice)

//     // 10. Discard - skip bytes
//     fmt.Println("\n9. Discard(n):")
//     reader = bufio.NewReader(strings.NewReader("Skip this: Important data"))
//     discarded, _ := reader.Discard(11)
//     fmt.Printf("   Discarded %d bytes\n", discarded)
//     remaining, _ := io.ReadAll(reader)
//     fmt.Printf("   Remaining: %q\n", remaining)
// }
// ```

// **Output:**
// ```
// === bufio.Reader METHODS ===

// 1. Read([]byte):
//    Read 5 bytes: "Hello"

// 2. ReadByte():
//    Read byte: H (value: 72)
//    After UnreadByte, read again: H

// 3. ReadRune():
//    Read rune: 日 (size: 3 bytes)

// 4. ReadString(delimiter):
//    Read until '\n': "Hello\n"

// 5. ReadBytes(delimiter):
//    Read bytes until '\n': "World\n"

// 6. Peek(n):
//    Peeked: "Hello"
//    Read after peek: "Hello" (same data!)

// 7. Buffered():
//    Bytes buffered: 11

// 8. ReadSlice(delimiter):
//    ReadSlice until ',': "Hello,"

// 9. Discard(n):
//    Discarded 11 bytes
//    Remaining: "Important data"

// package main

// import (
//     "bufio"
//     "fmt"
//     "strings"
// )

// func main() {
//     fmt.Println("=== bufio.Scanner ===")
//     fmt.Println()

//     text := "Line 1\nLine 2\nLine 3"

//     // 1. Scan lines (default)
//     fmt.Println("1. ScanLines (default):")
//     scanner := bufio.NewScanner(strings.NewReader(text))
//     lineNum := 1
//     for scanner.Scan() {
//         fmt.Printf("   Line %d: %s\n", lineNum, scanner.Text())
//         lineNum++
//     }

//     // 2. Scan words
//     fmt.Println("\n2. ScanWords:")
//     wordText := "Hello World Go Programming"
//     scanner = bufio.NewScanner(strings.NewReader(wordText))
//     scanner.Split(bufio.ScanWords)

//     wordNum := 1
//     for scanner.Scan() {
//         fmt.Printf("   Word %d: %s\n", wordNum, scanner.Text())
//         wordNum++
//     }

//     // 3. Scan bytes
//     fmt.Println("\n3. ScanBytes:")
//     scanner = bufio.NewScanner(strings.NewReader("Hello"))
//     scanner.Split(bufio.ScanBytes)

//     fmt.Print("   Bytes: ")
//     for scanner.Scan() {
//         fmt.Printf("%q ", scanner.Text())
//     }
//     fmt.Println()

//     // 4. Scan runes (UTF-8 characters)
//     fmt.Println("\n4. ScanRunes:")
//     scanner = bufio.NewScanner(strings.NewReader("Go日本"))
//     scanner.Split(bufio.ScanRunes)

//     fmt.Print("   Runes: ")
//     for scanner.Scan() {
//         fmt.Printf("%q ", scanner.Text())
//     }
//     fmt.Println()

//     // 5. Check for errors
//     fmt.Println("\n5. Error handling:")
//     scanner = bufio.NewScanner(strings.NewReader("test"))
//     for scanner.Scan() {
//         fmt.Printf("   Token: %s\n", scanner.Text())
//     }
//     if err := scanner.Err(); err != nil {
//         fmt.Printf("   Error: %v\n", err)
//     } else {
//         fmt.Println("   No errors!")
//     }

//     // 6. Scanner.Bytes() vs Scanner.Text()
//     fmt.Println("\n6. Bytes() vs Text():")
//     scanner = bufio.NewScanner(strings.NewReader("Hello"))
//     scanner.Split(bufio.ScanWords)
//     for scanner.Scan() {
//         fmt.Printf("   Text(): %s (type: string)\n", scanner.Text())
//         fmt.Printf("   Bytes(): %v (type: []byte)\n", scanner.Bytes())
//     }
// }