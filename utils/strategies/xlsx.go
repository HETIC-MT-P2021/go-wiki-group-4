package utils

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"packages.hetic.net/gomvc/models"
)

// XLSXStrategyMIMEType Mime type for XLSX documents
var XLSXStrategyMIMEType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

// ExportArticleXLSX creates a xlsx file with all aricles
func ExportArticleXLSX(datas []models.Article) (ExportedContent, error) {
	var exportedXLSXContent ExportedContent

	f := excelize.NewFile()
	// Set value of a cell.
	f.SetCellValue("Sheet", "A1", "ArticleID")
	f.SetCellValue("Sheet", "B1", "Title")
	f.SetCellValue("Sheet", "C1", "Content")

	for index, item := range datas {

		indexExcel := strconv.Itoa(index + 2)
		f.SetCellValue("Sheet", "A"+indexExcel, item.ArticleID)
		f.SetCellValue("Sheet", "B"+indexExcel, item.Title)
		f.SetCellValue("Sheet", "C"+indexExcel, item.Content)
	}

	// Save xlsx file
	err := f.SaveAs("article.xlsx")

	if err != nil {
		return exportedXLSXContent, err
	}

	data, err := ioutil.ReadFile("article.xlsx")
	defer os.Remove("article.xlsx")

	if err != nil {
		return exportedXLSXContent, err
	}

	exportedXLSXContent.Data = data

	return exportedXLSXContent, err
}
