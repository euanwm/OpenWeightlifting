"use client"

import { LifterSearchList } from './fetchLifterNamesTypes'

export default async function fetchLifterNames(
  params: { [key: string]: string },
): Promise<LifterSearchList> {
  if (params['name']?.length < 3) {
    return { names: [], total: 0 }
  }

  const URLParams = new URLSearchParams(params)

  const response = await fetch(`${process.env.API}/search?${URLParams}`)
  const jsonResponse = response.json()

  return jsonResponse as Promise<LifterSearchList>
}
