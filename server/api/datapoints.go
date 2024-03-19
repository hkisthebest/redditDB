package api

import (
  "encoding/json"
  "fmt"
  "net/http"
  "server/dao"
  "server/model"
  "strconv"
  "time"

  "github.com/go-chi/chi/v5"
)

func DatapointRouter() http.Handler {
  r := chi.NewRouter()
  r.Get("/datapoints/{subreddit}", getDatapoints)
  r.Get("/datapoints/top", getTopDatapoints)
  return r
}

func getDatapoints(w http.ResponseWriter, r *http.Request) {
  subreddit := chi.URLParam(r, "subreddit")
  duration, err := strconv.Atoi(r.URL.Query().Get("duration"))
  if len(subreddit) < 3 {
    res, _ := json.Marshal(make(map[string]string))
    w.Write(res)
    return
  }
  if (err != nil) {
    fmt.Println(err)
  }

  query := &dao.DatapointDao{
    TimeAfter: time.Now().Add(time.Duration(duration * -1) * time.Hour),
    Subreddit: subreddit,
  }

  datapoints := query.GetDatapoints()
  response := make(map[string][]model.Datapoint)
  for _, datapoint := range datapoints {
    if datapoint != nil {
      response[datapoint.Subreddit] = append(response[datapoint.Subreddit], *datapoint)
    }
  }
  responseJson, _ := json.Marshal(response)
  w.Write(responseJson)
}

func getTopDatapoints(w http.ResponseWriter, r *http.Request) {
  duration, err := strconv.Atoi(r.URL.Query().Get("duration"))
  if (err != nil) {
    fmt.Println(err)
  }

  query := &dao.DatapointDao{
    TimeAfter: time.Now().Add(time.Duration(duration * -1) * time.Hour),
  }

  datapoints := query.GetTopDatapoints()
  response := make(map[string][]model.Datapoint)
  for _, datapoint := range datapoints {
    if datapoint != nil {
      response[datapoint.Subreddit] = append(response[datapoint.Subreddit], *datapoint)
    }
  }
  responseJson, _ := json.Marshal(response)
  w.Write(responseJson)
}
