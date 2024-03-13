export type dataset = {
  x: string,
  y: number,
}

export type datapoint = {
  id: string,
  subreddit: string,
  time: string,
  users: number,
  subscribers: number,
  createdAt: string,
  updatedAt: string
}

export type datapointResponse = Record<string, datapoint[]>
