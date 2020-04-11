package cfg

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DBCfg struct {
	User     string
	Password string
	Name     string
	Host     string
	Port     string
}

type AppCfg struct {
	Host   string
	Port   string
	Static string
}

var Db DBCfg
var App AppCfg

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	Db.User = os.Getenv("MYSQL_USER")
	Db.Password = os.Getenv("MYSQL_PASSWORD")
	Db.Name = os.Getenv("MYSQL_DATABASE")
	Db.Host = os.Getenv("MYSQL_HOSTNAME")
	Db.Port = os.Getenv("MYSQL_PORT")

	App.Host = os.Getenv("APP_HOST")
	App.Port = os.Getenv("APP_PORT")
	App.Static = os.Getenv("APP_STATIC")
}
