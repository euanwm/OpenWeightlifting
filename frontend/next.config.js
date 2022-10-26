/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  env: {
    API: process.env.API ?? 'https://api.openweightlifting.org',
  },
}

module.exports = nextConfig
