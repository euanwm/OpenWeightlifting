import "@/styles/globals.css";
import { NextUIProvider } from "@nextui-org/react";
import { ThemeProvider as NextThemesProvider } from "next-themes";
import { fontSans, fontMono } from "@/config/fonts";
import type { AppProps } from "next/app";
import { useEffect } from "react";
import { hotjar } from "react-hotjar";

export default function App({ Component, pageProps }: AppProps) {
  useEffect(() => {
    hotjar.initialize(3147762, 6);
  }, [])

  return (
    <NextUIProvider>
      <NextThemesProvider>
        <Component {...pageProps} />
      </NextThemesProvider>
    </NextUIProvider>
  );
}

export const fonts = {
  sans: fontSans.style.fontFamily,
  mono: fontMono.style.fontFamily,
};