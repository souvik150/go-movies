package repository

import (
	"database/sql"
	"github.com/souvik150/go-movies-backend/internal/models"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*models.Movie, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id int) (*models.User, error)
}

// sudo docker-compose up -d
