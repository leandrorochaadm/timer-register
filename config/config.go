package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	//	StringConnectionDB with Mysql
	StringConnectionDB = ""

	//	Port
	Port = 0
)

// Load iniciabile variable
func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	StringConnectionDB = fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))
}
