package utils

import (
	"bytes"
	"encoding/csv"
	"strconv"

	"packages.hetic.net/gomvc/models"
)

var articleHeaders = []string{"ArticleID", "Title", "Content"}

// CSVStrategyMIMEType Mime type for CSV documents
var CSVStrategyMIMEType = "text/csv"

// ExportArticleCSV create a csv file and send it
func ExportArticleCSV(articles []models.Article) (ExportedContent, error) {
	var exportedCSVContent ExportedContent

	var content = new(bytes.Buffer)

	writer := csv.NewWriter(content)
	defer writer.Flush()

	err := writer.Write(articleHeaders)

	for i := 0; i < len(articles); i++ {
		thisArticleRow := []string{strconv.Itoa(articles[i].ArticleID), articles[i].Title, articles[i].Content}
		err = writer.Write(thisArticleRow)

		if err != nil {
			return exportedCSVContent, err
		}
	}

	exportedCSVContent.Data = content.Bytes()

	return exportedCSVContent, nil
}
