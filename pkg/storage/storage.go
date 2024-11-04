package storage

import (
	"context"
	"gonews/v2/cache"

	"github.com/jackc/pgx/v4/pgxpool"
)

const connection = "postgres://postgres:postgres@localhost:5432/gonews?sslmode=disable"

type Storage struct {
	db    *pgxpool.Pool
	cache *cache.Impl
}

type DB interface {
	AddPost(Post) error
	GetPosts(int, int) ([]Post, error)
	AddInfo(Info) error
	AddSource(Source) (int, error)
}

func New(cache *cache.Impl) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), connection)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db:    db,
		cache: cache,
	}
	return &s, nil
}

func (s *Storage) StoreNews(news []Post) error {
	for _, n := range news {
		// добавляем источник в базу
		_, err := s.AddSource(n.Source)
		if err != nil {
			return err
		}

		// добавляем новость в базу
		err = s.AddPost(n)
		if err != nil {
			return err
		}
	}
	return nil
}
