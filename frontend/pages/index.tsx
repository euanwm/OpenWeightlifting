import { SWRConfig } from 'swr'

import HomePage from '@/components/organisms/homePage'
import fetchLifterData from '@/api/fetchLifterData/fetchLifterData'
import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'

function Home({ fallback }: { fallback: { leaderBoard: LifterResult[] } }) {
  return (
    <SWRConfig value={{ fallback }}>
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
