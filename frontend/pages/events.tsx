import EventsListPage from '@/components/organisms/eventsListPage'
import { SWRConfig } from 'swr'
import fetchEventsList from '@/api/fetchEventsList/fetchEventsList'
import Head from "next/head";

function Events({ fallback }: { fallback: { eventsList: [] } }) {
  return (
  <SWRConfig value={{ fallback }}>
    <Head>
      <title>Recent Events</title>
      <meta
        name="description"
        content="The most recent events in the OpenWeightlifting database."
      />
    </Head>
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
