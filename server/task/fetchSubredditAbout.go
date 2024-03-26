package task

import (
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
          log.Println(err.Error(), currentSubreddit.Name)
          subreddits = append(subreddits, currentSubreddit)
        } else {
          if rs.Data.ActiveUserCount == 0 {
            log.Println("subredditData: ", string(subredditData[:]))
          }
          datapointDao.InsertDatapoints(currentSubreddit.Name, rs.Data.ActiveUserCount, rs.Data.Subscribers)
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
