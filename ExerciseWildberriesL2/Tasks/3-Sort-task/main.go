package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var regexpLeadingNumber = regexp.MustCompile(`^[+-]?\d+`)
var regexpHumanNumber = regexp.MustCompile(`^(?P<sign>[+-]?)(?P<number>\d+)(?P<siSuffix>[kKMGTPEZY]?)`)

type SortSettings struct {
	Column        int
	asNumbers     bool
	reverse       bool
	unique        bool
	behindSpace   bool
	checkToSorted bool
	Month         bool
	HumanSort     bool
	fileContent   string
}

func ParsSortSettings() SortSettings {
	sorted := SortSettings{}
	flag.IntVar(&sorted.Column, "k", -1, "sort by k'th column")
	flag.BoolVar(&sorted.asNumbers, "n", false, "sort as numbers")
	flag.BoolVar(&sorted.reverse, "r", false, "reverse sorting")
	flag.BoolVar(&sorted.unique, "u", false, "return only unique strings")
	flag.BoolVar(&sorted.Month, "M", false, "month sorting")
	flag.BoolVar(&sorted.behindSpace, "b", false, "delete spaces at the end")
	flag.BoolVar(&sorted.checkToSorted, "c", false, "Check to sorted")
	flag.BoolVar(&sorted.HumanSort, "h", false, "human sorting")
	flag.Parse()

	remainingArgs := flag.Args()

	fileName := strings.TrimSpace(strings.Join(remainingArgs, ""))

	if len(fileName) <= 0 {
		panic("Please write path to a file to sort")
	}

	file, _ := os.ReadFile(fileName)
	sorted.fileContent = string(file)

	return sorted
}

func extractColumn(s1 string, column int) string {
	splitOnColumns := strings.Split(s1, " ")
	if column >= 0 && column < len(splitOnColumns) {
		return splitOnColumns[column]
	}
	return s1
}

func getNumericValue(s string) int64 {
	matchedNumber := regexpLeadingNumber.FindString(s)

	if len(matchedNumber) == 0 {
		return 0
	}

	number, err := strconv.ParseInt(matchedNumber, 10, 64)
	if err != nil {
		return 0
	}

	return number
}

type SiNumber struct {
	Sign     int
	Number   int64
	Exponent int
}

func (number SiNumber) Equals(other SiNumber) bool {
	return number.Sign == other.Sign &&
		number.Exponent == other.Exponent &&
		number.Number == other.Number
}

func (number SiNumber) Less(other SiNumber) bool {
	if number.Sign < other.Sign {
		return true
	}

	if number.Sign > other.Sign {
		return false
	}

	if number.Exponent*number.Sign < other.Exponent*other.Sign {
		return true
	}

	if number.Exponent*number.Sign > other.Exponent*other.Sign {
		return false
	}

	return number.Number*int64(number.Sign) < other.Number*int64(other.Sign)
}

func getHumanValue(s string) SiNumber {
	match := regexpHumanNumber.FindStringSubmatch(s)

	if len(match) == 0 {
		return SiNumber{1, 0, 0}
	}

	sSign := match[regexpHumanNumber.SubexpIndex("sign")]
	sNumber := match[regexpHumanNumber.SubexpIndex("number")]
	sSuffix := match[regexpHumanNumber.SubexpIndex("siSuffix")]

	sign := a2sign(sSign)

	number, err := strconv.ParseInt(sNumber, 10, 64)
	if err != nil {
		return SiNumber{sign, 0, 0}
	}

	suffix := siSuffix2Exponent(sSuffix)

	return SiNumber{sign, number, suffix}
}

func a2sign(sSign string) int {
	var sign int
	switch sSign {
	case "-":
		sign = -1
	case "+":
		sign = 1
	case "":
		sign = 1
	}
	return sign
}

func getNumericMonth(s string) int {
	var orderOfMonths = map[string]int{"JAN": 1, "FEB": 2, "MAR": 3, "APR": 4, "MAY": 5, "JUN": 6, "JUL": 7, "AUG": 8,
		"SEP": 9, "OCT": 10, "NOV": 11, "DEC": 12}

	if len(s) < 3 {
		return 0
	}

	return orderOfMonths[s[:3]]
}

func siSuffix2Exponent(sSuffix string) int {
	switch sSuffix {
	case "k":
		return 3
	case "K":
		return 3
	case "M":
		return 6
	case "G":
		return 9
	case "T":
		return 12
	case "P":
		return 15
	case "E":
		return 18
	case "Z":
		return 21
	case "Y":
		return 24
	}
	return 0
}

func makeUnique(strings []string) (uniqueStrings []string) {
	var prevString string
	for i := 0; i < len(strings); i++ {
		if strings[i] == prevString {
			continue
		}

		uniqueStrings = append(uniqueStrings, strings[i])
		prevString = strings[i]
	}
	return
}

func (sortSettings SortSettings) checkToSort(text []string) bool {
	for i := 0; i < len(text); i++ {
		text[i] = strings.Trim(text[i], "\r")
	}

	lineLess := sortSettings.LineLess(text)

	if len(text) <= 1 {
		return true
	}

	for i := 1; i < len(text); i++ {
		if !lineLess(i-1, i) {
			return false
		}
	}
	return true
}

func numericSort(s1 string, s2 string) bool {
	ni := getNumericValue(s1)
	nj := getNumericValue(s2)

	if ni == nj {
		return s1 < s2
	}

	return ni < nj
}

func humanSort(s1 string, s2 string) bool {
	ni := getHumanValue(s1)
	nj := getHumanValue(s2)

	if ni.Equals(nj) {
		return s1 < s2
	}

	return ni.Less(nj)
}

func monthSort(s1 string, s2 string) bool {
	var numS1 = getNumericMonth(s1)
	var numS2 = getNumericMonth(s2)

	if numS1 == numS2 {
		return s1 < s2
	}

	return numS1 < numS2
}

func chooseSortFunc(settings SortSettings) func(string, string) bool {
	if settings.asNumbers {
		return numericSort
	}

	if settings.Month {
		return monthSort
	}

	if settings.HumanSort {
		return humanSort
	}

	return func(s1 string, s2 string) bool {
		return s1 < s2
	}
}

func (sortSettings SortSettings) LineLess(text []string) func(i int, j int) bool {
	return func(i, j int) bool {
		if sortSettings.behindSpace {
			text[i] = strings.TrimSpace(text[i])
			text[j] = strings.TrimSpace(text[j])
		}

		s1 := extractColumn(text[i], sortSettings.Column)
		s2 := extractColumn(text[j], sortSettings.Column)

		sortFunc := chooseSortFunc(sortSettings)
		numberSorted := sortFunc(s1, s2)

		if sortSettings.reverse {
			return !numberSorted
		}

		return numberSorted
	}
}

func (sortSettings SortSettings) theSort() string {

	if sortSettings.checkToSorted {
		text := strings.Split(sortSettings.fileContent, "\n")
		isSorted := sortSettings.checkToSort(text)
		if isSorted {
			return "Sorted"
		}
		return "Not sorted"
	}

	text := strings.Split(sortSettings.fileContent, "\n")
	for i := 0; i < len(text); i++ {
		text[i] = strings.Trim(text[i], "\r")
	}

	sort.Slice(text, sortSettings.LineLess(text))

	if sortSettings.unique {
		text = makeUnique(text)
	}

	endText := strings.Join(text, "\n")

	return endText
}

func main() {
	sortSettings := ParsSortSettings()

	TextToBeSorted := sortSettings.theSort()

	fmt.Printf("%s", TextToBeSorted)
}
