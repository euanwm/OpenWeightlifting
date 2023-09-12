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
}
