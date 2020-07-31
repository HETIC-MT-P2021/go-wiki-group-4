package main

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"packages.hetic.net/gomvc/models"
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

func TestENVExist(t *testing.T) {
	env, _ := godotenv.Read(".env")

	if len(env["DB_PASSWORD"]) == 0 {
		t.Error("Exepected .env to have DB_PASSWORD")
	}
}

func TestExportArticleInCSVFile(t *testing.T) {
	env, _ := godotenv.Read(".env")

	models.ConnectToDB(env["DB_HOST"], env["DB_NAME"], env["DB_USER"], env["DB_PASSWORD"], env["DB_PORT"])

	exportContext := utils.InitExportContext()

	exportContent, err := exportContext.ExportArticles()

	if err != nil {
		fmt.Println(err)
		t.Error("Expected no error at file creation")
	}

	if exportContent.Type != "text/csv" {
		t.Error("Expected file to be of type CSV")
	}

	if len(exportContent.Data) == 0 {
		t.Error("Expected file to have data")
	}
}
