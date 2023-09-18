import { useState } from 'react'
import { LifterSearchList } from "./fetchLifterNamesTypes"

import HeaderBar from '@/layouts/head'
import fetchLifterNames from "@/api/fetchLifterNames/fetchLifterNames";

import { Input, Listbox, ListboxItem, ListboxSection } from '@nextui-org/react'

function SearchPage() {
  const [searchResults, setLifterNames] = useState<LifterSearchList>(
    {
      names: []
    }
  )

  return (
    <>
      <HeaderBar />
      <div className="flex justify-center mt-4">
        <div className="max-w-lg mx-4 space-y-4">
      <center>
        <h1>Lifter search function</h1>
        <p>
          Enter 3 or more letters to start searching for a lifter.
        </p>
      </center>
      <Input
        placeholder="Search for a lifter..."
        onChange={async (e) => {
          const lifterNames = await fetchLifterNames(e.target.value)
          setLifterNames(lifterNames)
          }
        }
      />
      <Listbox>
        <ListboxSection title="Search Results">
          {searchResults.names.map((lifterName) => (
            <ListboxItem key={lifterName}
              onClick={() => {
                window.location.href = "/lifter?name=" + lifterName
              }
            }
            >{lifterName}</ListboxItem>
          ))}
        </ListboxSection>
      </Listbox>
      </div>
      </div>
    </>
  )
}

export default SearchPage;