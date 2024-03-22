import { useRef } from 'react'
import 'chartjs-adapter-luxon'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
  scales,
} from 'chart.js'
import { Line } from 'react-chartjs-2'
import { Dataset } from './types/dataset'
import zoomPlugin from 'chartjs-plugin-zoom'
import { IoRefresh } from "react-icons/io5"



ChartJS.register(
  scales,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
  zoomPlugin,
)

const options = {
  scales: {
    x: {
      type: 'time',
    }
  },
  responsive: true,
  plugins: {
    legend: {
      position: 'top' as const,
    },
    title: {
      display: true,
    },
    zoom: {
      zoom: {
        drag: {
          enabled: true,
        },
        wheel: {
          enabled: true,
        },
        pinch: {
          enabled: true,
        },
        mode: 'xy',
      }
    }
  },
}

function Chart({ title, datasets }: { title: string, datasets: Dataset[] }) {
  const chartRef = useRef(null)
  const data = {
    datasets: [
      {
        label: 'users',
        data: datasets,
        borderColor: 'rgb(56, 157, 245)',
        backgroundColor: 'rgba(56, 157, 245)',
        borderWidth: 2,
        radius: 3,
      },
    ],
  }
  const handleResetZoom = () => {
    if (chartRef && chartRef.current) {
      chartRef.current.resetZoom();
    }
  }
  return (
    <div>
      <Line ref={chartRef} width={400} height={300} options={{
        ...options,
        plugins: {
          ...options.plugins,
          title: {
            ...options.plugins.title,
            text: title,
          }
        }
      }} data={data} />
      <IoRefresh onClick={handleResetZoom} title="Reset zoom" />
    </div >
  )
}

export default Chart
