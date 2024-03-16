import { useEffect, useState } from 'react'
import { DatapointResponse } from './types/dataset'
const apiHost: string = import.meta.env.VITE_API_HOST || 'http://localhost:3000'
export default function useSearch(input: string, duration: string) {
  const [result, setResult] = useState<DatapointResponse>({})
  const [loading, setLoading] = useState<boolean>(false)

  useEffect(() => {
    setLoading(true)
    const fetchData = async () => {
      try {
        let response;
        if (input.length >= 3 && input) {
          response = await fetch(`${apiHost}/api/datapoints/${input}?duration=${duration}`)
        } else {
          response = await fetch(`${apiHost}/api/datapoints/top?duration=${duration}`)
        }
        if (!response.ok) {
          throw new Error(response.statusText)
        }
        const resultJson = await response.json() as DatapointResponse
        setResult(resultJson)
      } catch (error) {
        console.error('Error fetching data:', error);
      }
    }

    fetchData()
    .then(() => {
      setLoading(false)
    })
    .catch(() => {
      setLoading(false)
    }) 
  }, [input, duration])

  return {
    result,
    loading
  }
}

