package utils

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strconv"

	"packages.hetic.net/gomvc/models"
)

var articleHeaders = []string{"ArticleID", "Title", "Content"}

// InitCSVStrategy returns the strategy to create the CSV file
func InitCSVStrategy() ExportStrategyInterface {
	var CSVStrategy ExportStrategyInterface

	CSVStrategy.MIMEType = "text/csv"
	CSVStrategy.ExportArticlesFile = ExportArticleCSV

	return CSVStrategy
}

// ExportArticleCSV create a csv file and send it
func ExportArticleCSV(articles []models.Article) (ExportedContent, error) {
	var exportedCSVContent ExportedContent

	file, err := os.Create("article.csv")

	if err != nil {
		return exportedCSVContent, err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(articleHeaders)

	for i := 0; i < len(articles); i++ {
		thisArticleRow := []string{strconv.Itoa(articles[i].ArticleID), articles[i].Title, articles[i].Content}
		err = writer.Write(thisArticleRow)

		if err != nil {
			return exportedCSVContent, err
		}
	}

	data, err := ioutil.ReadFile("article.csv")
	defer os.Remove("article.csv")

	if err != nil {
		return exportedCSVContent, err
	}

	exportedCSVContent.Data = data

	return exportedCSVContent, nil
}
