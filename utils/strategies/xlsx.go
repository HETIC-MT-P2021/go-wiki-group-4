package utils

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
	"packages.hetic.net/gomvc/models"
)

// InitXLSXStrategy returns the strategy to create the XLX file
func InitXLSXStrategy() ExportStrategyInterface {
	var XLSXStrategy ExportStrategyInterface

	XLSXStrategy.MIMEType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	XLSXStrategy.ExportArticlesFile = ExportArticleXLSX

	return XLSXStrategy
}

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

	data, err := f.WriteToBuffer()

	if err != nil {
		return exportedXLSXContent, err
	}

	if err != nil {
		return exportedXLSXContent, err
	}

	exportedXLSXContent.Data = data.Bytes()

	return exportedXLSXContent, err
}
