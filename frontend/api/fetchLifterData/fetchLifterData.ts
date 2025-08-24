'use client'

import { LeaderboardResult } from './fetchLifterDataTypes'
import { getAPI } from '../loadSplitter'

export default async function fetchLifterData(params: {
  [key: string]: string
}): Promise<LeaderboardResult> {
  const URLParams = new URLSearchParams(params)

  const response = await fetch(`${getAPI()}/leaderboard?${URLParams}`, {
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
  })

  const jsonResponse = await response.json()
  return jsonResponse
}
