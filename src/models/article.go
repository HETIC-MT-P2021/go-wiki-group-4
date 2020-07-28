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
func CreateArticle(title string, content string) (Article, error) {
	insertSQLStatement := `
	INSERT INTO article (title, content)
	VALUES (?, ?);`

	var newArticle Article

	_, insertErr := db.Exec(insertSQLStatement, title, content)

	if insertErr != nil {
		return newArticle, insertErr
	}

	sqlStatement := `
	SELECT article_id FROM article WHERE title=?;`

	row := db.QueryRow(sqlStatement, title)
	err := row.Scan(&newArticle.ArticleID)

	newArticle.Content = content
	newArticle.Title = title

	if err != nil {
		return newArticle, err
	}

	return newArticle, nil
}

// DeleteArticle from the db
func DeleteArticle(articleID string, title string, content string, db *sql.DB) error {

	unlinkSQL := `
	DELETE * FROM article WHERE articleID=? ;`

	unlinkRow := db.QueryRow(unlinkSQL, articleID)
	unlinkErr := unlinkRow.Scan()

	if unlinkErr != nil {
		return unlinkErr
	}

	return nil
}

// GetArticle by ID
func GetArticle(articleID string) (Article, error) {

	sqlStatement := `SELECT * FROM article WHERE article_id=?;`

	var article Article

	row := db.QueryRow(sqlStatement, articleID)

	var err error

	err = row.Scan(&article.ArticleID, &article.Title, &article.Content)

	switch err {
	case sql.ErrNoRows:
		return article, errors.New("Notfound, no article found for this id")
	case nil:
		return article, nil

	default:
		return article, errors.New("Internal Server error")
	}
}

// GetCommentFromArticleID gets comments for this article
func GetCommentFromArticleID(articleID string) ([]Comment, error) {

	sqlStatement := `SELECT * FROM comment WHERE article_id=?;`

	var comments []Comment

	rows, queryErr := db.Query(sqlStatement, articleID)

	if queryErr != nil {
		return comments, queryErr
	}

	var err error

	for rows.Next() {
		var thisComment Comment

		if err = rows.Scan(&thisComment.commentID, &thisComment.articleID, &thisComment.content); err != nil {
			return comments, err
		}
		comments = append(comments, thisComment)
	}

	switch err {
	case sql.ErrNoRows:
		return comments, errors.New("Notfound, no article found for this id")
	case nil:
		return comments, nil

	default:
		return comments, errors.New("Internal Server error")
	}
}
