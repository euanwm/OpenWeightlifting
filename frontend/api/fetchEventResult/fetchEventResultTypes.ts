import { LifterResult } from "@/api/fetchLifterData/fetchLifterDataTypes";

export type EventResult = {
  size: number;
  data: LifterResult[];
}