package storage

import (
	"context"
	"fmt"
)

type Post struct {
	ID         int
	Title      string
	Content    string
	PubTime    int
	Link       string
	Source     *Source
	ExternalID string `json:"-"`
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
			post.id as id, 
			title, 
			content, 
			pub_time, 
			link,

			source.id as source_id,
			source.name as source_name,
			source.url as source_link,
			source.description as source_description
			
		FROM post
		LEFT JOIN source ON post.source_id = source.id 
		ORDER BY pub_time DESC
		LIMIT $1`, limit)

	defer rows.Close()

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return nil, err
	}

	for rows.Next() {
		var post Post
		post.Source = &Source{} // initialize the post.Source field
		err = rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PubTime,
			&post.Link,

			&post.Source.ID,
			&post.Source.Name,
			&post.Source.Link,
			&post.Source.Description,
		)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
