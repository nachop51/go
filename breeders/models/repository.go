package models

import "database/sql"

type Repository interface {
	AllDogBreeds() ([]*DogBreed, error)
	GetBreedByName(string) (*DogBreed, error)
}

type postgresRepo struct {
	db *sql.DB
}

func newPostgresRepo(conn *sql.DB) Repository {
	return &postgresRepo{
		db: conn,
	}
}

type testRepository struct {
	db *sql.DB
}

func newTestRepo(*sql.DB) Repository {
	return &testRepository{
		db: nil,
	}
}

func (repo *testRepository) AllDogBreeds() ([]*DogBreed, error) {
	return nil, nil
}

func (repo *testRepository) GetBreedByName(string) (*DogBreed, error) {
	return nil, nil
}
