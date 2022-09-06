import React from 'react';
import Document, { Html, Head, Main, NextScript } from 'next/document';
import { CssBaseline } from '@nextui-org/react';
import Script from 'next/script'

class MyDocument extends Document {
  static async getInitialProps(ctx) {
    const initialProps = await Document.getInitialProps(ctx);
    return {
      ...initialProps,
      styles: React.Children.toArray([initialProps.styles])
    };
  }
// G-NJZS5B4Y3R
  render() {
    return (
      <Html lang="en">
        <Head>{CssBaseline.flush()}
          <div className="container">
            <Script
              src="https://www.googletagmanager.com/gtag/js?id=G-NJZS5B4Y3R"
              strategy="afterInteractive"
            />
            <Script id="google-analytics" strategy="afterInteractive">
              {`
                window.dataLayer = window.dataLayer || [];
                function gtag(){window.dataLayer.push(arguments);}
                gtag('js', new Date());
      
                gtag('config', 'G-NJZS5B4Y3R');
              `}
            </Script>
          </div>
        </Head>
        <body>
          <Main />
          <NextScript />
        </body>
      </Html>
    );
  }
}

export default MyDocument;