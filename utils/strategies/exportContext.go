package utils

import (
	"packages.hetic.net/gomvc/models"
)

// ExportStrategyContext struct is used to define how should the data be exported
type ExportStrategyContext struct {
	selectedStrategy ExportStrategyInterface
}

// SetStrategy change the selected strategy
func (c *ExportStrategyContext) SetStrategy(newStrategy ExportStrategyInterface) {
	c.selectedStrategy = newStrategy
}

// GetSelectedStrategyMimeType returns the selected strategy mime type
func (c *ExportStrategyContext) GetSelectedStrategyMimeType() string {
	return c.selectedStrategy.MIMEType
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
	exportedFileContent, err = c.selectedStrategy.ExportArticlesFile(articles)

	if err != nil {
		return exportedFileContent, err
	}

	exportedFileContent.Type = c.GetSelectedStrategyMimeType()

	return exportedFileContent, nil
}

// InitExportContext is the function called in main to init the strategies
func InitExportContext() ExportStrategyContext {
	csvStrategy := InitCSVStrategy()
	// initXLSXStrategy()
	var exportStrategies ExportStrategyContext

	exportStrategies.SetStrategy(csvStrategy)

	return exportStrategies
}
