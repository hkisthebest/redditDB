import './App.css'
import Chart from './chart'
import { useState } from 'react'
import { DateTime } from 'luxon'
import useSearch from './hook'
import { Datapoint } from './types/dataset'
import GitHubButton from 'react-github-btn'

function App() {
  const [subredditInput, setsubredditInput] = useState<string>('')
  const [duration, setDuration] = useState<string>('24')
  const { result, loading } = useSearch(subredditInput, duration)
  const durationList = ['1','24','72','240','720','1200'];

  if (loading) return null

  return (
    <>
      <div style={{ marginTop: '10px', marginLeft: '10px', position: 'fixed', top: '0px', left: '0px' }}>
        <h2 >Active subreddit users</h2>
        r/ <input value={subredditInput} onChange={(e) => setsubredditInput(e.target.value)} />
        <select onChange={(e) => setDuration(e.target.value)}>
          {
            durationList.map((data) => (
              <option value={data}>{data} {`day${data === '1' ? '' : 's'}`}</option>
            ))
          }
        </select>
        <p style={{ fontWeight: 'bold' }}>All the time is shown in UTC</p>
        <div style={{ display: 'flex', marginTop: 20, justifyContent: 'start' }}>
          <GitHubButton href="https://github.com/hkisthebest/redditdb" data-color-scheme="no-preference: light; light: light; dark: dark;" data-size="large" data-show-count="true" aria-label="Star hkisthebest/redditdb on GitHub">Star</GitHubButton>
        </div>
      </div>
      <div style={{ display: 'flex', flexWrap: 'wrap', marginTop: 50 }}>
        {Object.entries(result).map(([subreddit, data]: [string, Datapoint[]]) => {
          return (
            <div key={data?.[0].id} style={{ boxSizing: 'border-box', }}>
              <Chart title={subreddit} datasets={data.map((d: Datapoint) => {
                const x = DateTime.fromISO(d.time).toFormat('MMM, d, h a')
                return { x, y: d.users }
              })
              } />
            </div>
          )
        })}
      </ div>
    </>
  )
}

export default App
