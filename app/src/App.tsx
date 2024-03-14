import './App.css'
import Chart from './chart'
import { useState, useEffect } from 'react'
import { datapointResponse } from './types/dataset'

function App() {
  const [result, setResult] = useState<datapointResponse>({})
  const [loading, setLoading] = useState(false)
  const [subredditInput, setsubredditInput] = useState<string>('')
  const [duration, setDuration] = useState<string>('24')

  useEffect(() => {
    if (subredditInput.length >= 3) return
    setLoading(true)
    const fetchData = async () => {
      try {
        const response = await fetch(`http://${window.location.host}/api/datapoints/top?duration=${duration}`)
        if (!response.ok) {
          throw new Error(response.statusText)
        }
        const resultJson = await response.json()
        setResult(resultJson)
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }

    fetchData()
    setLoading(false)
  }, [subredditInput, duration])

  useEffect(() => {
    if (subredditInput.length < 3) return
    setLoading(true)
    const fetchData = async () => {
      try {
        const response = await fetch(`http://${window.location.host}/api/datapoints/${subredditInput}?duration=${duration}`)
        if (!response.ok) {
          throw new Error(response.statusText)
        }
        const resultJson = await response.json()
        setResult(resultJson)
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }

    fetchData()
    setLoading(false)
  }, [subredditInput, duration])


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
        {Object.entries(result).map(([subreddit, data]) => {
          return (
            <div key={data?.[0].id} style={{ boxSizing: 'border-box', }}>
              <Chart title={subreddit} datasets={data.map(d => ({ x: d.time, y: d.users }))
              } />
            </div>
          )
        })}
      </ div>
    </>
  )
}

export default App
