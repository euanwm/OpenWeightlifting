import { LifterChartData } from '@/api/fetchLifterGraphData/fetchLifterGraphDataTypes'
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  ChartOptions,
} from 'chart.js'
import { Line } from 'react-chartjs-2'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
)

// todo: add a aspect ratio prop to the component
const LifterGraph = ({
  lifterHistory,
  setRatio,
}: {
  lifterHistory: LifterChartData | null
  setRatio: number
}) => {
  if (!lifterHistory) {
    return null
  }

  // todo: define each dataset as a type/interface instead of manually indexing into the array
  const processedData = {
    labels: lifterHistory.labels,
    datasets: [
      {
        ...lifterHistory.datasets[0],
        borderColor: '#0072F5',
        backgroundColor: '#3694FF',
      },
      {
        ...lifterHistory.datasets[1],
        borderColor: '#17C964',
        backgroundColor: '#78F2AD',
      },
      {
        ...lifterHistory.datasets[2],
        borderColor: '#F31260',
        backgroundColor: '#F75F94',
      },
      {
        ...lifterHistory.datasets[3],
        borderColor: '#F3A312',
        backgroundColor: '#F7C78F',
      },
    ],
  }

  const config = {
    plugins: {
      legend: {
        display: false,
      },
    },
    scales: {
      x: {
        grid: {
          display: false,
        },
        ticks: {
          color: '#A0A0A0',
          font: {
            size: 12,
          },
        },
      },
      y: {
        grid: {
          color: '#F0F0F0',
          borderColor: '#F0F0F0',
          borderWidth: 1,
        },
        ticks: {
          color: '#A0A0A0',
          font: {
            size: 12,
          },
        },
      },
    },
    maintainAspectRatio: false,
    aspectRatio: setRatio,
  }

  return <Line data={processedData} options={config} />
}

export default LifterGraph
