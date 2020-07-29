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

//  create a xlsx file and send it
func ExportArticleXLSX(datas []models.Article) (ExportedContent, error) {
	f := excelize.NewFile()
	// Set value of a cell.
	f.SetCellValue("Sheet", "A1", "Title")
	f.SetCellValue("Sheet", "B1", "Content")

	for index, item := range datas {

		index_excel := strconv.Itoa(index + 2)
		f.SetCellValue("Sheet", "A"+index_excel, item.Title)
		f.SetCellValue("Sheet", "B"+index_excel, item.Content)
	}

	// Save xlsx file
	f.SaveAs("article.xlsx")

	data, err := ioutil.ReadFile("article.xlsx")
	defer os.Remove("article.xlsx")

	return ExportedContent{
		Type: "",
		Data: data,
	}, err
}
