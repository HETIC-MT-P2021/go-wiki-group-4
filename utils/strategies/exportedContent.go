package utils

const (
	CsvApplicationType  = "text/csv"
	XlsxApplicationType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
)

// ExportedContent define the structure of the content we export
type ExportedContent struct {
	Type string
	Data []byte
}
