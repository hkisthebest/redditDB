export type Dataset = {
  x: string,
  y: number,
}

export type Datapoint = {
  id: string,
  subreddit: string,
  time: string,
  users: number,
  subscribers: number,
  createdAt: string,
  updatedAt: string
}

export type DatapointResponse = Record<string, Datapoint[]>
