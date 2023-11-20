import { LifterSearchList } from './fetchLifterNamesTypes'

export default async function fetchLifterNames(
  name: string,
): Promise<LifterSearchList> {
  if (name?.length < 3) {
    return { names: [] }
  }

  const response = await fetch(`${process.env.API}/search?name=${name}`)
  const jsonResponse = response.json()

  return jsonResponse
}
