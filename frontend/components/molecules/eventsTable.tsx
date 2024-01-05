import { EventsList } from '@/api/fetchEventsList/fetchEventsListTypes'
import { Link, Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@nextui-org/react'

export const EventsListTable = ({
  events,
}: {
  events: EventsList
}) => {
  return (
    <Table>
      <TableHeader>
        <TableColumn>Date</TableColumn>
        <TableColumn>Federation</TableColumn>
        <TableColumn>Name</TableColumn>
        <TableColumn>Raw Data</TableColumn>
      </TableHeader>
      <TableBody>
        {events.events.map((event, index) => {
          const { date, name, federation, id } = event
          const event_page = `events/show?fed=${federation}&id=${id}`

          return (
            <TableRow key={`event-${index}`}>
              <TableCell>{date}</TableCell>
              <TableCell>{federation}</TableCell>
              <TableCell>
                <Link href={event_page}>{name}</Link>
              </TableCell>
              <TableCell>
                <Link href={`https://github.com/euanwm/OpenWeightlifting/tree/development/backend/event_data/${federation}/${id}`}>Data</Link>
              </TableCell>
            </TableRow>
          )
        })}
      </TableBody>
    </Table>
  )
}