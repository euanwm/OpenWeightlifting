import { Select, SelectItem } from '@nextui-org/react'

// todo: convert to enums?
const sortByList = [
  { value: 'total', label: 'Total' },
  { value: 'sinclair', label: 'Sinclair' },
]

const federationList = [
  { value: 'allfeds', label: 'ALL' },
  { value: 'UK', label: 'UK' },
  { value: 'US', label: 'US' },
  { value: 'AUS', label: 'AUS' },
  { value: 'NVF', label: 'Norway' },
  { value: 'IWF', label: 'IWF' },
  { value: 'OPEN', label: 'OPEN'}
]

const weightClassList = [
  { value: 'MALL', label: 'Men\'s ALL' },
  { value: 'FALL', label: 'Women\'s ALL' },
  { value: 'M55', label: 'Men\'s 55kg' },
  { value: 'M61', label: 'Men\'s 61kg' },
  { value: 'M67', label: 'Men\'s 67kg' },
  { value: 'M73', label: 'Men\'s 73kg' },
  { value: 'M81', label: 'Men\'s 81kg' },
  { value: 'M89', label: 'Men\'s 89kg' },
  { value: 'M96', label: 'Men\'s 96kg' },
  { value: 'M102', label: 'Men\'s 102kg' },
  { value: 'M109', label: 'Men\'s 109kg' },
  { value: 'M109+', label: 'Men\'s +109kg' },
  { value: 'F45', label: 'Women\'s 45kg' },
  { value: 'F49', label: 'Women\'s 49kg' },
  { value: 'F55', label: 'Women\'s 55kg' },
  { value: 'F59', label: 'Women\'s 59kg' },
  { value: 'F64', label: 'Women\'s 64kg' },
  { value: 'F71', label: 'Women\'s 71kg' },
  { value: 'F76', label: 'Women\'s 76kg' },
  { value: 'F81', label: 'Women\'s 81kg' },
  { value: 'F87', label: 'Women\'s 87kg' },
  { value: 'F87+', label: 'Women\'s +87kg' }
]

const yearsList = [
  { value: 69, label: 'All Years' },
  { value: 2015, label: '2015' },
  { value: 2016, label: '2016' },
  { value: 2017, label: '2017' },
  { value: 2018, label: '2018' },
  { value: 2019, label: '2019' },
  { value: 2020, label: '2020' },
  { value: 2021, label: '2021' },
  { value: 2022, label: '2022' },
  { value: 2023, label: '2023' }
]
export const Filters = ({ sortBy, federation, handleGenderChange, weightClass, year }: {sortBy: string, federation: string, handleGenderChange: any, weightClass: string, year: number}) => (
  <div className="flex flex-col md:flex-row space-y-1 md:space-y-0 md:space-x-4 mt-4 mx-4">
    <Select
      items={sortByList}
      label="Total/Sinclair"
      placeholder={sortByList[0].label}
      fullWidth={false}
      onChange={
        (e) => handleGenderChange({ type: 'sortBy', value: e.target.value })
      }
      >
      {(sortBy) => <SelectItem key={sortBy.value} value={sortBy.value}>{sortBy.label}</SelectItem>}
    </Select>
    <Select
      items={federationList}
      label="Federation"
      placeholder={federationList[0].label}
      fullWidth={false}
      onChange={
        (e) => handleGenderChange({ type: 'federation', value: e.target.value })
      }
      >
      {(federation) => <SelectItem key={federation.value} value={federation.value}>{federation.label}</SelectItem>}
    </Select>
    <Select
      items={weightClassList}
      label="Weight Class"
      placeholder={weightClassList[0].label}
      fullWidth={false}
      onChange={
        (e) => handleGenderChange({ type: 'weightclass', value: e.target.value })
      }
      >
      {(weightClass) => <SelectItem key={weightClass.value} value={weightClass.value}>{weightClass.label}</SelectItem>}
    </Select>
    <Select
      items={yearsList}
      label="Year"
      placeholder={yearsList[0].label}
      fullWidth={false}
      onChange={
        (e) => handleGenderChange({ type: 'year', value: parseInt(e.target.value) })
      }
      >
      {(year) => <SelectItem key={year.value} value={year.value}>{year.label}</SelectItem>}
    </Select>
  </div>
)
