package storage

import "github.com/dobb2/zenTotem/internal/entity"

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name Storer
type Storer interface {
	Create(user entity.User) (entity.User, error)
}

//go:generate go run github.com/vektra/mockery/v2@v2.20.2 --name Cacher
type Cacher interface {
	Increment(element entity.Element) (entity.Element, error)
}
