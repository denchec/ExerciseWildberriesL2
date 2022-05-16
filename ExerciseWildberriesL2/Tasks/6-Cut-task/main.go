package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	stdinStrings []string
	fields       int
	delimiter    string
	separated    bool
}

func inputData() Data {

	data := Data{}

	flag.IntVar(&data.fields, "f", 0, "column selection")
	flag.StringVar(&data.delimiter, "d", "\t", "Use any delimiter")
	flag.BoolVar(&data.separated, "s", false, "only strings with delimiter")
	flag.Parse()

	rd := bufio.NewReader(os.Stdin)

	fmt.Println("Введите кол-во строк")
	num, _ := rd.ReadString('\n')

	numLines, _ := strconv.Atoi(strings.Trim(num, "\r\n"))

	for i := 0; i < numLines; i++ {
		fmt.Println("Введите строку")
		text, _ := rd.ReadString('\n')
		data.stdinStrings = append(data.stdinStrings, text)
	}

	return data
}

func (data Data) separatedCheck(text []string) []string {
	var separatedStrings []string

	for _, i := range text {
		if strings.Contains(i, data.delimiter) {
			separatedStrings = append(separatedStrings, i)
		}
	}

	return separatedStrings
}

func (data Data) theCut() string {

	text := data.stdinStrings

	if data.fields == 0 {
		log.Fatalln("Fields must be > 0")
	}

	if data.separated {
		text = data.separatedCheck(text)
	}

	var endText string

	for _, i := range text {
		splitStrings := strings.Split(i, data.delimiter)
		if data.fields <= len(splitStrings) {
			endText += splitStrings[data.fields-1] + "\n"
		} else {
			endText += i
		}
	}

	return endText
}

func main() {
	data := inputData()
	cutStrings := data.theCut()
	fmt.Printf("%s", cutStrings)
}
