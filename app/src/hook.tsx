import { Dispatch, SetStateAction } from 'react'
import { DatapointResponse } from './types/dataset'
const apiHost: string = import.meta.env.VITE_API_HOST || 'http://localhost:3000'

export default async function search(input: string, duration: number, setLoading: Dispatch<SetStateAction<boolean>>, setResult: Dispatch<SetStateAction<DatapointResponse>>) {
  try {
    setLoading(true)
    let url = `${apiHost}/api/datapoints/top?duration=${duration}`
    if (input.length >= 3) {
      url = `${apiHost}/api/datapoints/${input}?duration=${duration}`
    }
    const response = await fetch(url)
    const resultJson = await response.json() as DatapointResponse
    setResult(resultJson)
  } catch (error) {

    console.error('Error fetching data:', error);
  } finally {
    setLoading(false)
  }

}
