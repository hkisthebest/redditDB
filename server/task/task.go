package task

import (
	"net/http"
	"os"
	"server/service"

	"github.com/robfig/cron/v3"
)
func Init() {
  c := cron.New()
  client := &service.RedditService{
    Client: &http.Client{},
  }
  if os.Getenv("CRONJOB") == "true" {
    client.RefreshToken()
    c.AddFunc("@every 20h", client.RefreshToken)
    c.AddFunc("1 * * * *", fetchSubredditAbout)
    c.Start()
  }

}
