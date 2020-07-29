package utils

import (
	"encoding/csv"
	"os"
	"strconv"

	"packages.hetic.net/gomvc/models"
)

var articleHeaders = []string{"ArticleID", "Title", "Content"}

// CreateCSVFile create a csv file and save it
func CreateCSVFile(articles []models.Article) error {
	file, err := os.Create("aricles.csv")

	if err != nil {
		return err
	}

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(articleHeaders)

	for i := 0; i < len(articles); i++ {
		thisArticleRow := []string{strconv.Itoa(articles[i].ArticleID), articles[i].Title, articles[i].Content}
		err = writer.Write(thisArticleRow)

		if err != nil {
			return err
		}
	}

	return nil
}
