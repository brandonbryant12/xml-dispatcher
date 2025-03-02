
// <ai_context>
// This file provides example implementations of the Handler interface:
// - ReportHandler: Handles XML with a <report> root element.
// - InvoiceHandler: Handles XML with an <invoice type="sales"> root element.
// These examples demonstrate how to create custom handlers for specific XML formats.
// </ai_context>

package xmlprocessor

import (
    "encoding/xml"
    "fmt"
)

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

// InvoiceHandler handles XML with an <invoice type="sales"> root element.
type InvoiceHandler struct{}

// CanHandle checks if the XML is an <invoice> with type="sales".
func (i *InvoiceHandler) CanHandle(xmlData []byte) bool {
    type Root struct {
        XMLName xml.Name `xml:"invoice"`
        Type    string   `xml:"type,attr"`
    }
    var root Root
    err := xml.Unmarshal(xmlData, &root)
    return err == nil && root.XMLName.Local == "invoice" && root.Type == "sales"
}

// Handle processes the <invoice> XML data.
func (i *InvoiceHandler) Handle(xmlData []byte) error {
    type Invoice struct {
        Amount string `xml:"amount"`
    }
    var invoice Invoice
    if err := xml.Unmarshal(xmlData, &invoice); err != nil {
        return err
    }
    fmt.Println("Processing invoice amount:", invoice.Amount)
    return nil
}
      