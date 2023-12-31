export type EventsList = {
  events: EventMetaData[];
}

export type EventMetaData = {
  name: string;
  federation: string;
  date: string;
  id: string;
}

export type EventsListRequest = {
  startdate: string;
  enddate: string;
}