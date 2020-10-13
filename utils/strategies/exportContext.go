package utils

import (
	"packages.hetic.net/gomvc/models"
)

// ExportStrategyContext struct is used to define how should the data be exported
type ExportStrategyContext struct {
	SelectedStrategy ExportStrategyInterface
}

// SetStrategy change the selected strategy
func (c *ExportStrategyContext) SetStrategy(newStrategy ExportStrategyInterface) {
	c.SelectedStrategy = newStrategy
}

// GetSelectedStrategyMimeType returns the selected strategy mime type
func (c *ExportStrategyContext) GetSelectedStrategyMimeType() string {
	return c.SelectedStrategy.MIMEType
}

// ExportArticles create a file with all articles
func (c *ExportStrategyContext) ExportArticles() (ExportedContent, error) {
	var exportedFileContent ExportedContent

	articles, articleErr := models.GetArticles()

	if articleErr != nil {
		return exportedFileContent, articleErr
	}

	var err error

	// Run selected strategy
	exportedFileContent, err = c.SelectedStrategy.ExportArticlesFile(articles)

	if err != nil {
		return exportedFileContent, err
	}

	exportedFileContent.Type = c.GetSelectedStrategyMimeType()

	return exportedFileContent, nil
}
