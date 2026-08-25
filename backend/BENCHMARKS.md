# Gin Endpoint Benchmarks

`bench_test.go` benchmarks every gin endpoint against the real router from
`buildServer()`, using real data from `event_data/`. `LeaderboardColdCache`/
`LeaderboardWarmCache`/`LeaderboardSearchWarmCache` isolate the query cache's
effect specifically (cache cleared before every call vs. warmed once).

## Running

```bash
go test -run '^$' -bench . -benchmem -count=10 . > bench.txt
```

To compare two branches: `go install golang.org/x/perf/cmd/benchstat@latest`,
run the above on each branch into separate files, then `benchstat a.txt b.txt`.

## `development` vs `feat/split_events_and_lifts`, before → after

`n=10` per side.

| Benchmark | time before | time after | Δ time | mem before | mem after | Δ mem |
|---|---|---|---|---|---|---|
| ServerTime | 1.75µs | 1.86µs | +6% | 6.30KB | 6.27KB | -0.5% |
| SearchName | 38.0ms | 6.93ms | **-82%** | 9.75MB | 1.96MB | -80% |
| LifterGraph | 1.51ms | 2.44ms | +62% | 29.5KB | 12.3KB | -58% |
| LifterHistory | 1.56ms | 2.63ms | +69% | 59.8KB | 34.1KB | -43% |
| Rival | 36.3ms | 92.9µs | **-99.7%** | 7.4KB | 613KB | +8237%* |
| EventsList | 606µs | 715µs | +18% | 7.83KB | 7.80KB | -0.4% |
| SingleEvent | 20.6µs | 12.7µs | -38% | 37.0KB | 17.5KB | -53% |
| LeaderboardColdCache | 5.64s | 5.61s | -0.6% | 45.6MB | 10.4MB | -77% |
| LeaderboardWarmCache | 4.29ms | 343µs | **-92%** | 37.9MB | 2.69MB | -93% |
| LeaderboardSearchWarmCache | 38.3ms | 7.23ms | **-81%** | 9.75MB | 1.97MB | -80% |
| **geomean** | 3.03ms | 997µs | **-67%** | 290KB | 175KB | -40% |

\* `Rival` regressed hard mid-branch (up to 3.17MB/op) from an endpoint that
built a cache query and then ignored it, doing a full unfiltered scan twice
per request. Fixed by routing it through the same cache-backed
`dbtools.FilterLifts` path `/leaderboard` uses, and dropping the redundant
re-dedup in `lifter.Rivals`. Net result is still above `development`'s 7.4KB
because `development` never builds the `Lifter`/`Event` pointer graph in the
first place — not a like-for-like baseline — but it's a **-99.7% time** and
**-80%** improvement from `feat`'s own worst point.

`SingleEvent`'s regression was a masked correctness bug, not a perf issue: a
`Lifter.Lifts` back-reference caused a JSON-encode cycle, so `/events` was
silently returning `200 OK` with an empty body. Fixed with `json:"-"`.
