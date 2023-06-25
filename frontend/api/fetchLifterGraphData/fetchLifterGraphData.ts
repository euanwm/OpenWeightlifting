import { LifterChartData } from './fetchLifterGraphDataTypes';

export default async function fetchLifterGraphData(lifterName = ''): Promise<LifterChartData> {
  const bodyContent = JSON.stringify({"NameStr": lifterName });

  try {
    const response = await fetch(`${process.env.API}/lifter`, {
      method: 'POST',
      headers: {
        Accept: '*/*',
        'Content-Type': 'application/json',
      },
      body: bodyContent,
    });
    const jsonResponse = response.json();

    return jsonResponse;
  } catch(error) {
    console.error('error in fetchLifterGraphData', error);

    return undefined;
  }
};
