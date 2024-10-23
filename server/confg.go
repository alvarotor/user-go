package main

import (
	"log"
	"os"
	"strconv"

	"github.com/alvarotor/user-go/server/model"
)

func checkEnvVarsConf(conf *model.Config) {
	checkEnvVar("PROJECT_PORT")
	checkEnvVar("POSTGRES_HOST")
	checkEnvVar("POSTGRES_DB")
	checkEnvVar("POSTGRES_USER")
	checkEnvVar("POSTGRES_PASSWORD")
	checkEnvVar("POSTGRES_PORT")
	checkEnvVar("RandomStringValidation")
	checkEnvVar("SizeRandomStringValidation")
	checkEnvVar("Issuer")
	checkEnvVar("JWT_KEY")

	PROJECT_PORT, err := strconv.Atoi(os.Getenv("PROJECT_PORT"))
	if err != nil {
		log.Fatalf(`Missing PROJECT_PORT env var`)
	}
	conf.PROJECT_PORT = PROJECT_PORT
	conf.POSTGRES_HOST = os.Getenv("POSTGRES_HOST")
	conf.POSTGRES_DB = os.Getenv("POSTGRES_DB")
	conf.POSTGRES_USER = os.Getenv("POSTGRES_USER")
	conf.POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	conf.POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	conf.RandomStringValidation = os.Getenv("RandomStringValidation")
	SizeRandomStringValidation, err := strconv.Atoi(os.Getenv("SizeRandomStringValidation"))
	if err != nil {
		log.Fatalf(`Missing SizeRandomStringValidation env var`)
	}
	conf.SizeRandomStringValidation = SizeRandomStringValidation
	conf.Issuer = os.Getenv("Issuer")
	conf.JWTKey = []byte(os.Getenv("JWT_KEY"))
}

func checkEnvVar(envVar string) {
	if len(os.Getenv(envVar)) == 0 {
		log.Fatalf(`Missing %s env var`, envVar)
	}
}
