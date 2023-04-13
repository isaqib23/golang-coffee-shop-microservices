package user

import (
	"github.com/google/wire"
	"github.com/isaqib23/golang-coffee-shop-microservices/database"
	"github.com/isaqib23/golang-coffee-shop-microservices/domain"
	"sync"
)

var (
	hdl     *handler
	hdlOnce sync.Once

	svc     *service
	svcOnce sync.Once

	repo     *repository
	repoOnce sync.Once

	// ProviderSet => those providers are frequently use together, so provider set will be helpful
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHandler,
		ProvideService,
		ProvideRepository,

		// Service interface depends on Repository interface
		// and Repository provider returns pointer of the repository struct not interface
		// to fix this uses Interface binding is needed to bind an abstract interface to its concrete implementation
		// e.g: Bind *repository to the domain.UserRepository
		wire.Bind(new(domain.UserHandler), new(*handler)),
		wire.Bind(new(domain.UserService), new(*service)),
		wire.Bind(new(domain.UserRepository), new(*repository)),
	)
)

// ProvideHandler => initializer function for creating a single struct
func ProvideHandler(svc domain.UserService) *handler {
	hdlOnce.Do(func() {
		hdl = &handler{svc: svc}
	})
	return hdl
}

// ProvideService => initializer function for creating a single struct
func ProvideService(repo domain.UserRepository) *service {
	svcOnce.Do(func() {
		svc = &service{repo: repo}
	})
	return svc
}

// ProvideRepository => initializer function for creating a single struct
func ProvideRepository(db *database.DBClient) *repository {
	repoOnce.Do(func() {
		repo = &repository{db: db}
	})
	return repo
}
