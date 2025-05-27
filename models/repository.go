package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
}

type postgresRepository struct {
	DB *sql.DB
}

func newPostgresRepo(conn *sql.DB) Repository {
	return &postgresRepository{
		DB: conn,
	}
}

type testRepository struct {
	DB *sql.DB
}

func newTestRepo(conn *sql.DB) Repository {
	return &testRepository{
		DB: conn,
	}
}
