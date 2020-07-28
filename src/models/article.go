package models

import (
	"database/sql"
	"errors"
)

// Article wiki
type Article struct {
	ArticleID int
	Title     string
	Content   string
}

// CreateArticle handle request to add a new article to the db
func CreateArticle(Title string, Content string) (Article, error) {
	sqlStatement := `
	INSERT INTO article (Title, Content)
	VALUES ($1,$2) RETURNING article_id, Title, Content;`

	var newArticle Article

	row := db.QueryRow(sqlStatement, Title)
	err := row.Scan(&newArticle.ArticleID, &newArticle.Title)

	if err != nil {
		return newArticle, err
	}
	return newArticle, nil
}

// DeleteArticle by Id
func DeleteArticle(ArticleID string, Title string, Content string, db *sql.DB) error {

	unlinkSQL := `
	DELETE * FROM article WHERE ArticleID=$1 ;`

	customerRow := db.QueryRow(unlinkSQL, ArticleID)
	unlinkErr := customerRow.Scan()

	if unlinkErr != nil {
		return unlinkErr
	}

	return nil
}

// GetArticle by ID
func GetArticle(ArticleID string) (Article, error) {

	sqlStatement := `SELECT * FROM article WHERE article_id=$1;`

	var article Article

	row := db.QueryRow(sqlStatement, ArticleID)

	var err error

	err = row.Scan(&article.ArticleID, &article.Title)

	switch err {
	case sql.ErrNoRows:
		return article, errors.New("Notfound, no article found for this id")
	case nil:
		return article, nil

	default:
		return article, errors.New("Internal Server error")
	}
}
