import { LeaderboardResult } from './fetchLifterDataTypes'

export default async function fetchLifterData(
  params: { [key: string]: string },
): Promise<LeaderboardResult> {

  const URLParams = new URLSearchParams(params)

  const response = await fetch(`${process.env.API}/leaderboard?${URLParams}`, {
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
  })

  const jsonResponse = await response.json()
  return jsonResponse
}
