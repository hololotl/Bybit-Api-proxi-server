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
	db, err := sql.Open("postgres", storagePath)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return &Storage{db: db}, nil
}

func (storage *Storage) GetUserApiSecret(ctx context.Context, userId int64, ApiPublic string) (string, error) {
	row := storage.db.QueryRow("select privateapi from BybitApis where UserId = $1 and PublicApi = $2 ", userId, ApiPublic)
	var ApiSecret string
	err := row.Scan(&ApiSecret)
	if err != nil {
		return "", err
	}
	return ApiSecret, nil
}
