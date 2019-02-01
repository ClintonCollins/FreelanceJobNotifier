package discord

import (
	"FreelanceJobNotifier/models"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type Message struct {
	Content  string  `json:"content"`
	Username string  `json:"username"`
	Embeds   []Embed `json:"embeds"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type Embed struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	URL         string    `json:"url"`
	Color       int       `json:"color"`
	Timestamp   time.Time `json:"timestamp"`
	Footer      struct {
		IconURL string `json:"icon_url"`
		Text    string `json:"text"`
	} `json:"footer"`
	Thumbnail struct {
		URL string `json:"url"`
	} `json:"thumbnail"`
	Image struct {
		URL string `json:"url"`
	} `json:"image"`
	Author struct {
		Name    string `json:"name"`
		URL     string `json:"url"`
		IconURL string `json:"icon_url"`
	} `json:"author"`
	Fields []Field `json:"fields"`
}

func HandleJobGroup(data *models.Data, group models.JobGroup) {
	fields := createDiscordEmbedJobs(group.Jobs)
	if len(fields) > 0 {
		for len(fields) > 25 {
			var fieldsToSend []Field
			fieldsToSend, fields = fields[:25], fields[26:]
			executeDiscordWebhook(data, fieldsToSend, group.Name)
			time.Sleep(time.Second * 5)
		}
		if len(fields) <= 25 {
			executeDiscordWebhook(data, fields, group.Name)
		}
	}
}

func createDiscordEmbedJobs(jobs []*models.Job) []Field {
	var embedFields []Field
	for _, j := range jobs {
		field := Field{}
		field.Name = fmt.Sprintf("**%s**", j.Title)
		field.Value = j.URL
		embedFields = append(embedFields, field)
	}
	return embedFields
}

func executeDiscordWebhook(data *models.Data, fields []Field, q string) {
	rand.Seed(time.Now().UTC().UnixNano())
	newMessage := Message{}
	newDiscordEmbed := Embed{}
	newDiscordEmbed.Title = "New Jobs"
	newDiscordEmbed.Timestamp = time.Now()
	newDiscordEmbed.Fields = fields
	newDiscordEmbed.Color = rand.Intn(16777215)
	newDiscordEmbed.Title = fmt.Sprintf("-=- %s jobs -=-", strings.Title(q))
	newMessage.Embeds = append(newMessage.Embeds, newDiscordEmbed)
	discordJSON, err := json.Marshal(newMessage)
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("payload_json", string(discordJSON))
	_, err = http.PostForm(data.Configuration.DiscordHook, params)
	if err != nil {
		log.Fatal(err)
	}
}
