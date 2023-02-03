import { Table } from "@nextui-org/react"
import { AllDetails } from "./AllDetails"
import { IconContext } from 'react-icons'
import { FaInstagram } from "react-icons/fa";

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

const DataTable = ({ lifters }) => (
  <Table striped aria-label="Open weight lifting lifters results table" css={tableStyle}>
    <Table.Header>
      <Table.Column css={tableHeaderStyles}>Rank</Table.Column>
      <Table.Column css={tableHeaderStyles}>Lifter</Table.Column>
      <Table.Column></Table.Column>
      <Table.Column css={tableHeaderStyles}>Federation</Table.Column>
      <Table.Column css={tableHeaderStyles}>Top Snatch</Table.Column>
      <Table.Column css={tableHeaderStyles}>Top Clean & Jerk</Table.Column>
      <Table.Column css={tableHeaderStyles}>Total</Table.Column>
      <Table.Column css={tableHeaderStyles}>Sinclair</Table.Column>
      <Table.Column css={tableHeaderStyles}>Details</Table.Column>
    </Table.Header>
    <Table.Body>
      {lifters.map((lifter, i) => (
        <Table.Row key={`lifter-${i + 1}`}>
          <Table.Cell>{i + 1}</Table.Cell>
          <Table.Cell align='center'>{lifter.lifter_name}</Table.Cell>
            <Table.Cell>
                <IconContext.Provider value={{ style: { verticalAlign: 'middle' } }}>
                {lifter.instagram.length > 0 &&
                    <button onClick={event => window.open('https://www.instagram.com/' + lifter.instagram, '_blank')}>
                        <FaInstagram size={20} />
                    </button>
                }
                </IconContext.Provider>
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
      ))}
    </Table.Body>
    <Table.Pagination
      align="center"
      rowsPerPage={50}
      onPageChange={(page) => console.log({ page })}
      css={{ margin: '20px 0' }}
    />
  </Table>
)

export default DataTable