import './App.css'
import Chart from './chart'
import { useEffect, useState } from 'react'
import search from './hook'
import { Datapoint, DatapointResponse } from './types/dataset'
import GitHubButton from 'react-github-btn'

const durations: number[] = [24, 72, 240, 720, 1200]

function App() {
  const [subredditInput, setSubredditInput] = useState<string>('')
  const [duration, setDuration] = useState<number>(24)
  const [loading, setLoading] = useState<boolean>(false)
  const [result, setResult] = useState<DatapointResponse>({})

  useEffect(() => {
    search('', 24, setLoading, setResult)
  }, [])

  const onFormSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    await search(subredditInput, duration, setLoading, setResult)
  }

  if (loading) return null

  return (
    <div style={{ flexDirection: 'row', position: 'absolute', top: 0, left: 0 }}>
      <div style={{
        float: 'left',
        paddingLeft: 15
      }}>
        <h2 >Active subreddit users</h2>
        <form onSubmit={onFormSubmit}>
          r/ <input value={subredditInput}
            onChange={(e) => setSubredditInput(e.target.value)}
            placeholder='Three characters...'
            style={{ marginRight: 5 }}
          />
          <select
            onChange={(e) => setDuration(+e.target.value)}
            value={duration}
            style={{ marginRight: 5 }}
          >
            {durations.map(d => (
              <option key={d} value={d}>{d / 24} day{d === 24 ? '' : 's'}</option>
            ))}
          </select>
          <button type='submit' value="Search">GO!</button>
        </form>
        <p
          style={{ fontWeight: 'bold' }}>
          Time is shown in local time
        </p>
        <p>
          Scroll to zoom!
        </p>
        <div style={{ display: 'flex', marginTop: 20, justifyContent: 'start' }}>
          <GitHubButton href="https://github.com/hkisthebest/redditdb"
            data-color-scheme="no-preference: light; light: light; dark: dark;"
            data-size="large"
            data-show-count="true"
            aria-label="Star hkisthebest/redditdb on GitHub"
          >
            Star
          </GitHubButton>
        </div>
      </div>
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        {Object.entries(result).map(([subreddit, data]: [string, Datapoint[]]) => (
          <div key={subreddit} style={{ width: '375px', boxSizing: 'border-box', padding: '5px' }}>
            <Chart
              title={subreddit} datasets={data.map((d: Datapoint) => ({
                x: d.time,
                y: d.users,
              }))} />
          </div>
        ))}
      </div>
    </ div>
  )
}

export default App
