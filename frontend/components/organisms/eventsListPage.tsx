import HeaderBar from '@/components/molecules/head'
import { EventsListRequest } from '@/api/fetchEventsList/fetchEventsListTypes'
import fetchEventsList from '@/api/fetchEventsList/fetchEventsList'
import useSWR from 'swr'
import { EventsListTable } from '@/components/molecules/eventsTable'
import { useState } from 'react'
import { EventsFilters } from '@/components/molecules/eventsfilters'
import { Spinner } from '@nextui-org/react'

const defaultDaysPrevious = 15

function buildPayload(daysPrevious: number): EventsListRequest {
  // todo: fix this bug
  // it's not this function that's causing it, it's the fetchEventsList function
  // the conditional statement below is a workaround
  // judge me if you want, I have no shame
  if (daysPrevious < defaultDaysPrevious) {
    daysPrevious = defaultDaysPrevious
  }

  const today = new Date().toISOString().slice(0, 10)
  const daysPreviousDate = new Date()
  daysPreviousDate.setDate(daysPreviousDate.getDate() - daysPrevious)
  const daysPreviousString = daysPreviousDate.toISOString().slice(0, 10)

  return {
    startdate: daysPreviousString,
    enddate: today,
  }
}


function EventsListPage() {
  const [dayRange, setDayRange] = useState(defaultDaysPrevious)

  const { data, isLoading } = useSWR(
    buildPayload(dayRange),
    fetchEventsList,
    { keepPreviousData: false },
  )

  function handleFilterChange(newFilter: any) {
    const { type, value } = newFilter
    switch (type) {
      case 'dayRange':
        setDayRange(value)
        break
      default:
        setDayRange(value)
    }
  }

  return (
    <>
      {isLoading && (
        <div className="fixed w-screen h-screen z-10 flex justify-center items-center">
          <Spinner size="lg" label="Loading..." />
        </div>
      )}
      <HeaderBar />
      <div className={'flex flex-col content-center'}>
        <EventsFilters handleFilterChange={handleFilterChange} />
        {data ? (
          <div className={'flex flex-col content-center'}>
            <EventsListTable events={data} />
          </div>
        ) : (
          <div>{`No data for events list`}</div>
        )}
      </div>
    </>
  )
}

export default EventsListPage