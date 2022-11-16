package main

import (
	"errors"
	"github.com/prongbang/goenv"
	"gorm.io/gorm"
	"os"
)

type Config struct {
	DBType string
	DBDSN  string
	SQLDB  *gorm.DB
	PORT   string
	MODE   string
	USERID string
}

func initSettings() error {
	var err error

	config.MODE = os.Getenv("MODE")
	if config.MODE != "prod" {
		goenv.LoadEnv(".env")
	}

	config.PORT = ":" + os.Getenv("PORT")

	//Initialize DB type, it could be postgres or mongodb
	config.DBType = os.Getenv("DB_TYPE")
	config.USERID = os.Getenv("USERID")

	if config.USERID == "" {
		return errors.New("please define inside the .env which is the USERID field")
	}

	if config.DBType == "postgres" {
		//Instance database string connection
		config.DBDSN = "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT") + "/" + os.Getenv("POSTGRES_DB")
		if config.SQLDB, err = initDB(); err != nil {
			return err
		}

	} else if config.DBType == "mongodb" {

	} else {
		return errors.New("the passed DB_TYPE string, inside the .env file, is missing or invalid")
	}

	return nil
}
