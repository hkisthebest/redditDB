package task

import (
	"net/http"
	"server/dao"
	"server/service"
	"time"
)

func fetchSubredditAbout() {
  datapointDao := &dao.DatapointDao{}
  subredditDao := &dao.SubredditDao{}
  subreddits := subredditDao.GetAllSubreddits()
  service := &service.RedditService{
    Client: &http.Client{},
  }
  ticker := time.NewTicker(850 * time.Millisecond)
  counter := 0
  go func () {
    for {
      select {
      case <- ticker.C:
        currentSubreddit := subreddits[counter]
        subredditData := service.FetchDataFromReddit(currentSubreddit.Name)
        if subredditData.Data.ActiveUserCount == 0 {
          subreddits = append(subreddits, currentSubreddit)
        } else {
          datapointDao.InsertDatapoints(currentSubreddit.Name, subredditData.Data.ActiveUserCount, subredditData.Data.Subscribers)
        }
        counter++
      }
      if counter >= len(subreddits) {
        ticker.Stop()
        return
      }
    }
  }()

}
