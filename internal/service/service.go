package service

import (
	"context"
	"github.com/Ypxd/WebService/internal/models"
	"github.com/Ypxd/WebService/internal/repository"
	"github.com/jmoiron/sqlx"
)

type Item interface {
	GetItem(ctx context.Context, ids []string) ([]models.Items, error)
}

type Service struct {
	Item Item
}

func NewService(repo *repository.Repository, conn *sqlx.DB) *Service {
	item := NewItemService(repo, conn)

	return &Service{
		Item: item,
	}
}
