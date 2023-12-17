import useSWR from 'swr'
import { useState } from 'react'
import { Input, Listbox, ListboxItem, ListboxSection, Spinner } from '@nextui-org/react'

import HeaderBar from '@/components/molecules/head'
import fetchLifterNames from '@/api/fetchLifterNames/fetchLifterNames'

function SearchPage() {
  const [inputSearch, setInputSearch] = useState('')

  const { data, isLoading } = useSWR(
    inputSearch,
    fetchLifterNames,
  )

  return (
    <>
      {isLoading && (
        <div className="fixed w-screen h-screen z-10 flex justify-center items-center">
          <Spinner size="lg" label="Loading..." />
        </div>
      )}
      <HeaderBar />
      <div className="flex justify-center mt-4">
        <div className="max-w-lg mx-4 space-y-4">
          <center>
            <h1>Lifter search function</h1>
            <p>Enter 3 or more letters to start searching for a lifter.</p>
          </center>
          <Input
            placeholder="Search for a lifter..."
            onChange={async e => setInputSearch(e.target.value)}
          />
          {data && Array.isArray(data.names) && (
            <Listbox>
              <ListboxSection title="Search Results">
                {data.names.map(lifterName => (
                  <ListboxItem
                    key={lifterName}
                    onClick={() => {
                      window.location.href = '/lifter?name=' + lifterName
                    }}
                  >
                    {lifterName}
                  </ListboxItem>
                ))}
              </ListboxSection>
            </Listbox>
          )}
        </div>
      </div>
    </>
  )
}

export default SearchPage
