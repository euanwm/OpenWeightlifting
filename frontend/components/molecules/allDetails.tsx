import {
  Popover,
  Table,
  Button,
  TableRow,
  TableBody,
  TableCell,
  TableColumn,
  TableHeader,
  PopoverContent,
  PopoverTrigger,
  Link,
} from '@nextui-org/react'
import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'

import ReportPopout from '@/components/molecules/reportPopout'


export const AllDetails = ({ full_comp }: { full_comp: LifterResult }) => {
  return (
    <Popover placement="left">
      <PopoverTrigger>
        <Button>
          More details...
        </Button>
      </PopoverTrigger>
      <PopoverContent>
        <Table
          aria-label="Open weight lifting lifters extended results table"
        >
          <TableHeader>
            <TableColumn>
              <Link href={`/events/show?name=${full_comp.event}`}>
                {full_comp.event}
              </Link>
            </TableColumn>
          </TableHeader>
          <TableBody>
            <TableRow>
              <TableCell>Event Date: {full_comp.date}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Bodyweight: {full_comp.bodyweight}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Category: {full_comp.gender}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Opening Snatch: {full_comp.snatch_1}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Second Snatch: {full_comp.snatch_2}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Final Snatch: {full_comp.snatch_3}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Opening C&J: {full_comp.cj_1}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Second C&J: {full_comp.cj_2}</TableCell>
            </TableRow>
            <TableRow>
              <TableCell>Final C&J: {full_comp.cj_3}</TableCell>
            </TableRow>
          </TableBody>
        </Table>
        <ReportPopout singleLift={full_comp} page_origin="leaderboard" />
      </PopoverContent>
    </Popover>
  )
}
