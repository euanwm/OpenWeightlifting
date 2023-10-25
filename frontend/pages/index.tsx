import { useState, useEffect } from 'react'
import { Button, Modal, ModalContent, ModalHeader } from '@nextui-org/react'

import HeaderBar from '@/layouts/head'

import { DataTable } from '@/components/datatable'
import { Filters } from '@/components/filters'
import { LifterGraph } from '@/components/liftergraph'

import fetchLifterData from '@/api/fetchLifterData/fetchLifterData'
import fetchLifterGraphData from '@/api/fetchLifterGraphData/fetchLifterGraphData'

import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'
import { LifterChartData } from '@/api/fetchLifterGraphData/fetchLifterGraphDataTypes'

// I fucking hate this shit
// todo: fix this fucking shit
let maxLifters = 50
const defaultLifterQty = 50

function Home({ data }: { data: LifterResult[] }) {
  const [sortBy, setSortBy] = useState('total')
  const [federation, setFederation] = useState('allfeds')
  const [weightclass, setWeightclass] = useState('MALL')
  const [year, setYear] = useState(69)
  const [currentLifterList, setCurrentLifterList] =
    useState<LifterResult[]>(data)
  const [currentLifterName, setCurrentLifterName] = useState('')
  const [showLifterGraph, setShowLifterGraph] = useState(false)
  const [currentLifterGraph, setCurrentLifterGraph] =
    useState<LifterChartData>()
  const [isGraphLoading, setIsGraphLoading] = useState(true)

  useEffect(() => {
    async function callFetchLifterData() {
      if (maxLifters != defaultLifterQty) {
        maxLifters = defaultLifterQty
      }
      setCurrentLifterList(
        await fetchLifterData(
          0,
          defaultLifterQty,
          sortBy,
          federation,
          weightclass,
          parseInt(String(year)),
        ),
      )
    }

    callFetchLifterData()
  }, [sortBy, federation, weightclass, year])

  // todo: define newFilter type/interface
  const handleGenderChange = (newFilter: any) => {
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

  const openLifterGraphHandler = (lifterName: string) => {
    setIsGraphLoading(true)
    setCurrentLifterName(lifterName)
    fetchLifterGraphData(lifterName)
      .then(data => setCurrentLifterGraph(data))
      .then(() => setShowLifterGraph(true))
      .then(() => setIsGraphLoading(false))
  }

  const closeLifterGraphHandler = () => setShowLifterGraph(false)

  async function updateLifterList(maxLifters: number) {
    setCurrentLifterList(
      await fetchLifterData(
        0,
        maxLifters,
        sortBy,
        federation,
        weightclass,
        parseInt(String(year)),
      ),
    )
  }

  return (
    <div className={'flex flex-col content-center'}>
      <HeaderBar />
      <Filters
        sortBy={sortBy}
        federation={federation}
        handleGenderChange={handleGenderChange}
        weightClass={weightclass}
        year={year}
      />
      {currentLifterList && (
        <DataTable
          lifters={currentLifterList}
          openLifterGraphHandler={openLifterGraphHandler}
        />
      )}

      <Modal
        closeButton
        isOpen={showLifterGraph}
        onClose={closeLifterGraphHandler}
        size={'4xl'}
        placement={'center'}
      >
        <ModalContent>
          <ModalHeader>{currentLifterName}</ModalHeader>
          {isGraphLoading ? (
            <h4>Loading...</h4>
          ) : (
            <LifterGraph lifterHistory={currentLifterGraph} setRatio={1} />
          )}
        </ModalContent>
      </Modal>
      <Button
        className={'flex justify-center'}
        aria-label={'Load more results'}
        color={'primary'}
        onClick={() => updateLifterList((maxLifters += defaultLifterQty))}
        isDisabled={currentLifterList.length < maxLifters}
      >
        Load more results
      </Button>
    </div>
  )
}

export async function getServerSideProps() {
  const data = await fetchLifterData()
  return { props: { data } }
}

export default Home
