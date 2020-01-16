package animenames

import (
	"regexp"
)

var (
	regexpVolume        = regexp.MustCompile(`^[Vv]ol\.?([0-9]{1,2})$`)
	regexpSeason        = regexp.MustCompile(`^S([0-9]+)$`)
	regexpEpisode       = regexp.MustCompile(`^([0-9]+)(v[0-9]+)?$`)
	regexpSeasonEpisode = regexp.MustCompile(`^S([0-9]+)E([0-9]+)$`)
	regexpWordSplit     = regexp.MustCompile(`[\s_]`)
	regexpYear          = regexp.MustCompile(`^[0-9]{4}$`)
	regexpBatch         = regexp.MustCompile(`([0-9]+)\-([0-9]+)`)

	regexpSeriesTrim = regexp.MustCompile(`[\s\&\-]+$`)
)

func isCRC32(s string) bool {
	if len(s) != 8 {
		return false
	}

	for _, c := range s {
		if c >= '0' && c <= '9' {
			continue
		}
		if c >= 'a' && c <= 'f' {
			continue
		}
		if c >= 'A' && c <= 'F' {
			continue
		}

		return false
	}

	return true
}

func splitByWords(s string) []string {
	words := regexpWordSplit.Split(s, -1)

	return words
}
