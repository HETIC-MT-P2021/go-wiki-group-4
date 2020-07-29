package utils

import (
	"io/ioutil"
	"os"
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"packages.hetic.net/gomvc/models"
)

// InitXLSXStrategy Init
func InitXLSXStrategy() ExportStrategyInterface {
	var XLSXStrategy ExportStrategyInterface

	XLSXStrategy.MIMEType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	XLSXStrategy.ExportArticlesFile = ExportArticleXLSX

	return XLSXStrategy
}

// ExportArticleXLSX create a xlsx file and send it
func ExportArticleXLSX(datas []models.Article) (ExportedContent, error) {
	f := excelize.NewFile()
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "Title")
	f.SetCellValue("Sheet1", "B1", "Content")

	for index, item := range datas {
		// first index = 0 and does not exit in excel
		// second index = 1 is already use before
		index_excel := strconv.Itoa(index + 2)
		f.SetCellValue("Sheet1", "A"+index_excel, item.Title)
		f.SetCellValue("Sheet1", "B"+index_excel, item.Content)
	}

	// Save xlsx file by the given path.
	f.SaveAs("article.xlsx")

	data, err := ioutil.ReadFile("article.xlsx")
	defer os.Remove("article.xlsx")

	return ExportedContent{
		Type: "",
		Data: data,
	}, err
}
