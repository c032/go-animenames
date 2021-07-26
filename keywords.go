package animenames

import (
	"strings"
)

var aliases = map[string]string{
	"blu-ray": "bd",
	"bluray":  "bd",
	"h.264":   "h264",
	"x264":    "h264",
}

var resolutions = []string{
	"1080p",
	"360p",
	"480p",
	"720p",
}

var quality = []string{
	"10bit",
	"8bit",
	"bd",
	"dvd",
	"tv",
}

var videoCodecs = []string{
	"h264",
	"hevc",
}

var audioCodecs = []string{
	"aac",
	"ac3",
	"flac",
	"opus",
}

var extensions = []string{
	"ass",
	"mkv",
	"mp4",
}

var otherProperties = []string{
	"censored",
	"complete",
	"dual",
	"dub",
	"english",
	"simuldub",
	"uncensored",
}

var keywordsMap = map[string]bool{}

func init() {
	lists := [][]string{
		resolutions,
		quality,
		videoCodecs,
		audioCodecs,
		extensions,
		otherProperties,
	}

	for _, list := range lists {
		for _, keyword := range list {
			keywordsMap[keyword] = true
		}
	}
}

// isKeyword returns true when word is a known keyword and false otherwise.
func isKeyword(word string) bool {
	if _, ok := aliases[word]; ok {
		return true
	}

	return keywordsMap[word]
}

// removeKeywords returns a slice with words from text, without any keywords.
func removeKeywords(text string) []string {
	words := make([]string, 0)
	allWords := splitByWords(text)

	for _, word := range allWords {
		if isKeyword(strings.ToLower(word)) {
			continue
		}
		words = append(words, word)
	}

	return words
}

// parseKeywords tries to extract information from the keywords present in the
// chunk, and updates `*anime`.
//
// TODO: Return something indicating what was changed.
func parseKeywords(chunk string, anime *Anime) {
	words := splitByWords(chunk)

	for _, word := range words {
		lword := strings.ToLower(word)

		if lword == "bd" || lword == "bdrip" || lword == "blu-ray" || lword == "bluray" {
			anime.IsBD = true

			continue
		}
	}
}
