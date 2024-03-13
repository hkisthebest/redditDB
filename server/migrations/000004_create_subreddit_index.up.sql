CREATE INDEX IF NOT EXISTS trgm_subreddit_idx ON datapoint USING GIST (subreddit gist_trgm_ops)
