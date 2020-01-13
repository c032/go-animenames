package animenames

import (
	"container/list"
	"errors"
	"strconv"
	"strings"

	"git.wark.io/lib/textutil-go"
)

// Parse returns anime information from a file name.
func Parse(name string) (anime *Anime, err error) {
	var (
		chunk   string
		year    int
		episode int
	)

	anime = new(Anime)

	for _, ext := range extensions {
		if strings.HasSuffix(name, "."+ext) {
			name = strings.TrimSuffix(name, "."+ext)
			break
		}
	}

	chunks := textutil.SplitParens(name)

	// There's nothing to parse.
	if len(chunks) == 0 {
		return
	}

	if len(chunks) == 1 {
		chunk = strings.TrimSpace(chunks[0])

		// If the whole chunk is wrapped inside parens, remove them.
		chunk = textutil.StripParens(chunk)
		chunk = strings.TrimSpace(chunks[0])

		chunk = strings.Join(removeKeywords(chunk), " ")

		err = parseMain(anime, chunk)
		if err != nil {
			return
		}

		return
	}

	// More than 1 chunk.

	// Convert the chunk slice to a doubly-linked list.
	l := list.New()
	for _, chunk := range chunks {
		l.PushBack(chunk)
	}

	// Find CRC32.
	//
	// Search from right to left.
	for e := l.Back(); e != nil; e = e.Prev() {
		chunk, err = elementToString(e)
		if err != nil {
			return
		}

		noparens := textutil.StripParens(chunk)

		// CRC32 must be inside parens.
		if chunk == noparens {
			continue
		}

		chunk = noparens

		// CRC32 is only one word, so we ignore chunks with more
		// than that.
		words := regexpWordSplit.Split(chunk, -1)
		if len(words) != 1 {
			continue
		}

		checksum := words[0]

		// Check whether CRC32 has valid syntax.
		if !regexpCRC32.MatchString(checksum) {
			continue
		}

		anime.CRC32 = checksum

		// Remove the CRC32 node from the list.
		l.Remove(e)

		break
	}

	// Remove extension.
	//
	// Search from right to left.
	for e := l.Back(); e != nil; e = e.Prev() {
		chunk, err = elementToString(e)
		if err != nil {
			return
		}

		noparens := textutil.StripParens(chunk)

		// Extension can't be inside parens.
		if chunk != noparens {
			continue
		}

		chunk = strings.TrimSpace(chunk)

		// Compare against a list of common extensions.
		for _, ext := range extensions {
			if strings.HasSuffix(chunk, "."+ext) {
				// Remove the extension.
				chunk = strings.TrimSuffix(chunk, "."+ext)
				break
			}
		}

		// Remove the extension from the current node.
		e.Value = chunk

		break
	}

	// Try to find the group/fansub at the leftmost chunk.
	if e := l.Front(); e != nil {
		chunk, err = elementToString(e)
		if err != nil {
			return
		}

		// Group is usually inside parens.
		if noparens := textutil.StripParens(chunk); chunk != noparens {
			anime.Group = noparens

			// Remove the "group" node from the list.
			l.Remove(e)
		}
	}

	// We looked the most common info in their common places.

	// Try to parse everything else.
	//
	// Read from right to left.
	for e := l.Back(); e != nil; e = e.Prev() {
		chunk, err = elementToString(e)
		if err != nil {
			return
		}

		chunk = strings.TrimSpace(chunk)
		if chunk == "" {
			continue
		}

		noparens := textutil.StripParens(chunk)

		// If text is inside parens, try to find useful info.
		if chunk != noparens {
			parseKeywords(anime, noparens)
		}

		// Keywords should be parsed already. We don't need them.
		words := removeKeywords(noparens)

		// No words left. Nothing to do.
		if len(words) == 0 {
			continue
		}

		// If there's at least one keyword, and we're inside parens,
		// assume all words are keywords.
		if allwords := regexpWordSplit.Split(chunk, -1); chunk != noparens && len(allwords) > len(words) {
			continue
		}

		// Year.
		//
		// Ignore if we already have it.
		if anime.Year == 0 && len(words) == 1 && regexpYear.MatchString(noparens) {
			year, err = strconv.Atoi(noparens)
			if err != nil {
				return
			}
			anime.Year = year
			continue
		}

		// Episode number.
		//
		// Ignore if we already have it.
		//
		// Some shows have an episode 0. In those cases it should be
		// parsed to 0 again (must confirm this).
		//
		// FIXME: Use -1 to indicate "no episode"?
		if anime.Episode == 0 && len(words) == 1 && regexpEpisode.MatchString(noparens) {
			episode, err = strconv.Atoi(noparens)
			if err != nil {
				return
			}
			anime.Episode = episode
			continue
		}

		// Group.
		//
		// Ignore if we already have it.
		if anime.Group == "" && len(words) == 1 {
			anime.Group = words[0]
			continue
		}

		// We got a text inside parens, but we don't know how to parse
		// it.
		//
		// Just ignore.
		if chunk != noparens {
			continue
		}

		// `chunk` is text outside parens. Most likely containing
		// anime title and episode number.

		err = parseMain(anime, chunk)
		if err != nil {
			return
		}
	}

	return
}

