import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
} from 'chart.js'
import { Line } from 'react-chartjs-2'
import { Popover } from "@nextui-org/react"

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
)

export const LifterGraph = ({ data }) => {
  const { labels, datasets } = data
  const processedData = {
    labels: labels,
    datasets: [{
      ...datasets[0],
      borderColor: '#0072F5',
      backgroundColor: '#3694FF'
    }, {
      ...datasets[1],
      borderColor: '#17C964',
      backgroundColor: '#78F2AD'
    }, {
      ...datasets[2],
      borderColor: '#F31260',
      backgroundColor: '#F75F94'
    }]
  }

  const config = {
    color: 'white',
    layout: {
      padding: 20
    },
    elements: {
      point: {
        radius: 4,
        borderWidth: 0,
        hitRadius: 2
      },
      line: {
        tension: 0.1,
        borderCapStyle: 'round',
        fill: false,
        borderWidth: 2
      }
    },
    scales: {
      x: {
        grid: {
          color: '#313538'
        }
      },
      y: {
        grid: {
          color: '#313538'
        }
      }
    },
    plugins: {
      legend: {
        position: 'bottom',
        labels: {
          boxHeight: 15,
          boxWidth: 15
        },
      }
    },
  }

  return (
    <>
      <Line data={processedData} options={config} />
    </>
  )
};