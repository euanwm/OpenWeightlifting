import { NextUIProvider } from '@nextui-org/react'
import { hotjar } from 'react-hotjar'

import Layout from '../components/layout/index.component'
import { useEffect } from "react";

// todo: declare types for this
function MyApp({ Component, pageProps }: any) {
    useEffect(() => {
        hotjar.initialize(3147762, 6);
    }, [])
  return (
      <NextUIProvider>
        <Layout>
          <Component {...pageProps} />
        </Layout>
      </NextUIProvider>
  )
}

export default MyApp
