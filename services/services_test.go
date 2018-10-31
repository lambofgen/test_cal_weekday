package services

import (
	"strconv"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	testCase := []float64{2000, 1940, 1864, 2104, 1960, 2408, 2116}
	for _, year := range testCase {
		result := isLeapYear(year)
		var expected = true
		if result != expected {
			t.Errorf("Value %v Expected %s but got %s", year, strconv.FormatBool(expected), strconv.FormatBool(result))
		}
	}
}

func TestIsNotLeapYear(t *testing.T) {
	testCase := []float64{2100, 1900, 1943, 1865, 1981, 1969, -2013}
	for _, year := range testCase {
		result := isLeapYear(year)
		var expected = false
		if result != expected {
			t.Errorf("Value %v Expected %s but got %s", year, strconv.FormatBool(expected), strconv.FormatBool(result))
		}
	}
}

func TestSumDayOfTargetYear(t *testing.T) {
	testCase := []struct {
		year     float64
		month    float64
		day      float64
		expected float64
	}{
		{2018, 5, 10, 130},
		{2100, 10, 31, 304},
		{1991, 7, 4, 185},
		{1924, 3, 26, 86},
		{1900, 1, 1, 1},
		{123, 1, 1, 1},
		{-12321, 21, 2, 0},
		{2000, 31, 32, 0},
	}
	for _, date := range testCase {
		result, _ := sumDayOfTargetYear(date.year, date.month, date.day)
		var expected = date.expected
		if result != expected {
			t.Errorf("Year: %v Month: %v Day: %v || Expected %v but got %v", date.year, date.month, date.day, expected, result)
		}
	}
}

func TestCalCountOfLeapYear(t *testing.T) {
	testCase := []struct {
		baseYear   float64
		targetYear float64
		expected   float64
	}{
		{1900, 2005, 26},
		{1900, 1948, 12},
		{-2093, -1233, 0},
		{2000, 1923, 0},
	}
	for _, year := range testCase {
		result, _ := calCountOfLeapYear(year.baseYear, year.targetYear)
		var expected = year.expected
		if result != expected {
			t.Errorf("BaseYear: %v TargetYear: %v || Expected %v but got %v", year.baseYear, year.targetYear, expected, result)
		}
	}

}

func TestCalWeekDay(t *testing.T) {
	testCase := []struct {
		year     float64
		month    float64
		day      float64
		expected string
	}{
		{1900, 1, 1, "Monday"},
		{1900, 1, 2, "Tuesday"},
		{1991, 10, 5, "Saturday"},
		{2018, 10, 2, "Tuesday"},
		{1994, 1, 18, "Tuesday"},
		{2117, 1, 1, "Friday"},
		{2118, 12, 31, "Saturday"},
	}
	for _, date := range testCase {
		result, _ := CalWeekDay(date.year, date.month, date.day)
		var expected = date.expected
		if result != expected {
			t.Errorf("Year: %v Month: %v Day: %v || Expected %v but got %v", date.year, date.month, date.day, expected, result)
		}
	}
}
