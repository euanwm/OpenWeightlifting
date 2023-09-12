import { Html, Head, Main, NextScript } from 'next/document'
import HeaderBar from "@/layouts/head";

export default function Document() {
  return (
    <Html lang="en">
      <Head />
      <title>OpenWeightlifting</title>
        <body className="min-h-screen bg-background font-sans antialiased">
      <Main />
      <NextScript />
      </body>
    </Html>
  )
}