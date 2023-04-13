package user

import (
	"context"
	"github.com/isaqib23/golang-coffee-shop-microservices/database"
	"github.com/isaqib23/golang-coffee-shop-microservices/domain"
)

type repository struct {
	db *database.DBClient
}

func (r *repository) FetchByName(ctx context.Context, name string) (*domain.UserEntity, error) {
	/*dbClient := r.db.DB
	dbClient.Exec("YOUR QUERY")*/

	return &domain.UserEntity{
		ID:       "1",
		Name:     "Saqib Rao",
		Password: "123",
	}, nil
}
