package main

import (
	"fmt"

	"github.com/leandrorochaadm/time-register/config"
	date_deadline "github.com/leandrorochaadm/time-register/date-deadline"
)

func main() {
	config.Load()
	fmt.Println("API executando na porta ", config.Port)

	date_deadline.DeadlineYear()
	date_deadline.DeadlineLife()

	// router.Generate()

}
