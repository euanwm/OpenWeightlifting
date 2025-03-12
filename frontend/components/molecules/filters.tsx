import { Select, SelectItem } from '@nextui-org/react'
import years from '../../autobuild/filter_years.json'

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
  { value: 'FFH', label: 'France' },
  { value: 'IWF', label: 'IWF' },
  { value: 'OPEN', label: 'OPEN'},
  { value: 'CH', label: 'Switzerland' },
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

const yearsList: {label: string, value: string}[] = years.years.map(year => {
    const [label, value] = Object.entries(year)[0]
    return { value, label }
})

export const Filters = ({ sortBy, federation, handleFilterChange, weightClass, year }: {sortBy: string, federation: string, handleFilterChange: any, weightClass: string, year: string}) => (
  <div className="flex flex-col md:flex-row space-y-1 md:space-y-0 md:space-x-4 mt-4 mx-4">
    <Select
      items={sortByList}
      label="Total/Sinclair"
      placeholder={sortBy.charAt(0).toUpperCase() + sortBy.slice(1)}
      fullWidth={false}
      onChange={
        (e) => handleFilterChange({ type: 'sortBy', value: e.target.value })
      }
      >
      {(sortBy) => <SelectItem key={sortBy.value} value={sortBy.value}>{sortBy.label}</SelectItem>}
    </Select>
    <Select
      items={federationList}
      label="Federation"
      placeholder={federationList.find((fed) => fed.value === federation)?.label}
      fullWidth={false}
      onChange={
        (e) => handleFilterChange({ type: 'federation', value: e.target.value })
      }
      >
      {(federation) => <SelectItem key={federation.value} value={federation.value}>{federation.label}</SelectItem>}
    </Select>
    <Select
      items={weightClassList}
      label="Weight Class"
      placeholder={weightClassList.find((wc) => wc.value === weightClass)?.label}
      fullWidth={false}
      onChange={
        (e) => handleFilterChange({ type: 'weightclass', value: e.target.value })
      }
      >
      {(weightClass) => <SelectItem key={weightClass.value} value={weightClass.value}>{weightClass.label}</SelectItem>}
    </Select>
    <Select
      items={yearsList}
      label="Year"
      placeholder={year === '69' ? 'All Years' : year}
      fullWidth={false}
      onChange={
        (e) => handleFilterChange({ type: 'year', value: parseInt(e.target.value) })
      }
      >
      {(year) => <SelectItem key={year.value} value={year.value}>{year.label}</SelectItem>}
    </Select>
  </div>
)
