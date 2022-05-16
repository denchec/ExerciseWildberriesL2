package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type GrepSettings struct {
	after      int
	before     int
	context    int
	count      bool
	ignoreCase bool
	invert     bool
	fixed      bool
	lineNum    bool
	countFlag  int
	textFile   string
}

var regexpSample *regexp.Regexp

func parsGrepSettings() GrepSettings {
	grep := GrepSettings{}
	flag.IntVar(&grep.after, "A", -1, "Output after matching")
	flag.IntVar(&grep.before, "B", -1, "Output before matching")
	flag.IntVar(&grep.context, "C", -1, "Output around the lines matching")
	flag.BoolVar(&grep.count, "c", false, "Count the lines matching")
	flag.BoolVar(&grep.ignoreCase, "i", false, "Ignore case the matching lines")
	flag.BoolVar(&grep.invert, "v", false, "Ignore case the matching lines")
	flag.BoolVar(&grep.fixed, "F", false, "Exact string match")
	flag.BoolVar(&grep.lineNum, "n", false, "Output string number")
	flag.Parse()

	remainingArgs := flag.Args()

	sample := strings.TrimSpace(remainingArgs[0])
	textFile := strings.TrimSpace(remainingArgs[1])

	regexpSample = regexp.MustCompile(grep.getRXFlag(sample))

	if len(textFile) <= 0 {
		panic("Please write path to a file to sort")
	}

	file, _ := os.ReadFile(textFile)
	grep.textFile = string(file)

	grep.countFlag = flag.NFlag()

	return grep
}

func (grep GrepSettings) getRXFlag(sample string) string {
	if grep.fixed {
		sample = regexp.QuoteMeta(sample)
	}

	if grep.ignoreCase {
		sample = "(?i)" + sample
	}

	return sample
}

func (grep GrepSettings) afterGrep(text []string, sliceWithPatternLines []int) []string {
	var afterSlice []string
	var checkingFlag bool

	for _, i := range sliceWithPatternLines {
		if checkingFlag == true {
			afterSlice = append(afterSlice, "--")
			checkingFlag = false
		}
		afterSlice = append(afterSlice, text[i])
		for o := 1; o < grep.after+1; o++ {
			if regexpSample.MatchString(text[i+o]) {
				break
			}
			if len(text) >= i+o {
				afterSlice = append(afterSlice, text[i+o])
			}
			if o == grep.after && !regexpSample.MatchString(text[i+grep.after+1]) {
				checkingFlag = true
			}
		}
	}

	return afterSlice
}

func (grep GrepSettings) beforeGrep(text []string, sliceWithPatternLines []int) []string {
	var beforeSlice []string
	var checkingFlag bool

	for i := len(sliceWithPatternLines) - 1; i >= 0; i-- {
		if checkingFlag == true {
			beforeSlice = append(beforeSlice, "--")
			checkingFlag = false
		}
		beforeSlice = append(beforeSlice, text[sliceWithPatternLines[i]])
		for o := 1; o < grep.before+1; o++ {
			if sliceWithPatternLines[i]-o < 0 || regexpSample.MatchString(text[sliceWithPatternLines[i]-o]) {
				break
			}
			if sliceWithPatternLines[i]-o >= 0 {
				beforeSlice = append(beforeSlice, text[sliceWithPatternLines[i]-o])
			}
			if sliceWithPatternLines[i]-o-1 >= 0 && o == grep.before &&
				!regexpSample.MatchString(text[sliceWithPatternLines[i]-grep.before-1]) {
				checkingFlag = true
			}
		}
	}

	return beforeSlice
}

func countStringGrep(text []int) int {
	return len(text)
}

func (grep GrepSettings) contextGrep(text []string, sliceWithPatternLines []int) []string {
	var contextSlice []string
	var checkingFlag bool
	var lastPrinted = -1

	for _, i := range sliceWithPatternLines {
		for b := grep.context; b >= 0; b-- {
			if checkingFlag == true && lastPrinted+1 < i-b && b == grep.context {
				contextSlice = append(contextSlice, "--")
				checkingFlag = false
			}
			if i-b > lastPrinted && !regexpSample.MatchString(text[i-b]) {
				contextSlice = append(contextSlice, text[i-b])
			}
		}
		contextSlice = append(contextSlice, text[i])
		for a := 1; a < grep.context+1; a++ {
			if regexpSample.MatchString(text[i+a]) {
				break
			}
			contextSlice = append(contextSlice, text[i+a])
			lastPrinted = i + a
			if a == grep.context && !regexpSample.MatchString(text[i+grep.context+1]) {
				checkingFlag = true
			}
		}
	}

	return contextSlice
}

func (grep GrepSettings) invertGrep(text []string) []string {
	var invertSlice []string

	for i := 0; i < len(text); i++ {
		if regexpSample.MatchString(text[i]) {
			invertSlice = append(invertSlice, text[i])
		}
	}

	return invertSlice
}

func (grep GrepSettings) lineNumGrep(text []string, sliceWithPatternLines []int) []string {
	var lineNumSlice []string

	for _, i := range sliceWithPatternLines {
		numString := strconv.Itoa(i+1) + ":" + text[i]
		lineNumSlice = append(lineNumSlice, numString)
	}

	return lineNumSlice
}

func (grep GrepSettings) theGrep() string {

	textFile := grep.textFile

	if grep.ignoreCase {
		strings.ToLower(textFile)
	}

	text := strings.Split(strings.Trim(textFile, "\r"), "\n")

	var sliceWithPatternLines []int

	if grep.invert {
		text = grep.invertGrep(text)
	}

	for i := 0; i < len(text); i++ {
		if regexpSample.MatchString(text[i]) {
			sliceWithPatternLines = append(sliceWithPatternLines, i)
		}
	}

	if grep.after != -1 {
		text = grep.afterGrep(text, sliceWithPatternLines)
	}

	if grep.before != -1 {
		var reverse []string
		array := grep.beforeGrep(text, sliceWithPatternLines)
		for i := len(array) - 1; i >= 0; i-- {
			reverse = append(reverse, array[i])
		}
		text = array
	}

	if grep.context != -1 {
		text = grep.contextGrep(text, sliceWithPatternLines)
	}

	if grep.count {
		text = strings.Split(strconv.Itoa(countStringGrep(sliceWithPatternLines)), "")

	}

	if grep.lineNum {
		text = grep.lineNumGrep(text, sliceWithPatternLines)
	}

	if grep.countFlag == 0 || grep.fixed {
		var noFlagGrep []string

		for _, i := range sliceWithPatternLines {
			noFlagGrep = append(noFlagGrep, text[i])
		}
		text = noFlagGrep
	}

	endText := strings.Join(text, "\r\n")

	return endText
}

func main() {

	grepSettings := parsGrepSettings()

	grep := grepSettings.theGrep()

	fmt.Printf("%s\n", grep)
}
