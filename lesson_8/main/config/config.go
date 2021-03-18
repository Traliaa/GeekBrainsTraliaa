package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Address string
	Port    int
	DbHost  string
}

func NewConfig() (*Config, error) {
	c := Config{}
	address := os.Getenv("HOST")
	dbHost := os.Getenv("DBHOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	c.Port = port
	if strings.Contains(address, "http://") {
		c.Address = address
	} else {
		return nil, fmt.Errorf("Url is not valid %s", address)
	}

	if strings.Contains(dbHost, "postgres://") {
		c.DbHost = dbHost
	} else {
		return nil, fmt.Errorf("DB host is not valid %s", dbHost)
	}
	return &c, err
}
