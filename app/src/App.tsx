import './App.css'
import Chart from './chart'
import { useState } from 'react'
import { DateTime } from 'luxon'
import useSearch from './hook'
import { Datapoint } from './types/dataset'

function App() {
  const [subredditInput, setsubredditInput] = useState<string>('')
  const [duration, setDuration] = useState<string>('24')
  const { result, loading } = useSearch(subredditInput, duration)

  if (loading) return null

  return (
    <>
      <div style={{ marginTop: '10px', marginLeft: '10px', position: 'fixed', top: '0px', left: '0px' }}>
        <h2 >Active subreddit users</h2>
        r/ <input value={subredditInput} onChange={(e) => setsubredditInput(e.target.value)} />
        <select onChange={(e) => setDuration(e.target.value)}>
          <option value="24">1 day</option>
          <option value="72">3 days</option>
          <option value="240">10 days</option>
          <option value="720">30 days</option>
          <option value="1200">50 days</option>
        </select>
      </div>
      <div style={{ display: 'flex', flexWrap: 'wrap' }}>
        {Object.entries(result).map(([subreddit, data]: [string, Datapoint[]]) => {
          console.log(data)
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
