package utils

import (
	"packages.hetic.net/gomvc/models"
)

// ExportStrategy struct is used to define how should the data be exported
type ExportStrategy struct {
	SelectedStrategyName string
}

// SelectStrategy change the selected strategy
func (s ExportStrategy) SelectStrategy(strategyName string) {
	s.SelectedStrategyName = strategyName
}

// ExportArticles create a file with all articles
func (s ExportStrategy) ExportArticles() string, error {
	// Fetch articles

	articles, err := models.GetArticles()

	if err != nil {
		return nil, err
	}

	// Run selected strategy

	if s.SelectedStrategyName == "XLSX" {
		fileName, err := createXLSXFile(articles)
	} else {
		fileName, err := CreateCSVFile(articles)
	}

	return fileName, nil
}
