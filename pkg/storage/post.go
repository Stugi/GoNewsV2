package storage

import (
	"context"
)

type Post struct {
	ID         int
	Title      string
	Content    string
	PubTime    int `db:"pub_time"`
	Link       string
	Source     *Source
	ExternalID string
}

// Добавление новости.
func (s *Storage) AddPost(post Post) error {
	_, err := s.db.Exec(context.Background(),
		`INSERT INTO post (title, content, pub_time, link, source_id, external_id) 
		VALUES ($1, $2, $3, $4, $5, $6) 
		ON CONFLICT DO NOTHING`,
		post.Title, post.Content, post.PubTime, post.Link, post.Source.ID, post.ExternalID)
	return err
}

// Получение последних новостей.
func (s *Storage) GetPosts(sourceID int, limit int) ([]Post, error) {
	var posts []Post
	rows, err := s.db.Query(context.Background(),
		`SELECT 
			id, 
			title, 
			content, 
			pub_time, 
			link, 
			source_id 
			FROM post  
			LIMIT $1
		ORDER BY pub_time DESC`, limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.PubTime, &post.Link, &post.Source.ID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
