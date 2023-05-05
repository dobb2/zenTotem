package storage

import "github.com/dobb2/zenTotem/internal/entity"

type Storer interface {
	Create(user entity.User) (entity.User, error)
}

type Cacher interface {
	Increment(element entity.Element) (entity.Element, error)
}
