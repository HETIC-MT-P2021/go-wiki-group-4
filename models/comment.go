package models

// Comment wiki
type Comment struct {
	CommentID int
	ArticleID int
	Content   string
}

// CreateComment will add a new comment to a article in DB
func CreateComment(content string, articleID string) (Comment, error) {
	var newComment Comment

	createCommentSQL := `
	INSERT INTO comment (article_id, content)
	VALUES (?, ?);`

	_, err := db.Exec(createCommentSQL, content, articleID)

	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

// GetComment will get information from a comment in DB
func GetComment(commentID string) (Comment, error) {
	var Comment Comment

	commentSQL := `SELECT * FROM comment WHERE comment_id=?;`

	commentRow := db.QueryRow(commentSQL, commentID)
	err := commentRow.Scan(&Comment.CommentID, &Comment.ArticleID, &Comment.Content)

	if err != nil {
		return Comment, err
	}

	return Comment, nil
}
