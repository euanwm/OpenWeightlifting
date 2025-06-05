import { useState } from 'react'
import { useSearchParams, useRouter } from "next/navigation";

import {
  Input,
  Radio,
  Button,
  RadioGroup,
  Select,
  SelectItem,
  Tabs,
  Tab
} from '@nextui-org/react'

import HeaderBar from '@/components/molecules/head'

import { QPointsCalculator } from '@/components/molecules/qpoints';

const Coefficients = {
  AMale2001: 0.938573813,
  BMale2001: 135.390,
  AFemale2001: 1.005487664,
  BFemale2001: 112.811,

  AMale2005: 0.845716976,
  BMale2005: 168.091,
  AFemale2005: 1.316081431,
  BFemale2005: 107.844,

  AMale2009: 0.784780654,
  BMale2009: 173.961,
  AFemale2009: 1.056683941,
  BFemale2009: 125.441,

  AMale2013: 0.794358141,
  BMale2013: 174.393,
  AFemale2013: 0.89726074,
  BFemale2013: 148.026,

  AMale2017: 0.75194503,
  BMale2017: 175.508,
  AFemale2017: 0.783497476,
  BFemale2017: 153.655,

  AMale2021: 0.722762521,
  BMale2021: 193.609,
  AFemale2021: 0.787004341,
  BFemale2021: 153.757,
}

interface CoefficientSettings {
  ACoefficient: number
  BCoefficient: number
}

const SinclairCalculator = {
  getSinclairCoefficient: function (
    bodyweight: number,
    total: number,
    coeffSettings: CoefficientSettings,
  ) {
    const x = Math.log10(bodyweight / coeffSettings.BCoefficient)
    const ax2 = coeffSettings.ACoefficient * Math.pow(x, 2)
    return total * Math.pow(10, ax2)
  },

  doSinclairCalc: function (
    bodyweight: number,
    totalKg: number,
    coeffSettings: CoefficientSettings,
  ) {
    return this.getSinclairCoefficient(bodyweight, totalKg, coeffSettings)
  },

  getSinclair: function (
    year: string,
    gender: string,
    bodyWeightKg: number,
    total: number,
  ) {
    let isMale = gender == 'male'

    let coeffSettings: CoefficientSettings
    switch (year) {
      case "2001":
        coeffSettings = {
          ACoefficient: isMale
            ? Coefficients.AMale2001
            : Coefficients.AFemale2001,
          BCoefficient: isMale
            ? Coefficients.BMale2001
            : Coefficients.BFemale2001,
        }
        break
      case "2005":
        coeffSettings = {
          ACoefficient: isMale
            ? Coefficients.AMale2005
            : Coefficients.AFemale2005,
          BCoefficient: isMale
            ? Coefficients.BMale2005
            : Coefficients.BFemale2005,
        }
        break
      case "2009":
        coeffSettings = {
          ACoefficient: isMale
            ? Coefficients.AMale2009
            : Coefficients.AFemale2009,
          BCoefficient: isMale
            ? Coefficients.BMale2009
            : Coefficients.BFemale2009,
        }
        break
      case "2013":
        coeffSettings = {
          ACoefficient: isMale
            ? Coefficients.AMale2013
            : Coefficients.AFemale2013,
          BCoefficient: isMale
            ? Coefficients.BMale2013
            : Coefficients.BFemale2013,
        }
        break
      case "2017":
        coeffSettings = {
          ACoefficient: isMale
            ? Coefficients.AMale2017
            : Coefficients.AFemale2017,
          BCoefficient: isMale
            ? Coefficients.BMale2017
            : Coefficients.BFemale2017,
        }
        break
      case "2021":
        coeffSettings = {
          ACoefficient: isMale
            ? Coefficients.AMale2021
            : Coefficients.AFemale2021,
          BCoefficient: isMale
            ? Coefficients.BMale2021
            : Coefficients.BFemale2021,
        }
        break
      default:
        coeffSettings = {
          ACoefficient: isMale
            ? Coefficients.AMale2021
            : Coefficients.AFemale2021,
          BCoefficient: isMale
            ? Coefficients.BMale2021
            : Coefficients.BFemale2021,
        }
    }

    return this.doSinclairCalc(bodyWeightKg, total, coeffSettings)
  },
}

function CalculatorPage() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const option = searchParams && searchParams.size === 1 ? Array.from(searchParams.keys())[0]?.toString() : null;
  const pathname = option ? option : 'sinclair';

  const [sinclair, setSinclair] = useState<number>(0)
  const [bodyweight, setBodyweight] = useState<number>(0)
  const [total, setTotal] = useState<number>()
  const [selected, setSelected] = useState<string>('male')
  const [sinclairYear, setSinclairYear] = useState("2021")
  const [lifterAge, setLifterAge] = useState<number>()
  const [qPoints, setQPoints] = useState<number>(0)
  const [qPointsMasters, setQPointsMasters] = useState<number>(0)

  function handleQPoints(masters: boolean) {
    const gender = selected === "male" ? 1 : 0
    setQPoints(QPointsCalculator(
      gender,
      bodyweight,
      total,
    ))
    setQPointsMasters(QPointsCalculator(
      gender,
      bodyweight,
      total,
      lifterAge
    ))
  }

