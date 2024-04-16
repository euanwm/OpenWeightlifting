import fetchLifterGraphData from './fetchLifterGraphData'

describe('fetchLifterGraphData', () => {
  it('should return a list of events', async () => {
    const params = { name: 'Euan Meston' }
    const result = await fetchLifterGraphData(params);

    // check that each attribute is present and not null
    expect(result?.labels.length).toBeGreaterThan(0)
    expect(result?.datasets.length).toBeGreaterThan(0)
  })
})