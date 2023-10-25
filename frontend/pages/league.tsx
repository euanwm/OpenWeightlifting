import {
  Button,
  Input,
  Popover,
  PopoverContent,
  PopoverTrigger,
} from '@nextui-org/react'
import HeaderBar from '@/layouts/head'
import { useState } from 'react'

type RegistrationForm = {
  username: string
  clubname: string
  location: string
  email: string
  comments: string
}

type FormResponse = {
  success: boolean
  message: string
}

async function submitForm(data: RegistrationForm): Promise<FormResponse> {
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

function checkEmail(email: string) {
  const re = /\S+@\S+\.\S+/
  return re.test(email)
}

function validateForm(data: RegistrationForm) {
  const { username, clubname, location, email } = data
  if (username.length < 1) {
    return false
  }
  if (clubname.length < 1) {
    return false
  }
  if (location.length < 1) {
    return false
  }
  return checkEmail(email)
}

function League() {
  const [username, setUsername] = useState('')
  const [clubname, setClubname] = useState('')
  const [location, setLocation] = useState('')
  const [email, setEmail] = useState('')
  const [comments, setComments] = useState('')
  const [submitted, setSubmitted] = useState(false)
  const [error, setError] = useState(false)
  const [errorMessage, setErrorMessage] = useState('')
  const [loading, setLoading] = useState(false)
  const [disableButton, setDisableButton] = useState(false)

  const handleSubmit = async (e: any) => {
    e.preventDefault()
    setLoading(true)

    if (!validateForm({ username, clubname, location, email, comments })) {
      setError(true)
      setErrorMessage('Please fill out all fields correctly.')
      setLoading(false)
      return
    }
    const data = {
      username,
      clubname,
      location,
      email,
      comments,
    }

    const response = await submitForm(data)
    if (response.success && response.message === 'Thanks for registering!') {
      setDisableButton(true)
      setSubmitted(true)
      setLoading(false)
    } else if (
      response.success &&
      response.message === 'You have already registered!'
    ) {
      setDisableButton(true)
      setSubmitted(true)
      setLoading(false)
    } else {
      setError(true)
      setErrorMessage(response.message)
      setLoading(false)
    }
  }

  return (
    <>
      <HeaderBar />
      <div className="flex justify-center mt-4 mb-8">
        <div className="w-4/5 md:w-2/3 xl:w-1/2 ">
          <div className="mb-6">
            <h1 className="text-3xl text-center font-bold mb-2">The League</h1>
            <p className="text-center">
              OpenWeightlifting are pleased to announce that we are now planning
              our first - club run - league. We aim to run the league during the
              first 2 weeks of December 2023 and look to run it in a similar
              format to virtual competitions hosted by other federations.
            </p>
          </div>

          <div className="mb-6">
            <h2 className="text-2xl text-center font-bold  mb-2">
              How it works
            </h2>
            <p className="text-center">
              The league will be run over 2 weeks, and each registered affiliate
              will be responsible for recording and uploading their gym /club
              scores. In a typical competition format, you have 6 attempts in
              total, 3 for snatch and 3 for clean & jerk. Each attempt will be
              recorded from the front and the registered club affiliate will be
              responsible for your weigh-in and logging your scores. You do NOT
              need to pay for a membership with us or be registered with your
              National Governing Body.
            </p>
          </div>

          <div className="mb-6">
            <h2 className="text-2xl text-center font-bold mb-2">The rules</h2>
            <p className="text-center">
              Typical weightlifting competition rules apply but we are NOT
              enforcing a strict press-out rule, a lift will only be disallowed
              should you exceed approximately 30degrees of elbow flexion. You do
              not need a singlet, although you can wear one if you think it adds
              some KGs to your total. Weigh-ins are held by the club affiliate
              and it is up to yourself how light you want to be on the scale. As
              always, no straps.
            </p>
          </div>

          <div className="mb-6">
            <h2 className="text-2xl text-center font-bold mb-2">The cost</h2>
            <p className="text-center">
              We are looking to charge around £10GBP per entry which will be
              invoiced to the club affiliate post-submission of scores and prior
              to winners being drawn. The breakdown of the fee is as follows,
              80% of all fees billed will be put towards the overall prize pot
              with the remainder going to further development of the
              OpenWeightlifting League platform and future events.
            </p>
          </div>

          <div className="mb-6">
            <h1 className="text-3xl text-center font-bold mb-2">
              League Affiliate Registration
            </h1>
            <p className="text-center">
              Please fill out the form below to register your interest as an
              affiliate.
            </p>
            <p className="italic text-center">
              Reminder, you do not need to hold an NGB affiliation.
            </p>
          </div>

          {!submitted && (
            <form
              className="flex flex-col space-y-4 mt-4"
              onSubmit={handleSubmit}
            >
              <Input
                aria-label="Username"
                type="text"
                placeholder="Username"
                onChange={e => setUsername(e.target.value)}
                required
              />
              <Input
                aria-label="Club Name"
                type="text"
                placeholder="Club Name"
                onChange={e => setClubname(e.target.value)}
                required
              />
              <Input
                aria-label="Location"
                type="text"
                placeholder="Location"
                onChange={e => setLocation(e.target.value)}
                required
              />
              <Input
                aria-label="Email"
                type="email"
                placeholder="Email"
                onChange={e => setEmail(e.target.value)}
                required
              />
              <Input
                aria-label="Comments"
                type="number"
                placeholder="How many lifters are you looking to enter?"
                onChange={e => setComments(e.target.value)}
              />

              <Popover>
                <PopoverTrigger>
                  <Button
                    className="flex justify-center"
                    type="submit"
                    aria-label="Submit"
                    color="primary"
                    disabled={disableButton}
                  >
                    Submit
                  </Button>
                </PopoverTrigger>
                {error && (
                  <PopoverContent>
                    <p className="text-center">{errorMessage}</p>
                  </PopoverContent>
                )}
                {loading && (
                  <PopoverContent>
                    <p className="text-center">Loading...</p>
                  </PopoverContent>
                )}
              </Popover>
            </form>
          )}

          {submitted && (
            <p className="text-center border-primary border-3 rounded-xl p-4">
              Thank you for registering your interest. We will be in touch
              shortly with more information.
            </p>
          )}
        </div>
      </div>
    </>
  )
}

export default League
