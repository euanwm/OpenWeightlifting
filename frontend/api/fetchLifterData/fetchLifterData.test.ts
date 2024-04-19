import fetchLifterData from './fetchLifterData'

describe('fetchLifterData', () => {
  it('should return a list of rankings', async () => {
    const params = {
      start: '0',
      stop: '10',
      sortby: 'total',
      federation: 'allfeds',
      weightclass: 'MALL',
      year: '69',
    }

    const result = await fetchLifterData(params);
    expect(result.size).toBeGreaterThan(40000)
    expect(result.data.length).toBe(10)
  })
})