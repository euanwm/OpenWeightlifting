import { Popover, Table, Button, Text } from "@nextui-org/react"

const tableHeaderStyles = {
    textAlign: "center",
    whiteSpace: "nowrap",
    padding: "5px 10px"
}

export const AllDetails = ({full_comp}) => {
    return (
        <Popover>
            <Popover.Trigger>
                <Button auto flat>More details...</Button>
            </Popover.Trigger>
            <Popover.Content>
                <Table selectionMode="single" striped aria-label="Open weight lifting lifters extended results table" css={{ tableLayout: 'fixed', textAlign: 'center', width: "auto", paddingLeft: "0", paddingRight: "0" }}>
                    <Table.Header>
                        <Table.Column css={tableHeaderStyles}>Event - {full_comp.event}</Table.Column>
                    </Table.Header>
                    <Table.Body>
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
};