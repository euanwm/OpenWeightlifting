export type LifterResult = {
    event: string;
    date: string;
    gender: string;
    lifter_name: string;
    bodyweight: number;
    snatch_1: number;
    snatch_2: number;
    snatch_3: number;
    cj_1: number;
    cj_2: number;
    cj_3: number;
    best_snatch: number;
    best_cj: number;
    total: number;
    sinclair: number;
    country: string;
    instagram: string;
  };

export type LeaderboardResult = {
    size: number;
    data: LifterResult[];
}