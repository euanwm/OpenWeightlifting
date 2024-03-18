import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'

export type ReportErrorForm = {
  lift: LifterResult
  comments: string
}

export type FormResponse = {
  success: boolean
  message: string
}