import { LifterSearchList } from "./fetchLifterNamesTypes"

export default async function fetchLifterNames(nameQuery: string): Promise<LifterSearchList> {
  if (nameQuery.length < 3) {
    return {names: []}
  }

  try {
    const response = await fetch(`https://api.openweightlifting.org/search?name=${nameQuery}`);
    const jsonResponse = response.json();

    return jsonResponse;
  } catch (error) {
    console.error('error in searching names', error)

    return {names: []}
  }
}