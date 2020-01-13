package animenames

import (
	"regexp"
)

var (
	regexpCRC32         = regexp.MustCompile(`^([0-9a-f]{8}|[0-9A-F]{8})$`)
	regexpVolume        = regexp.MustCompile(`^[Vv]ol\.?([0-9]{1,2})$`)
	regexpSeason        = regexp.MustCompile(`^S([0-9]+)$`)
	regexpEpisode       = regexp.MustCompile(`^([0-9]+)(v[0-9]+)?$`)
	regexpSeasonEpisode = regexp.MustCompile(`^S([0-9]+)E([0-9]+)$`)
	regexpWordSplit     = regexp.MustCompile(`[\s_]`)
	regexpYear          = regexp.MustCompile(`^[0-9]{4}$`)
	regexpBatch         = regexp.MustCompile(`([0-9]+)\-([0-9]+)`)

	regexpSeriesTrim = regexp.MustCompile(`[\s\&\-]+$`)
)
