package main

import (
	"fmt"
	date_deadline "github.com/leandrorochaadm/time-register/date-deadline"
	"github.com/leandrorochaadm/time-register/meta"
	"time"
)

func main() {
	//config.Load()
	//database.InitialiseDBConnection()
	//router.Generate()

	//fmt.Println("API executando na porta ", config.Port)

	fmt.Println("\n-------------------------------------------------------------------------------------")
	date_deadline.DeadlineYear()
	date_deadline.DeadlineLife()
	fmt.Println("\n-------------------------------------------------------------------------------------")
	//fmt.Println(time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC).Add(time.Hour*24*(365/2)))

	metaPecorrida, metaPrazo, _ := meta.Meta(
		time.Now().Add(-time.Minute*60*24*30),
		time.Now().Add(time.Minute*60*24*30),
		time.Now().Add(time.Minute*60*24*7),
		1.0,
		2.0,
		1.4,
	)

	fmt.Println(metaPecorrida, metaPrazo)

}
