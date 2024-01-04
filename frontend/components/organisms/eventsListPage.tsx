import HeaderBar from '@/components/molecules/head'
import { useRouter } from 'next/router'
import { EventsListRequest, EventMetaData } from '@/api/fetchEventsList/fetchEventsListTypes'
import fetchEventsList from '@/api/fetchEventsList/fetchEventsList'
import useSWR from 'swr'
import { Spinner } from '@nextui-org/react'

function EventsListPage() {
  const router = useRouter()
  const { fed } = router.query
  const { id } = router.query

  const today = new Date().toISOString().slice(0, 10)
  const fifteenDaysAgo = new Date()
  fifteenDaysAgo.setDate(fifteenDaysAgo.getDate() - 30)
  const fifteenDaysAgoString = fifteenDaysAgo.toISOString().slice(0, 10)

  const eventsListPayload: EventsListRequest = {
    startdate: fifteenDaysAgoString,
    enddate: today,
  }

  const { data: eventsList, isLoading } = useSWR(
    eventsListPayload,
    fetchEventsList,
  )

  return (
    <>
      {isLoading && (
        <div className="fixed w-screen h-screen z-10 flex justify-center items-center">
          <Spinner size="lg" label="Loading..." />
        </div>
      )}
    <HeaderBar />
    <div className={'flex flex-col content-center'}>
      <h1>Events Page</h1>
      {eventsList ? (
        <div className={'flex flex-col content-center'}>
          <h2>Events List</h2>
          <ul>
            {eventsList.events.map((event: EventMetaData) => (
              <li key={event.id}>
                <a href={`/events/show?fed=${event.federation}&id=${event.id}`}>{event.name}</a>
              </li>
            ))}
          </ul>
        </div>
      ) : (
        <div>{`No data for events list`}</div>
      )}
    </div>
    </>
  )
}

export default EventsListPage