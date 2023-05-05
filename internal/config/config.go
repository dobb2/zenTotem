package config

import (
	"flag"
)

type Config struct {
	Host      string
	User      string
	Password  string
	Port      string
	Db        string
	Address   string
	HostRedis string
	PortRedis string
}

func CreateAgentConfig() Config {
	var cfg Config

	flag.StringVar(&cfg.User, "user", "dobb2", "a user postgress")
	flag.StringVar(&cfg.Host, "hostDB", "localhost", "a host of postgress")
	flag.StringVar(&cfg.Password, "password", "root", "a password user postgress")
	flag.StringVar(&cfg.Port, "portDB", "54320", "a port of postgress")
	flag.StringVar(&cfg.Address, "address", "127.0.0.1:8080", "address of postgress")
	flag.StringVar(&cfg.Db, "db", "testWB", "name db")
	flag.StringVar(&cfg.HostRedis, "host", "127.0.0.1", "a host redis")
	flag.StringVar(&cfg.PortRedis, "port", "6380", "a port redis")

	flag.Parse()

	return cfg
}
