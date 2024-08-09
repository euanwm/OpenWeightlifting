/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  env: {
    API: process.env.API ?? 'https://api.openweightlifting.org',
  },
  async headers() {
    return [
      {
        source: '/:all*',
        headers: [
          {
            key: 'Cache-Control',
            value: 'public, max-age=2628000',
          }
        ],
      },
    ]
  }
}

module.exports = nextConfig
