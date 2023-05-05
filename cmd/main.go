package main

import (
	"github.com/dobb2/zenTotem/internal/config"
	"github.com/dobb2/zenTotem/internal/driver"
	"github.com/dobb2/zenTotem/internal/handler"
	"github.com/dobb2/zenTotem/internal/logging"
	"github.com/dobb2/zenTotem/internal/storage/postgres"
	"github.com/dobb2/zenTotem/internal/storage/redis"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	logger := logging.CreateLogger()
	conf := config.CreateAgentConfig()
	// connect to postgres
	db, err := driver.ConnectToPostgres(conf)
	if err != nil {
		logger.Fatal().Err(err).Msg("Cannot connect to postgres")
	}
	defer db.Close()

	// connect to postgres
	dbcache := driver.ConnectToRedis(conf)
	defer dbcache.Close()

	cache := redis.Create(dbcache)
	postgrestore := postgres.Create(db)
	handler := handler.New(postgrestore, cache, logger)

	logger.Info().Msg("All connect is done")

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Compress(5))

	r.Post("/redis/incr", handler.IncrementVal)
	r.Post("/sign/hmacsha512", handler.PostSign)
	r.Post("/postgres/users", handler.CreateUser)

	if err := http.ListenAndServe(conf.Address, r); err != nil {
		logger.Fatal().Err(err).Msg("server failed")
	}
}
