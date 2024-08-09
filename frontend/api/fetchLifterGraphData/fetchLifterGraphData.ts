"use client"

import { LifterChartData } from './fetchLifterGraphDataTypes'

export default async function fetchLifterGraphData(
  params: { [key: string]: string },
): Promise<LifterChartData> {
  if (!params['name']) {
    return
  }
  if (!params['federation'] || params['federation'] === 'allfeds') {
    delete params['federation']
  }

  const URLParams = new URLSearchParams(params)

  const response = await fetch(`${process.env.API}/graph?${URLParams}`, {
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
  })
  const jsonResponse = response.json()

  return jsonResponse
}
