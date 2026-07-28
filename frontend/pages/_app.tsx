import '@/styles/globals.css'
import { NextUIProvider } from '@nextui-org/react'
import { ThemeProvider as NextThemesProvider } from 'next-themes'
import { fontSans, fontMono } from '@/config/fonts'
import type { AppProps } from 'next/app'
import { useEffect } from 'react'
import posthog from 'posthog-js'
import type { CaptureResult } from 'posthog-js'
import { PostHogProvider } from 'posthog-js/react'

// Messages that come from browser extensions / injected wallet-provider scripts
// rather than our own code. These get swept up by posthog-js exception
// autocapture and add noise (and quota cost) to error tracking.
const EXTENSION_EXCEPTION_PATTERNS = [
  /ParityDigital/i,
  /web3/i,
  /ethereum/i,
  /metamask/i,
  /solana/i,
  /evmAsk/i,
]

// Drop $exception events that don't originate from our own bundle: either
// frameless synthetic exceptions (no stack trace to attribute to our code) or
// ones whose message matches a known extension pattern.
function dropExtensionExceptions(event: CaptureResult | null): CaptureResult | null {
  if (!event || event.event !== '$exception') {
    return event
  }

  const exceptionList = event.properties?.$exception_list
  if (!Array.isArray(exceptionList) || exceptionList.length === 0) {
    // No structured exception info to attribute — treat as non-actionable noise.
    return null
  }

  const isExtensionNoise = exceptionList.every((exception: any) => {
    const message = `${exception?.type ?? ''} ${exception?.value ?? ''}`
    if (EXTENSION_EXCEPTION_PATTERNS.some((pattern) => pattern.test(message))) {
      return true
    }
    // Synthetic exceptions injected by extensions typically arrive with no
    // stack frames, so there's nothing tying them to our code.
    const frames = exception?.stacktrace?.frames
    return !Array.isArray(frames) || frames.length === 0
  })

  return isExtensionNoise ? null : event
}

export default function App({ Component, pageProps }: AppProps) {
  useEffect(() => {
    if (process.env.NODE_ENV === 'production') {
      posthog.init('phc_gi3XPh5YpNuzgy5uailSyNKuEjuwny8hu5LjK5t7AGx', {
        api_host: 'https://us.i.posthog.com',
        persistence: 'localStorage',
        person_profiles: 'always', // 'identified_only' or 'always' to create profiles for anonymous users as well
        before_send: dropExtensionExceptions,
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
