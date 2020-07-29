package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"packages.hetic.net/gomvc/models"
	"strconv"
)
/* Example
func utils() {
	f := excelize.NewFile()
	// Create a new sheet.
	index := f.NewSheet("Sheet2")
	// Set value of a cell.
	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save xlsx file by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}*/
type ExportXLSX struct {}

func (exportXLSX ExportXLSX) Export(datas []model.Article) (error) {
	f := excelize.NewFile()
	// Set value of a cell.
	f.SetCellValue("Sheet1", "A1", "Title")
	f.SetCellValue("Sheet1", "B1", "Content")

	for index, item := range datas {
		// first index = 0 and does not exit in excel
		// second index = 1 is already use before
		index_excel := strconv.Itoa(index + 2)
		f.SetCellValue("Sheet1", "A" + index_excel, item.Title)
		f.SetCellValue("Sheet1", "B" + index_excel, item.Content)
	}

	// Save xlsx file by the given path.
	if err := f.SaveAs("assets/xlsx/articles_" + formatCurrentDate() + ".xlsx"); err != nil {
		println(err.Error())
		return err
	}

	return nil
}