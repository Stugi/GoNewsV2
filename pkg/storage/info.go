package storage

import "context"

type Info struct {
	ID      int
	Message string
	Time    int
	Type    string
}

// Добавление ошибки.
func (s *Storage) AddInfo(info Info) error {
	_, err := s.db.Exec(context.Background(),
		"INSERT INTO info (message, time, type) VALUES ($1, $2, $3)",
		info.Message, info.Time, info.Type)
	return err
}
