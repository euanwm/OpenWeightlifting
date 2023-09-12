import { Table, TableCell, TableRow, TableHeader, TableColumn, TableBody, Link, Divider } from "@nextui-org/react"
import { FaInstagram } from "react-icons/fa"
import { VscGraphLine } from 'react-icons/vsc'
import { CgProfile } from 'react-icons/cg'

import { AllDetails } from "./alldetails"

import { LifterResult } from "@/api/fetchLifterData/fetchLifterDataTypes"

export const DataTable = ({ lifters, openLifterGraphHandler }: { lifters: LifterResult[], openLifterGraphHandler: (lifterName: string) => void }) => {
  const generateLifterRow = (lifter: LifterResult, lifterNo: number) => {
    const lifter_page = "lifter?name=" + lifter.lifter_name
    return (
      <TableRow key={`lifter-${lifterNo}`}>
        <TableCell>{lifterNo}</TableCell>
        <TableCell className="flex flex-row items-center">
          {lifter.lifter_name}
          <Divider orientation="vertical"
            style={{ marginLeft: '0.5rem', marginRight: '0.5rem' }}
          />
          {lifter.instagram.length > 0 && (
            <a href={`https://www.instagram.com/${lifter.instagram}`} ><FaInstagram size={25} /></a>
          )}
          <Divider orientation="vertical" />
          <button onClick={() => openLifterGraphHandler(lifter.lifter_name)}>
            <VscGraphLine size={25} />
          </button>
          <Divider orientation="vertical" />
          <Link href={lifter_page}><CgProfile size={25} /></Link>
        </TableCell>
        <TableCell>{lifter.country}</TableCell>
        <TableCell>{lifter.best_snatch}</TableCell>
        <TableCell>{lifter.best_cj}</TableCell>
        <TableCell>{lifter.total}</TableCell>
        <TableCell>{lifter.sinclair}</TableCell>
        <TableCell>
          <AllDetails full_comp={lifter}></AllDetails>
        </TableCell>
      </TableRow>
    )
  }

  return (
    <Table aria-label="Open weight lifting lifters results table">
      <TableHeader>
        <TableColumn>Rank</TableColumn>
        <TableColumn>Lifter</TableColumn>
        <TableColumn>Federation</TableColumn>
        <TableColumn>Top Snatch</TableColumn>
        <TableColumn>Top Clean & Jerk</TableColumn>
        <TableColumn>Total</TableColumn>
        <TableColumn>Sinclair</TableColumn>
        <TableColumn>Details</TableColumn>
      </TableHeader>
      <TableBody>
        {lifters.map((lifter, i) => generateLifterRow(lifter, i + 1))}
      </TableBody>
    </Table>
  )
}