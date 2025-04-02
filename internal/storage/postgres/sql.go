package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(storagePath string) (*Storage, error) {
	const op = "storage.NewStorage"
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &Storage{db: db}, nil
}

func (storage *Storage) GetUserApiSecret(ctx context.Context, userId int64, ApiPublic string) (string, error) {
	const op = "storage.GetUserApiSecret"

	row := storage.db.QueryRow("select privateapi from BybitApis where UserId = $1 and PublicApi = $2 ", userId, ApiPublic)
	var ApiSecret string
	err := row.Scan(&ApiSecret)
	if err != nil {
		return "", fmt.Errorf("%s: %w", op, err)
	}
	return ApiSecret, nil
}
