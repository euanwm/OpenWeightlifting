import { Select, SelectItem } from '@nextui-org/react'

const dayRanges = [
  { value: 15, label: '15 days' },
  { value: 30, label: '30 days' },
  { value: 60, label: '60 days' },
]

export const EventsFilters = ({
  handleFilterChange,
}: {
  handleFilterChange: any
}) => (
  <div className="flex flex-col md:flex-row space-y-1 md:space-y-0 md:space-x-4 mt-4 mx-4">
    <Select
      items={dayRanges}
      label="Day Range"
      placeholder={dayRanges[0].label}
      fullWidth={false}
      onChange={e =>
        handleFilterChange({ type: 'dayRange', value: e.target.value })
      }
    >
      {dayRange => (
        <SelectItem key={dayRange.value} value={dayRange.value}>
          {dayRange.label}
        </SelectItem>
      )}
    </Select>
  </div>
)