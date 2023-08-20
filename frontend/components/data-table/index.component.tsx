import { Table, TableCell, TableRow, TableHeader, TableColumn, TableBody, Pagination, Link } from "@nextui-org/react"
import { FaInstagram } from "react-icons/fa"
import { VscGraphLine } from 'react-icons/vsc'
import { CgProfile } from 'react-icons/cg'

import { AllDetails } from "../all-details/index.component"

import styles from './data-table.module.css'
import { LifterResult } from "api/fetchLifterData/fetchLifterDataTypes"

const tableStyle = {
  tableLayout: 'fixed',
  textAlign: 'center',
  width: "auto",
  height: "auto",
  paddingLeft: "0",
  paddingRight: "0",
}

const tableHeaderStyles = {
  textAlign: "center",
  whiteSpace: "nowrap",
  padding: "5px 10px"
}

export const DataTable = ({ lifters, openLifterGraphHandler }: { lifters: LifterResult[], openLifterGraphHandler: (lifterName: string) => void }) => {
  const generateLifterRow = (lifter: LifterResult, lifterNo: number) => {
    const lifter_page = "lifter?name=" + lifter.lifter_name
    return (
      <TableRow key={`lifter-${lifterNo}`}>
        <TableCell>{lifterNo}</TableCell>
        <TableCell>{lifter.lifter_name}</TableCell>
        <TableCell>
          <div className={styles.iconContainer}>
            {lifter.instagram.length > 0 && (
              <a href={`https://www.instagram.com/${lifter.instagram}`} target="_blank" rel="noreferrer noopener"><FaInstagram size={25} /></a>
            )}
            <button onClick={() => openLifterGraphHandler(lifter.lifter_name)} className={styles.graphButton}>
              <VscGraphLine size={25} />
            </button>
            <Link href={lifter_page}><CgProfile size={25} /></Link>
          </div>
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
        <TableColumn> </TableColumn>
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