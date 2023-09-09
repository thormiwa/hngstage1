package main

import (
  "fmt"
  "net/http"
  "time"
  "encoding/json"
)

func handler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
		case "GET":
			slackName := r.URL.Query().Get("slack_name")
			track := r.URL.Query().Get("track")

			if slackName == "" || track == "" {
				http.Error(w, "Please provide a slack_name and track", http.StatusBadRequest)
				return
			}
			currentDay := time.Now().Weekday().String()
			utcTime := time.Now().UTC()
			githubFileUrl := "https://github.com/thormiwa/hngtask1/main.go"
			githubRepoUrl := "https://github.com/thormiwa/hngtask1"
			statusCode := http.StatusOK

			response := map[string]interface{}{
				"slack_name": slackName,
				"current_day": currentDay,
				"utc_time": utcTime,
				"track": track,
				"github_file_url": githubFileUrl,
				"github_repo_url": githubRepoUrl,
				"status_code": statusCode,
			}

			jsonResponse, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "Failed to marshal JSON response", http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonResponse)
		default:
			fmt.Fprintf(w, "Sorry, only GET method are supported.")
	}
}

func main() {
    http.HandleFunc("/api", handler)
    http.ListenAndServe(":8080", nil)
}