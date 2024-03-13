CREATE TABLE IF NOT EXISTS datapoint(
  id VARCHAR(36) PRIMARY KEY,
  subreddit varchar(50) NOT NULL REFERENCES subreddit (name),
  time TIMESTAMPTZ NOT NULL,
  users INTEGER NOT NULL,
  subscribers INTEGER NOT NULL,
  created_at TIMESTAMPTZ NOT NULL,
  updated_at TIMESTAMPTZ NOT NULL
);
