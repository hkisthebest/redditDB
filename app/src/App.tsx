import './App.css'
import Chart from './chart'
import { useEffect, useState } from 'react'
import { DateTime } from 'luxon'
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

  if (loading) return null

  return (
    <div style={{ flexDirection: 'row' }}>
      <div style={{
        width: '20%',
        float: 'left',
        paddingLeft: 15
      }}>
        <h2 >Active subreddit users</h2>
        r/ <input value={subredditInput}
          onChange={(e) => setSubredditInput(e.target.value)}
          placeholder='Three characters...'
        />
        <select onChange={(e) => setDuration(+e.target.value)} value={duration}>
          {durations.map(d => (
            <option key={d} value={d}>{d / 24} day{d === 24 ? '' : 's'}</option>
          ))}
        </select>
        <input type='button'
          onClick={() => {
            search(subredditInput, duration, setLoading, setResult)
          }}
          value="Search" />
        <p
          style={{ fontWeight: 'bold' }}>
          Time is shown in UTC
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
            <Chart title={subreddit} datasets={data.map((d: Datapoint) => ({
              x: DateTime.fromISO(d.time).toFormat('MMM, d, h a'),
              y: d.users,
            }))} />
          </div>
        ))}
      </div>
    </ div>
  )
}

export default App
