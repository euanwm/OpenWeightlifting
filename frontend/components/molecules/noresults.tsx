import { Button } from '@nextui-org/react'

// I have no fucking clue why I had to implement this, but it worked
// If you don't like it, raise a PR ya loser
export const NoResults = () => (
  <div className="flex flex-col items-center justify-center space-y-4 mt-4 mx-4">
    <h1 className="text-2xl">No results found</h1>
    <Button
      className="ml-4"
      color="primary"
      onClick={() => window.history.back()}
      >
      Go back
    </Button>
  </div>
)