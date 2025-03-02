package xmldispatcher

import (
	"fmt"
)

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

func (p *XMLProcessor) ProcessXML(xmlData []byte) error {
	for _, handler := range p.handlers {
		if handler.CanHandle(xmlData) {
			return handler.Handle(xmlData)
		}
	}
	return fmt.Errorf("no handler found for the given XML")
}
