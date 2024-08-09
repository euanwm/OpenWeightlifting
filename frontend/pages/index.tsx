import { SWRConfig } from 'swr'

import HomePage from '@/components/organisms/homePage'
import fetchLifterData from '@/api/fetchLifterData/fetchLifterData'
import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'
import Head from "next/head";

function Home({ fallback }: { fallback: { leaderBoard: LifterResult[] } }) {
  return (
    <SWRConfig value={{ fallback }}>
      <Head>
        <title>OpenWeightlifting Rankings</title>
        <meta
          name="description"
          content="The OpenWeightlifting project aims to create a permanent, accurate, convenient, accessible, open archive of the world's olympic weightlifting data."
        />
      </Head>
      <HomePage />
    </SWRConfig>
  )
}

export async function getServerSideProps() {
  const data = await fetchLifterData({})
  return {
    props: {
      fallback: {
        leaderboard: data,
      },
    },
  }
}

export default Home
