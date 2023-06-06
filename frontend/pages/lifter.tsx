import { useRouter } from "next/router";

import { useEffect, useState } from "react";

import { LifterGraph } from "../components/lifter-graph/index.component";
import { HistoryTable } from "../components/history-table/index.components";

import { LifterHistory } from "../models/api_endpoint";
import { Card, Container, Row, Text } from "@nextui-org/react";

const sampleLifterHistory: LifterHistory = {
  "name": "Euan Meston",
  "lifts": [
    {
      "event": "BWL South Open Series 3 - 2019",
      "date": "2019-09-28",
      "gender": "Men's Senior 81Kg",
      "lifter_name": "Euan Meston",
      "bodyweight": 80.9,
      "snatch_1": -82,
      "snatch_2": -82,
      "snatch_3": 82,
      "cj_1": 100,
      "cj_2": -105,
      "cj_3": 105,
      "best_snatch": 82,
      "best_cj": 105,
      "total": 187,
      "sinclair": 227.46225,
      "country": "UK",
      "instagram": ""
    },
    {
      "event": "P10 Christmas Open 2019",
      "date": "2019-12-14",
      "gender": "Men's Senior 81Kg",
      "lifter_name": "Euan Meston",
      "bodyweight": 79.3,
      "snatch_1": 90,
      "snatch_2": -95,
      "snatch_3": -95,
      "cj_1": 110,
      "cj_2": -120,
      "cj_3": -120,
      "best_snatch": 90,
      "best_cj": 110,
      "total": 200,
      "sinclair": 245.77774,
      "country": "UK",
      "instagram": ""
    },
    {
      "event": "Virtual Winter Open 2020",
      "date": "2020-12-31",
      "gender": "Men's Senior 81Kg",
      "lifter_name": "Euan Meston",
      "bodyweight": 79.8,
      "snatch_1": 95,
      "snatch_2": 100,
      "snatch_3": 105,
      "cj_1": 105,
      "cj_2": 115,
      "cj_3": 120,
      "best_snatch": 105,
      "best_cj": 120,
      "total": 225,
      "sinclair": 275.60324,
      "country": "UK",
      "instagram": ""
    },
    {
      "event": "SCOTTISH OPEN 2 2021",
      "date": "2021-04-30",
      "gender": "Men's Senior 81Kg",
      "lifter_name": "Euan Meston",
      "bodyweight": 80.4,
      "snatch_1": 95,
      "snatch_2": 101,
      "snatch_3": 106,
      "cj_1": 112,
      "cj_2": 119,
      "cj_3": 121,
      "best_snatch": 106,
      "best_cj": 121,
      "total": 227,
      "sinclair": 276.988,
      "country": "UK",
      "instagram": ""
    },
    {
      "event": "2021 British Weightlifting Open",
      "date": "2021-11-13",
      "gender": "Men's Senior 81Kg",
      "lifter_name": "Euan Meston",
      "bodyweight": 78.8,
      "snatch_1": 105,
      "snatch_2": 111,
      "snatch_3": -114,
      "cj_1": 125,
      "cj_2": -130,
      "cj_3": -135,
      "best_snatch": 111,
      "best_cj": 125,
      "total": 236,
      "sinclair": 290.97495,
      "country": "UK",
      "instagram": ""
    },
    {
      "event": "Crystal Palace Spring Open 22",
      "date": "2022-03-19",
      "gender": "Men's Senior 81Kg",
      "lifter_name": "Euan Meston",
      "bodyweight": 81,
      "snatch_1": -103,
      "snatch_2": -107,
      "snatch_3": 107,
      "cj_1": 126,
      "cj_2": 133,
      "cj_3": -140,
      "best_snatch": 107,
      "best_cj": 133,
      "total": 240,
      "sinclair": 291.74796,
      "country": "UK",
      "instagram": ""
    },
    {
      "event": "England Senior and Masters Championships 2022",
      "date": "2022-10-29",
      "gender": "Men's Senior 81Kg",
      "lifter_name": "Euan Meston",
      "bodyweight": 81,
      "snatch_1": 100,
      "snatch_2": 110,
      "snatch_3": -120,
      "cj_1": 120,
      "cj_2": 125,
      "cj_3": -135,
      "best_snatch": 110,
      "best_cj": 125,
      "total": 235,
      "sinclair": 285.66986,
      "country": "UK",
      "instagram": ""
    }
  ],
  "graph": {
    "labels": [
      "2019-09-28",
      "2019-12-14",
      "2020-12-31",
      "2021-04-30",
      "2021-11-13",
      "2022-03-19",
      "2022-10-29"
    ],
    "datasets": [
      {
        "label": "Competition Total",
        "data": [
          187,
          200,
          225,
          227,
          236,
          240,
          235
        ]
      },
      {
        "label": "Best Snatch",
        "data": [
          82,
          90,
          105,
          106,
          111,
          107,
          110
        ]
      },
      {
        "label": "Best C&J",
        "data": [
          105,
          110,
          120,
          121,
          125,
          133,
          125
        ]
      },
      {
        "label": "Bodyweight",
        "data": [
          80.9,
          79.3,
          79.8,
          80.4,
          78.8,
          81,
          81
        ]
      }
    ]
  }
}

const fetchLifterHistory = async (name: string) => {
  const response = await fetch('http://localhost:8080/history', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({ NameStr: name })
  }).then((res) => res.json()).catch(error => console.error('error in fetchLifterHistory', error))

  return await response as LifterHistory
}


const Lifter = () => {
  const router = useRouter()
  const { name } = router.query

  const [lifterHistory, setLifterHistory] = useState({} as LifterHistory)

  useEffect(() => {
    async function fetchLifterHistoryFromAPI() {
      setLifterHistory(await fetchLifterHistory(name as string))
    }

  fetchLifterHistory(name as string).then()
  }, [name])

  return (
    <div>
      <Container alignContent={'center'} alignItems={'center'}>
        <Card>
          <Card.Body>
            <Row justify="center" align="center">
              <Text h1>{sampleLifterHistory['name']}</Text>
            </Row>
          </Card.Body>
        </Card>
      </Container>
      <LifterGraph lifterHistory={sampleLifterHistory['graph']} />
      <HistoryTable history={sampleLifterHistory['lifts']} />
    </div>
  )
}

export default Lifter