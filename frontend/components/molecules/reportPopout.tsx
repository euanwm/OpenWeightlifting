import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'
import { Button, Popover, PopoverContent, PopoverTrigger, Input } from '@nextui-org/react'
import { MdOutlineFmdBad } from "react-icons/md";
import submitResultIssue from '@/api/submitResultIssue/submitResultIssueForm'

export const ReportPopout = ({ singleLift }: { singleLift: LifterResult }) => {
  const form_data = {
    lift_data: singleLift,
    description: ''
  }
  return (
    <div>
      <Popover>
        <PopoverTrigger>
          <Button color="danger">Report Issue</Button>
        </PopoverTrigger>
        <PopoverContent>
          <div className="flex flex-col items-center justify-center space-y-2 p-4">
            <MdOutlineFmdBad size={24} />
            <Input placeholder="Issue Description" />
            <Button color="success" onClick={() => submitResultIssue(form_data)}>Submit</Button>
          </div>
        </PopoverContent>
      </Popover>
    </div>
  )
}

export default ReportPopout