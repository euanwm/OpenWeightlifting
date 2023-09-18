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
    <div>
      <HeaderBar />
      <center>
        <h1>Search</h1>
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
  );
}

export default SearchPage;