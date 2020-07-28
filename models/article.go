package models

import (
	"database/sql"
	"errors"
)

// Article wiki
type Article struct {
	articleID int
	title     string
	content   string
}

// CreateArticle handle request to add a new article to the db
func CreateArticle(title string, content string) (Article, error) {
	sqlStatement := `
	INSERT INTO article (title, content)
	VALUES ($1,$2) RETURNING article_id, title, content;`

	var newArticle Article

	row := db.QueryRow(sqlStatement, title)
	err := row.Scan(&newArticle.articleID, &newArticle.title)

	if err != nil {
		return newArticle, err
	}
	return newArticle, nil
}

// DeleteArticle from the db
func DeleteArticle(articleID string, title string, content string, db *sql.DB) error {

	unlinkSQL := `
	DELETE * FROM article WHERE articleID=$1 ;`

	unlinkRow := db.QueryRow(unlinkSQL, articleID)
	unlinkErr := unlinkRow.Scan()

	if unlinkErr != nil {
		return unlinkErr
	}

	return nil
}

// GetArticle by ID
func GetArticle(articleID string) (Article, error) {

	sqlStatement := `SELECT * FROM article WHERE article_id=$1;`

	var article Article

	row := db.QueryRow(sqlStatement, articleID)

	var err error

	err = row.Scan(&article.articleID, &article.title, &article.content)

	switch err {
	case sql.ErrNoRows:
		return article, errors.New("Notfound, no article found for this id")
	case nil:
		return article, nil

	default:
		return article, errors.New("Internal Server error")
	}
}
