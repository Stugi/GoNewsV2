package storage

import (
	"context"
	"fmt"
	"time"
)

type Source struct {
	ID          int
	Link        string
	Name        string
	Description string
}

// Добавление источника новостей, если его нет.
func (s *Storage) AddSource(source *Source) (int, error) {
	var sources []Source
	if value, ok := s.cache.Get("sources"); ok {
		sources = value.([]Source)
	}
	var id int
	// проверяем есть ли источник в кэше
	for _, s := range sources {
		if s.Link == source.Link {
			source.ID = s.ID
			return s.ID, nil
		}
	}

	// добавляем источник в базу
	err := s.db.QueryRow(context.Background(),
		`INSERT INTO source (url, name, description) 
		VALUES ($1, $2, $3) 
		ON CONFLICT DO NOTHING 
		RETURNING id`,
		source.Link,
		source.Name,
		source.Description).Scan(&id)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return 0, err
	}

	// добавляем источник в кэш
	source.ID = id
	err = s.cache.Set(
		"sources",
		append(sources, *source),
		time.Hour*24)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	}

	return id, err
}

func (s *Storage) GetSources() ([]Source, error) {
	var sources []Source
	rows, err := s.db.Query(context.Background(), "SELECT id, url, name, description FROM source")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var source Source
		err = rows.Scan(&source.ID, &source.Link, &source.Name, &source.Description)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return sources, nil
}
