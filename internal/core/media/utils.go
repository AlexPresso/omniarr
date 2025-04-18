package media

import (
	"strings"
	"unicode"
)

var separators = []string{":", "-", ".", "_", "'"}

func MakeAlternateTitles(title string) []string {
	title = strings.TrimSpace(title)
	titles := map[string]struct{}{
		title: {},
	}

	normalized := title
	for _, sep := range separators {
		normalized = strings.ReplaceAll(normalized, sep, " ")
	}

	cleaned := reduceMultipleSpaces(normalized)
	titles[cleaned] = struct{}{}
	titles[strings.ToLower(cleaned)] = struct{}{}

	collapsed := collapseToAlphanum(title)
	titles[collapsed] = struct{}{}
	titles[strings.ToLower(collapsed)] = struct{}{}

	dehyphenated := mergeHyphenatedWords(title)
	titles[dehyphenated] = struct{}{}
	titles[strings.ToLower(dehyphenated)] = struct{}{}

	spaceReplaced := strings.ReplaceAll(title, "-", " ")
	titles[spaceReplaced] = struct{}{}
	titles[strings.ToLower(spaceReplaced)] = struct{}{}

	if abbr := createAcronymFromTitle(normalized); len(abbr) > 2 {
		titles[strings.Replace(normalized, getLastWords(normalized, 3), abbr, 1)] = struct{}{}
	}

	var alternateTitles []string
	for k := range titles {
		if k != "" {
			alternateTitles = append(alternateTitles, strings.TrimSpace(k))
		}
	}

	return alternateTitles
}

func mergeHyphenatedWords(s string) string {
	words := strings.Split(s, " ")
	for i, word := range words {
		if strings.Contains(word, "-") {
			words[i] = strings.ReplaceAll(word, "-", "")
		}
	}
	return strings.Join(words, " ")
}

func reduceMultipleSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func collapseToAlphanum(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			b.WriteRune(r)
		}
	}
	return reduceMultipleSpaces(b.String())
}

func createAcronymFromTitle(title string) string {
	words := strings.Fields(title)
	start := len(words) - 3
	if start < 0 {
		start = 0
	}
	acronym := ""
	for _, w := range words[start:] {
		if len(w) > 0 {
			acronym += strings.ToUpper(string(w[0]))
		}
	}
	return acronym
}

func getLastWords(s string, n int) string {
	words := strings.Fields(s)
	if len(words) < n {
		return strings.Join(words, " ")
	}
	return strings.Join(words[len(words)-n:], " ")
}
