import fetchLifterHistory from './fetchLifterHistory'

describe('fetchLifterHistory', () => {
  it('should return a list of events', async () => {
    const result = await fetchLifterHistory('Euan Meston');

    // check that each attribute is present and not null
    expect(result?.name).toBe('Euan Meston')
    expect(result?.lifts.length).toBeGreaterThan(0)
    expect(result?.graph.labels.length).toBeGreaterThan(0)
    expect(result?.graph.datasets.length).toBeGreaterThan(0)
    expect(result?.stats.best_snatch).toBeGreaterThan(0)
    expect(result?.stats.best_cj).toBeGreaterThan(0)
    expect(result?.stats.best_total).toBeGreaterThan(0)
    expect(result?.stats.make_rate_snatches.length).toBeGreaterThan(0)
    expect(result?.stats.make_rate_cj.length).toBeGreaterThan(0)
  })
})