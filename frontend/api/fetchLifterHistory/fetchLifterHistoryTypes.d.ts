export type MainLeaderboard = {
  lifts: LifterResult[]
}

export type LifterChartData = {
  labels: string[]
  datasets: ChartSubData[]
}

export type ChartSubData = {
  label: string
  data: number[]
}

export type LifterHistory = {
  name: string
  lifts: LifterResult[]
  graph: LifterChartData
  stats: LifterStats
}

export type LifterStatistics = {
  best_snatch: number
  best_cj: number
  best_total: number
  make_rate_snatches: number[]
  make_rate_cj: number[]
}