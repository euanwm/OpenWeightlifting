import EventsListPage from '@/components/organisms/eventsListPage'
import { SWRConfig } from 'swr'
import fetchEventsList from '@/api/fetchEventsList/fetchEventsList'

function Events({ fallback }: { fallback: { eventsList: [] } }) {
  return (
  <SWRConfig value={{ fallback }}>
    <EventsListPage />
  </SWRConfig>
  )}

export async function getServerSideProps() {
  const data = await fetchEventsList(null)
  return {
    props: {
      fallback: {
        eventsList: data,
      },
    },
  }
}

export default Events
