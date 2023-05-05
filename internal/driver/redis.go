package driver

import (
	"fmt"
	"github.com/dobb2/zenTotem/internal/config"
	"github.com/redis/go-redis/v9"
)

func ConnectToRedis(conf config.Config) *redis.Client {
	addrString := fmt.Sprintf("%v:%v", conf.HostRedis, conf.PortRedis)
	client := redis.NewClient(&redis.Options{
		Addr:     addrString,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	return client
}
