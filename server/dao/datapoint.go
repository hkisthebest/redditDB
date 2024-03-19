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
    `
    WITH similar_subreddits AS (
      SELECT s.*, SIMILARITY(s.name, $1) AS similarity
      FROM subreddit s
      WHERE s.name % $1 AND SIMILARITY(s.name, $1) > 0.4
      ORDER BY similarity DESC
      LIMIT 10
    ),
    top_datapoints AS (
      SELECT dp.*
      FROM datapoint dp
      JOIN similar_subreddits ss ON dp.subreddit = ss.name
      WHERE dp.time >= $2
      ORDER BY ss.similarity DESC, dp.time ASC
    )
    SELECT *
    FROM top_datapoints;
    `,
    fmt.Sprintf("r/%s", s.Subreddit),
    s.TimeAfter,
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
