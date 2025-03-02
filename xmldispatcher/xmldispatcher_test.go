// <ai_context>
// This file contains unit tests for the xmldispatcher package:
// - Tests the XMLProcessor with a custom TestCodeChangesHandler.
// - Verifies behavior when no handler is found.
// - Tests the ReportHandler's CanHandle method.
// </ai_context>

package xmldispatcher

import (
	"encoding/xml"
	"testing"
)

// TestCodeChangesHandler is a test handler for a specific XML structure.
type TestCodeChangesHandler struct {
	parsedBranchName string
}

// CanHandle checks if the XML matches the expected format for code changes.
func (t *TestCodeChangesHandler) CanHandle(xmlData []byte) bool {
	type Root struct {
		XMLName xml.Name `xml:"code_changes"`
	}
	var root Root
	err := xml.Unmarshal(xmlData, &root)
	return err == nil && root.XMLName.Local == "code_changes"
}

// Handle parses the code changes XML and stores data for testing.
func (t *TestCodeChangesHandler) Handle(xmlData []byte) error {
	type CodeChanges struct {
		BranchName string `xml:"branch_name"`
	}
	var cc CodeChanges
	if err := xml.Unmarshal(xmlData, &cc); err != nil {
		return err
	}
	t.parsedBranchName = cc.BranchName
	return nil
}

// TestProcessXML tests processing XML with a registered handler.
func TestProcessXML(t *testing.T) {
	processor := NewXMLProcessor()
	handler := &TestCodeChangesHandler{}
	processor.RegisterHandler(handler)

	// Simulate XML data for testing
	xmlData := []byte(`<code_changes><branch_name>feature/update-docs</branch_name></code_changes>`)

	err := processor.ProcessXML(xmlData)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if handler.parsedBranchName != "feature/update-docs" {
		t.Errorf("Expected parsedBranchName to be 'feature/update-docs', got '%s'", handler.parsedBranchName)
	}
}

// TestNoHandlerFound tests the case where no handler matches the XML.
func TestNoHandlerFound(t *testing.T) {
	processor := NewXMLProcessor()
	xmlData := []byte(`<unknown></unknown>`)
	err := processor.ProcessXML(xmlData)
	if err == nil {
		t.Error("Expected an error, got nil")
	}
	if err.Error() != "no handler found for the given XML" {
		t.Errorf("Expected error message 'no handler found for the given XML', got '%s'", err.Error())
	}
}

// TestReportHandlerCanHandle tests the ReportHandler's CanHandle method.
func TestReportHandlerCanHandle(t *testing.T) {
	handler := &ReportHandler{}
	xmlData := []byte(`<report></report>`)
	if !handler.CanHandle(xmlData) {
		t.Error("Expected CanHandle to return true for report XML")
	}
	xmlDataInvalid := []byte(`<invoice></invoice>`)
	if handler.CanHandle(xmlDataInvalid) {
		t.Error("Expected CanHandle to return false for non-report XML")
	}
}
