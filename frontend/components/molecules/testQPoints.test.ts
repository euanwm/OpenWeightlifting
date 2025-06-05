import { Gender, QPointsCalculator } from '@/components/molecules/qpoints'

describe('with correct parameters', () => {
  it('Normal Female QPoints - 3 Decimal Places', () => {
    const testInput = {
      gender: Gender.Female,
      bodyweight: 60.00,
      total: 150.00,
      age: 25
    }
    const expectedScore = 209.768;
    const returnedResult = QPointsCalculator(
      testInput.gender,
      testInput.bodyweight,
      testInput.total,
      testInput.age,
    )
    expect(returnedResult).toEqual(expectedScore)
  })
  it('Normal Male QPoints - 3 Decimal Places', () => {
    const testInput = {
      gender: Gender.Male,
      bodyweight: 80.00,
      total: 200.00,
      age: 25
    }
    const expectedScore = 261.716;
    const returnedResult = QPointsCalculator(
        testInput.gender,
        testInput.bodyweight,
        testInput.total,
        testInput.age,
    )
    expect(returnedResult).toEqual(expectedScore)
  })
})