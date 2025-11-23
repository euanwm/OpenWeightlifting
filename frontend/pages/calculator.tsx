import CalculatorPage from '@/components/organisms/calculatorPage'
import Head from "next/head";

function Calculator() {
  return (
    <>
      <Head>
        <title>Calculator Calculator</title>
        <meta
          name="description"
          content="Olympic-cycle selectable Sinclair and Q-Points calculator."
        />
      </Head>
      <CalculatorPage />
    </>
  )
}

export default Calculator
