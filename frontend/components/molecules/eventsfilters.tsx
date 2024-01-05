import { Select, SelectItem } from '@nextui-org/react'

const dayRanges = [
  { value: 15, label: 'The last 15 days' },
  { value: 30, label: 'The last 30 days' },
  { value: 60, label: 'The last 60 days' },
  { value: 120, label: 'The last 120 days' },
]

export const EventsFilters = ({
  handleFilterChange,
}: {
  handleFilterChange: any
}) => (
  <div className="flex flex-col md:flex-row space-y-1 md:space-y-0 md:space-x-4 mt-4 mx-4">
    <Select
      items={dayRanges}
      label="Range"
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