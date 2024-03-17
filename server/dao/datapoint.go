package dao

import (
  "context"
  "fmt"
  "server/db"
  "server/model"
  "time"

  "github.com/georgysavva/scany/v2/pgxscan"
  "github.com/google/uuid"
)

type DatapointDao struct {
  TimeAfter time.Time
  Subreddit string
}

func(s DatapointDao) GetDatapoints() (datapoints []*model.Datapoint) {
  if s.Subreddit == ""{
    return nil
  }
  pgxscan.Select(
    context.Background(),
    db.Pool,
    &datapoints,
    `SELECT id, subreddit, time, users, subscribers, created_at, updated_at FROM datapoint WHERE time >= $1 AND subreddit % $2 ORDER BY subreddit <-> $2 ASC, time ASC`,
    s.TimeAfter,
    fmt.Sprintf("r/%%%s%%", s.Subreddit),
  )
  return datapoints
}

func(s DatapointDao) GetTopDatapoints() (datapoints []*model.Datapoint) {
  pgxscan.Select(
    context.Background(),
    db.Pool,
    &datapoints,
    `SELECT id, subreddit, time, users, subscribers, created_at, updated_at
    FROM datapoint
    WHERE time >= $1 AND
    subreddit IN 
    (
      'r/funny',
      'r/AskReddit',
      'r/gaming',
      'r/aww',
      'r/worldnews',
      'r/todayilearned',
      'r/Music',
      'r/movies',
      'r/science',
      'r/pics',
      'r/Showerthoughts',
      'r/memes'
    )
    ORDER BY time ASC`,
    s.TimeAfter,
  )
  return datapoints
}

func(s DatapointDao) InsertDatapoints(subreddit string, users int, subscribers int) {
  db.Pool.Exec(
    context.Background(),
    "INSERT INTO datapoint(id, subreddit, time, users, subscribers, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
    uuid.NewString(),
    subreddit,
    time.Now(),
    users,
    subscribers,
    time.Now(),
    time.Now(),
  )
}
