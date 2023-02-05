import { Line } from 'react-chartjs-2'
import { Popover } from "@nextui-org/react"

/*
* See below link for tutorial shite
* https://itnext.io/chartjs-tutorial-with-react-nextjs-with-examples-2f514fdc130
* */
export const LifterGraph = async ({ lifter_name }) => {
  const res = await fetch(`${process.env.API}/lifter`, {
    method: 'POST',
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
    body: { 'NameStr': lifter_name },
  }).catch(error => console.error(error))


  return (
    // NONE OF THIS WORKS YET
    <Popover>
      <h4>{lifter_name}: History (Total)</h4>
      <Line data={res}
        width={200}
        height={200} />
    </Popover>
  )
};