import fetchLifterNames from './fetchLifterNames'

describe('fetchLifterNames', () => {
  it('should return a list of lifter names', async () => {
    const result = await fetchLifterNames('Euan Meston');

    expect(result?.names.length).toBe(1)
    expect(result?.names[0]).toEqual(
      {"Federation": "UK", "Name": "Euan Meston"}
    )
  })
})