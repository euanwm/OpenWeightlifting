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
} from '@nextui-org/react'
import { LifterResult } from 'api/fetchLifterData/fetchLifterDataTypes'

const tableHeaderStyles = {
  fontSize: '$sm',
  textAlign: 'center',
  whiteSpace: 'nowrap',
}

const tableBodyStyle = {
  fontSize: '$sm',
}

const tableLayoutStyle = {
  tableLayout: 'fixed',
  textAlign: 'center',
  width: 'auto',
  paddingLeft: '0',
  paddingRight: '0',
}

const tableButtonStyle = {
  margin: '0 auto',
}

// todo: this need to be called on the button click
// todo: clean and jerks arent showing up
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
              {full_comp.event}
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
      </PopoverContent>
    </Popover>
  )
}
