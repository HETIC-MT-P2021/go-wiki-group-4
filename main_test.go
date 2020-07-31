package main

import (
	"testing"

	utils "packages.hetic.net/gomvc/utils/strategies"
)

func TestInitExportContext(t *testing.T) {
	exportContext := utils.InitExportContext()

	if exportContext.GetSelectedStrategyMimeType() != "text/csv" {
		t.Error("Expected selected strategy to be CSV at init")
	}
}

func TestSetXLSXStrategies(t *testing.T) {
	exportContext := utils.InitExportContext()

	xlsxStrategy := utils.InitXLSXStrategy()

	exportContext.SetStrategy(xlsxStrategy)

	if exportContext.GetSelectedStrategyMimeType() != "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" {
		t.Error("Expected selected strategy to be XLSX after SetStrategy")
	}
}
