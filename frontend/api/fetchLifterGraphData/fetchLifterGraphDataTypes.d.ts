export type LifterChartData = {
  labels: string[];
  datasets: ChartSubData[];
} | undefined

export type ChartSubData = {
  label: string;
  data: number[];
}