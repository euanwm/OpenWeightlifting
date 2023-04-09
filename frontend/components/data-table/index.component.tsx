import { Table } from "@nextui-org/react"
import { FaInstagram } from "react-icons/fa"
import { VscGraphLine } from 'react-icons/vsc'
import { useTheme } from '@nextui-org/react'

import { AllDetails } from "../all-details/index.component"

import styles from './data-table.module.css'
import { LifterResult } from "../../models/api_endpoint";

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
    const { lifter_name, instagram, country, best_snatch, best_cj, total, sinclair } = lifter

    return (
      <Table.Row key={`lifter-${lifterNo}`}>
        <Table.Cell>{lifterNo}</Table.Cell>
        <Table.Cell>{lifter_name}</Table.Cell>
        <Table.Cell>
          <div className={styles.iconContainer}>
            {instagram.length > 0 && (
              <a href={`https://www.instagram.com/${instagram}`} target="_blank" rel="noreferrer noopener"><FaInstagram size={25} className={themeIconClass} /></a>
            )}
            <button onClick={() => openLifterGraphHandler(lifter_name)} className={styles.graphButton}>
              <VscGraphLine size={25} className={themeIconClass} />
            </button>
          </div>
        </Table.Cell>
        <Table.Cell>{country}</Table.Cell>
        <Table.Cell>{best_snatch}</Table.Cell>
        <Table.Cell>{best_cj}</Table.Cell>
        <Table.Cell>{total}</Table.Cell>
        <Table.Cell>{sinclair}</Table.Cell>
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
        onPageChange={(page) => console.log({ page })}
        css={{ margin: '20px 0' }}
      />
    </Table>
  )
}