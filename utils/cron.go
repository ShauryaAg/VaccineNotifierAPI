package utils

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

func AddCronJobs(CURRENT_HOST string) {
	c := cron.New()
	c.AddFunc("@every 1h30m", func() {
		fmt.Println("RAN CRON")
		SendVaccineInfo(CURRENT_HOST)
	})
	c.Start()
}
