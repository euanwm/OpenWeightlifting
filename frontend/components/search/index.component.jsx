import { useState, useEffect } from 'react'

async function searchLifters(lifterName = '') {
  const urlQuery = `http://localhost:8080/search?name=` + lifterName

  const res = await fetch(urlQuery, {
    method: 'GET',
    headers: {
      Accept: '*/*',
      'Content-Type': 'application/json',
    },
  }).catch(error => console.error(error))

  return await res.json()
}

export default function SearchFilter() {
  const [searchInput, setSearchInput] = useState('')
  const [filteredData, setFilterdData] = useState({ names: [] })

  useEffect(() => {
    async function fetchLifters() {
      if (searchInput.length > 2) {
        const response = await searchLifters(searchInput)
        setFilterdData(response)
      } else if (searchInput.length === 0) {
        setFilterdData({ names: [] })
      }
    }

    fetchLifters()
  }, [searchInput])

  return (
    <div>
      <input
        type="text"
        className="form-control"
        placeholder="Search"
        value={searchInput}
        onChange={e => setSearchInput(e.target.value)}
      />
      {Boolean(filteredData?.names?.length) && (
        <>
          <h5 className="text-primary">
            {filteredData?.names?.length} results found.
          </h5>
          <div>Names:</div>
          <div>
            {filteredData.names.map(item => (
              <p>{item}</p>
            ))}
          </div>
        </>
      )}
    </div>
  )
}
