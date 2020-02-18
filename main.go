package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getDataSliceFromFile(name string) []string {
	stringData, err := ioutil.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Split(string(stringData), "\n")
}

func printSlice(data []string) {
	for _, word := range data {
		fmt.Println(word)
	}
}

func getReversedSlice(data []string) []string {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
	return data
}

func writeSliceToFile(data []string, fileName string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for _, word := range data {
		if _, err = f.WriteString(word + "\n"); err != nil {
			log.Fatal(err)
		}
	}

}

func makeUniq(data []string) []string {
	res := make([]string, 0)
	i := 0
	for i < len(data)-1 {
		if data[i] != data[i+1] {
			res = append(res, data[i])
			i++
		} else {
			res = append(res, data[i])
			for i < len(data)-1 && data[i] == data[i+1] {
				i++
			}
			i++
		}
	}

	if data[len(data)-1] != data[len(data)-2] {
		res = append(res, data[len(data)-1])
	}
	return res
}

func sortNumbers(data []string) []string {
	arrNumbers := make([]int, len(data))
	for i, numberStr := range data {
		num, _ := strconv.Atoi(numberStr)
		arrNumbers[i] = num
	}
	sort.Ints(arrNumbers)
	arrStrs := make([]string, len(data))
	for i, word := range arrNumbers {
		arrStrs[i] = strconv.Itoa(word)
	}
	return arrStrs
}

func sortDataByColumn(data []string, colNum int) []string {
	sort.Slice(data, func(i, j int) bool { return data[i][colNum] < data[j][colNum] })
	return data
}

func makeUniqWithF(data []string) []string {
	res := make([]string, 0)
	i := 0
	for i < len(data)-1 {
		if strings.ToLower(data[i]) != strings.ToLower(data[i+1]) {
			res = append(res, data[i])
			i++
		} else {
			res = append(res, data[i])
			for i < len(data)-1 && strings.ToLower(data[i]) == strings.ToLower(data[i+1]) {
				i++
			}
			i++

		}
	}

	if data[len(data)-1] != data[len(data)-2] {
		res = append(res, data[len(data)-1])
	}
	return res
}

func sortWithF(data []string) []string {
	sort.Slice(data, func(i, j int) bool { return strings.ToLower(data[i]) < strings.ToLower(data[j]) })
	return data
}

func sortWithFlags() {
	var fFlag bool
	flag.BoolVar(&fFlag, "f", false, "Flag f")
	var uFlag bool
	flag.BoolVar(&uFlag, "u", false, "Flag u")
	var rFlag bool
	flag.BoolVar(&rFlag, "r", false, "Flag r")
	var nFlag bool
	flag.BoolVar(&nFlag, "n", false, "Flag n")
	var kFlag int
	flag.IntVar(&kFlag, "k", 0, "Flag k")
	var oFlag string
	flag.StringVar(&oFlag, "o", "", "Flag o")

	flag.Parse()

	filePath := os.Args[len(os.Args)-1]

	result := getDataSliceFromFile(filePath)
	sort.Strings(result)
	flags := os.Args[1 : len(os.Args)-1]

	for _, el := range flags {
		if el == "-f" {
			result = sortWithF(result)
		}
		if el == "-r" {
			result = getReversedSlice(result)
		}
		if el == "-n" {
			result = sortNumbers(result)
		}
		if el == "-u" {
			if fFlag == true {
				result = makeUniqWithF(result)
			} else {
				result = makeUniq(result)
			}
		}
		if el == "-k" {
			result = sortDataByColumn(result, kFlag)
		}
	}
	if oFlag != "" {
		writeSliceToFile(result, oFlag)
	} else {
		printSlice(result)
	}

}

func main() {
	sortWithFlags()
}