function resetValues() {
  setBodyweight(0)
  setTotal(undefined)
  setSinclair(0)
  setSelected('male')
  setSinclairYear('2021')
  setLifterAge(undefined)
  setQPoints(0)
  setQPointsMasters(0)
}

  return (
    <>
      <HeaderBar />
      <div className="flex justify-center mt-8">
        <div className="max-w-lg rounded-lg shadow-md p-8 space-y-8">
          <Tabs
            selectedKey={pathname}
            onSelectionChange={key => {
              router.replace(`?${key}`)
            }}
          >
            <Tab key="sinclair" title="Sinclair">
              <div className="items-center space-y-2">
                <Input
                  aria-label="Bodyweight"
                  type="number"
                  placeholder="Bodyweight"
                  value={bodyweight ? bodyweight : "Bodyweight"}
                  onChange={e => setBodyweight(parseFloat(e.target.value))}
                />
                <Input
                  aria-label="Total"
                  type="number"
                  placeholder="Total"
                  value={total}
                  onChange={e => setTotal(parseFloat(e.target.value))}
                />
                <Select
                  aria-label="Calculator Year"
                  placeholder="Jan 2021 - 2024 Dec"
                  onChange={e => setSinclairYear(e.target.value)}
                >
                  <SelectItem key="2001">Jan 2001 - 2004 Dec</SelectItem>
                  <SelectItem key="2005">Jan 2005 - 2008 Dec</SelectItem>
                  <SelectItem key="2009">Jan 2009 - 2012 Dec</SelectItem>
                  <SelectItem key="2013">Jan 2013 - 2016 Dec</SelectItem>
                  <SelectItem key="2017">Jan 2017 - 2020 Dec</SelectItem>
                  <SelectItem key="2021">Jan 2021 - 2024 Dec</SelectItem>
                </Select>
                <RadioGroup
                  aria-label="Gender"
                  value={selected}
                  onValueChange={setSelected}
                  orientation="horizontal"
                >
                  <Radio value="male" color="primary">
                    Male
                  </Radio>
                  <Radio value="female" color="danger">
                    Female
                  </Radio>
                </RadioGroup>
                <div className="flex gap-4 mt-2">
                  <Button
                    onClick={resetValues}
                    color="warning"
                    fullWidth
                  >Reset</Button>
                  <Button
                    fullWidth
                    onClick={() =>
                      setSinclair(
                        SinclairCalculator.getSinclair(
                          sinclairYear,
                          selected,
                          bodyweight,
                          total,
                        ),
                      )
                    }
                  >
                    Calculate
                  </Button>
                </div>
                <h2 className="text-lg font-semibold text-center mt-4">
                  Calculator Score: {sinclair === 0 ? 0 : sinclair.toFixed(3)}
                </h2>
              </div>
            </Tab>
            <Tab key="qpointsmasters" title="Masters QPoints">
              <div className="items-center space-y-2">
              <Input
                aria-label="Bodyweight"
                type="number"
                placeholder="Bodyweight"
                value={bodyweight ? bodyweight : "Bodyweight"}
                onChange={e => setBodyweight(parseFloat(e.target.value))}
              />
              <Input
                aria-label="Total"
                type="number"
                placeholder="Total"
                value={total}
                onChange={e => setTotal(parseFloat(e.target.value))}
              />
              <Input
                aria-label="Age"
                type="number"
                placeholder="Age (30 and over)"
                value={lifterAge}
                onChange={e => setLifterAge(parseFloat(e.target.value))}
              />
              <RadioGroup
                aria-label="Gender"
                value={selected}
                onValueChange={setSelected}
                orientation="horizontal"
              >
                <Radio value="male" color="primary">
                  Male
                </Radio>
                <Radio value="female" color="danger">
                  Female
                </Radio>
              </RadioGroup>
              <div className="flex gap-4 mt-2">
                <Button
                  fullWidth
                  onClick={resetValues}
                  color="warning"
                >Reset</Button>
                <Button
                  fullWidth
                  onClick={() => handleQPoints(true)}>
                  Calculate
                </Button>
              </div>
              <h2 className="text-lg font-semibold text-center mt-4">
                Calculator Score: {qPointsMasters}
              </h2>
              </div>
            </Tab>
            <Tab key="qpoints" title="QPoints">
              <div className="items-center space-y-2">
                <Input
                  aria-label="Bodyweight"
                  type="number"
                  placeholder="Bodyweight"
                  value={bodyweight ? bodyweight : "Bodyweight"}
                  onChange={e => setBodyweight(parseFloat(e.target.value))}
                />
                <Input
                  aria-label="Total"
                  type="number"
                  placeholder="Total"
                  value={total}
                  onChange={e => setTotal(parseFloat(e.target.value))}
                />
                <RadioGroup
                  aria-label="Gender"
                  value={selected}
                  onValueChange={setSelected}
                  orientation="horizontal"
                >
                  <Radio value="male" color="primary">
                    Male
                  </Radio>
                  <Radio value="female" color="danger">
                    Female
                  </Radio>
                </RadioGroup>
                <div className="flex gap-4 mt-2">
                  <Button
                    onClick={resetValues}
                    color="warning"
                    fullWidth
                  >Reset</Button>
                  <Button
                    fullWidth
                    onClick={() => handleQPoints(false)}>
                    Calculate
                  </Button>
                </div>
                <h2 className="text-lg font-semibold text-center mt-4">
                  Calculator Score: {qPoints}
                </h2>
              </div>
            </Tab>
          </Tabs>
        </div>
      </div>
    </>
  )
}

export default CalculatorPage
