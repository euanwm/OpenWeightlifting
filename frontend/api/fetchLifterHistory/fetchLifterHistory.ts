import { LifterHistory } from './fetchLifterHistoryTypes'

export default async function fetchLifterHistory(
  name: string,
): Promise<LifterHistory | null> {
  if (!name) {
    return null
  }

  const response = await fetch(`${process.env.API}/history`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ "name": name }),
  })

  const jsonResponse = response.json()
  return jsonResponse
}
