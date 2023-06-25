import { LifterResult } from './fetchLifterDataTypes';

export default async function fetchLifterData(
  start = 0,
  stop = 500,
  sortby = 'total',
  federation = 'allfeds',
  weightclass = 'MALL',
  year = 69,
): Promise<LifterResult> {
  const bodyContent = JSON.stringify({
    start,
    stop,
    sortby,
    federation,
    weightclass,
    year,
  })

  try {
    const response = await fetch(`${process.env.API}/leaderboard`, {
        method: 'POST',
        headers: {
          Accept: '*/*',
          'Content-Type': 'application/json',
        },
        body: bodyContent,
      });

      const jsonResponse = await response.json();
      return jsonResponse;
  } catch (error) {
    console.error('error in fetchLifterData', error);

    return [];
  }
}
