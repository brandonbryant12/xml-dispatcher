// Package Structure:
// /Users/brandon/Development/xml-dispatcher
// ├── .github
// │   └── worflows
// │       └── go.yml
// ├── handler.go            // New file containing the Handler interface
// ├── processor.go          // New file containing the XMLProcessor implementation
// ├── report_handler.go     // New file containing the ReportHandler implementation
// ├── xml_dispatcher_test.go // Renamed from xmldispatcher/xmldispatcher_test.go
// ├── go.mod
// └── readme.md

package xmldispatcher

type Handler interface {
	CanHandle(xmlData []byte) bool
	Handle(xmlData []byte) error
}