// elementToString returns a list element as a string.
func elementToString(e *list.Element) (value string, err error) {
	var ok bool

	value, ok = e.Value.(string)

	if !ok {
		err = errors.New("element is not string")
	}

	return
}

// IsKeyword returns true when word is a known keyword and false otherwise.
func IsKeyword(word string) bool {
	if _, ok := aliases[word]; ok {
		return true
	}

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
			if word == keyword {
				return true
			}
		}
	}

	return false
}

// removeKeywords returns a slice with words from text, ignoring any keywords.
func removeKeywords(text string) []string {
	words := make([]string, 0)
	allWords := regexpWordSplit.Split(text, -1)

	for _, word := range allWords {
		if IsKeyword(strings.ToLower(word)) {
			continue
		}
		words = append(words, word)
	}

	return words
}

// parseKeywords tries to extract information from the keywords present in
// the chunk.
//
// anime will be updated accordingly.
func parseKeywords(anime *Anime, chunk string) {
	words := regexpWordSplit.Split(chunk, -1)
	for _, word := range words {
		lword := strings.ToLower(word)

		if lword == "bd" || lword == "bdrip" || lword == "blu-ray" || lword == "bluray" {
			anime.IsBD = true
			continue
		}
	}
}

// parseMain parses a chunk of text outside parens.
func parseMain(anime *Anime, chunk string) error {
	words := regexpWordSplit.Split(chunk, -1)
	split := -1

	for i, word := range words {
		m := regexpSeasonEpisode.FindStringSubmatch(word)
		if m == nil {
			continue
		}

		season, err := strconv.Atoi(m[1])
		if err != nil {
			return err
		}

		episode, err := strconv.Atoi(m[2])
		if err != nil {
			return err
		}

		anime.Season = season
		anime.Episode = episode

		split = i
		break
	}

	if split > -1 {
		chunk = strings.Join(words[:split], " ")
	}

	ignore := struct {
		Season  bool
		Episode bool
		Volume  bool
	}{
		Season:  false,
		Episode: false,
		Volume:  false,
	}

	title := ""

	iterationCompleted := true

	words = regexpWordSplit.Split(chunk, -1)
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]

		// If we don't complete an iteration it means that some
		// special word was found in the previous iteration
		// (e.g. episode number).
		if !iterationCompleted {
			// Anime title can't contain special words.
			title = ""
		}

		iterationCompleted = false

		// Episode number.
		if !ignore.Episode {

			// Simple episode number.
			if m := regexpEpisode.FindStringSubmatch(word); m != nil {
				episode, err := strconv.Atoi(m[1])
				if err != nil {
					return err
				}

				anime.Episode = episode

				ignore.Episode = true
				continue
			}

			// More than one episode.
			if m := regexpBatch.FindStringSubmatch(word); m != nil {
				start, err := strconv.Atoi(m[1])
				if err != nil {
					return err
				}

				end, err := strconv.Atoi(m[2])
				if err != nil {
					return err
				}

				anime.Batch = &Batch{
					Start: start,
					End:   end,
				}

				ignore.Episode = true
				continue
			}
		}

		// Season.
		if !ignore.Season {
			if m := regexpSeason.FindStringSubmatch(word); m != nil {
				season, err := strconv.Atoi(m[1])
				if err != nil {
					return err
				}
				anime.Season = season
				ignore.Season = true

				continue
			}
		}

		// Check if name contains more than one season (e.g. a batch).
		if ignore.Season && anime.Season != 0 {
			if m := regexpSeason.FindStringSubmatch(word); m != nil {
				anime.Season = 0
				continue
			}
		}

		// Volume.
		if !ignore.Volume {
			if m := regexpVolume.FindStringSubmatch(word); m != nil {
				volume, err := strconv.Atoi(m[1])
				if err != nil {
					return err
				}
				anime.Volume = volume
				ignore.Volume = true
				continue
			}
		}

		// Case sensitive.
		if word == "OVA" {
			anime.IsOVA = true
			continue
		}

		// Case sensitive.
		if word == "BD" {
			anime.IsBD = true
			continue
		}

		// If nothing matches, just add the word to the title.
		//
		// Remember we're reading from right to left.
		title = word + " " + title

		iterationCompleted = true
	}

	// Remove some useless characters from the title.
	title = strings.TrimSpace(title)
	title = regexpSeriesTrim.ReplaceAllString(title, "")

	// Finally.
	anime.Title = title

	return nil
}
