import { EventsList, EventsListRequest } from "@/api/fetchEventsList/fetchEventsListTypes";

export default  async function fetchEventsList(
  eventsListRequest: EventsListRequest,
): Promise<EventsList> {
  const response = await fetch(`${process.env.API}/events`, {
    method: 'OPTIONS',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(eventsListRequest),
  })

  const jsonResponse = response.json()
  return jsonResponse as Promise<EventsList>
}