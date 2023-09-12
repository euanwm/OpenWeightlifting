import {nextui} from '@nextui-org/react'

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './pages/**/*.{js,ts,jsx,tsx,mdx}',
    './components/**/*.{js,ts,jsx,tsx,mdx}',
    './app/**/*.{js,ts,jsx,tsx,mdx}',
    './node_modules/@nextui-org/theme/dist/**/*.{js,ts,jsx,tsx}'
  ],
  darkMode: "class",
  plugins: [nextui(
    {
      prefix: 'nextui',
      defaultTheme: 'dark',
      defaultExtendTheme: 'dark',
      layout: {
        spacingUnit: 4, // in px
        disabledOpacity: ".5", // this value is applied as opacity-[value] when the component is disabled
        dividerWeight: "1px", // h-divider the default height applied to the divider component
        fontSize: {
          tiny: "0.75rem", // text-tiny
          small: "0.875rem", // text-small
          medium: "1rem", // text-medium
          large: "1.125rem", // text-large
        },
        lineHeight: {
          tiny: "1rem", // text-tiny
          small: "1.25rem", // text-small
          medium: "1.5rem", // text-medium
          large: "1.75rem", // text-large
        },
        radius: {
          small: "8px", // rounded-small
          medium: "12px", // rounded-medium
          large: "14px", // rounded-large
        },
        borderWidth: {
          small: "1px", // border-small
          medium: "2px", // border-medium (default)
          large: "3px", // border-large
        },
      },
      themes: {
        dark: {
          colors: {
            background: '#000000',
            foreground: '#ffffff', // font color
            content1: '#000000', // main table background
            content2: '#9d3d3d',
            content3: '#598138',
            content4: '#ffce00',
            default: {
              DEFAULT: '#484848',
              50: '#0a4a6b',
              100: '#252c2d', // dropdowns
              200: '#364041',
              300: '#ffce00',
              400: '#ffce00',
              500: '#ffce00',
              600: '#ffce00',
              700: '#ffce00',
              800: '#ffce00',
              900: '#ffce00'
            },
            primary: {
              DEFAULT: '#00B0F0',
              foreground: '#000000',
            }
          },
          }
        }
      }
  )],
}