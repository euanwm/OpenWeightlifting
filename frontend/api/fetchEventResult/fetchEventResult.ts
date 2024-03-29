import { EventResult } from './fetchEventResultTypes'
import { EventMetaData } from "@/api/fetchEventsList/fetchEventsListTypes";

export default  async function fetchEventResult(
  eventMetaData: EventMetaData,
): Promise<EventResult> {

  const response = await fetch(`${process.env.API}/events`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      "federation": eventMetaData.federation,
      "id": eventMetaData.id,
    }),
  })

  const jsonResponse = response.json()
  return jsonResponse
}