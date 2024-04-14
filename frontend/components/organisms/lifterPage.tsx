import useSWR from 'swr'
import { useRouter } from 'next/router'
import { Spinner } from '@nextui-org/react'

import { LifterStats } from '@/components/molecules/lifterStats'
import { LifterGraph } from '@/components/molecules/lifterGraph'
import { HistoryTable } from '@/components/molecules/historyTable'
import fetchLifterHistory from '../../api/fetchLifterHistory/fetchLifterHistory'
import HeaderBar from '@/components/molecules/head'

function LifterPage() {
  const router = useRouter()
  const params: { [key: string]: string } = {}
  for (const key in router.query) {
    params[key] = router.query[key]?.toString() || ''
  }

  const { data, isLoading } = useSWR(
    params,
    fetchLifterHistory,
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
          <center className={'text-4xl'}>
            <h1>{data.name}</h1>
          </center>
          <LifterStats stats={data.stats} />
          <LifterGraph lifterHistory={data.graph} setRatio={1.5} />
          <HistoryTable history={data.lifts} />
        </>
      ) : (
        <div>{`No data for lifter '${params['name']}'`}</div>
      )}
    </>
  )
}

export default LifterPage
