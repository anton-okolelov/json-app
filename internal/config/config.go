package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const DEFAULT_PORT = "8080"

type DBConf struct {
	Host     string
	Port     uint16
	User     string
	Password string
	DBName   string
}

type Config struct {
	DBConf DBConf
	Port   int
}

// LoadFromEnv - initialize configs
func LoadFromEnv() (Config, error) {
	c := Config{}

	var err error

	httpBind := os.Getenv("HTTP_BIND")
	if len(httpBind) == 0 {
		fmt.Println("port is empty, use default port")
		httpBind = DEFAULT_PORT
	}

	c.Port, err = strconv.Atoi(httpBind)
	if err != nil {
		return c, err
	}

	c.DBConf, err = loadDBConf()
	if err != nil {
		return c, fmt.Errorf("cannot read DB Config: %w", err)
	}

	return c, nil
}

func loadDBConf() (DBConf, error) {
	var c DBConf
	dbHost := os.Getenv("DB_HOST")
	if len(dbHost) == 0 {
		return c, errors.New("env DB_HOST is empty")
	}
	c.Host = dbHost

	dbStrPort := os.Getenv("DB_PORT")
	if len(dbStrPort) == 0 {
		return c, errors.New("env DB_PORT is empty")
	}

	dbPort, err := strconv.ParseInt(dbStrPort, 10, 64)
	if err != nil {
		return c, err
	}

	c.Port = uint16(dbPort)

	dbUser := os.Getenv("DB_USER")
	if len(dbUser) == 0 {
		return c, errors.New("env DB_USER is empty")
	}
	c.User = dbUser

	dbPass := os.Getenv("DB_PASS")
	if len(dbPass) == 0 {
		fmt.Println("env DB_PASS is empty. Use empty value")
	}
	c.Password = dbPass

	dbName := os.Getenv("DB_NAME")
	if len(dbName) == 0 {
		return c, errors.New("env DB_NAME is empty")
	}
	c.DBName = dbName
	return c, nil
}
