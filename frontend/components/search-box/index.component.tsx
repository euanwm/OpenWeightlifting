// todo: make this actually work
import { useState } from 'react'
import { Combobox } from '@headlessui/react'
import fetchLifterNames from 'api/fetchLifterNames/fetchLifterNames'

// Decided to use this combobox because the nextui doesn't have one
// https://headlessui.com/react/combobox

// This is going to be a component on the main leaderboard page that will allow you to search for a lifter.
// When you select a lifter, it will take you to the lifter page and show you their data.
// The lifter page will be generated based upon the query string.
// This component will live in the layout component somewhere.

async function SearchBox() {
  const [selectedPerson, setSelectedPerson] = useState('')
  const [query, setQuery] = useState('')

  const { names: lifterNames } = await fetchLifterNames(query)

  return (
    <Combobox value={selectedPerson} onChange={setSelectedPerson}>
      <Combobox.Input onChange={event => setQuery(event.target.value)} />
      <Combobox.Options>
        {lifterNames?.map(name => (
          <Combobox.Option key={name} value={name}>
            {name}
          </Combobox.Option>
        ))}
      </Combobox.Options>
    </Combobox>
  )
}

export default SearchBox;
