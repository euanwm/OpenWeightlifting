import { Html, Head, Main, NextScript } from 'next/document'
import Script from 'next/script'

export default function Document() {
  return (
    <Html lang="en">
      <Head>
        <script data-name="BMC-Widget" data-cfasync="false" src="https://cdnjs.buymeacoffee.com/1.0.0/widget.prod.min.js" data-id="openweightlifting" data-description="Support us on Buy me a coffee!" data-message="" data-color="#00B0F0" data-position="Right" data-x_margin="18" data-y_margin="18" async />
      </Head>
      <title>OpenWeightlifting</title>
      <body className="min-h-screen bg-background font-sans antialiased">
      <Main />
      <NextScript />
      </body>
    </Html>
  )
}