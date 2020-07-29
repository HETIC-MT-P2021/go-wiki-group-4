package utils

import (
	"packages.hetic.net/gomvc/models"
)

// ExportStrategy struct is used to define how should the data be exported
type ExportStrategyInterface struct {
	ExportArticlesFile func(articles []models.Article)
}
