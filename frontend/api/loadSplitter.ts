import posthog from 'posthog-js'

export function getAPI() {
  let api = process.env.API
  if (posthog.isFeatureEnabled('railway_backend')) {
    api = process.env.ALTERNATE_API
  }
  return api
}
