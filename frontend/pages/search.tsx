import SearchPage from '@/components/organisms/searchPage'
import Head from "next/head";

function Search() {
  return (
    <>
      <Head>
        <title>Lifter Search</title>
        <meta
          name="description"
          content="Search for a lifter in our database."/>
      </Head>
      <SearchPage />
    </>
  )
}

export default Search
