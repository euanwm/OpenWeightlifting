import { Select, SelectItem, SelectSection } from '@nextui-org/react'
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
  { value: 'OPEN', label: 'OPEN' },
  { value: 'CH', label: 'Switzerland' },
]

const weightClassList2 = {
  '': [
    { value: 'MALL', label: "Men's ALL" },
    { value: 'FALL', label: "Women's ALL" },
  ],
  'New 2025 Categories': [
    { value: 'M60', label: "Men's 60kg" },
    { value: 'M65', label: "Men's 65kg" },
    { value: 'M71', label: "Men's 71kg" },
    { value: 'M79', label: "Men's 79kg" },
    { value: 'M88', label: "Men's 88kg" },
    { value: 'M94', label: "Men's 94kg" },
    { value: 'M110', label: "Men's 110kg" },
    { value: 'M110+', label: "Men's +110kg" },
    { value: 'F48', label: "Women's 48kg" },
    { value: 'F53', label: "Women's 53kg" },
    { value: 'F58', label: "Women's 58kg" },
    { value: 'F63', label: "Women's 63kg" },
    { value: 'F69', label: "Women's 69kg" },
    { value: 'F77', label: "Women's 77kg" },
    { value: 'F86', label: "Women's 86kg" },
    { value: 'F86+', label: "Women's +86kg" },
  ],
  '2018 - 2025': [
    { value: 'M55', label: "Men's 55kg" },
    { value: 'M61', label: "Men's 61kg" },
    { value: 'M67', label: "Men's 67kg" },
    { value: 'M73', label: "Men's 73kg" },
    { value: 'M81', label: "Men's 81kg" },
    { value: 'M89', label: "Men's 89kg" },
    { value: 'M96', label: "Men's 96kg" },
    { value: 'M102', label: "Men's 102kg" },
    { value: 'M109', label: "Men's 109kg" },
    { value: 'M109+', label: "Men's +109kg" },
    { value: 'F45', label: "Women's 45kg" },
    { value: 'F49', label: "Women's 49kg" },
    { value: 'F55', label: "Women's 55kg" },
    { value: 'F59', label: "Women's 59kg" },
    { value: 'F64', label: "Women's 64kg" },
    { value: 'F71', label: "Women's 71kg" },
    { value: 'F76', label: "Women's 76kg" },
    { value: 'F81', label: "Women's 81kg" },
    { value: 'F87', label: "Women's 87kg" },
    { value: 'F87+', label: "Women's +87kg" },
  ],
  '1998 - 2018': [
    { value: 'M56', label: "Men's 56kg" },
    { value: 'M62', label: "Men's 62kg" },
    { value: 'M69', label: "Men's 69kg" },
    { value: 'M77', label: "Men's 77kg" },
    { value: 'M85', label: "Men's 85kg" },
    { value: 'M94-98', label: "Men's 94kg" },
    { value: 'M105', label: "Men's 105kg" },
    { value: 'M105+', label: "Men's +105kg" },
    { value: 'F48-98', label: "Women's 48kg" },
    { value: 'F53-98', label: "Women's 53kg" },
    { value: 'F58-98', label: "Women's 58kg" },
    { value: 'F63-98', label: "Women's 63kg" },
    { value: 'F69-98', label: "Women's 69kg" },
    { value: 'F75', label: "Women's 75kg" },
    { value: 'F75+', label: "Women's +75kg" },
    { value: 'F90', label: "Women's 90kg" },
    { value: 'F90+', label: "Women's +90kg" },
  ],
  '1993 - 1998': [
    { value: 'M54', label: "Men's 54kg" },
    { value: 'M59', label: "Men's 59kg" },
    { value: 'M64', label: "Men's 64kg" },
    { value: 'M70', label: "Men's 70kg" },
    { value: 'M76', label: "Men's 76kg" },
    { value: 'M83', label: "Men's 83kg" },
    { value: 'M91', label: "Men's 91kg" },
    { value: 'M99', label: "Men's 99kg" },
    { value: 'M108', label: "Men's 108kg" },
    { value: 'M108+', label: "Men's +108kg" },
    { value: 'F46', label: "Women's 46kg" },
    { value: 'F50', label: "Women's 50kg" },
    { value: 'F54', label: "Women's 54kg" },
    { value: 'F59-93', label: "Women's 59kg" },
    { value: 'F64-93', label: "Women's 64kg" },
    { value: 'F70', label: "Women's 70kg" },
    { value: 'F76', label: "Women's 76kg" },
    { value: 'F83', label: "Women's 83kg" },
    { value: 'F83+', label: "Women's +83kg" },
  ],
}

const yearsList: { label: string; value: string }[] = years.years.map(year => {
  const [label, value] = Object.entries(year)[0]
  return { value, label }
})

export const Filters = ({
  sortBy,
  federation,
  handleFilterChange,
  weightClass,
  year,
}: {
  sortBy: string
  federation: string
  handleFilterChange: any
  weightClass: string
  year: string
}) => (
  <div className="flex flex-col md:flex-row space-y-1 md:space-y-0 md:space-x-4 mt-4 mx-4">
    <Select
      items={sortByList}
      label="Total/Sinclair"
      placeholder={sortBy.charAt(0).toUpperCase() + sortBy.slice(1)}
      fullWidth={false}
      onChange={e =>
        handleFilterChange({ type: 'sortBy', value: e.target.value })
      }
    >
      {sortBy => (
        <SelectItem key={sortBy.value} value={sortBy.value}>
          {sortBy.label}
        </SelectItem>
      )}
    </Select>
    <Select
      items={federationList}
      label="Federation"
      placeholder={federationList.find(fed => fed.value === federation)?.label}
      fullWidth={false}
      onChange={e =>
        handleFilterChange({ type: 'federation', value: e.target.value })
      }
    >
      {federation => (
        <SelectItem key={federation.value} value={federation.value}>
          {federation.label}
        </SelectItem>
      )}
    </Select>
    <Select
      items={Object.entries(weightClassList2)}
      label="Weight Class"
      placeholder={
        Object.values(weightClassList2)
          .flat()
          .find(wc => wc.value === weightClass)?.label
      }
      fullWidth={false}
      onChange={e =>
        handleFilterChange({ type: 'weightclass', value: e.target.value })
      }
    >
      {([section, items]) => (
        <SelectSection key={section} title={section} showDivider>
          {items.map(weightClass => (
            <SelectItem key={weightClass.value} value={weightClass.value}>
              {weightClass.label}
            </SelectItem>
          ))}
        </SelectSection>
      )}
    </Select>
    <Select
      items={yearsList}
      label="Year"
      placeholder={year === '69' ? 'All Years' : year}
      fullWidth={false}
      onChange={e =>
        handleFilterChange({ type: 'year', value: parseInt(e.target.value) })
      }
    >
      {year => (
        <SelectItem key={year.value} value={year.value}>
          {year.label}
        </SelectItem>
      )}
    </Select>
  </div>
)
