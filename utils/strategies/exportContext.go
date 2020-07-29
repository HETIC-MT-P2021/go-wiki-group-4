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
	articles, err := models.GetArticles()

	if err != nil {
		// return file, err
	}

	// Run selected strategy
	exportedContent, err := c.selectedStrategy.ExportArticlesFile(articles)

	return exportedContent, err
}

// InitStrategies is the function called in main to init the strategies
func InitStrategies() ExportStrategyContext {
	csvStrategy := InitCSVStrategy()
	// initXLSXStrategy()
	var exportStrategies ExportStrategyContext

	exportStrategies.SetStrategy(csvStrategy)

	return exportStrategies
}
