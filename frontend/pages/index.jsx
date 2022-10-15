import { useState } from "react"
import DataTable from "../components/data-table/index.component"
import Filters from "../components/filters/index.component"

const fetchLifterData = async (gender, start, stop, sortby, federation) => {
  const bodyContent = JSON.stringify({
    "gender": gender || "male",
    "start": start || 0,
    "stop": stop || 500,
    "sortby": sortby || "total",
    "federation": federation || "allfeds"
  })

  const res = await fetch('https://api.openweightlifting.org/leaderboard', {
    method: "POST",
    headers: {
      "Accept": "*/*",
      "Content-Type": "application/json"
    },
    body: bodyContent
  }).catch((error) => console.error(error))

  return await res.json()
}

const Home = ({ data }) => {
  const [currentGender, setCurrentGender] = useState("male")
  const [sortBy, setSortBy] = useState("total")
  const [federation, setFederation] = useState("allfeds")
  const [currentLifterList, setCurrentLifterList] = useState(data)

  //todo: refactor this bit below
  const handleGenderChange = async (event) => {
    const newFilter = event.target.value
    let newLifters;

    if (newFilter === "female" || newFilter === "male") {
        setCurrentGender(newFilter)
        newLifters = await fetchLifterData(newFilter, 0, 500, sortBy, federation);
        setCurrentLifterList(newLifters)
    } else if (newFilter === "total" || newFilter === "sinclair") {
        setSortBy(newFilter)
        newLifters = await fetchLifterData(currentGender, 0, 500, newFilter, federation);
        setCurrentLifterList(newLifters)
    } else if (newFilter.match("UK|US|AUS|allfeds|IWF")) {
        setFederation(newFilter)
        newLifters = await fetchLifterData(currentGender, 0, 500, sortBy, newFilter);
        setCurrentLifterList(newLifters)
    }
  }

  return (
    <>
      <Filters currentGender={currentGender} sortBy={sortBy} federation={federation} handleGenderChange={handleGenderChange} />
      {currentLifterList && <DataTable lifters={currentLifterList} />}
    </>
  )
}


export async function getServerSideProps() {
  const data = await fetchLifterData()

  return { props: { data } }
}

export default Home
