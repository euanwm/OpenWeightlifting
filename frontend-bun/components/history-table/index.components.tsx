import { Table } from '@nextui-org/react'
import { LifterHistory } from '../../api/fetchLifterHistory/fetchLifterHistoryTypes'

export const HistoryTable = ({ history }: { history: LifterHistory['lifts'] }) => {
  return (
    <Table striped aria-label="OpenWeightLifting complete history table">
      <Table.Header>
        <Table.Column>Date</Table.Column>
        <Table.Column>Event</Table.Column>
        <Table.Column>Bodyweight</Table.Column>
        <Table.Column>1st Snatch</Table.Column>
        <Table.Column>2nd Snatch</Table.Column>
        <Table.Column>3rd Snatch</Table.Column>
        <Table.Column>1st C&J</Table.Column>
        <Table.Column>2nd C&J</Table.Column>
        <Table.Column>3rd C&J</Table.Column>
        <Table.Column>Total</Table.Column>
        <Table.Column>Sinclair</Table.Column>
      </Table.Header>
      <Table.Body>
        {history.map((lift, index) => {
          return (
            <Table.Row key={`history-${index}`}>
              <Table.Cell>{lift.date}</Table.Cell>
              <Table.Cell>{lift.event}</Table.Cell>
              <Table.Cell>{lift.bodyweight}</Table.Cell>
              <Table.Cell>{lift.snatch_1}</Table.Cell>
              <Table.Cell>{lift.snatch_2}</Table.Cell>
              <Table.Cell>{lift.snatch_3}</Table.Cell>
              <Table.Cell>{lift.cj_1}</Table.Cell>
              <Table.Cell>{lift.cj_2}</Table.Cell>
              <Table.Cell>{lift.cj_3}</Table.Cell>
              <Table.Cell>{lift.total}</Table.Cell>
              <Table.Cell>{lift.sinclair}</Table.Cell>
            </Table.Row>
          )
        })}
      </Table.Body>
    </Table>
  )
}

