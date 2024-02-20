import fetchEventsList from './fetchEventsList';
import { EventsListRequest } from './fetchEventsListTypes';

describe('fetchEventsList', () => {
  it('should return a list of events', async () => {
    const eventsListRequest: EventsListRequest = {
      startdate: '2023-01-01',
      enddate: '2023-01-15',
    }
    const result = await fetchEventsList(eventsListRequest);

    // check that it returns a list of events (length 172)
    expect(result?.events.length).toBe(172);
  })
})
