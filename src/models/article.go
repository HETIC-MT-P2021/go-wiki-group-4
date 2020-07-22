package models
<<<<<<< Updated upstream
=======

import (
"database/sql"

"github.com/gin-gonic/gin"
)
// article wiki
type article struct {
	article_id      int
	title           string
	content         string
}

func CreateArticle(title string, db *sql.DB) (article, error) {
	sqlStatement := `
	INSERT INTO article (title)
	VALUES ($1) RETURNING article_id, title;`

	var newArticle article

	row := db.QueryRow(sqlStatement, title)
	err := row.Scan(&newArticle.article_id, &newArticle.title)

	if err != nil {
		return newArticle, err
	}
	return newArticle, nil
}
>>>>>>> Stashed changes
