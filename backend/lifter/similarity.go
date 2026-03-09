package lifter

import (
	"math"
	"sort"
	"strings"
	"unicode"
)

const similarityThreshold = 0.6

// combinedNameScore combines token sort ratio and token-level Jaro-Winkler.
// Token sort ratio handles "LASTNAME Firstname" vs "Firstname LASTNAME" format flips.
// Token-level JW handles typos and prefix matches like "Chris" vs "Christopher".
func combinedNameScore(a, b string) float64 {
	tsr := tokenSortRatio(a, b)
	tjw := tokenJaroWinkler(a, b)
	return 0.4*tsr + 0.6*tjw
}

// tokenSortRatio sorts tokens alphabetically then computes Levenshtein similarity.
// This neutralises the IWF "LASTNAME Firstname" vs "Firstname Lastname" format difference.
func tokenSortRatio(a, b string) float64 {
	return levenshteinSimilarity(sortedTokenString(a), sortedTokenString(b))
}

func sortedTokenString(s string) string {
	tokens := strings.Fields(strings.ToLower(s))
	sort.Strings(tokens)
	return strings.Join(tokens, " ")
}

// tokenJaroWinkler computes, for each token in a, the best Jaro-Winkler match
// among tokens in b (with a phonetic boost via Soundex), then averages the scores.
// This handles partial name matches (e.g. "Chris" ≈ "Christopher") and varied spellings.
func tokenJaroWinkler(a, b string) float64 {
	tokensA := strings.Fields(strings.ToLower(a))
	tokensB := strings.Fields(strings.ToLower(b))
	if len(tokensA) == 0 || len(tokensB) == 0 {
		return 0
	}

	var total float64
	for _, ta := range tokensA {
		soundA := soundex(ta)
		best := 0.0
		for _, tb := range tokensB {
			score := jaroWinkler(ta, tb)
			// Phonetic boost: matching Soundex code guarantees at least 0.75 similarity.
			if soundA != "" && soundA == soundex(tb) && score < 0.75 {
				score = 0.75
			}
			if score > best {
				best = score
			}
		}
		total += best
	}
	return total / float64(len(tokensA))
}

// jaroWinkler returns the Jaro-Winkler similarity in [0, 1].
func jaroWinkler(s1, s2 string) float64 {
	j := jaro(s1, s2)
	prefix := 0
	for i := 0; i < 4 && i < len(s1) && i < len(s2); i++ {
		if s1[i] == s2[i] {
			prefix++
		} else {
			break
		}
	}
	return j + float64(prefix)*0.1*(1-j)
}

// jaro returns the Jaro similarity in [0, 1].
func jaro(s1, s2 string) float64 {
	if s1 == s2 {
		return 1.0
	}
	if len(s1) == 0 || len(s2) == 0 {
		return 0.0
	}

	matchDist := int(math.Max(float64(len(s1)), float64(len(s2)))/2) - 1
	if matchDist < 0 {
		matchDist = 0
	}

	s1Matches := make([]bool, len(s1))
	s2Matches := make([]bool, len(s2))
	matches := 0

	for i := 0; i < len(s1); i++ {
		start := i - matchDist
		if start < 0 {
			start = 0
		}
		end := i + matchDist + 1
		if end > len(s2) {
			end = len(s2)
		}
		for j := start; j < end; j++ {
			if s2Matches[j] || s1[i] != s2[j] {
				continue
			}
			s1Matches[i] = true
			s2Matches[j] = true
			matches++
			break
		}
	}

	if matches == 0 {
		return 0.0
	}

	transpositions := 0
	k := 0
	for i := 0; i < len(s1); i++ {
		if !s1Matches[i] {
			continue
		}
		for !s2Matches[k] {
			k++
		}
		if s1[i] != s2[k] {
			transpositions++
		}
		k++
	}

	return (float64(matches)/float64(len(s1)) +
		float64(matches)/float64(len(s2)) +
		float64(matches-transpositions/2)/float64(matches)) / 3.0
}

// levenshteinSimilarity returns a normalised [0, 1] similarity score.
func levenshteinSimilarity(a, b string) float64 {
	dist := levenshteinDistance(a, b)
	maxLen := len(a)
	if len(b) > maxLen {
		maxLen = len(b)
	}
	if maxLen == 0 {
		return 1.0
	}
	return 1.0 - float64(dist)/float64(maxLen)
}

func levenshteinDistance(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	la, lb := len(ra), len(rb)

	dp := make([][]int, la+1)
	for i := range dp {
		dp[i] = make([]int, lb+1)
		dp[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			if ra[i-1] == rb[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min3(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}
	return dp[la][lb]
}

func min3(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}

// soundex returns a 4-character American Soundex code for phonetic matching.
func soundex(s string) string {
	if len(s) == 0 {
		return ""
	}

	var letters []rune
	for _, r := range strings.ToUpper(s) {
		if unicode.IsLetter(r) {
			letters = append(letters, r)
		}
	}
	if len(letters) == 0 {
		return ""
	}

	table := map[rune]byte{
		'B': '1', 'F': '1', 'P': '1', 'V': '1',
		'C': '2', 'G': '2', 'J': '2', 'K': '2', 'Q': '2', 'S': '2', 'X': '2', 'Z': '2',
		'D': '3', 'T': '3',
		'L': '4',
		'M': '5', 'N': '5',
		'R': '6',
	}

	result := make([]byte, 4)
	result[0] = byte(letters[0])
	pos := 1
	prevCode := table[letters[0]]

	for i := 1; i < len(letters) && pos < 4; i++ {
		r := letters[i]
		if r == 'H' || r == 'W' {
			continue // H and W don't separate same-coded letters
		}
		code, ok := table[r]
		if !ok {
			prevCode = 0 // vowel resets the duplicate-suppression context
			continue
		}
		if code != prevCode {
			result[pos] = code
			pos++
		}
		prevCode = code
	}

	for pos < 4 {
		result[pos] = '0'
		pos++
	}
	return string(result)
}
