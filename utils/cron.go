package utils

import (
	"github.com/robfig/cron/v3"
)

func AddCronJobs() {
	c := cron.New()
	c.AddFunc("@every 1h30m", func() {})
}
