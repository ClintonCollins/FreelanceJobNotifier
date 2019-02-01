package models

import (
	"sync"
	"time"
)

type Configuration struct {
	SearchQueries           []string
	FreelancerEnabled       bool
	StandardOutNotification bool
	DiscordNotifications    bool
	DiscordHook             string
}

type Data struct {
	LastRunTimeUnix int64
	LastRunTime     time.Time
	JobsChannel     chan JobGroup
	Configuration   Configuration
	sync.RWMutex
}

type JobGroup struct {
	Name string
	Jobs []*Job
}

type Job struct {
	URL         string
	Title       string
	Description string
	Created     time.Time
	Query       string
}

func (d *Data) GetRunTimeUnix() int64 {
	d.RLock()
	timeUnix := d.LastRunTimeUnix
	d.RUnlock()
	return timeUnix
}

func (d *Data) GetRunTime() time.Time {
	d.RLock()
	timeTime := d.LastRunTime
	d.RUnlock()
	return timeTime
}

func (d *Data) UpdateLastRunTime() {
	d.Lock()
	d.LastRunTime = time.Now()
	d.LastRunTimeUnix = time.Now().Unix()
	d.Unlock()
	return
}
