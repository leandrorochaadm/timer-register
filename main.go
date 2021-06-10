package main

import (
	"fmt"
	"github.com/leandrorochaadm/time-register/config"
	"github.com/leandrorochaadm/time-register/database"
	date_deadline "github.com/leandrorochaadm/time-register/date-deadline"
	"github.com/leandrorochaadm/time-register/router"
)

func main() {
	config.Load()
	database.InitialiseDBConnection()
	router.Generate()

	fmt.Println("API executando na porta ", config.Port)

	date_deadline.DeadlineYear()
	date_deadline.DeadlineLife()
	fmt.Println("---------")
	//fmt.Println(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour*24*(365/2)))

}
