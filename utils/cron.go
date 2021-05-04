package utils

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func AddCronJobs() {
	c := cron.New()
	c.AddFunc("@every 1h30m", func() {
		fmt.Println("RAN CRON")
		SendVaccineInfo()
	})
	c.Start()
}
