import { Table, useTheme, Link } from "@nextui-org/react"
import { FaInstagram } from "react-icons/fa"
import { VscGraphLine } from 'react-icons/vsc'
import { CgProfile } from 'react-icons/cg'

import { AllDetails } from "../all-details/index.component"

import styles from './data-table.module.css'
import { LifterResult } from "../../api/fetchLifterData/fetchLifterDataTypes"

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
  const { isDark } = useTheme();
  const themeIconClass = isDark ? styles.themeIconDark : styles.themeIconLight

  const generateLifterRow = (lifter: LifterResult, lifterNo: number) => {
    const lifter_page = "lifter?name=" + lifter.lifter_name
    return (
      <Table.Row key={`lifter-${lifterNo}`}>
        <Table.Cell>{lifterNo}</Table.Cell>
        <Table.Cell>{lifter.lifter_name}</Table.Cell>
        <Table.Cell>
          <div className={styles.iconContainer}>
            {lifter.instagram.length > 0 && (
              <a href={`https://www.instagram.com/${lifter.instagram}`} target="_blank" rel="noreferrer noopener"><FaInstagram size={25} className={themeIconClass} /></a>
            )}
            <button onClick={() => openLifterGraphHandler(lifter.lifter_name)} className={styles.graphButton}>
              <VscGraphLine size={25} className={themeIconClass} />
            </button>
            <Link href={lifter_page}><CgProfile size={25} className={themeIconClass} /></Link>
          </div>
        </Table.Cell>
        <Table.Cell>{lifter.country}</Table.Cell>
        <Table.Cell>{lifter.best_snatch}</Table.Cell>
        <Table.Cell>{lifter.best_cj}</Table.Cell>
        <Table.Cell>{lifter.total}</Table.Cell>
        <Table.Cell>{lifter.sinclair}</Table.Cell>
        <Table.Cell>
          <AllDetails full_comp={lifter}></AllDetails>
        </Table.Cell>
      </Table.Row>
    )
  }

  return (
    <Table striped aria-label="Open weight lifting lifters results table" css={tableStyle}>
      <Table.Header>
        <Table.Column css={tableHeaderStyles}>Rank</Table.Column>
        <Table.Column css={tableHeaderStyles}>Lifter</Table.Column>
        <Table.Column> </Table.Column>
        <Table.Column css={tableHeaderStyles}>Federation</Table.Column>
        <Table.Column css={tableHeaderStyles}>Top Snatch</Table.Column>
        <Table.Column css={tableHeaderStyles}>Top Clean & Jerk</Table.Column>
        <Table.Column css={tableHeaderStyles}>Total</Table.Column>
        <Table.Column css={tableHeaderStyles}>Sinclair</Table.Column>
        <Table.Column css={tableHeaderStyles}>Details</Table.Column>
      </Table.Header>
      <Table.Body>
        {lifters.map((lifter, i) => generateLifterRow(lifter, i + 1))}
      </Table.Body>
      <Table.Pagination
        align="center"
        rowsPerPage={50}
        onPageChange={(page: any) => console.log({ page })}
        css={{ margin: '20px 0' }}
      />
    </Table>
  )
}