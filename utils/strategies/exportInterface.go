package utils

import (
	"packages.hetic.net/gomvc/models"
)

// ExportStrategyInterface struct is used to define how should the data be exported
type ExportStrategyInterface struct {
	MIMEType string

	ExportArticlesFile (func(articles []models.Article) ([]byte, error))
}
