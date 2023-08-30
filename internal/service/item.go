package service

import (
	"context"
	"github.com/Ypxd/WebService/internal/models"
	"github.com/Ypxd/WebService/internal/repository"
	"github.com/jmoiron/sqlx"
)

type ItemService struct {
	repo *repository.Repository
	conn *sqlx.DB
}

func (i *ItemService) GetItem(ctx context.Context, ids []string) ([]models.Items, error) {
	return i.repo.ItemRepo.GetByID(ctx, ids)
}

func NewItemService(repo *repository.Repository, conn *sqlx.DB) *ItemService {
	return &ItemService{
		repo: repo,
		conn: conn,
	}
}
