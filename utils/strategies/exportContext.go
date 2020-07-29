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

// ExportArticles create a file with all articles
func (c *ExportStrategyContext) ExportArticles() ([]byte, error) {
	// Fetch articles
	var file []byte

	articles, err := models.GetArticles()

	if err != nil {
		return file, err
	}

	// Run selected strategy

	file, err = c.selectedStrategy.ExportArticlesFile(articles)

	if err != nil {
		return file, err
	}

	return file, nil
}

// InitStrategies is the function called in main to init the strategies
func InitStrategies() ExportStrategyContext {
	csvStrategy := initCSVStrategy()
	// initXLSXStrategy()
	var exportStrategies ExportStrategyContext

	exportStrategies.SetStrategy(csvStrategy)

	return exportStrategies
}
