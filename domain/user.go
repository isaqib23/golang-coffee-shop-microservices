// Package domain put your structs and interface
package domain

import (
	"context"
	"net/http"
)

type (
	User struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	UserEntity struct {
		ID       string
		Name     string
		Password string
	}

	UserRepository interface {
		FetchByName(ctx context.Context, name string) (*UserEntity, error)
	}

	UserService interface {
		FetchByName(ctx context.Context, name string) (*User, error)
	}

	UserHandler interface {
		FetchByName() http.HandlerFunc
	}
)
