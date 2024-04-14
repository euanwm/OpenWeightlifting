import fetchLifterNames from './fetchLifterNames'
import { LifterSearchList } from '@/api/fetchLifterNames/fetchLifterNamesTypes'

describe('fetchLifterNames', () => {
  it('should return a list of lifter names', async () => {
    const result = await fetchLifterNames({ 'name': 'Euan Meston' });
    const expectedReturn: LifterSearchList = {
      names: [
        {
          Name: "Euan Meston",
          Federation: "UK"
        }
      ],
      total: 1
    }

    expect(result?.total).toEqual(1)
    expect(result).toEqual(expectedReturn)
  })
  it('should return an empty list if the name is less than 3 characters', async () => {
    const result = await fetchLifterNames({ 'name': 'Eu' });

    expect(result?.names.length).toBe(0)
  })
  it('should limit the number of names returned', async () => {
    const result = await fetchLifterNames({ 'name': 'dave', 'limit': '5' });

    expect(result?.names.length).toEqual(5)
    expect(result?.total).toBeGreaterThan(5)
  })
})