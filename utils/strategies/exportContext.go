package utils

import (
	"packages.hetic.net/gomvc/models"
)

// ExportStrategy struct is used to define how should the data be exported
type ExportStrategyContext struct {
	selectedStrategy ExportStrategyInterface

	SelectStrategy func(newStrategy ExportStrategyInterface)
	ExportArticles func()
}


// SelectStrategy change the selected strategy
func (c ExportStrategyContext) SelectStrategy(newStrategy ExportStrategyInterface) {
	c.selectedStrategy = newStrategy
}

// ExportArticles create a file with all articles
func (c ExportStrategyContext) ExportArticles() string, error {
	// Fetch articles

	articles, err := models.GetArticles()

	if err != nil {
		return nil, err
	}

	// Run selected strategy

	
	filePath, err := c.selectedStrategy.ExportArticlesFile(articles)

	if err != nil {
		return nil, err
	}

	return filePath, nil
}
