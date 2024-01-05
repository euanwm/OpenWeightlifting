import { EventsList, EventsListRequest } from "@/api/fetchEventsList/fetchEventsListTypes";

export default  async function fetchEventsList(
  eventsListRequest: EventsListRequest | null,
): Promise<EventsList> {
  const response = await fetch(`${process.env.API}/events/list`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(eventsListRequest),
  })

  const jsonResponse = response.json()
  return jsonResponse as Promise<EventsList>
}