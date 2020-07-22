package models

// Comment wiki
type Comment struct {
	commentID int
	articleID int
	content   string
}

// CreateComment will add a new comment to a article in DB
func CreateComment(content string, articleID string) (Comment, error) {
	var newComment Comment

	createCommentSQL := `
	INSERT INTO comment (article_id, content)
	VALUES ($1, $2) RETURNING *;`

	commentRow := db.QueryRow(createCommentSQL, content, articleID)
	err := commentRow.Scan(&newComment.commentID, &newComment.articleID, &newComment.content)

	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

// GetComment will get information from a comment in DB
func GetComment(commentID string) (Comment, error) {
	var thisComment Comment

	// Get information from this mailingLit
	commentSQL := `SELECT * FROM comment WHERE comment_id=$1;`

	commentRow := db.QueryRow(commentSQL, commentID)
	err := commentRow.Scan(&thisComment.commentID, &thisComment.articleID, &thisComment.content)

	if err != nil {
		return thisComment, err
	}

	return thisComment, nil
}
