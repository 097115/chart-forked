package main

import (
	"reflect"
	"testing"
	"time"
)

func TestParseLineFormat(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		sep      rune
		df       string
		expected string
	}{
		{
			name:     "empty case",
			s:        "",
			sep:      '\t',
			expected: "s",
		},
		{
			name:     "basic tab",
			s:        "1\t2\t3",
			sep:      '\t',
			expected: "fff",
		},
		{
			name:     "basic space",
			s:        "1 2 3",
			sep:      ' ',
			expected: "fff",
		},
		{
			name:     "basic comma",
			s:        "1,2,3",
			sep:      ',',
			expected: "fff",
		},
		{
			name:     "basic semicolon",
			s:        "1;2;3",
			sep:      ';',
			expected: "fff",
		},
		{
			name:     "space with extras before, in between and after",
			s:        "  1   2  3  ",
			sep:      ' ',
			expected: "fff",
		},
		{
			name:     "commas with complete floating numbers",
			s:        "-1,2.0e3,-3.239847E-1",
			sep:      ',',
			expected: "fff",
		},
		{
			name:     "subsequent commas",
			s:        ",,",
			sep:      ',',
			expected: "ss",
		},
		{
			name:     "string and float",
			s:        "a,1",
			sep:      ',',
			expected: "sf",
		},
		{
			name:     "float and string",
			s:        "1,a",
			sep:      ',',
			expected: "fs",
		},
		{
			name:     "float and string; ignore other separators",
			s:        "1,a;b c\td",
			sep:      ',',
			expected: "fs",
		},
	}

	for _, ts := range tests {
		result := parseLineFormat(ts.s, ts.sep, ts.df)
		if !reflect.DeepEqual(result, ts.expected) {
			t.Errorf("'%v' failed: %v was not equal to %v", ts.name, result, ts.expected)
		}
	}
}

func TestParseFormat(t *testing.T) {
	tests := []struct {
		name     string
		i        []string
		sep      rune
		df       string
		expected string
	}{
		{
			name:     "empty case",
			i:        []string{},
			sep:      '\t',
			expected: "",
		},
		{
			name:     "string, float",
			i:        []string{"a\t1", "b\t2", "c\t3"},
			sep:      '\t',
			expected: "sf",
		},
		{
			name:     "string, float with one outlier (minority)",
			i:        []string{"a\t1", "onlystring", "c\t3"},
			sep:      '\t',
			expected: "sf",
		},
	}

	for _, ts := range tests {
		result := parseFormat(ts.i, ts.sep, ts.df)
		if !reflect.DeepEqual(result, ts.expected) {
			t.Errorf("'%v' failed: %v was not equal to %v", ts.name, result, ts.expected)
		}
	}
}

func TestParseLine(t *testing.T) {
	tests := []struct {
		name            string
		i               string
		sep             rune
		df              string
		format          string
		expectedFloats  []float64
		expectedStrings []string
		expectedTimes   []time.Time
		fails           bool
	}{
		{
			name:            "base case",
			i:               "1,2,3",
			sep:             ',',
			format:          "fff",
			expectedFloats:  []float64{1, 2, 3},
			expectedStrings: []string{},
			expectedTimes:   []time.Time{},
			fails:           false,
		},
		{
			name:            "base failing case",
			i:               "1,a,3",
			sep:             ',',
			format:          "fff",
			expectedFloats:  []float64{},
			expectedStrings: []string{},
			expectedTimes:   []time.Time{},
			fails:           true,
		},
		{
			name:            "with strings",
			i:               "a,1",
			sep:             ',',
			format:          "sf",
			expectedFloats:  []float64{1},
			expectedStrings: []string{"a"},
			expectedTimes:   []time.Time{},
			fails:           false,
		},
		{
			name:            "strings and extra whitespace",
			i:               "    a   ,   1   ",
			sep:             ',',
			format:          "sf",
			expectedFloats:  []float64{1},
			expectedStrings: []string{"a"},
			expectedTimes:   []time.Time{},
			fails:           false,
		},
	}

	for _, ts := range tests {
		fs, ss, ds, err := parseLine(ts.i, ts.format, ts.sep, ts.df)
		if err != nil && !ts.fails {
			t.Errorf("'%v' failed: should not have failed but did! With [%v]", ts.name, err)
		}
		if err == nil && ts.fails {
			t.Errorf("'%v' failed: should have failed but didn't!", ts.name)
		}
		if !ts.fails && !reflect.DeepEqual(fs, ts.expectedFloats) {
			t.Errorf("'%v' failed: %v != %v", ts.name, fs, ts.expectedFloats)
		}
		if !ts.fails && !reflect.DeepEqual(ss, ts.expectedStrings) {
			t.Errorf("'%v' failed: %v != %v", ts.name, ss, ts.expectedStrings)
		}
		if !ts.fails && !reflect.DeepEqual(ds, ts.expectedTimes) {
			t.Errorf("'%v' failed: %v != %v", ts.name, ds, ts.expectedTimes)
		}
	}
}
