package storage

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

const connection = "postgresql://postgres:password@localhost:5432/postgres?sslmode=disable"

type Storage struct {
	db *pgxpool.Pool
}

func New() (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

type Post struct {
	ID       int
	Title    string
	Content  string
	PubTime  int64
	Link     string
	SourceID int
}

// Добавление новости
func (s *Storage) AddPost(post Post) error {
	_, err := s.db.Exec(context.Background(),
		"INSERT INTO posts (title, content, pub_time, link, source_id) VALUES ($1, $2, $3, $4, $5)",
		post.Title, post.Content, post.PubTime, post.Link, post.SourceID)
	return err
}

// Получение последних новостей
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
			FROM posts  
			LIMIT $1
		ORDER BY pub_time DESC`, limit)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post Post
		err = rows.Scan(&post.ID, &post.Title, &post.Content, &post.PubTime, &post.Link, &post.SourceID)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

type Info struct {
	ID      int
	Message string
	Time    int
	Type    string
}

// Добавление ошибки
func (s *Storage) AddInfo(info Info) error {
	_, err := s.db.Exec(context.Background(),
		"INSERT INTO info (message, time, type) VALUES ($1, $2, $3)",
		info.Message, info.Time, info.Type)
	return err
}

type Sourse struct {
	ID   int
	Link string
}

// Добавление источника новостей, если его нет.
func (s *Storage) AddSourceIfNotExists(source Sourse) (int, error) {
	var id int
	err := s.db.QueryRow(context.Background(),
		"INSERT INTO sourses (url) VALUES ($1) ON CONFLICT DO NOTHING RETURNING id",
		source.Link).Scan(&id)
	return id, err
}
