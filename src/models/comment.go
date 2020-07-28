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
	var thisComment Comment

	commentSQL := `SELECT * FROM comment WHERE comment_id=?;`

	commentRow := db.QueryRow(commentSQL, commentID)
	err := commentRow.Scan(&thisComment.CommentID, &thisComment.ArticleID, &thisComment.Content)

	if err != nil {
		return thisComment, err
	}

	return thisComment, nil
}
