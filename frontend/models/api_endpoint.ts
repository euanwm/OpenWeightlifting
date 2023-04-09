// Description: This file contains the interface for all API endpoints
// These should mirror the backend API endpoints as closely as possible
// Reference: /backend/structs/structs.go

export interface LifterResult {
  event: string;
  date: string;
  gender: string;
  name: string;
  bodyweight: number;
  snatch_1: number;
  snatch_2: number;
  snatch_3: number;
  clean_jerk_1: number;
  clean_jerk_2: number;
  clean_jerk_3: number;
  best_snatch: number;
  best_clean_jerk: number;
  total: number;
  sinclair: number;
  federation: string;
  instagram: string;
}

export interface MainLeaderboard {
  lifts: LifterResult[];
}

export interface LifterChartData {
  labels: string[];
  datasets: ChartSubData[];
}

export interface ChartSubData {
  label: string;
  data: number[];
}

export interface LifterHistory {
  name: string;
  lifts: LifterResult[];
}

export interface LeaderboardPayload {
  start: number;
  stop: number;
  sortby: string;
  federation: string;
  weightclass: string;
  year: number;
}

export interface LifterNamePayload {
  name: string;
}