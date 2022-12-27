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
                        <Table.Column css={tableHeaderStyles}>Event</Table.Column>
                        <Table.Column css={tableHeaderStyles}>Bodyweight</Table.Column>
                    </Table.Header>
                    <Table.Body>
                        <Table.Row>
                            <Table.Cell>{full_comp.event}</Table.Cell>
                            <Table.Cell>{full_comp.bodyweight}</Table.Cell>
                        </Table.Row>
                    </Table.Body>
                </Table>
            </Popover.Content>
        </Popover>
    );
};