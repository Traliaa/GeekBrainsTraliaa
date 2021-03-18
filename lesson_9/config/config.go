package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Address string `yaml:"host"`
	Port    int    `yaml:"port"`
	DbHost  string `yaml:"hostDB"`
}

func NewConfigForENV() (*Config, error) {
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

func NewConfigForYaml(filename string) (*Config, error) {
	c := Config{}
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	if !strings.Contains(c.Address, "http://") {
		return nil, fmt.Errorf("Url is not valid %s", c.Address)
	}

	if !strings.Contains(c.DbHost, "postgres://") {
		return nil, fmt.Errorf("DB host is not valid %s", c.DbHost)
	}
	return &c, err
}
