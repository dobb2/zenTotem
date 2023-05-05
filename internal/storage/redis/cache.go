package redis

import (
	"context"
	"github.com/dobb2/zenTotem/internal/entity"
	"github.com/redis/go-redis/v9"
)

type Cacher struct {
	db *redis.Client
}

func Create(db *redis.Client) *Cacher {
	return &Cacher{db: db}
}

func (c Cacher) Increment(element entity.Element) (entity.Element, error) {
	ctx := context.TODO()
	res := c.db.IncrBy(ctx, element.Key, element.Value)
	var outElement entity.Element

	outElement.Value = res.Val()
	return outElement, nil

}
