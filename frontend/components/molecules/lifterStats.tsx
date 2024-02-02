import { LifterStatistics } from '@/api/fetchLifterHistory/fetchLifterHistoryTypes'
import { Table, TableBody, TableCell, TableColumn, TableHeader, TableRow } from '@nextui-org/react'

export const LifterStats = ({ stats }: { stats: LifterStatistics }) => {
  return (
    <Table>
      <TableHeader>
        <TableColumn>Best Snatch</TableColumn>
        <TableColumn>Best C&J</TableColumn>
        <TableColumn>Best Total</TableColumn>
        <TableColumn>Snatch Opener Rate</TableColumn>
        <TableColumn>Snatch 2nd Attempt Rate</TableColumn>
        <TableColumn>Snatch Final Attempt Rate</TableColumn>
        <TableColumn>C&J Opener Rate</TableColumn>
        <TableColumn>C&J 2nd Attempt Rate</TableColumn>
        <TableColumn>C&J Final Attempt Rate</TableColumn>
      </TableHeader>
      <TableBody>
        <TableRow>
          <TableCell>{stats.best_snatch}</TableCell>
          <TableCell>{stats.best_cj}</TableCell>
          <TableCell>{stats.best_total}</TableCell>
          <TableCell>{stats.make_rate_snatches[0]}%</TableCell>
          <TableCell>{stats.make_rate_snatches[1]}%</TableCell>
          <TableCell>{stats.make_rate_snatches[2]}%</TableCell>
          <TableCell>{stats.make_rate_cj[0]}%</TableCell>
          <TableCell>{stats.make_rate_cj[1]}%</TableCell>
          <TableCell>{stats.make_rate_cj[2]}%</TableCell>
        </TableRow>
      </TableBody>
    </Table>
  )
}