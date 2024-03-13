package model

import "time"

type Datapoint struct {
  Id string `json:"id"`
  Subreddit string `json:"subreddit"`
  Time time.Time `json:"time"`
  Users int `json:"users"`
  Subscribers int `json:"subscribers"`
  CreatedAt time.Time `json:"createdAt"`
  UpdatedAt time.Time `json:"updatedAt"`
}
