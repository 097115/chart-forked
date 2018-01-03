package main

import (
	"reflect"
	"testing"
)

func TestResolveOptions(t *testing.T) {
	tests := []struct {
		args     []string
		expected options
		fails    bool
	}{
		{
			args: []string{},
			expected: options{
				title:     "",
				separator: '\t',
			},
		},
		{
			args: []string{"-t", "title"},
			expected: options{
				title:     "title",
				separator: '\t',
			},
		},
		{
			args: []string{"-title", "title"},
			expected: options{
				title:     "title",
				separator: '\t',
			},
		},
		{
			args:  []string{"-t"},
			fails: true,
		},
		{
			args:  []string{"-title"},
			fails: true,
		},
		{
			args: []string{"bar", "log", "invert", ","},
			expected: options{
				title:     "",
				separator: ',',
				scaleType: logarithmic,
				chartType: bar,
			},
		},
		{
			args: []string{"bar", ";"},
			expected: options{
				title:     "",
				separator: ';',
				chartType: bar,
			},
		},
		{
			args: []string{" "},
			expected: options{
				title:     "",
				separator: ' ',
			},
		},
		{
			args: []string{" ", "pie"},
			expected: options{
				title:     "",
				separator: ' ',
				chartType: pie,
			},
		},
		{
			args: []string{"line"},
			expected: options{
				title:     "",
				separator: '\t',
				chartType: line,
			},
		},
		{
			args: []string{"scatter"},
			expected: options{
				title:     "",
				separator: '\t',
				chartType: scatter,
			},
		},
		{
			args: []string{"-title", "title", "legacy-color", "1"},
			expected: options{
				title:     "title",
				separator: '\t',
				colorType: legacyColor,
			},
		},
		{
			args: []string{"-title", "title", "gradient", "1"},
			expected: options{
				title:     "title",
				separator: '\t',
				colorType: gradient,
			},
		},
	}

	for _, ts := range tests {
		result, err := resolveOptions(ts.args)

		if ts.fails && err == nil {
			t.Errorf("should have failed with %v", result)
		}

		if !ts.fails && err != nil {
			t.Errorf("should not have failed resolving options with %v", result)
		}

		if !ts.fails && !reflect.DeepEqual(result, ts.expected) {
			t.Errorf("options are incorrect: %v was not equal to %v", result, ts.expected)
		}
	}
}
