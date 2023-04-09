import { useState, useEffect } from 'react'
import { useTheme, Modal } from '@nextui-org/react'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'

import { DataTable } from '../components/data-table/index.component'
import { Filters } from '../components/filters/index.component'
import { LifterGraph } from '../components/lifter-graph/index.component'

import { LifterChartData, MainLeaderboard } from "../models/api_endpoint";

const fetchLifterData = async (
  // todo: implement enums for these
  start = 0,
  stop = 500,
  sortby = 'total',
  federation = 'allfeds',
  weightclass = 'MALL',
  year = 69
) => {
  const bodyContent = JSON.stringify({
    start,
    stop,
    sortby,
    federation,
    weightclass,
    year,
  })

  const res = await fetch(`${process.env.API}/leaderboard`, {
    method: 'POST',
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
    body: bodyContent,
  }).then((res) => res.json()).catch(error => console.error('error in fetchLifterData', error))

  return await res as MainLeaderboard
}

const fetchLifterGraphData = async (lifterName: string) => {
  const bodyContent = JSON.stringify({
    "NameStr": lifterName
  })

  const res = await fetch(`${process.env.API}/lifter`, {
    method: 'POST',
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
    body: bodyContent,
  }).then((res) => res.json()).catch(error => console.error('error in fetchLifterGraphData', error))

  return await res as LifterChartData
}

const darkTheme = createTheme({
  palette: {
    mode: 'dark',
  },
})

const lightTheme = createTheme({
  palette: {
    mode: 'light',
  },
})

const Home = ({ data }: { data: MainLeaderboard }) => {
  const [sortBy, setSortBy] = useState('total')
  const [federation, setFederation] = useState('allfeds')
  const [weightclass, setWeightclass] = useState('MALL')
  const [year, setYear] = useState(69)
  const [currentLifterList, setCurrentLifterList] = useState(data)
  const [currentLifterName, setCurrentLifterName] = useState('')
  const [showLifterGraph, setShowLifterGraph] = useState(false)
  const [currentLifterGraph, setCurrentLifterGraph] = useState({} as LifterChartData)
  const [isGraphLoading, setIsGraphLoading] = useState(true)
  const { isDark } = useTheme()

  useEffect(() => {
    async function callFetchLifterData() {
      setCurrentLifterList(
        await fetchLifterData(0, 500, sortBy, federation, weightclass, parseInt(String(year))),
      )
    }

    callFetchLifterData().then()
  }, [sortBy, federation, weightclass, year])

  // todo: define newFilter type/interface
  const handleGenderChange = (newFilter: any) => {
    const { type, value } = newFilter
    console.log(type, value)
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
      .then((data) => setCurrentLifterGraph(data))
      .then(() => setShowLifterGraph(true))
      .then(() => setIsGraphLoading(false))
  }

  const closeLifterGraphHandler = () => setShowLifterGraph(false)

  return (
    <>
      <ThemeProvider theme={isDark ? darkTheme : lightTheme}>
        <CssBaseline />
        <Filters
          sortBy={sortBy}
          federation={federation}
          handleGenderChange={handleGenderChange}
          weightClass={weightclass}
          year={year}
        />
        {currentLifterList && <DataTable lifters={currentLifterList} openLifterGraphHandler={openLifterGraphHandler} />}
      </ThemeProvider>
      <Modal closeButton blur open={showLifterGraph} onClose={closeLifterGraphHandler} width='1000'>
        <h3>{currentLifterName}: History (Total)</h3>
        {isGraphLoading ? (
          <h4>Loading...</h4>
        ) : (
          <LifterGraph lifterHistory={currentLifterGraph} />
        )}
      </Modal>
    </>
  )
}

export async function getServerSideProps() {
  let data: MainLeaderboard
  try {
    data = await fetchLifterData()
  } catch {
    data = { lifts: [] }
  }

  return { props: { data } }
}

export default Home
