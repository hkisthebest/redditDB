package task

import (
	"fmt"
	"log"
	"net/http"
	"server/dao"
	"server/service"
	"time"
)

func fetchSubredditAbout() {
  log.Println("start fetching...")
  datapointDao := &dao.DatapointDao{}
  subredditDao := &dao.SubredditDao{}
  subreddits := subredditDao.GetAllSubreddits()
  redditService := &service.RedditService{
    Client: &http.Client{},
  }
  ticker := time.NewTicker(850 * time.Millisecond)
  counter := 0
  go func () {
    for {
      select {
      case <- ticker.C:
        currentSubreddit := subreddits[counter]
        subredditData := redditService.FetchDataFromReddit(currentSubreddit.Name)

        var rs = &service.SubRedditAboutResponse{}
        err := redditService.UnMarshalAboutResponseToStruct(subredditData, rs)

        if err == service.SubRedditAboutErrorResponseBanned {
          log.Println(err.Error(), currentSubreddit.Name)
        } else if err == service.SubRedditAboutErrorResponseTooMany {
          subreddits = append(subreddits, currentSubreddit)
        } else {
          datapointDao.InsertDatapoints(currentSubreddit.Name, rs.Data.ActiveUserCount, rs.Data.Subscribers)
          fmt.Println("subreddits to go: ", len(subreddits))
        }
        counter++
        if counter >= len(subreddits) {
          log.Println("fetching ended...")
        }
      }
      if counter >= len(subreddits) {
        ticker.Stop()
        return
      }
    }
  }()

}
