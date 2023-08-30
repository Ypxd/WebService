package postgres

import (
	"context"
	"github.com/Ypxd/WebService/internal/models"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type ItemRepository struct {
	db *sqlx.DB
}

func (i *ItemRepository) GetByID(ctx context.Context, id []string) ([]models.Items, error) {
	query := `
				SELECT * FROM public.items
				WHERE id in (?);
`
	query, args, err := sqlx.In(query, id)
	if err != nil {
		return nil, errors.WithMessage(err, "sqlx in")
	}
	query = i.db.Rebind(query)
	resp := make([]models.Items, 0)
	err = i.db.SelectContext(ctx, &resp, query, args...)
	if err != nil {
		return nil, errors.WithMessage(err, "select items by ids")
	}

	return resp, nil
}

func NewItemRepo(db *sqlx.DB) *ItemRepository {
	return &ItemRepository{
		db: db,
	}
}
