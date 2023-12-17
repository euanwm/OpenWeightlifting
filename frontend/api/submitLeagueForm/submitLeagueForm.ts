import { RegistrationForm, FormResponse } from './submitLeagueFormTypes'

export default async function submitLeagueForm(
  data: RegistrationForm,
): Promise<FormResponse> {
  try {
    const response = await fetch('https://v2.openweightlifting.org/register', {
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
