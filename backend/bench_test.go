package main //nolint:typecheck

import (
	"backend/dbtools"
	"backend/enum"
	"backend/structs"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
)

// Benchmarks against the real gin server, built from the real event_data CSVs
// (same as production). Run `go test -bench=. -benchmem ./...` on this branch
// and on development, then `benchstat` the two outputs to see the net delta.

var (
	benchOnce   sync.Once
	benchEngine *gin.Engine
	benchLifter structs.NameSearch
	benchSex    string
	benchFed    string
	benchEvent  string
)

func setupBenchmark(tb testing.TB) *gin.Engine {
	benchOnce.Do(func() {
		gin.SetMode(gin.TestMode)
		gin.DefaultWriter = io.Discard
		log.SetOutput(io.Discard)
		benchEngine = buildServer()

		if len(LeaderboardData.AllTotals) == 0 {
			tb.Fatal("no lifts loaded, cannot benchmark")
		}
		sample := LeaderboardData.AllTotals[0]
		benchLifter = structs.NameSearch{NameStr: sample.Lifter.Name, Federation: sample.Event.Federation}
		benchSex = sample.Lifter.Gender

		if len(EventsData.Events) > 0 {
			benchFed = EventsData.Events[0].Federation
			benchEvent = EventsData.Events[0].CSVID
		}
	})
	return benchEngine
}

func doRequest(engine *gin.Engine, method, target string, body []byte) *httptest.ResponseRecorder {
	var reader io.Reader
	if body != nil {
		reader = bytes.NewReader(body)
	}
	req := httptest.NewRequest(method, target, reader)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	rec := httptest.NewRecorder()
	engine.ServeHTTP(rec, req)
	return rec
}

func resetQueryCache() {
	QueryCache = dbtools.QueryCache{}
}

func defaultLeaderboardQuery() string {
	return "/leaderboard?federation=" + enum.ALLFEDS + "&weightclass=MALL&sortBy=" + enum.Total + "&start=0&stop=50"
}

func withQuery(path string, params map[string]string) string {
	q := url.Values{}
	for k, v := range params {
		q.Set(k, v)
	}
	return path + "?" + q.Encode()
}

func BenchmarkServerTime(b *testing.B) {
	engine := setupBenchmark(b)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, "/time", nil)
	}
}

func BenchmarkSearchName(b *testing.B) {
	engine := setupBenchmark(b)
	target := withQuery("/search", map[string]string{"name": benchLifter.NameStr, "limit": "50"})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

func BenchmarkSimilarNameSearch(b *testing.B) {
	engine := setupBenchmark(b)
	target := withQuery("/search/similarity", map[string]string{"name": benchLifter.NameStr, "federation": benchLifter.Federation})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

func BenchmarkLifterGraph(b *testing.B) {
	engine := setupBenchmark(b)
	target := withQuery("/graph", map[string]string{"name": benchLifter.NameStr, "federation": benchLifter.Federation})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

func BenchmarkLifterHistory(b *testing.B) {
	engine := setupBenchmark(b)
	target := withQuery("/history", map[string]string{"name": benchLifter.NameStr, "federation": benchLifter.Federation})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

func BenchmarkRival(b *testing.B) {
	engine := setupBenchmark(b)
	target := withQuery("/rivals", map[string]string{"name": benchLifter.NameStr, "sex": benchSex})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

func BenchmarkEventsList(b *testing.B) {
	engine := setupBenchmark(b)
	target := withQuery("/events/list", map[string]string{"startdate": enum.ZeroDate, "enddate": enum.MaxDate})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

func BenchmarkSingleEvent(b *testing.B) {
	engine := setupBenchmark(b)
	if benchEvent == "" {
		b.Skip("no events loaded")
	}
	target := withQuery("/events", map[string]string{"fed": benchFed, "id": benchEvent})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

// BenchmarkLeaderboardColdCache measures the leaderboard endpoint with the
// query cache cleared before every call, i.e. what response times look like
// without caching.
func BenchmarkLeaderboardColdCache(b *testing.B) {
	engine := setupBenchmark(b)
	target := defaultLeaderboardQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		resetQueryCache()
		b.StartTimer()
		doRequest(engine, http.MethodGet, target, nil)
	}
}

// BenchmarkLeaderboardWarmCache warms the cache once, then repeats the same
// query, i.e. the steady-state cached response time.
func BenchmarkLeaderboardWarmCache(b *testing.B) {
	engine := setupBenchmark(b)
	target := defaultLeaderboardQuery()
	resetQueryCache()
	doRequest(engine, http.MethodGet, target, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodGet, target, nil)
	}
}

// BenchmarkLeaderboardSearchWarmCache benchmarks /leaderboard/search, which
// requires the matching leaderboard query to already be cached.
func BenchmarkLeaderboardSearchWarmCache(b *testing.B) {
	engine := setupBenchmark(b)
	resetQueryCache()
	doRequest(engine, http.MethodGet, defaultLeaderboardQuery(), nil)

	body, _ := json.Marshal(structs.SearchLeaderboardRequest{
		ActiveQuery: structs.LeaderboardPayload{SortBy: enum.Total, Federation: enum.ALLFEDS, WeightClass: "MALL"},
		LifterData:  benchLifter,
	})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		doRequest(engine, http.MethodPost, "/leaderboard/search", body)
	}
}
