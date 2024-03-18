import { LifterResult } from '@/api/fetchLifterData/fetchLifterDataTypes'


export default async function submitResultIssue(
  data: {
    lift_data: LifterResult
    description: string
  },
): Promise<{ success: boolean; message: string }> {
  try {
    const response = await fetch(`${process.env.API}/issue`, {
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