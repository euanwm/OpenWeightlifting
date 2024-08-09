import SinclairPage from '@/components/organisms/sinclairPage'
import Head from "next/head";

function Sinclair() {
  return (
    <>
      <Head>
        <title>Sinclair Calculator</title>
        <meta
          name="description"
          content="Olympic-cycle selectable Sinclair calculator."
        />
      </Head>
      <SinclairPage />
    </>
  )
}

export default Sinclair
