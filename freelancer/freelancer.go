package freelancer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"FreelanceJobNotifier/models"
)

type ProjectJSON struct {
	Status string `json:"status"`
	Result struct {
		TotalCount   int         `json:"total_count"`
		SelectedBids interface{} `json:"selected_bids"`
		Users        interface{} `json:"users"`
		Projects     []struct {
			Hidebids    bool        `json:"hidebids"`
			Files       interface{} `json:"files"`
			DriveFiles  interface{} `json:"drive_files"`
			Attachments interface{} `json:"attachments"`
			Bidperiod   int         `json:"bidperiod"`
			Currency    struct {
				Code         string  `json:"code"`
				Name         string  `json:"name"`
				Country      string  `json:"country"`
				Sign         string  `json:"sign"`
				ExchangeRate float64 `json:"exchange_rate"`
				ID           int     `json:"id"`
				IsExternal   bool    `json:"is_external"`
			} `json:"currency"`
			Featured           bool   `json:"featured"`
			PreviewDescription string `json:"preview_description"`
			Upgrades           struct {
				ActivePrepaidMilestone interface{} `json:"active_prepaid_milestone"`
				SuccessBundle          interface{} `json:"success_bundle"`
				NonCompete             bool        `json:"non_compete"`
				ProjectManagement      bool        `json:"project_management"`
				NDA                    bool        `json:"NDA"`
				Assisted               interface{} `json:"assisted"`
				Urgent                 bool        `json:"urgent"`
				Featured               bool        `json:"featured"`
				Nonpublic              bool        `json:"nonpublic"`
				Fulltime               bool        `json:"fulltime"`
				Qualified              bool        `json:"qualified"`
				Sealed                 bool        `json:"sealed"`
				PfOnly                 bool        `json:"pf_only"`
				IPContract             bool        `json:"ip_contract"`
				Recruiter              interface{} `json:"recruiter"`
				Listed                 interface{} `json:"listed"`
			} `json:"upgrades"`
			InvitedFreelancers     interface{} `json:"invited_freelancers"`
			ID                     int         `json:"id"`
			ActivePrepaidMilestone interface{} `json:"active_prepaid_milestone"`
			Negotiated             bool        `json:"negotiated"`
			Title                  string      `json:"title"`
			Assisted               interface{} `json:"assisted"`
			SupportSessions        interface{} `json:"support_sessions"`
			Submitdate             int         `json:"submitdate"`
			NdaSignatures          interface{} `json:"nda_signatures"`
			ProjectCollaborations  interface{} `json:"project_collaborations"`
			Nonpublic              bool        `json:"nonpublic"`
			Location               struct {
				AdministrativeArea interface{} `json:"administrative_area"`
				City               interface{} `json:"city"`
				Country            struct {
					HighresFlagURL    interface{} `json:"highres_flag_url"`
					Code              interface{} `json:"code"`
					Name              interface{} `json:"name"`
					SeoURL            interface{} `json:"seo_url"`
					FlagURLCdn        interface{} `json:"flag_url_cdn"`
					HighresFlagURLCdn interface{} `json:"highres_flag_url_cdn"`
					PhoneCode         interface{} `json:"phone_code"`
					LanguageCode      interface{} `json:"language_code"`
					Demonym           interface{} `json:"demonym"`
					LanguageID        interface{} `json:"language_id"`
					Person            interface{} `json:"person"`
					Iso3              interface{} `json:"iso3"`
					Sanction          interface{} `json:"sanction"`
					FlagURL           interface{} `json:"flag_url"`
					RegionID          interface{} `json:"region_id"`
				} `json:"country"`
				Vicinity    interface{} `json:"vicinity"`
				Longitude   interface{} `json:"longitude"`
				FullAddress interface{} `json:"full_address"`
				Latitude    interface{} `json:"latitude"`
			} `json:"location"`
			RecommendedFreelancers interface{} `json:"recommended_freelancers"`
			Type                   string      `json:"type"`
			Hireme                 bool        `json:"hireme"`
			OwnerID                int         `json:"owner_id"`
			Status                 string      `json:"status"`
			Jobs                   interface{} `json:"jobs"`
			Description            interface{} `json:"description"`
			CanPostReview          interface{} `json:"can_post_review"`
			Deleted                bool        `json:"deleted"`
			Qualifications         interface{} `json:"qualifications"`
			TimeFreeBidsExpire     int         `json:"time_free_bids_expire"`
			TrackIds               interface{} `json:"track_ids"`
			FrontendProjectStatus  string      `json:"frontend_project_status"`
			HourlyProjectInfo      interface{} `json:"hourly_project_info"`
			TrueLocation           interface{} `json:"true_location"`
			SubStatus              interface{} `json:"sub_status"`
			TimeUpdated            int64       `json:"time_updated"`
			Language               string      `json:"language"`
			SeoURL                 string      `json:"seo_url"`
			Urgent                 bool        `json:"urgent"`
			UserDistance           interface{} `json:"user_distance"`
			Local                  bool        `json:"local"`
			TimeSubmitted          int64       `json:"time_submitted"`
			Budget                 struct {
				CurrencyID  interface{} `json:"currency_id"`
				Minimum     float64     `json:"minimum"`
				Maximum     float64     `json:"maximum"`
				ProjectType interface{} `json:"project_type"`
				Name        interface{} `json:"name"`
			} `json:"budget"`
			NegotiatedBid interface{} `json:"negotiated_bid"`
			NdaDetails    struct {
				Signatures        interface{} `json:"signatures"`
				HiddenDescription interface{} `json:"hidden_description"`
			} `json:"nda_details"`
			BidStats struct {
				BidCount int     `json:"bid_count"`
				BidAvg   float64 `json:"bid_avg"`
			} `json:"bid_stats"`
			HiremeInitialBid interface{} `json:"hireme_initial_bid"`
			FromUserLocation interface{} `json:"from_user_location"`
		} `json:"projects"`
	} `json:"result"`
	RequestID string `json:"request_id"`
}

