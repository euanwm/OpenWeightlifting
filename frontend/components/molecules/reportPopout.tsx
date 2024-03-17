import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'
import { Button, Popover, PopoverContent, PopoverTrigger, Input } from '@nextui-org/react'
import { MdOutlineFmdBad } from "react-icons/md";
import submitResultIssue from '@/api/submitResultIssue/submitResultIssueForm'
import { useState } from 'react'
import { ReportErrorForm } from '@/api/submitResultIssue/submitResultIssueFormTypes'
import useSWR from 'swr'

export const ReportPopout = ({ singleLift }: { singleLift: LifterResult }) => {
  const [popout, setPopout] = useState(false)
  const [comments, setComments] = useState('')
  const [validationErrorMessage, setValidationErrorMessage] = useState('')
  const [previouslySubmitted, setPreviouslySubmitted] = useState(false)
  const [formSubmissionData, setFormSubmissionData] =
    useState<null | ReportErrorForm>(null)

  const { data, isLoading, error } = useSWR(formSubmissionData, submitResultIssue, {
    keepPreviousData: true,
  })

  const handleSubmit = async (e: any) => {
    e.preventDefault()

    if (!comments) {
      setValidationErrorMessage('Please fill out all fields correctly.')
      return
    }

    const data = {
      lift: singleLift,
      comments,
    }
    setFormSubmissionData(data)
    setPreviouslySubmitted(true)
    setPopout(false)
  }

  return (
    <div>
      <Popover isOpen={popout} placement="left" onOpenChange={(e) => setPopout(e)}>
        <PopoverTrigger>
          <Button color="danger" isDisabled={previouslySubmitted}>
            Report Issue
          </Button>
        </PopoverTrigger>
        <PopoverContent>
          <div className="flex flex-col items-center justify-center space-y-2 p-4">
            <MdOutlineFmdBad size={24} />
            <Input placeholder="Issue Description" onChange={(e) => setComments(e.target.value)} />
            <Button onClick={handleSubmit}>Submit</Button>
          </div>
        </PopoverContent>
      </Popover>
    </div>
  )
}

export default ReportPopout