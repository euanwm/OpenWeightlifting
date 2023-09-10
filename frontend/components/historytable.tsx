import { Table, TableHeader, TableColumn, TableCell, TableRow, TableBody } from '@nextui-org/react'
import { LifterHistory } from '@/api/fetchLifterHistory/fetchLifterHistoryTypes'

export const HistoryTable = ({ history }: { history: LifterHistory['lifts'] }) => {
  return (
    <Table>
      <TableHeader>
        <TableColumn>Date</TableColumn>
        <TableColumn>Event</TableColumn>
        <TableColumn>Bodyweight</TableColumn>
        <TableColumn>1st Snatch</TableColumn>
        <TableColumn>2nd Snatch</TableColumn>
        <TableColumn>3rd Snatch</TableColumn>
        <TableColumn>1st C&J</TableColumn>
        <TableColumn>2nd C&J</TableColumn>
        <TableColumn>3rd C&J</TableColumn>
        <TableColumn>Total</TableColumn>
        <TableColumn>Sinclair</TableColumn>
      </TableHeader>
      <TableBody>
        {history.map((lift, index) => {
          return (
            <TableRow key={`history-${index}`}>
              <TableCell>{lift.date}</TableCell>
              <TableCell>{lift.event}</TableCell>
              <TableCell>{lift.bodyweight}</TableCell>
              <TableCell>{lift.snatch_1}</TableCell>
              <TableCell>{lift.snatch_2}</TableCell>
              <TableCell>{lift.snatch_3}</TableCell>
              <TableCell>{lift.cj_1}</TableCell>
              <TableCell>{lift.cj_2}</TableCell>
              <TableCell>{lift.cj_3}</TableCell>
              <TableCell>{lift.total}</TableCell>
              <TableCell>{lift.sinclair}</TableCell>
            </TableRow>
          )
        })}
      </TableBody>
    </Table>
  )
}

