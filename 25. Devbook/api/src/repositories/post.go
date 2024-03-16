package repositories

import (
	"api/src/models"
	"database/sql"
)

// Posts is a struct that represents the posts repository
type Posts struct {
	db *sql.DB
}

// NewPostsRepository creates a posts repository
func NewPostsRepository(db *sql.DB) *Posts {
	return &Posts{db}
}

// CreatePostOnDb creates a new post in the database
func (repository Posts) CreatePostOnDB(post models.Post) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO posts (title, content, author_id) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	var lastInsertedID uint64
	err = stmt.QueryRow(post.Title, post.Content, post.AuthorID).Scan(&lastInsertedID)
	if err != nil {
		return 0, err
	}

	return lastInsertedID, nil
}

// GetPostByIdFromDB returns a post ID from the user ID
func (repository Posts) GetPostByIdFromDB(postID uint64) (models.Post, error) {
	row, err := repository.db.Query(`
		SELECT p.*, u.nick 
		FROM posts p
		INNER JOIN users u 
		ON u.id = p.author_id
		WHERE p.id = $1`, postID)
	if err != nil {
		return models.Post{}, err
	}
	defer row.Close()

	var post models.Post
	if row.Next() {
		if err = row.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorID,
			&post.AuthorNick,
		); err != nil {
			return models.Post{}, err
		}
	}

	return post, nil
}

// GetPostsFromDB returns all posts from the database
func (repository Posts) GetPostsFromDB(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT DISTINCT ON (p.id) p.*, u.nick 
		FROM posts p
		INNER JOIN users u 
		ON u.id = p.author_id
		LEFT JOIN followers f 
		ON p.author_id = f.user_id
		WHERE u.id = $1 OR f.follower_id = $1
		ORDER BY p.id DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorID,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// UpdatePostOnDB updates a post in the database
func (repository Posts) UpdatePostOnDB(postID uint64, post models.Post) error {
	stmt, err := repository.db.Prepare("UPDATE posts SET title = $1, content = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(post.Title, post.Content, postID); err != nil {
		return err
	}

	return nil
}

// DeletePostFromDB deletes a post from the database
func (repository Posts) DeletePostFromDB(postID uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM posts WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

// GetPostByUserIdFromDB returns all posts from a user
func (repository Posts) GetPostsByUserIdFromDB(userID uint64) ([]models.Post, error) {
	rows, err := repository.db.Query(`
		SELECT p.*, u.nick 
		FROM posts p
		INNER JOIN users u 
		ON u.id = p.author_id
		WHERE p.author_id = $1
		ORDER BY p.id DESC`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.Likes,
			&post.CreatedAt,
			&post.AuthorID,
			&post.AuthorNick,
		); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

// LikePostOnDB likes a post in the database
func (repository Posts) LikePostOnDB(postID uint64) error {
	stmt, err := repository.db.Prepare("UPDATE posts SET likes = likes + 1 WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}

// UnlikePostOnDB unlikes a post in the database until it reaches 0
func (repository Posts) UnlikePostOnDB(postID uint64) error {
	stmt, err := repository.db.Prepare(`
	UPDATE posts SET likes = 
	CASE 
		WHEN likes > 0 THEN likes - 1 
		ELSE 0 
	END
	WHERE id = $1
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	if _, err = stmt.Exec(postID); err != nil {
		return err
	}

	return nil
}
