package utils

import (
	"errors"
	"github.com/go-playground/validator"
	"github.com/prongbang/goenv"
	"gorm.io/gorm"
	"os"
)

var (
	Cfg Config
)

type Config struct {
	DBType    string
	DBDSN     string
	SQLDB     *gorm.DB
	PORT      string
	MODE      string
	USERID    string
	VALIDATOR *validator.Validate
}

func InitSettings() error {
	var err error

	Cfg.MODE = os.Getenv("MODE")
	if Cfg.MODE != "prod" {
		goenv.LoadEnv(".env")
	}

	Cfg.PORT = ":" + os.Getenv("PORT")

	//Initialize DB type, it could be postgres or mongodb
	Cfg.DBType = os.Getenv("DB_TYPE")
	Cfg.USERID = os.Getenv("USER_ID")

	if Cfg.USERID == "" {
		return errors.New("please define inside the .env which is the USER_ID field")
	}

	if Cfg.DBType == "postgres" {
		//Instance database string connection
		Cfg.DBDSN = "postgres://" + os.Getenv("POSTGRES_USER") + ":" + os.Getenv("POSTGRES_PASSWORD") + "@" + os.Getenv("POSTGRES_HOST") + ":" + os.Getenv("POSTGRES_PORT") + "/" + os.Getenv("POSTGRES_DB")
		if Cfg.SQLDB, err = initDB(); err != nil {
			return err
		}

	} else if Cfg.DBType == "mongodb" {

	} else {
		return errors.New("the passed DB_TYPE string, inside the .env file, is missing or invalid")
	}

	//Instance the validator
	Cfg.VALIDATOR = validator.New()

	return nil
}
