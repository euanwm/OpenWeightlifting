export type LeaderboardPayload = {
  start: number;
  stop: number;
  sortby: string;
  federation: string;
  weightclass: string;
  year: number;
}

export type LifterNamePayload = {
  name: string;
}

export type LifterSearchList = {
  names: Array<{
    Federation: string;
    Name: string;
  }>;
  total: number;
}