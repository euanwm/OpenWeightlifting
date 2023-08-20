import { useRouter } from 'next/router'
import { useEffect, useState } from 'react'
import { TableRow } from '@nextui-org/react'

import { LifterGraph } from '../components/lifter-graph/index.component'
import { HistoryTable } from '../components/history-table/index.components'
import fetchLifterHistory from 'api/fetchLifterHistory/fetchLifterHistory'
import { LifterHistory } from 'api/fetchLifterHistory/fetchLifterHistoryTypes'

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
        <TableRow>
          <u>
            <h1>{lifterHistory['name']}</h1>
          </u>
        </TableRow>
      <LifterGraph lifterHistory={lifterHistory['graph']} />
      <HistoryTable history={lifterHistory['lifts']} />
    </div>
  )
}

export default Lifter
