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

func CreateArticle(title string, content string, db *sql.DB) (article, error) {
	sqlStatement := `
	INSERT INTO article (title, content)
	VALUES ($1,$2) RETURNING article_id, title, content;`

	var newArticle article

	row := db.QueryRow(sqlStatement, title)
	err := row.Scan(&newArticle.article_id, &newArticle.title)

	if err != nil {
		return newArticle, err
	}
	return newArticle, nil
}
// GetArticle by ID
func GetArticle(article_id string, db *sql.DB) (article, error) {

	sqlStatement := `SELECT * FROM article WHERE article_id=$1;`

	var article article

	row := db.QueryRow(sqlStatement, article_id)

	var err error

	err = row.Scan(&article.article_id, &article.title)

	switch err {
	case sql.ErrNoRows:
		return article, errors.New("Notfound, no article found for this id")
	case nil:
		return article, nil

	default:
		return article, errors.New("Internal Server error")
	}
}
