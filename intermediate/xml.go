package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age,omitempty"`
	Email   string   `xml:"-"` // Never include
	Address Address  `xml:"address"`
}

type Address struct {
	City  string `xml:"city"`
	State string `xml:"state"`
}

func main() {
	person := Person{
		Name:  "John",
		Email: "john@example.com", // Won't appear!
		Address: Address{
			City:  "Oakland",
			State: "CA",
		},
	}

	// Pretty-printed XML
	xmlData, _ := xml.MarshalIndent(person, "", "  ")
	fmt.Println(string(xmlData))
}


package main

import (
    "encoding/xml"
    "fmt"
)

type Item struct {
    XMLName xml.Name `xml:"item"`
    Name    string   `xml:"name"`
    Price   float64  `xml:"price"`
}

func main() {
    item := Item{Name: "Widget", Price: 9.99}
    
    // Compact — single line
    compact, _ := xml.Marshal(item)
    fmt.Println("Compact:")
    fmt.Println(string(compact))
    
    // Pretty — with indentation
    // MarshalIndent(value, prefix, indent)
    pretty, _ := xml.MarshalIndent(item, "", "  ")
    fmt.Println("\nPretty:")
    fmt.Println(string(pretty))
    
    // With prefix on every line
    prefixed, _ := xml.MarshalIndent(item, ">>> ", "  ")
    fmt.Println("\nWith prefix:")
    fmt.Println(string(prefixed))
}
```

**Output:**
```
Compact:
<item><name>Widget</name><price>9.99</price></item>

Pretty:
<item>
  <name>Widget</name>
  <price>9.99</price>
</item>

With prefix:
>>> <item>
>>>   <name>Widget</name>
>>>   <price>9.99</price>
>>> </item>


package main

import (
    "encoding/xml"
    "fmt"
)

type Book struct {
    XMLName    xml.Name `xml:"book"`
    ISBN       string   `xml:"isbn,attr"`       // Attribute
    Title      string   `xml:"title,attr"`      // Attribute
    Author     string   `xml:"author,attr"`     // Attribute
    Publisher  string   `xml:"publisher"`       // Element
    Year       int      `xml:"year"`            // Element
}

func main() {
    book := Book{
        ISBN:      "978-0-13-468599-1",
        Title:     "The Go Programming Language",
        Author:    "Donovan & Kernighan",
        Publisher: "Addison-Wesley",
        Year:      2015,
    }
    
    xmlData, _ := xml.MarshalIndent(book, "", "  ")
    fmt.Println(string(xmlData))
}