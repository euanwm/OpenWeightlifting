import { useRouter } from 'next/router'
import HeaderBar from '@/components/molecules/head'
import { EventMetaData } from '@/api/fetchEventsList/fetchEventsListTypes'
import fetchEventResult from '@/api/fetchEventResult/fetchEventResult'
import useSWR from 'swr'
import { Spinner } from '@nextui-org/react'
import { EventTable } from '@/components/molecules/eventTable'

function ShowEvent(){
  const router = useRouter()
  const { fed } = router.query
  const { id } = router.query

  const requestPayload: EventMetaData = {
    federation: fed as string,
    id: id as string,
    name: '',
    date: ''
  }

  const { data, isLoading } = useSWR(
    requestPayload,
    fetchEventResult,
  )

  return (
    <>
      {isLoading && (
        <div className="fixed w-screen h-screen z-10 flex justify-center items-center">
          <Spinner size="lg" label="Loading..." />
        </div>
      )}
      <HeaderBar />
      {data ? (
        <>
          <center className={'flex flex-col content-center'}>
            <p>Event: {data.data[0].event}</p>
            <p>Federation: {data.data[0].country}</p>
          </center>
          <EventTable history={data.data} />
        </>
      ) : (
        <div>{`No data for event`}</div>
      )}
    </>
  )
}

export default ShowEvent