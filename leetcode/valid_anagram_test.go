package leetcode

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	anagramPairs := map[string]string{
		"listen": "silent",
		"evil":   "vile",
		"god":    "dog",
		"rat":    "tar",
		"bare":   "bear",
		"stop":   "tops",
		"night":  "thing",
		"part":   "trap",
		"stare":  "tears",
		"leap":   "peal",
		"cinema": "iceman",
		"save":   "vase",
		"read":   "dear",
		"form":   "from",
		"drawer": "reward",
		"desire": "reside",
		"angle":  "glean",
		"cheat":  "teach",
		"fail":   "life",
	}
	for s1, s2 := range anagramPairs {
		t.Run(fmt.Sprintf("anagram-pair: (%q, %q)", s1, s2), func(t *testing.T) {
			MatchValuesTrueAnagrams(t, s1, s2)
		})
	}

	//t.Run("'dog' and 'god' are anagrams", func(t *testing.T) {
	//	got := ValidAnagram("god", "dog")
	//	want := true
	//	if want != got {
	//		t.Errorf("want: %t, got: %t", want, got)
	//	}
	//})
	//t.Run("'ancestries' and 'resistance' are anagrams", func(t *testing.T) {
	//	got := ValidAnagram("ancestries", "resistance")
	//	want := true
	//	if want != got {
	//		t.Errorf("want: %t, got: %t", want, got)
	//	}
	//})
}

func MatchValuesTrueAnagrams(tb testing.TB, s1, s2 string) {
	tb.Helper()
	got := ValidAnagram(s1, s2)
	want := true
	if want != got {
		tb.Errorf("want: %t, got: %t [for string-paris: (%q, %q)]", want, got, s1, s2)
	}
}
