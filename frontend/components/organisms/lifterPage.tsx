import useSWR from 'swr'
import { useRouter } from 'next/router'
import { Spinner } from '@nextui-org/react'

import { LifterGraph } from '@/components/molecules/lifterGraph'
import { HistoryTable } from '@/components/molecules/historyTable'
import fetchLifterHistory from '../../api/fetchLifterHistory/fetchLifterHistory'
import HeaderBar from '@/components/molecules/head'

function LifterPage() {
  const router = useRouter()
  const { name } = router.query

  const { data, isLoading } = useSWR(
    name,
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
          <center>
            <h1>{data.name}</h1>
          </center>
          <LifterGraph lifterHistory={data.graph} setRatio={1.5} />
          <HistoryTable history={data.lifts} />
        </>
      ) : (
        <div>{`No data for lifter '${name}'`}</div>
      )}
    </>
  )
}

export default LifterPage
