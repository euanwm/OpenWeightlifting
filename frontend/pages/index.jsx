import { useState, useEffect } from 'react'
import { useTheme } from '@nextui-org/react'
import { ThemeProvider, createTheme } from '@mui/material/styles'
import CssBaseline from '@mui/material/CssBaseline'

import DataTable from '../components/data-table/index.component'
import Filters from '../components/filters/index.component'

const fetchLifterData = async (
  gender = 'male',
  start = 0,
  stop = 500,
  sortby = 'total',
  federation = 'allfeds',
) => {
  const bodyContent = JSON.stringify({
    gender,
    start,
    stop,
    sortby,
    federation,
  })

  const res = await fetch(`${process.env.API}/leaderboard`, {
    method: 'POST',
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
    body: bodyContent,
  }).catch(error => console.error(error))

  return await res.json()
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

const Home = ({ data }) => {
  const [currentGender, setCurrentGender] = useState('male')
  const [sortBy, setSortBy] = useState('total')
  const [federation, setFederation] = useState('allfeds')
  const [currentLifterList, setCurrentLifterList] = useState(data)
  const { isDark } = useTheme()

  useEffect(() => {
    async function callFetchLisfterData() {
      setCurrentLifterList(
        await fetchLifterData(currentGender, 0, 500, sortBy, federation),
      )
    }

    callFetchLisfterData()
  }, [currentGender, sortBy, federation])

  const handleGenderChange = newFilter => {
    const { type, value } = newFilter
    switch (type) {
      case 'gender':
        setCurrentGender(value)
        break
      case 'sortBy':
        setSortBy(value)
        break
      default:
        setFederation(value)
    }
  }

  return (
    <ThemeProvider theme={isDark ? darkTheme : lightTheme}>
      <CssBaseline />
      <Filters
        currentGender={currentGender}
        sortBy={sortBy}
        federation={federation}
        handleGenderChange={handleGenderChange}
      />
      {currentLifterList && <DataTable lifters={currentLifterList} />}
    </ThemeProvider>
  )
}

export async function getServerSideProps() {
  let data;
  try {
    data = await fetchLifterData()
  } catch {
    data = []
  }

  return { props: { data } }
}

export default Home
