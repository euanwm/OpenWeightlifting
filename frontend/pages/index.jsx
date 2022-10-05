import { useState } from "react"
import DataTable from "../components/data-table/index.component"
import Filters from "../components/filters/index.component"

const fetchLifterData = async (gender, start, stop, sortby) => {
  const bodyContent = JSON.stringify({
    "gender": gender || "male",
    "start": start || 0,
    "stop": stop || 500,
    "sortby": sortby || "total"
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
  const [currentLifterList, setCurrentLifterList] = useState(data)

  const handleGenderChange = async (event) => {
    const newFilter = event.target.value
    let newLifters;

    if (newFilter === "female" || newFilter === "male") {
        setCurrentGender(newFilter)
        newLifters = await fetchLifterData(newFilter, 0, 500, sortBy);
        setCurrentLifterList(newLifters)
    } else if (newFilter === "total" || newFilter === "sinclair") {
        setSortBy(newFilter)
        newLifters = await fetchLifterData(currentGender, 0, 500, newFilter);
        setCurrentLifterList(newLifters)
    }
  }

  return (
    <>
      <Filters currentGender={currentGender} sortBy={sortBy} handleGenderChange={handleGenderChange} />
      {currentLifterList && <DataTable lifters={currentLifterList} />}
    </>
  )
}


export async function getServerSideProps() {
  const data = await fetchLifterData()

  return { props: { data } }
}

export default Home
