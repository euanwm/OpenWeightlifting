import { Table } from "@nextui-org/react"

const tableHeaderStyles = {
  textAlign: "center",
  whiteSpace: "nowrap",
  padding: "5px 10px"
}

const DataTable = ({ lifters }) => (
  <Table selectionMode="single" striped aria-label="Open weight lifting lifters results table" css={{ tableLayout: 'fixed', textAlign: 'center', width: "auto", paddingLeft: "0", paddingRight: "0" }}>
    <Table.Header>
      <Table.Column css={tableHeaderStyles}>Position</Table.Column>
      <Table.Column css={tableHeaderStyles}>Lifter</Table.Column>
      <Table.Column css={tableHeaderStyles}>Country</Table.Column>
      <Table.Column css={tableHeaderStyles}>Bodyweight</Table.Column>
      <Table.Column css={tableHeaderStyles}>Snatch 1</Table.Column>
      <Table.Column css={tableHeaderStyles}>Snatch 2</Table.Column>
      <Table.Column css={tableHeaderStyles}>Snatch 3</Table.Column>
      <Table.Column css={tableHeaderStyles}>Clean & Jerk 1</Table.Column>
      <Table.Column css={tableHeaderStyles}>Clean & Jerk 2</Table.Column>
      <Table.Column css={tableHeaderStyles}>Clean & Jerk 3</Table.Column>
      <Table.Column css={tableHeaderStyles}>Total</Table.Column>
    </Table.Header>
    <Table.Body>
      {lifters.map((lifter, i) => (
        <Table.Row key={`lifter-${i + 1}`}>
          <Table.Cell>{i + 1}</Table.Cell>
          <Table.Cell>{lifter.lifter_name}</Table.Cell>
          <Table.Cell>{lifter.country}</Table.Cell>
          <Table.Cell>{lifter.bodyweight}</Table.Cell>
          <Table.Cell>{lifter.snatch_1}</Table.Cell>
          <Table.Cell>{lifter.snatch_2}</Table.Cell>
          <Table.Cell>{lifter.snatch_3}</Table.Cell>
          <Table.Cell>{lifter.cj_1}</Table.Cell>
          <Table.Cell>{lifter.cj_2}</Table.Cell>
          <Table.Cell>{lifter.cj_3}</Table.Cell>
          <Table.Cell>{lifter.total}</Table.Cell>
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