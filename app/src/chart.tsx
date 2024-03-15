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

ChartJS.register(
  scales,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler
)

const options = {
  responsive: true,
  plugins: {
    legend: {
      position: 'top' as const,
    },
    title: {
      display: true,
    },
  },
}

function Chart({ title, datasets }: { title: string, datasets: Dataset[] }) {
  const data = {
    datasets: [
      {
        label: 'users',
        data: datasets,
        borderColor: 'rgb(56, 157, 245)',
        backgroundColor: 'rgba(56, 157, 245)',
        borderWidth: 2
      },
    ],
  }
  return (
    <div>
      <Line width={400} height={300} options={{
        ...options,
        plugins: {
          ...options.plugins,
          title: {
            ...options.plugins.title,
            text: title,
          }
        }
      }} data={data} />
    </div >
  )
}

export default Chart