type ErrorJSON struct {
	Status    string `json:"status"`
	Message   string `json:"message"`
	ErrorCode string `json:"error_code"`
	RequestID string `json:"request_id"`
}

func QueryFreelancer(data *models.Data) {
	httpClient := http.Client{
		Timeout: 15 * time.Second,
	}
	throttle := time.Tick(10 * time.Second)
	for _, query := range data.Configuration.SearchQueries {
		getParams := url.Values{}
		if strings.ToLower(query) != "all" {
			getParams.Add("query", query)
		}
		getParams.Add("from_time", fmt.Sprintf("%d", data.GetRunTimeUnix()))
		if len(data.Configuration.AllowedCountries) > 0 {
			for _, cur := range data.Configuration.AllowedCountries {
				getParams.Add("countries[]", cur)
			}
		}
		if len(data.Configuration.AllowedLanguages) > 0 {
			for _, lang := range data.Configuration.AllowedLanguages {
				getParams.Add("languages[]", lang)
			}
		}
		formattedLink := fmt.Sprintf("https://www.freelancer.com/api/projects/0.1/projects/active/?%s", getParams.Encode())
		resp, err := httpClient.Get(formattedLink)
		if err != nil {
			log.Println(err)
			continue
		}
		projectJSON := ProjectJSON{}
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		closeErr := resp.Body.Close()
		if closeErr != nil {
			fmt.Println(closeErr)
		}
		if err != nil {
			log.Println(err)
			continue
		}
		err = json.Unmarshal(bodyBytes, &projectJSON)
		if err != nil {
			log.Printf("Unable to unmarshal JSON. Got this response text instead:\n%v\n", string(bodyBytes))
			continue
		}
		var jobs []*models.Job
		for _, r := range projectJSON.Result.Projects {
			if data.Configuration.MinimumFixed != 0 {
				if r.Type == "fixed" {
					if r.Budget.Maximum < data.Configuration.MinimumFixed {
						continue
					}
				}
			}
			if data.Configuration.MinimumHourly != 0 {
				if r.Type == "hourly" {
					if r.Budget.Maximum < data.Configuration.MinimumHourly {
						continue
					}
				}
			}
			timeSubmitted := time.Unix(r.TimeSubmitted, 0)
			newFreelancerJob := models.Job{}
			newFreelancerJob.URL = fmt.Sprintf("https://freelancer.com/projects/%s", r.SeoURL)
			newFreelancerJob.Description = r.PreviewDescription
			newFreelancerJob.Created = timeSubmitted
			newFreelancerJob.Title = r.Title
			newFreelancerJob.Query = query
			jobs = append(jobs, &newFreelancerJob)
		}
		if len(jobs) > 0 {
			jobGroup := models.JobGroup{
				Name: query,
				Jobs: jobs,
			}
			data.JobsChannel <- jobGroup
		}
		<-throttle
	}
	data.UpdateLastRunTime()
}
