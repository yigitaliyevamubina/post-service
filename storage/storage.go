package storage

import (
	"database/sql"
	"template-post-service/storage/postgres"
	"template-post-service/storage/repo"
)

type IStorage interface {
	Post() repo.PostStorageI
}

type storagePg struct {
	db       *sql.DB
	postRepo repo.PostStorageI
}

func NewStoragePg(db *sql.DB) *storagePg {
	return &storagePg{
		db:       db,
		postRepo: postgres.NewPostRepo(db),
	}
}

func (s storagePg) Post() repo.PostStorageI {
	return s.postRepo
}
