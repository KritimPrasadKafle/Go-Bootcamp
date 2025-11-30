// Reader — anything you can READ from
type Reader interface {
    Read(p []byte) (n int, err error)
}

// Writer — anything you can WRITE to
type Writer interface {
    Write(p []byte) (n int, err error)
}

// Closer — anything you can CLOSE
type Closer interface {
    Close() error
}
```
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   io.Reader          io.Writer          io.Closer               │
│   ══════════         ══════════         ══════════              │
│                                                                 │
│   "I can be          "I can be          "I can be               │
│    read from"         written to"        closed"                │
│                                                                 │
│   Examples:          Examples:          Examples:               │
│   • *os.File         • *os.File         • *os.File              │
│   • strings.Reader   • bytes.Buffer     • net.Conn              │
│   • bytes.Buffer     • os.Stdout        • http.Response.Body    │
│   • http.Request.Body• net.Conn         • Database conn         │
│   • net.Conn         • http.ResponseWriter                      │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘



// ReadWriter — can both read AND write
type ReadWriter interface {
    Reader
    Writer
}

// ReadCloser — can read AND close
type ReadCloser interface {
    Reader
    Closer
}

// WriteCloser — can write AND close
type WriteCloser interface {
    Writer
    Closer
}

// ReadWriteCloser — can do all three
type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}


package main

import (
    "fmt"
    "io"
    "log"
    "strings"
)

// Works with ANY io.Reader!
func readFromReader(r io.Reader) {
    buf := make([]byte, 1024)  // Buffer to hold data
    n, err := r.Read(buf)      // Read into buffer
    
    if err != nil && err != io.EOF {
        log.Fatalln("Error reading:", err)
    }
    
    fmt.Println(string(buf[:n]))  // Only use bytes that were read
}

func main() {
    // strings.Reader implements io.Reader
    sr := strings.NewReader("Hello from strings.Reader!")
    readFromReader(sr)
}
```

**Output:**
```
Hello from strings.Reader!



package main

import (
    "bytes"
    "fmt"
    "io"
    "os"
    "strings"
)

func processReader(r io.Reader) {
    data, _ := io.ReadAll(r)
    fmt.Printf("Read: %s\n", data)
}

func main() {
    // 1. strings.Reader
    processReader(strings.NewReader("From string"))
    
    // 2. bytes.Reader
    processReader(bytes.NewReader([]byte("From bytes")))
    
    // 3. bytes.Buffer
    buf := bytes.NewBufferString("From buffer")
    processReader(buf)
    
    // 4. *os.File
    file, _ := os.Open("example.txt")
    if file != nil {
        processReader(file)
        file.Close()
    }
    
    // 5. os.Stdin (user input)
    // processReader(os.Stdin)
}


package main

import (
    "bytes"
    "fmt"
    "io"
    "log"
    "os"
)

// Works with ANY io.Writer!
func writeToWriter(w io.Writer, data string) {
    n, err := w.Write([]byte(data))
    if err != nil {
        log.Fatalln("Error writing:", err)
    }
    fmt.Printf("Wrote %d bytes\n", n)
}

func main() {
    // 1. Write to bytes.Buffer
    var buf bytes.Buffer
    writeToWriter(&buf, "Hello Buffer!")
    fmt.Println("Buffer contains:", buf.String())
    
    fmt.Println()
    
    // 2. Write to os.Stdout (console)
    writeToWriter(os.Stdout, "Hello Console!\n")
    
    // 3. Write to file
    file, _ := os.Create("output.txt")
    if file != nil {
        writeToWriter(file, "Hello File!")
        file.Close()
    }
}
```

**Output:**
```
Wrote 13 bytes
Buffer contains: Hello Buffer!

Wrote 15 bytes
Hello Console!
Wrote 11 bytes
```

### What Implements io.Writer?
```
┌─────────────────────────────────────────────────────────────────┐
│                                                                 │
│   COMMON io.Writer IMPLEMENTATIONS                              │
│                                                                 │
│   Type                │ Use Case                                │
│   ────────────────────┼────────────────────────────────────     │
│   *os.File            │ Write to files                          │
│   os.Stdout           │ Print to console                        │
│   os.Stderr           │ Print errors to console                 │
│   *bytes.Buffer       │ Build strings/bytes in memory           │
│   net.Conn            │ Send data over network                  │
│   http.ResponseWriter │ Send HTTP response                      │
│   *bufio.Writer       │ Buffered writing                        │
│   io.Discard          │ Throw away data (like /dev/null)        │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘



