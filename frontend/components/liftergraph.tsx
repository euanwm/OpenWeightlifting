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
  ChartOptions
} from 'chart.js'
import { Line } from 'react-chartjs-2'

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend
)

export const LifterGraph = ({ lifterHistory }: { lifterHistory: LifterChartData }) => {
  if (!lifterHistory) {
    return null;
  }

  console.log(lifterHistory)
  // todo: define each dataset as a type/interface instead of manually indexing into the array
  const processedData = {
    labels: lifterHistory.labels,
    datasets: [{
      ...lifterHistory.datasets[0],
      borderColor: '#0072F5',
      backgroundColor: '#3694FF'
    }, {
      ...lifterHistory.datasets[1],
      borderColor: '#17C964',
      backgroundColor: '#78F2AD'
    }, {
      ...lifterHistory.datasets[2],
      borderColor: '#F31260',
      backgroundColor: '#F75F94'
    }, {
      ...lifterHistory.datasets[3],
      borderColor: '#F3A312',
      backgroundColor: '#F7C78F'
    }]
  }

  /* todo: implement scales gridlines colour into the config
      scales: {
        x: {grid: {color: '#313538'}},
        y: {grid: {color: '#313538'}}
      },
   */
  const config: ChartOptions = {
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