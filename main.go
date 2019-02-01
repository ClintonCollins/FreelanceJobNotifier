package main

import (
	"FreelanceJobNotifier/models"
	"FreelanceJobNotifier/scraper"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var (
	doneMain = make(chan struct{})
	data     = new(models.Data)
	closeChan = make(chan os.Signal, 1)
)

func init() {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	configFile, err := ioutil.ReadFile(currentDirectory + "/config.toml")
	if err != nil {
		log.Fatal(err)
	}
	_, err = toml.Decode(string(configFile), &data.Configuration)
	if err != nil {
		log.Fatal(err)
	}
}

func gracefulShutdown() {
	<-closeChan
	doneMain <- struct{}{}
	fmt.Println("Shutting down...")
}

func main() {
	signal.Notify(closeChan, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go gracefulShutdown()
	data.UpdateLastRunTime()
	data.JobsChannel = make(chan models.JobGroup)
	if data.Configuration.FreelancerEnabled {
		go scraper.QueryFreelancerJob(data)
	}
	go scraper.HandleJobs(data)
	fmt.Println("Fully started.")
	<-doneMain
}
