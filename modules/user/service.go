package user

import (
	"context"
	"github.com/isaqib23/golang-coffee-shop-microservices/domain"
)

type service struct {
	repo domain.UserRepository
}

func (s *service) FetchByName(ctx context.Context, name string) (*domain.User, error) {
	resp, _ := s.repo.FetchByName(ctx, name)
	return &domain.User{
		ID:   resp.ID,
		Name: resp.Name,
	}, nil
}
