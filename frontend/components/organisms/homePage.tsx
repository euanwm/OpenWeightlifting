import useSWR from 'swr'
import { useState } from 'react'
import { Button, Spinner } from '@nextui-org/react'

import HeaderBar from '@/components/molecules/head'
import fetchLifterData from '@/api/fetchLifterData/fetchLifterData'
import { Filters } from '../molecules/filters'
import { DataTable } from '../molecules/dataTable'
import LifterGraphModal from '../molecules/lifterGraphModal'

const lifterLoadMoreQty = 50

function HomePage() {
  const [start] = useState(0)
  const [stop, setStop] = useState(10)
  const [sortby, setSortBy] = useState('total')
  const [federation, setFederation] = useState('allfeds')
  const [weightclass, setWeightclass] = useState('MALL')
  const [showLifterGraph, setShowLifterGraph] = useState('')
  const [year, setYear] = useState(69)

  const { data, isLoading } = useSWR(
    {
      start,
      stop,
      sortby,
      federation,
      weightclass,
      year,
    },
    fetchLifterData,
    { keepPreviousData: true },
  )

  function handleFilterChange(newFilter: any) {
    const { type, value } = newFilter
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
          Load more results
        </Button>
      </div>
      {showLifterGraph && (
        <LifterGraphModal
          lifterName={showLifterGraph}
          onClose={() => setShowLifterGraph('')}
        />
      )}
    </>
  )
}

export default HomePage
