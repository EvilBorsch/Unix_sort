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

func getDataSliceFromFile(name string) ([]string, error) {
	stringData, err := ioutil.ReadFile(name)

	return strings.Split(string(stringData), "\n"), err
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

func writeSliceToFile(data []string, fileName string) error {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	for i, word := range data {
		if i == len(data)-1 {
			f.WriteString(word)
			break
		}
		if _, err = f.WriteString(word + "\n"); err != nil {
			return err
		}
	}
	return nil
}

func makeUniq(data []string) []string {
	res := make([]string, 0)
	i := 0
	for i < len(data)-1 {
		if data[i] != data[i+1] {
			res = append(res, data[i])
			i++
		}
		res = append(res, data[i])
		for i < len(data)-1 && data[i] == data[i+1] {
			i++
		}
		i++

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
		}
		res = append(res, data[i])
		for i < len(data)-1 && strings.ToLower(data[i]) == strings.ToLower(data[i+1]) {
			i++
		}
		i++

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

func sortWithFlags(result []string) ([]string, string) {
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

	sort.Strings(result)
	if fFlag == true {
		result = sortWithF(result)
	}
	if rFlag == true {
		result = getReversedSlice(result)
	}
	if nFlag == true {
		result = sortNumbers(result)
	}
	if uFlag == true {
		if fFlag == true {
			result = makeUniqWithF(result)
		} else {
			result = makeUniq(result)
		}
	}
	if kFlag != 0 {
		result = sortDataByColumn(result, kFlag)
	}

	return result, oFlag
}

func getDataFromFileByStdIn() ([]string, error) {
	filePath := os.Args[len(os.Args)-1]
	return getDataSliceFromFile(filePath)
}

func printOrWriteToFile(data []string, oFlag string) error {
	if oFlag != "" {
		err := writeSliceToFile(data, oFlag)
		return err
	}
	printSlice(data)
	return nil
}

func main() {
	data, _ := getDataFromFileByStdIn()
	sortedData, oFlag := sortWithFlags(data)
	err := printOrWriteToFile(sortedData, oFlag)
	if err != nil {
		log.Fatal(err)
	}

}
