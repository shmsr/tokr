package tokr

import (
	"fmt"
	"strings"

	"github.com/kljensen/snowball/english"
)

type document []string

func (d document) PrettyPrint() {
	for i := range d {
		fmt.Print("[", d[i], "]")
	}
}

func (d document) Lowercase() document {
	r := make(document, 0, len(d))
	for i := range d {
		r = append(r, strings.ToLower(d[i]))
	}

	return r
}

func (d document) Stemmer() document {
	r := make([]string, len(d))
	for i := range d {
		r[i] = english.Stem(d[i], false)
	}

	return r
}

func (d document) StopWord() document {
	alnum := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k",
		"l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v",
		"w", "x", "y", "z", "1", "2", "3", "4", "5", "6", "7",
		"8", "9", "0"}

	words := []string{
		"a", "and", "be", "have", "i", "in", "of", "that",
		"the", "to", "for", "etc", "as", "an", "is", "by",
		"on",
	}

	words = append(words, alnum...)

	wordsmap := make(map[string]struct{})
	for _, word := range words {
		wordsmap[word] = struct{}{}
	}

	r := make([]string, 0, len(d))
	for i := range d {
		if _, ok := wordsmap[d[i]]; !ok {
			r = append(r, d[i])
		}
	}

	return r
}
