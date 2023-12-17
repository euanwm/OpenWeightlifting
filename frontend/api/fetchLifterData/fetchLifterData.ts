import { LeaderboardResult } from './fetchLifterDataTypes'

export default async function fetchLifterData(
  args: {
    start: number
    stop: number
    sortby: string
    federation: string
    weightclass: string
    year: number
  } | null,
): Promise<LeaderboardResult> {

  const {
    start = 0,
    stop = 50,
    sortby = 'total',
    federation = 'allfeds',
    weightclass = 'MALL',
    year = 69,
  } = args || {}

  const bodyContent = JSON.stringify({
    start,
    stop,
    sortby,
    federation,
    weightclass,
    year,
  })

  const response = await fetch(`${process.env.API}/leaderboard`, {
    method: 'POST',
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
    body: bodyContent,
  })

  const jsonResponse = await response.json()
  return jsonResponse
}
