package tokr

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	// Debug is a set of single-letter flag(s):
	//  r   show [r]esult logging
	//	v	show [v]erbose logging
	//
	Debug = "r"
)

func dbg(b byte) bool { return strings.IndexByte(Debug, b) >= 0 }

func empty(text string) bool {
	return strings.TrimSpace(text) == ""
}

type keymap map[int]document
type keyhit map[int]int

func Run(doc, keyword string) error {
	kd, err := os.Open(keyword)
	if err != nil {
		return err
	}
	defer kd.Close()

	kscanner := bufio.NewScanner(kd)
	kscanner.Split(bufio.ScanLines)

	var text string

	kmap, khit := make(keymap), make(keyhit)

	lineno := 0
	for kscanner.Scan() {
		lineno++
		if text = kscanner.Text(); empty(text) {
			continue
		}
		kmap[lineno] = tokenize(text).Lowercase().StopWord().Stemmer()
	}

	dd, err := os.Open(doc)
	if err != nil {
		return err
	}
	defer dd.Close()

	dscanner := bufio.NewScanner(dd)
	dscanner.Split(bufio.ScanLines)

	var tokens document
	for dscanner.Scan() {
		if text = dscanner.Text(); empty(text) {
			continue
		}

		tokens = tokenize(text).Lowercase().StopWord().Stemmer()

		for lineno, keywords := range kmap {
			khit[lineno] += tokens.Count(keywords)
		}

		if dbg('v') {
			fmt.Println(">> Line:", text)
			fmt.Print(">> Tokens: ")
			tokens.PrettyPrint()
			fmt.Println()
		}
	}

	if dbg('r') {
		_, err := kd.Seek(0, 0)
		if err != nil {
			return err
		}

		kscanner = bufio.NewScanner(kd)

		lineno = 0
		for kscanner.Scan() {
			lineno++
			fmt.Printf("%s = %d hits\n", kscanner.Text(), khit[lineno])
		}
	}

	return nil
}
