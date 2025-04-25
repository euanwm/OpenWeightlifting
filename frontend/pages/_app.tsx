import '@/styles/globals.css'
import { NextUIProvider } from '@nextui-org/react'
import { ThemeProvider as NextThemesProvider } from 'next-themes'
import { fontSans, fontMono } from '@/config/fonts'
import type { AppProps } from 'next/app'
import { useEffect } from 'react'
import posthog from 'posthog-js'

export default function App({ Component, pageProps }: AppProps) {
  useEffect(() => {
    //disables analytics in a local environment
    if (process.env.NODE_ENV === 'production') {
      posthog.init('phc_gi3XPh5YpNuzgy5uailSyNKuEjuwny8hu5LjK5t7AGx', {
        api_host: 'https://us.i.posthog.com',
        person_profiles: 'always', // 'identified_only' or 'always' to create profiles for anonymous users as well
      })
    }
  }, [])
  return (
    <NextUIProvider>
      <NextThemesProvider attribute="class" defaultTheme="dark">
        <Component {...pageProps} />
      </NextThemesProvider>
    </NextUIProvider>
  )
}

export const fonts = {
  sans: fontSans.style.fontFamily,
  mono: fontMono.style.fontFamily,
}
