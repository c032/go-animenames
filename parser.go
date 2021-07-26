package animenames

import (
	"container/list"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/c032/go-textutil"
)

var (
	ErrCouldNotParseName = errors.New("could not parse name")
	ErrNotAString        = errors.New("not a string")
	ErrInvalidSeason     = errors.New("invalid season")
	ErrInvalidEpisode    = errors.New("invalid episode")
	ErrInvalidYear       = errors.New("invalid year")
)

// Parse returns anime information from a file name.
func Parse(name string) (Anime, error) {
	anime := Anime{}

	for _, ext := range extensions {
		if strings.HasSuffix(name, "."+ext) {
			name = strings.TrimSuffix(name, "."+ext)

			break
		}
	}

	chunks := textutil.SplitParens(name)
	if len(chunks) == 0 {
		err := fmt.Errorf("%w: %#v", ErrCouldNotParseName, name)

		return anime, err
	}

	// Either the name has no parens, or the outer parens are wrappinge
	// everything.
	if len(chunks) == 1 {
		var (
			err error

			chunk string
		)

		chunk = strings.TrimSpace(chunks[0])

		// Remove outer parens wrapping the whole chunk.
		chunk = textutil.StripParens(chunk)
		chunk = strings.TrimSpace(chunk)

		chunk = strings.Join(removeKeywords(chunk), " ")

		// Usually when there's parens wrapping everything, there's no inner
		// parens.
		//
		// If names with mixed outer + inner parens start appearing, we
		// probably should change this for a recursive call to `Parse`.
		err = parseMain(chunk, &anime)
		if err != nil {
			return anime, err
		}

		return anime, nil
	}

	return parseMultipleChunks(chunks)
}

