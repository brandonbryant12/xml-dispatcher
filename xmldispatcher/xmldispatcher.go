// <ai_context>
// This file defines the core components of the XML dispatcher system:
// - Handler interface: Defines the contract for XML handlers.
// - XMLProcessor struct: Manages a collection of handlers and processes XML data.
// - ReportHandler: An example handler for <report> XML, referenced in tests.
// </ai_context>

package xmldispatcher

import (
	"encoding/xml"
	"fmt"
)

// Handler defines the interface for processing specific XML formats.
type Handler interface {
	CanHandle(xmlData []byte) bool
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

// ReportHandler handles XML with a <report> root element.
type ReportHandler struct{}

// CanHandle checks if the XML has a <report> root.
func (r *ReportHandler) CanHandle(xmlData []byte) bool {
	type Root struct {
		XMLName xml.Name `xml:"report"`
	}
	var root Root
	err := xml.Unmarshal(xmlData, &root)
	return err == nil && root.XMLName.Local == "report"
}

// Handle processes the <report> XML data.
func (r *ReportHandler) Handle(xmlData []byte) error {
	type Report struct {
		Data string `xml:"data"`
	}
	var report Report
	if err := xml.Unmarshal(xmlData, &report); err != nil {
		return err
	}
	fmt.Println("Processing report:", report.Data)
	return nil
}
