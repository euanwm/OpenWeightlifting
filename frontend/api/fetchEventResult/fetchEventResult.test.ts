import fetchEventResult from './fetchEventResult';

describe('fetchEventResult', () => {
  it('should return a valid event result', async () => {
    const eventMetaData = {
      "fed": "NVF",
      "id": "58.csv"
    }
    const result = await fetchEventResult(eventMetaData);
    expect(result.size).toEqual(8)
  })
})
