
// <ai_context>
// This file defines the core components of the XML dispatcher system:
// - Handler interface: Defines the contract for XML handlers.
// - XMLProcessor struct: Manages a collection of handlers and processes XML data.
// </ai_context>

package xmlprocessor

import (
    "fmt"
)

// Handler defines the interface for processing specific XML formats.
type Handler interface {
    // CanHandle determines if this handler can process the given XML data.
    CanHandle(xmlData []byte) bool
    // Handle processes the XML data and returns an error if processing fails.
    Handle(xmlData []byte) error
}

// XMLProcessor manages a collection of handlers and processes XML data.
type XMLProcessor struct {
    handlers []Handler
}

// NewXMLProcessor creates a new XMLProcessor instance.
func NewXMLProcessor() *XMLProcessor {
    return &XMLProcessor{}
}

// RegisterHandler adds a handler to the processor.
func (p *XMLProcessor) RegisterHandler(h Handler) {
    p.handlers = append(p.handlers, h)
}

// ProcessXML processes the given XML data by delegating to the appropriate handler.
// Returns an error if no suitable handler is found or if processing fails.
func (p *XMLProcessor) ProcessXML(xmlData []byte) error {
    for _, handler := range p.handlers {
        if handler.CanHandle(xmlData) {
            return handler.Handle(xmlData)
        }
    }
    return fmt.Errorf("no handler found for the given XML")
}
      