import { LifterHistory } from './fetchLifterHistoryTypes'

const blankLifterHistory: LifterHistory = {
  name: '',
  graph: {
    labels: [],
    datasets: [],
  },
  lifts: [],
}

export default async function fetchLifterHistory(
  name: string | string[] | undefined,
): Promise<LifterHistory> {
  if (!name) {
    return blankLifterHistory
  }

  try {
    const response = await fetch(`${process.env.API}/history`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({ NameStr: name }),
    })

    const jsonResponse = response.json()
    return jsonResponse
  } catch (error) {
    console.error('error in fetchLifterHistory', error)

    return blankLifterHistory
  }
}
