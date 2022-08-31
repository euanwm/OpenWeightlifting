import { Table } from "@nextui-org/react"

const DataTable = ({ lifters }) => (
  <Table selectionMode="single" striped aria-label="Open weight lifting lifters results table">
    <Table.Header>
      <Table.Column>Position</Table.Column>
      <Table.Column>Lifter</Table.Column>
      <Table.Column>Country</Table.Column>
      <Table.Column>Bodyweight</Table.Column>
      <Table.Column>Snatch 1</Table.Column>
      <Table.Column>Snatch 2</Table.Column>
      <Table.Column>Snatch 3</Table.Column>
      <Table.Column>Clean & Jerk 1</Table.Column>
      <Table.Column>Clean & Jerk 2</Table.Column>
      <Table.Column>Clean & Jerk 3</Table.Column>
      <Table.Column>Total</Table.Column>
    </Table.Header>
    <Table.Body>
      { lifters.map((lifter, i) => (
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
  </Table>
)

export default DataTable