func parseMultipleChunks(chunks []string) (Anime, error) {
	anime := Anime{}

	l := chunksToList(chunks)

	// Search CRC32, from right to left.
	for e := l.Back(); e != nil; e = e.Prev() {
		var (
			err error

			chunk string
		)

		chunk, err = elementToString(e)
		if err != nil {
			return anime, err
		}

		noparens := textutil.StripParens(chunk)

		// CRC32 must be inside parens.
		if chunk == noparens {
			continue
		}

		chunk = noparens

		// CRC32 is only one word, so we ignore chunks with more
		// than that.
		words := splitByWords(chunk)
		if len(words) != 1 {
			continue
		}

		checksum := words[0]
		if !isCRC32(checksum) {
			continue
		}

		anime.CRC32 = checksum

		l.Remove(e)

		break
	}

	// Remove extension, searching from right to left.
	for e := l.Back(); e != nil; e = e.Prev() {
		var (
			err error

			chunk string
		)

		chunk, err = elementToString(e)
		if err != nil {
			return anime, err
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

		// Remove the extension from the current element, but keep the element
		// in the list because it might contain additional information.
		e.Value = chunk

		break
	}

	// Try to find the group/fansub at the leftmost chunk.
	if e := l.Front(); e != nil {
		var (
			err error

			chunk string
		)

		chunk, err = elementToString(e)
		if err != nil {
			return anime, err
		}

		// Group is usually inside parens.
		if noparens := textutil.StripParens(chunk); chunk != noparens {
			anime.Group = noparens

			l.Remove(e)
		}
	}

	// At this point we have looked the most common info in their common
	// places. From here onwards the guesswork becomes harder.

	// Used to keep track of text inside parens that's actually part of the
	// title.
	titleSuffix := ""
	keepTitleSuffix := false

	// When naming files, there's a tendency to put the series name at the left
	// side, and the additional information at the right side.
	//
	// Since the additional information consists mostly of well-known keywords,
	// it's better to read chunks from right to left to decrease our chances of
	// accidentally trimming the series name (especially when it contains
	// numbers in the middle).
	for e := l.Back(); e != nil; e = e.Prev() {
		if keepTitleSuffix {
			keepTitleSuffix = false
		} else {
			titleSuffix = ""
		}

		var (
			err error

			chunk string
		)

		chunk, err = elementToString(e)
		if err != nil {
			return anime, fmt.Errorf("invalid chunk %#v: %w", chunk, err)
		}

		chunk = strings.TrimSpace(chunk)
		if chunk == "" {
			continue
		}

		noparens := textutil.StripParens(chunk)

		if chunk != noparens {
			// `chunk` was surrounded by parens. Chances are there's some
			// keywords in here.
			parseKeywords(noparens, &anime)
		}

		// Keywords should be parsed already. We don't need them.
		//
		// NOTE: Maybe combine with `parseKeywords`.
		words := removeKeywords(noparens)

		// No words left. Nothing to do.
		if len(words) == 0 {
			continue
		}

		// If `chunk` is surrounded by parens, and we have found at least one
		// keyword, assume all words are keywords (known or unknown).
		//
		// TODO: Simplify.
		if allwords := splitByWords(chunk); chunk != noparens && len(allwords) > len(words) {
			continue
		}

		// Year.
		//
		// Ignore if we already have it.
		if anime.Year == 0 && len(words) == 1 && regexpYear.MatchString(noparens) {
			var year int

			year, err = strconv.Atoi(noparens)
			if err != nil {
				return anime, fmt.Errorf("could not parse %#v: %w", noparens, ErrInvalidYear)
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
		if anime.Episode == 0 && len(words) == 1 && regexpEpisode.MatchString(noparens) {
			var (
				err error

				episode int
			)

			episode, err = strconv.Atoi(noparens)
			if err != nil {
				return anime, fmt.Errorf("could not parse %#v: %w", noparens, ErrInvalidEpisode)
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
			if len(chunk) >= 2 && chunk[0] == '(' && chunk[len(chunk)-1] == ')' {
				titleSuffix = chunk
				keepTitleSuffix = true
			}

			continue
		}

		// `chunk` is text outside parens. Most likely containing anime title
		// and episode number.

		if titleSuffix != "" {
			chunk += " " + titleSuffix
		}

		err = parseMain(chunk, &anime)
		if err != nil {
			return anime, fmt.Errorf("could not parse chunk %#v: %w", chunk, err)
		}
	}

	return anime, nil
}

// elementToString returns a list element as a string.
func elementToString(e *list.Element) (string, error) {
	var (
		ok bool

		value string
	)

	value, ok = e.Value.(string)

	if !ok {
		return "", ErrNotAString
	}

	return value, nil
}

// parseMain parses a chunk of text outside parens, and updates `*anime`.
func parseMain(chunk string, anime *Anime) error {
	words := splitByWords(chunk)

	// `split` is the index of either the season number or the episode number,
	// whichever has lower value.
	//
	// Since by convention the series name is usually located at the left of
	// the season number or the episode number, we take advantage of this index
	// to discard everything from this index until the end of the chunk.
	split := -1

	// Look for the season number or episode number.
	for i, word := range words {
		m := regexpSeasonEpisode.FindStringSubmatch(word)
		if m == nil {
			continue
		}

		season, err := strconv.Atoi(m[1])
		if err != nil {
			return fmt.Errorf("could not parse %#v: %w", m[1], ErrInvalidSeason)
		}

		episode, err := strconv.Atoi(m[2])
		if err != nil {
			return fmt.Errorf("could not parse %#v: %w", m[2], ErrInvalidEpisode)
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

	// Series title.
	title := ""

	iterationCompleted := true

	words = splitByWords(chunk)
	for i := len(words) - 1; i >= 0; i-- {
		word := words[i]

		// If we don't complete an iteration it means that some special word
		// was found in the previous iteration (e.g. episode number).
		if !iterationCompleted {
			// Reset because anime title can't contain special words.
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

		// Assume "+" is a separator.
		if word == "+" {
			parseKeywords(title, anime)

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

func chunksToList(chunks []string) *list.List {
	l := list.New()

	for _, chunk := range chunks {
		l.PushBack(chunk)
	}

	return l
}
