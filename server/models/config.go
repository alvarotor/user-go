package models

import "os"

type Config struct {
	PROJECT_PORT_USER          int    `validate:"required"`
	POSTGRES_HOST              string `validate:"required"`
	POSTGRES_PORT              string `validate:"required"`
	POSTGRES_USER              string `validate:"required"`
	POSTGRES_DB                string `validate:"required"`
	POSTGRES_PASSWORD          string `validate:"required"`
	RandomStringValidation     string `validate:"required"`
	SizeRandomStringValidation int    `validate:"required"`
	RandomStringValidationRefresh string `validate:"required"`
	SizeRandomStringValidationRefresh int    `validate:"required"`
	Issuer                             string `validate:"required"`
	JWTKey                             []byte
	JWTKey_Refresh 	[]byte
	TokenExpirationTime                int    `validate:"required"`
	TokenExpirationTimeRefresh         int    `validate:"required"`
}

func (c *Config) IsLocalENV() bool {
	return os.Getenv("ENV") == "" || os.Getenv("ENV") == "local"
}

func (c *Config) IsDevENV() bool {
	return os.Getenv("ENV") == "dev"
}

func (c *Config) IsProdENV() bool {
	return os.Getenv("ENV") == "prod"
}

func (c *Config) IsStagingENV() bool {
	return os.Getenv("ENV") == "staging"
}