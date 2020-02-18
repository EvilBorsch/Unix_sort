package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetDataSliceFromFile(t *testing.T) {
	assert := assert.New(t)
	dataSlice := getDataSliceFromFile("test")
	assert.Equal(dataSlice, []string{"aba", "aaa", "paz"})

}

func TestGetReverseSlice(t *testing.T) {
	assert := assert.New(t)
	slice := []string{"1", "2", "3"}
	assert.Equal(getReversedSlice(slice), []string{"3", "2", "1"})
}

func TestWriteSliceToFile(t *testing.T) {
	assert := assert.New(t)
	slice := []string{"asd", "bsd", "csd"}

	writeSliceToFile(slice, "writeSlice")

	assert.Equal(getDataSliceFromFile("writeSlice"), getDataSliceFromFile("writeSlice"))

	slice = []string{"1", "2", "3"}

	writeSliceToFile(slice, "writeSlice")

	assert.Equal(getDataSliceFromFile("writeSlice"), getDataSliceFromFile("writeSlice"))

}

func TestMakeUniq(t *testing.T) {
	assert := assert.New(t)
	slice := []string{"asd", "bsd", "asd", "csd", "csd", "csd", "lsd"}
	assert.Equal([]string{"asd", "bsd", "asd", "csd", "lsd"}, makeUniq(slice),
		"Функкция удаляет только подряд идущие повторяющиеся элементы")
	slice = []string{"2", "2", "1"}
	assert.Equal([]string{"2", "1"}, makeUniq(slice),
		"Функкция удаляет только подряд идущие повторяющиеся элементы")

}

func TestSortNumbers(t *testing.T) {
	assert := assert.New(t)
	slice := []string{"2", "500", "1", "123", "12345"}
	assert.Equal([]string{"1", "2", "123", "500", "12345"}, sortNumbers(slice))

	slice = []string{"123", "124", "125", "121"}
	assert.Equal([]string{"121", "123", "124", "125"}, sortNumbers(slice))

}

func TestSortDataFromColumn(t *testing.T) {
	assert := assert.New(t)
	slice := getDataSliceFromFile("test")
	assert.Equal([]string{"aaa", "paz", "aba"}, sortDataByColumn(slice, 1))

}

func TestSortDataWithF(t *testing.T) {
	assert := assert.New(t)
	slice := getDataSliceFromFile("data")
	assert.Equal([]string{"Apple", "Book", "BOOK", "Go", "Hauptbahnhof", "January", "January", "Napkin"}, sortWithF(slice))
}

func TestMakeUniqWithF(t *testing.T) {
	assert := assert.New(t)
	data := []string{"Napkin", "Apple", "APPLE", "BOOK", "January", "Hauptbahnhof"}
	result := []string{"Napkin", "Apple", "BOOK", "January", "Hauptbahnhof"}
	assert.Equal(result, makeUniqWithF(data))
}
