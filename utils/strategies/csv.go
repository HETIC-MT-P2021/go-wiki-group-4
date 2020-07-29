package utils

import (
	"encoding/csv"
	"io/ioutil"
	"os"
	"strconv"

	"packages.hetic.net/gomvc/models"
)

var articleHeaders = []string{"ArticleID", "Title", "Content"}

func initCSVStrategy() ExportStrategyInterface {
	var CSVStrategy ExportStrategyInterface

	CSVStrategy.MIMEType = "text/csv"
	CSVStrategy.ExportArticlesFile = ExportArticleCSV

	return CSVStrategy
}

// ExportArticleCSV create a csv file and send it
func ExportArticleCSV(articles []models.Article) ([]byte, error) {
	file, err := os.Create("article.csv")

	if err != nil {
		return ioutil.ReadFile("article.csv")
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(articleHeaders)

	for i := 0; i < len(articles); i++ {
		thisArticleRow := []string{strconv.Itoa(articles[i].ArticleID), articles[i].Title, articles[i].Content}
		err = writer.Write(thisArticleRow)

		if err != nil {
			return ioutil.ReadFile("article.csv")
		}
	}
	defer os.Remove("article.csv")
	return ioutil.ReadFile("article.csv")
}
