# XML Dispatcher Integration Guide

The `xml-dispatcher` package provides a modular and reusable system for processing XML data in Go. It allows you to define custom handlers for specific XML formats and process XML inputs by delegating to the appropriate handler. This guide explains how to integrate `xml-dispatcher` into your Go project.

## Prerequisites

- **Go Version**: 1.22.4 or later (as specified in `go.mod`).
- **Module Path**: The package is hosted at `github.com/brandonbryant12/xml-dispatcher`. Ensure your project uses Go modules.

## Step 1: Import the Package

Add the `xmldispatcher` package to your project by importing it in your Go code and fetching it with `go get`.

```go
import "github.com/brandonbryant12/xml-dispatcher"
```

Run the following command to download the package:

```bash
go get github.com/brandonbryant12/xml-dispatcher
```

This will add the dependency to your go.mod file.

## Step 2: Implement Custom Handlers

Create custom handlers by implementing the Handler interface, which requires two methods:

- `CanHandle(xmlData []byte) bool`: Determines if the handler can process the given XML data.
- `Handle(xmlData []byte) error`: Processes the XML data and returns an error if processing fails.

Here's an example of a custom handler for an `<order>` XML structure:

```go
package mypackage

import (
	"encoding/xml"
	"fmt"
	"github.com/brandonbryant12/xml-dispatcher"
)

type OrderHandler struct{}

func (o *OrderHandler) CanHandle(xmlData []byte) bool {
	type Root struct {
		XMLName xml.Name `xml:"order"`
	}
	var root Root
	err := xml.Unmarshal(xmlData, &root)
	return err == nil && root.XMLName.Local == "order"
}

func (o *OrderHandler) Handle(xmlData []byte) error {
	type Order struct {
		ID string `xml:"id"`
	}
	var order Order
	if err := xml.Unmarshal(xmlData, &order); err != nil {
		return err
	}
	fmt.Println("Processing order ID:", order.ID)
	// Add your processing logic here (e.g., save to database)
	return nil
}
```

Notes:
- Use encoding/xml to parse XML into Go structs.
- CanHandle should validate the XML structure (e.g., root tag, attributes, or nested elements).
- Handle performs the actual processing (e.g., logging, storing data).

## Step 3: Set Up the XML Processor

Create an instance of XMLProcessor and register your custom handlers.

```go
package main

import (
	"log"
	"github.com/brandonbryant12/xml-dispatcher/xmldispatcher"
)

func main() {
	// Initialize the XML processor
	processor := xmldispatcher.NewXMLProcessor()

	// Register handlers
	processor.RegisterHandler(&xmldispatcher.ReportHandler{}) // Built-in example handler
	processor.RegisterHandler(&OrderHandler{})                // Custom handler

	// Example usage
	xmlData := []byte(`<order><id>12345</id></order>`)
	if err := processor.ProcessXML(xmlData); err != nil {
		log.Printf("Error processing XML: %v", err)
	}
}
```

- `NewXMLProcessor()`: Creates a new processor instance.
- `RegisterHandler(h Handler)`: Adds a handler to the processor's list. Handlers are checked in registration order.

## Step 4: Process XML Data

Call ProcessXML with your XML data (as a []byte). Handle any errors returned:

```go
xmlData := []byte(`<report><data>Hello</data></report>`)
if err := processor.ProcessXML(xmlData); err != nil {
	log.Printf("Error: %v", err)
}
```

Behavior:
- The processor iterates through registered handlers, calling CanHandle on each.
- The first handler that returns true processes the XML via its Handle method.
- If no handler matches, it returns an error: "no handler found for the given XML".

## Built-in Example Handler

The package includes a ReportHandler as an example:

- Matches: XML with a `<report>` root tag.
- Processes: Prints the `<data>` element's content.

You can use it as a reference or directly in your application:

```go
processor.RegisterHandler(&xmldispatcher.ReportHandler{})
```

## Error Handling

- **Invalid XML**: If CanHandle or Handle encounters parsing errors, the handler should return false or an error, respectively.
- **No Matching Handler**: ProcessXML returns an error if no handler can process the XML. Always check the return value.

## Testing Your Integration

To verify your setup:

1. Create a test handler (like OrderHandler above).
2. Register it and process sample XML.
3. Check logs or outputs to confirm the handler executed correctly.

You can also run the package's unit tests to ensure the core functionality works:

```bash
go test -v github.com/brandonbryant12/xml-dispatcher
```

## Extending the System

- **Add More Handlers**: Implement additional Handler types for new XML formats and register them.
- **Complex Validation**: Use nested structs in CanHandle to validate deeper XML structures (e.g., `<header><type>`).

Example of nested validation:

```go
func (o *OrderHandler) CanHandle(xmlData []byte) bool {
	type Header struct {
		Type string `xml:"type"`
	}
	type Root struct {
		XMLName xml.Name `xml:"order"`
		Header  Header   `xml:"header"`
	}
	var root Root
	err := xml.Unmarshal(xmlData, &root)
	return err == nil && root.XMLName.Local == "order" && root.Header.Type == "priority"
}
```

## Troubleshooting

- **Build Errors**: Ensure the package is correctly imported and fetched (go get).
- **Handler Not Triggering**: Verify CanHandle logic matches your XML structure.
- **Performance**: For large XML files, xml.Unmarshal loads everything into memory. Consider streaming with xml.Decoder for optimization (requires custom handler logic).

## Conclusion

The xmldispatcher package is lightweight and flexible, making it ideal for projects needing modular XML processing. By following this guide, you can integrate it into your application, customize it with handlers, and process XML data efficiently.