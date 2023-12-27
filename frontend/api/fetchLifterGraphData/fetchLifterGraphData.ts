import { LifterChartData } from './fetchLifterGraphDataTypes'

export default async function fetchLifterGraphData(
  lifterName: string,
): Promise<LifterChartData> {
  if (!lifterName) {
    return
  }

  const bodyContent = JSON.stringify({ "name": lifterName })

  const response = await fetch(`${process.env.API}/lifter`, {
    method: 'POST',
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
    body: bodyContent,
  })
  const jsonResponse = response.json()

  return jsonResponse
}
