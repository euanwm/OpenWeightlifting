import useSWR from 'swr'
import { useState } from 'react'
import { Input, Listbox, ListboxItem, ListboxSection, Spinner } from '@nextui-org/react'

import HeaderBar from '@/components/molecules/head'
import fetchLifterNames from '@/api/fetchLifterNames/fetchLifterNames'

function SearchPage() {
  const [inputSearch, setInputSearch] = useState('')
  const params = { name: inputSearch, limit: '50' }

  const { data, isLoading } = useSWR(
    params,
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
          Number of lifters found: {data?.total}, showing {data?.names.length}
          <Listbox aria-label='Lifter search results'>
            <ListboxSection>
              {data?.names.map(({Name, Federation }) => (
                <ListboxItem key={Name} onClick={() => {window.location.href = `/lifter?name=${Name}&federation=${Federation}`}}>

                  {Name}, {Federation}
                </ListboxItem>
              ))}
            </ListboxSection>
          </Listbox>
        </div>
      </div>
    </>
  )
}

export default SearchPage
