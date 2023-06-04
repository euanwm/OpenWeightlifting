import { useState } from "react";
import { Combobox } from "@headlessui/react";
import { LifterSearchList } from "../../models/api_endpoint";

// Decided to use this combobox because the nextui doesn't have one
// https://headlessui.com/react/combobox

// This is going to be a component on the main leaderboard page that will allow you to search for a lifter.
// When you select a lifter, it will take you to the lifter page and show you their data.
// The lifter page will be generated based upon the query string.
// This component will also be on the lifter page.

const fetchLifterNames = async (nameQuery: string) => {
  if (nameQuery.length < 3) {
    return ({
      names: []
    } as any) as LifterSearchList
  }
  const res = await fetch(`https://api.openweightlifting.org/search?name=${nameQuery}`
  ).then((res) => res.json()).catch(error => console.error('error in searching names', error))

  return await res as LifterSearchList
}

export const SearchBox = async () => {
  const [selectedPerson, setSelectedPerson] = useState('')
  const [query, setQuery] = useState('')

  const lifterNames = await fetchLifterNames(query).then((res) => res?.names).catch(error => console.error('error in fetchLifterNames', error))

  return (
    <Combobox value={selectedPerson} onChange={setSelectedPerson}>
      <Combobox.Input onChange={(event) => setQuery(event.target.value)} />
      <Combobox.Options>
        {lifterNames.map((name) => (
          <Combobox.Option key={name} value={name}>
            {name}
          </Combobox.Option>
        ))}
      </Combobox.Options>
    </Combobox>
  )
}
