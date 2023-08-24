package cronjob

import (
	"time"

	"github.com/go-co-op/gocron"
)

type CronJob struct {
	time   int
	handle func()
	s      *gocron.Scheduler
}

func NewCronJob(timeInSeconds int, handle func()) *CronJob {
	s := gocron.NewScheduler(time.UTC)
	cronJob := &CronJob{
		time:   timeInSeconds,
		handle: handle,
		s:      s,
	}
	return cronJob
}

func (c *CronJob) Start() {
	c.s.Every(c.time).Seconds().Do(c.handle)
	c.s.StartAsync()
}

func (c *CronJob) Stop() {
	c.s.Stop()
}

func (c *CronJob) Restart() {
	c.Stop()
	c.Start()
}

func (c *CronJob) SetTime(timeInSeconds int) {
	c.time = timeInSeconds
}
