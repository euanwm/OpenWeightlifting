import DataTable from "../components/data-table/index.component"

const Home = ({ data }) =>
  <DataTable lifters={data} />

export async function getServerSideProps() {
  const res = await fetch('https://owl-production-backend.herokuapp.com/api/leaderboard')
  const data = await res.json()

  return { props: { data } }
}

export default Home
