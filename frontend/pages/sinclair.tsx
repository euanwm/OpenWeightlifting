import { Input, Radio, Dropdown, DropdownItem, DropdownMenu, Button, RadioGroup, Select, SelectItem } from '@nextui-org/react'
import { useState } from "react";
import HeaderBar from "@/layouts/head";

const Coefficients = {
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
  ACoefficient: number,
  BCoefficient: number
}

const SinclairCalculator = {
  getSinclairCoefficient: function (bodyweight : number, total : number, coeffSettings : CoefficientSettings) {
    const x = Math.log10(bodyweight / coeffSettings.BCoefficient);
    const ax2 = coeffSettings.ACoefficient * Math.pow(x, 2);
    return total * Math.pow(10, ax2);
  },

  doSinclairCalc: function (bodyweight : number, totalKg: number, coeffSettings : CoefficientSettings) {
    return this.getSinclairCoefficient(bodyweight, totalKg, coeffSettings)
  },

  getSinclair: function (year : string, gender : string, bodyWeightKg : number, total : number) {
    let isMale = gender == "male";

    let coeffSettings : CoefficientSettings;
    switch (parseInt(year)) {
      case 2009:
        coeffSettings = { ACoefficient: isMale ? Coefficients.AMale2009 : Coefficients.AFemale2009, BCoefficient: isMale ? Coefficients.BMale2009 : Coefficients.BFemale2009 };
        break
      case 2013:
        coeffSettings = { ACoefficient: isMale ? Coefficients.AMale2013 : Coefficients.AFemale2013, BCoefficient: isMale ? Coefficients.BMale2013 : Coefficients.BFemale2013 };
        break
      case 2017:
        coeffSettings = { ACoefficient: isMale ? Coefficients.AMale2017 : Coefficients.AFemale2017, BCoefficient: isMale ? Coefficients.BMale2017 : Coefficients.BFemale2017 };
        break
      case 2021:
        coeffSettings = { ACoefficient: isMale ? Coefficients.AMale2021 : Coefficients.AFemale2021, BCoefficient: isMale ? Coefficients.BMale2021 : Coefficients.BFemale2021 };
        break
      default:
        coeffSettings = { ACoefficient: isMale ? Coefficients.AMale2021 : Coefficients.AFemale2021, BCoefficient: isMale ? Coefficients.BMale2021 : Coefficients.BFemale2021 };
    }

    return this.doSinclairCalc(bodyWeightKg, total, coeffSettings);
  }
};


function Sinclair() {
  const [sinclair, setSinclair] = useState<number>(0)
  const [bodyweight, setBodyweight] = useState<number>(0)
  const [total, setTotal] = useState<number>(0)
  const [selected, setSelected] = useState<string>("male")
  const [sinclairYear, setSinclairYear] = useState(new Set(["2021"]))

  // @ts-ignore
  return (
    <div>
      <HeaderBar />
      <h1>Sinclair Calculator</h1>

        <Input
          aria-label="Bodyweight"
          type="number"
          placeholder="Bodyweight"
          onChange={(e) => setBodyweight(parseFloat(e.target.value))}
        />

        <Input
          aria-label="Total"
          type="number"
          placeholder="Total"
          onChange={(e) => setTotal(parseFloat(e.target.value))}
        />

        <RadioGroup aria-label="Gender" value={selected} onValueChange={setSelected}>
          <Radio value="male" color="primary">Male</Radio>
          <Radio value="female" color="danger">Female</Radio>
        </RadioGroup>

        <Select
          aria-label="Sinclair Year"
          placeholder="Sinclair Year"
          onChange={(e) => setSinclairYear(new Set(e.target.value))}
        >
          <SelectItem key="2009">Jan 2009 - 2012 Dec</SelectItem>
          <SelectItem key="2013">Jan 2013 - 2016 Dec</SelectItem>
          <SelectItem key="2017">Jan 2017 - 2020 Dec</SelectItem>
          <SelectItem key="2021">Jan 2021 - 2024 Dec</SelectItem>
        </Select>

        <Button onClick={() => setSinclair(SinclairCalculator.getSinclair(Array.from(sinclairYear).join(''), selected, bodyweight, total))}>Calculate</Button>

        <h2>Sinclair Score: {sinclair.toFixed(3)}</h2>
    </div>

  )
}

export default Sinclair