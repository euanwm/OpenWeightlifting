import { EventsList, EventsListRequest } from "@/api/fetchEventsList/fetchEventsListTypes";

export default  async function fetchEventsList(
  eventsListRequest: EventsListRequest,
): Promise<EventsList> {
  if (!eventsListRequest) {
    return undefined
  }

  const response = await fetch(`${process.env.API}/events`, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(eventsListRequest),
  })

  const jsonResponse = response.json()
  return jsonResponse as Promise<EventsList>
}