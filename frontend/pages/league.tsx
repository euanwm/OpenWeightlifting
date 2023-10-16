import {
  Input,
  Button,
  Popover,
  PopoverTrigger,
  PopoverContent,
} from '@nextui-org/react'
import HeaderBar from "@/layouts/head";
import { useState } from 'react'
import { cookies } from "next/headers";

type RegistrationForm = {
  username: string
  clubname: string
  location: string
  email: string
  comments: string
}

// https://owl-mongo-86f8b66fdf19.herokuapp.com/register
async function submitForm(data: RegistrationForm) {
  try {
    const response = await fetch('https://v2.openweightlifting.org/register', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data),
    })
    const json = response.json()
    return json
  } catch (error) {
    console.error('Error:', error)
    return false
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
  return checkEmail(email);
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

  const handleSubmit = async (e: any) => {
    e.preventDefault()
    setLoading(true)
    const data = {
      username,
      clubname,
      location,
      email,
      comments,
    }

    const response = await submitForm(data)
    if (response) {
      setSubmitted(true)
    } else {
      setError(true)
      setErrorMessage('There was an error submitting your form. Please try again.')
    }
    setLoading(false)
  }

  return (
    <>
      <HeaderBar />
      <div className="flex justify-center mt-4">
        <div className="flex flex-col w-1/2">
          <h1 className="text-3xl font-bold text-center">League Registration</h1>
          <p className="text-center">Please fill out the form below to register your club for the league.</p>
          <form className="flex flex-col space-y-4" onSubmit={handleSubmit}>
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
              type="text"
              placeholder="Comments"
              onChange={e => setComments(e.target.value)}
            />
            <Button
              className="flex justify-center"
              type="submit"
              aria-label="Submit"
            >
              Submit
            </Button>
          </form>
          {submitted && (
            <p className="text-center">Thank you for registering your club! We will be in touch soon.</p>
          )}
          {error && (
            <p className="text-center">{errorMessage}</p>
          )}
        </div>
      </div>
    </>
  )
}

export default League