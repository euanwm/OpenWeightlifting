import useSWR from 'swr'
import { useState } from 'react'
import { Button, Spinner } from '@nextui-org/react'

import HeaderBar from '@/components/molecules/head'
import fetchLifterData from '@/api/fetchLifterData/fetchLifterData'
import { Filters } from '../molecules/filters'
import { DataTable } from '../molecules/dataTable'
import LifterGraphModal from '../molecules/lifterGraphModal'
import { useRouter } from 'next/router'

const lifterLoadMoreQty = 50

function HomePage() {
  const router = useRouter()
  const params: { [key: string]: string } = {}
  for (const key in router.query) {
    params[key] = router.query[key]?.toString() || ''
  }

  const [sortby, setSortBy] = useState(params.sortby || 'total')
  const [federation, setFederation] = useState(params.federation || 'allfeds')
  const [weightclass, setWeightclass] = useState(params.weightclass || 'MALL')
  const [year, setYear] = useState(params.year || '69')
  const [stop, setStop] = useState(parseInt(params.stop) || lifterLoadMoreQty)
  const [showLifterGraph, setShowLifterGraph] = useState('')


  const { data, isLoading } = useSWR(params, fetchLifterData, { keepPreviousData: true })

  function handleFilterChange(newFilter: any) {
    const { type, value } = newFilter
    router.push({ query: { ...params, [type]: value } })
    switch (type) {
      case 'sortBy':
        setSortBy(value)
        break
      case 'weightclass':
        setWeightclass(value)
        break
      case 'year':
        setYear(value)
        break
      default:
        setFederation(value)
    }
  }

  function updateLifterList() {
    setStop(previous => previous + lifterLoadMoreQty)
    router.push({ query: { ...params, stop: stop + lifterLoadMoreQty } })
  }

  return (
    <>
      {isLoading && (
        <div className="fixed w-screen h-screen z-10 flex justify-center items-center">
          <Spinner size="lg" label="Loading..." />
        </div>
      )}
      <div className={'flex flex-col content-center'}>
        <HeaderBar />
        <Filters
          sortBy={sortby}
          federation={federation}
          weightClass={weightclass}
          year={year}
          handleFilterChange={handleFilterChange}
        />
        {data && (
          <DataTable
            lifters={data}
            openLifterGraphHandler={lifterName =>
              setShowLifterGraph(lifterName)
            }
          />
        )}
        <Button
          className={'flex justify-center'}
          aria-label={'Load more results'}
          color={'primary'}
          onClick={updateLifterList}
          isDisabled={false}
        >
          Showing {stop} / {data?.size} lifters...
        </Button>
      </div>
      {showLifterGraph && (
        <LifterGraphModal
          lifterName={showLifterGraph}
          federation={federation}
          onClose={() => setShowLifterGraph('')}
        />
      )}
    </>
  )
}

export default HomePage
