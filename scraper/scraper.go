package scraper

import (
	"FreelanceJobNotifier/discord"
	"FreelanceJobNotifier/freelancer"
	"FreelanceJobNotifier/models"
	"FreelanceJobNotifier/standard_out"
	"time"
)

func QueryFreelancerJob(data *models.Data) {
	throttle := time.Tick(3 * time.Minute)
	for {
		freelancer.QueryFreelancer(data)
		<-throttle
	}
}

func HandleJobs(data *models.Data) {
	for group := range data.JobsChannel {
		if data.Configuration.DiscordNotifications {
			discord.HandleJobGroup(data, group)
		}
		if data.Configuration.StandardOutNotification {
			standard_out.HandleJobGroups(group)
		}
	}
}