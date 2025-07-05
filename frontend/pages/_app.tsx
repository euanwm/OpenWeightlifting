import '@/styles/globals.css'
import { NextUIProvider } from '@nextui-org/react'
import { ThemeProvider as NextThemesProvider } from 'next-themes'
import { fontSans, fontMono } from '@/config/fonts'
import type { AppProps } from 'next/app'
import { useEffect } from 'react'
import posthog from 'posthog-js'
import { PostHogProvider } from 'posthog-js/react'

export default function App({ Component, pageProps }: AppProps) {
  useEffect(() => {
    if (process.env.NODE_ENV === 'production') {
      posthog.init('phc_gi3XPh5YpNuzgy5uailSyNKuEjuwny8hu5LjK5t7AGx', {
        api_host: 'https://us.i.posthog.com',
        persistence: 'localStorage',
        person_profiles: 'always', // 'identified_only' or 'always' to create profiles for anonymous users as well
      })
    }

    let bmcButton: HTMLElement | null = null
    let interval: NodeJS.Timeout

    function handleClick() {
      posthog.capture('buy_me_a_coffee_clicked')
    }

    interval = setInterval(() => {
      bmcButton = document.getElementById('bmc-wbtn')
      if (bmcButton) {
        bmcButton.addEventListener('click', handleClick)
        clearInterval(interval)
      }
    }, 500)

    return () => {
      clearInterval(interval)
      if (bmcButton) {
        bmcButton.removeEventListener('click', handleClick)
      }
    }
  }, [])
  return (
    <NextUIProvider>
      <NextThemesProvider attribute="class" defaultTheme="dark">
        <PostHogProvider client={posthog}>
          <Component {...pageProps} />
        </PostHogProvider>
      </NextThemesProvider>
    </NextUIProvider>
  )
}

export const fonts = {
  sans: fontSans.style.fontFamily,
  mono: fontMono.style.fontFamily,
}
