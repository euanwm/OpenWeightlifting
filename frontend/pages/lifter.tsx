import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'

import { LifterGraph } from "@/components/liftergraph"
import { HistoryTable } from "@/components/historytable"
import fetchLifterHistory from '../api/fetchLifterHistory/fetchLifterHistory'
import { LifterHistory } from "@/api/fetchLifterHistory/fetchLifterHistoryTypes"
import HeaderBar from "@/layouts/head";

function Lifter() {
  const router = useRouter()
  const { name } = router.query

  const [lifterHistory, setLifterHistory] = useState<LifterHistory>({
    name: '',
    graph: {
      labels: [],
      datasets: [],
    },
    lifts: [],
  })

  useEffect(() => {
    async function fetchLifterHistoryFromAPI() {
      const response = await fetchLifterHistory(name)
      setLifterHistory(response)
    }

    fetchLifterHistoryFromAPI()
  }, [name])

  return (
    <div>
      <HeaderBar />
      <center>
        <h1>{lifterHistory['name']}</h1>
      </center>
      <LifterGraph lifterHistory={lifterHistory['graph']} setRatio={1.5} />
      <HistoryTable history={lifterHistory['lifts']} />
    </div>
  )
}

export default Lifter
