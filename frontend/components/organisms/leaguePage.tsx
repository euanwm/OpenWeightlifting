import useSWR from 'swr'
import { useState } from 'react'
import {
  Button,
  Input,
  Popover,
  PopoverContent,
  PopoverTrigger,
  Spinner,
} from '@nextui-org/react'

import HeaderBar from '@/components/molecules/head'
import submitLeagueForm from '@/api/submitLeagueForm/submitLeagueForm'
import { RegistrationForm } from '@/api/submitLeagueForm/submitLeagueFormTypes'

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

function LeaguePage() {
  const [username, setUsername] = useState('')
  const [clubname, setClubname] = useState('')
  const [location, setLocation] = useState('')
  const [email, setEmail] = useState('')
  const [comments, setComments] = useState('')
  const [validationErrorMessage, setValidationErrorMessage] = useState('')
  const [formSubmissionData, setFormSubmissionData] =
    useState<null | RegistrationForm>(null)

  const { data, isLoading, error } = useSWR(formSubmissionData, submitLeagueForm, {
    keepPreviousData: true,
  })

  const handleSubmit = async (e: any) => {
    e.preventDefault()

    if (!validateForm({ username, clubname, location, email, comments })) {
      setValidationErrorMessage('Please fill out all fields correctly.')
      return
    }

    const data = {
      username,
      clubname,
      location,
      email,
      comments,
    }

    setFormSubmissionData({
      username,
      clubname,
      location,
      email,
      comments,
    })
  }

  return (
    <>
      {isLoading && (
        <div className="fixed w-screen h-screen z-10 flex justify-center items-center">
          <Spinner size="lg" label="Loading..." />
        </div>
      )}

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

          {!data && (
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
                    disabled={isLoading}
                  >
                    Submit
                  </Button>
                </PopoverTrigger>
                {!data && validationErrorMessage && (
                  <PopoverContent>
                    <p className="text-center">{validationErrorMessage}</p>
                  </PopoverContent>
                )}
                {error && (
                  <PopoverContent>
                    <p className="text-center">Sorry there was a problem submitting the form</p>
                  </PopoverContent>
                )}
              </Popover>
            </form>
          )}

          {data && (
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

export default LeaguePage
