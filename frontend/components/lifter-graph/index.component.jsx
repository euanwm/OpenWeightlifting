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

export const LifterGraph = ({ lifter_name }) => {
  const data = {
    labels: ["2019-09-28", "2019-12-14", "2020-12-31", "2021-04-30", "2021-11-13", "2022-03-19", "2022-10-29"],
    datasets: [{
      label: "Competition Total",
      data: [187, 200, 225, 227, 236, 240, 235],
      fill: false,
      borderColor: 'blue',
      tension: 0.1
    },
    {
      "label": "Best Snatch",
      "data": [82, 90, 105, 106, 111, 107, 110],
      fill: false,
      borderColor: 'red',
      tension: 0.1
    },
    {
      "label": "Best C&J",
      "data": [105, 110, 120, 121, 125, 133, 125],
      fill: false,
      borderColor: 'rgb(75, 192, 192)',
      tension: 0.1
    }
    ]
  }

  return (
    <div style={{
      padding: '2rem'
    }}>
      {/* <h4>{lifter_name}: History (Total)</h4> */}
      < Line data={data} width={200} height={200} />
    </div >
  )
};