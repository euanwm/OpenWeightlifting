"use client"

import { LifterHistory } from './fetchLifterHistoryTypes'

export default async function fetchLifterHistory(
  params: { [key: string]: string },
): Promise<LifterHistory | null> {
  if (!params['name']) {
    return null
  }

  const URLParams = new URLSearchParams(params)

  const response = await fetch(`${process.env.API}/history?${URLParams}`, {
    headers: {
      'Content-Type': 'application/json',
    }
  })

  const jsonResponse = response.json()
  return jsonResponse
}
