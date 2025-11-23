'use client'

import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'
import { getAPI } from '../loadSplitter'

export default async function submitResultIssue(data: {
  lift_data: LifterResult
  description: string
}): Promise<{ success: boolean; message: string }> {
  try {
    const response = await fetch(`${getAPI()}/issue`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
    return response.json()
  } catch (error) {
    console.error('Error:', error)
    return {
      success: false,
      message: 'There was an error submitting your form. Please try again.',
    }
  }
}
