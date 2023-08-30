package repository

import (
	"context"
	"github.com/Ypxd/WebService/internal/models"
	"github.com/Ypxd/WebService/internal/repository/postgres"
	"github.com/Ypxd/WebService/utils"
	"github.com/jmoiron/sqlx"
)

type ItemRepo interface {
	GetByID(ctx context.Context, id []string) ([]models.Items, error)
}

type Repository struct {
	ItemRepo ItemRepo
}

func NewRepo() (*Repository, *sqlx.DB, error) {
	db, err := postgres.Connect(utils.GetConfig().DB)
	if err != nil {
		return nil, nil, err
	}

	itemRepo := postgres.NewItemRepo(db)

	return &Repository{
		ItemRepo: itemRepo,
	}, db, nil
}
