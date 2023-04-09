/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  typescript: {
    tsconfigPath: 'tsconfig.json',
  },
  env: {
    API: process.env.API ?? 'https://api.openweightlifting.org',
  },
}

module.exports = nextConfig
