import { EventResult } from './fetchEventResultTypes'
import { EventMetaData } from "@/api/fetchEventsList/fetchEventsListTypes";

export default  async function fetchEventResult(
  eventMetaData: EventMetaData,
): Promise<EventResult> {
  if (!eventMetaData) {
    return undefined
  }


  const response = await fetch(`${process.env.API}/events`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      "federation": eventMetaData.federation,
      "name": eventMetaData.name,
    }),
  })

  const jsonResponse = response.json()
  return jsonResponse as Promise<EventResult>
}