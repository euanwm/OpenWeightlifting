export type RegistrationForm = {
  username: string
  clubname: string
  location: string
  email: string
  comments: string
}

export type FormResponse = {
  success: boolean
  message: string
}