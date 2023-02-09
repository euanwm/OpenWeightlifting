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
      fill: false,
      borderColor: 'blue',
      tension: 0.1
    }, {
      ...datasets[1],
      fill: false,
      borderColor: 'red',
      tension: 0.1
    }, {
      ...datasets[2],
      fill: false,
      borderColor: 'green',
      tension: 0.1
    }]
  }

  console.log(processedData)

  return (
    <div style={{
      padding: '2rem'
    }}>
      {/* <h4>{lifter_name}: History (Total)</h4> */}
      < Line data={processedData} width={200} height={200} />
    </div >
  )
};