import { useState, useEffect } from 'react'
import { Modal } from '@nextui-org/react'
import CssBaseline from '@mui/material/CssBaseline'

import HeaderBar from '@/layouts/head'

import { DataTable } from "@/components/datatable"
import { Filters } from "@/components/filters"
import { LifterGraph } from "@/components/liftergraph"

import fetchLifterData from '@/api/fetchLifterData/fetchLifterData'
import fetchLifterGraphData from '@/api/fetchLifterGraphData/fetchLifterGraphData'

import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes';
import { LifterChartData } from '@/api/fetchLifterGraphData/fetchLifterGraphDataTypes';

function Home({ data }: { data: LifterResult[] }) {
  const [sortBy, setSortBy] = useState('total')
  const [federation, setFederation] = useState('allfeds')
  const [weightclass, setWeightclass] = useState('MALL')
  const [year, setYear] = useState(69)
  const [currentLifterList, setCurrentLifterList] = useState<LifterResult[]>(data)
  const [currentLifterName, setCurrentLifterName] = useState('')
  const [showLifterGraph, setShowLifterGraph] = useState(false)
  const [currentLifterGraph, setCurrentLifterGraph] = useState<LifterChartData>()
  const [isGraphLoading, setIsGraphLoading] = useState(true)

  useEffect(() => {
    async function callFetchLifterData() {
      setCurrentLifterList(
        await fetchLifterData(0, 500, sortBy, federation, weightclass, parseInt(String(year))),
      )
    }

    callFetchLifterData();
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
    <div>
    <HeaderBar />
      <Filters sortBy={
        sortBy
      } federation={
        federation
      } handleGenderChange={
        handleGenderChange
      } weightClass={
        weightclass
      } year={
        year
      }/>
    </div>
  )
}

export async function getServerSideProps() {
  const data = await fetchLifterData()
  return { props: { data } }
}

export default Home
