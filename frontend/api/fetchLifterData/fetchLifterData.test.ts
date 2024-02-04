import fetchLifterData from './fetchLifterData'

describe('fetchLifterData', () => {
  it('should return a list of rankings', async () => {
    const result = await fetchLifterData(
      {
        start: 0,
        stop: 10,
        sortby: 'total',
        federation: 'allfeds',
        weightclass: 'MALL',
        year: 69,
      });
    expect(result.size).toBeGreaterThan(40000)
    expect(result.data.length).toBe(10)
  })
})