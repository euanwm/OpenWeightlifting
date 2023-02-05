import { Popover, Table, Button } from "@nextui-org/react"

const tableHeaderStyles = {
  fontSize: '$sm',
  textAlign: "center",
  whiteSpace: "nowrap",
}

const tableBodyStyle = {
  fontSize: '$sm'
}

const tableLayoutStyle = {
  tableLayout: 'fixed',
  textAlign: 'center',
  width: "auto",
  paddingLeft: "0",
  paddingRight: "0"
}

const tableButtonStyle = {
  margin: "0 auto"
}

export const AllDetails = ({ full_comp }) => (
  <Popover placement="left">
    <Popover.Trigger>
      <Button auto flat css={tableButtonStyle}>More details...</Button>
    </Popover.Trigger>
    <Popover.Content>
      <Table striped aria-label="Open weight lifting lifters extended results table" css={tableLayoutStyle}>
        <Table.Header>
          <Table.Column css={tableHeaderStyles}>{full_comp.event}</Table.Column>
        </Table.Header>
        <Table.Body css={tableBodyStyle}>
          <Table.Row>
            <Table.Cell>Event Date: {full_comp.date}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Bodyweight: {full_comp.bodyweight}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Category: {full_comp.gender}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Opening Snatch: {full_comp.snatch_1}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Second Snatch: {full_comp.snatch_2}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Final Snatch: {full_comp.snatch_3}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Opening C&J: {full_comp.cj_1}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Second C&J: {full_comp.cj_2}</Table.Cell>
          </Table.Row><Table.Row>
            <Table.Cell>Final C&J: {full_comp.cj_3}</Table.Cell>
          </Table.Row>
        </Table.Body>
      </Table>
    </Popover.Content>
  </Popover>
);
