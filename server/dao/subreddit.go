package dao

import (
  "context"
  "server/db"
  "server/model"
  "github.com/georgysavva/scany/v2/pgxscan"
)

type SubredditDao struct {
  Id string
  Subreddit string
}

func(s SubredditDao) GetAllSubreddits() (subreddits []*model.Subreddit) {
  pgxscan.Select(
    context.Background(),
    db.Pool,
    &subreddits,
    "select id, name, created_at, updated_at from subreddit",
  )
  return subreddits
}
