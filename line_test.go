package main

import (
	"strings"
	"testing"
)

func TestLine(t *testing.T) {
	tests := []struct {
		name      string
		fss       [][]float64
		sss       [][]string
		title     string
		scaleType scaleType
		fails     bool
	}{
		{
			name:  "empty case; should fail",
			fss:   [][]float64{},
			sss:   [][]string{},
			title: "",
			fails: true,
		},
		{
			name:  "inconsistent number of values between data points and labels",
			fss:   [][]float64{[]float64{1}},
			sss:   [][]string{},
			title: "",
			fails: true,
		},
		{
			name:  "basic working example",
			fss:   [][]float64{[]float64{1}, []float64{2}, []float64{3}},
			sss:   [][]string{[]string{"a"}, []string{"b"}, []string{"c"}},
			title: "Basic working example",
		},
	}

	for _, ts := range tests {
		templateData, resultLineTemplate, err := setupLine(ts.fss, ts.sss, ts.title, linear)
		if ts.fails && err == nil {
			t.Errorf("'%v' should have failed", ts.name)
		}

		if !ts.fails {
			if err != nil {
				t.Errorf("'%v' shouldn't have failed, but did with %v", ts.name, err)
			}
			if resultLineTemplate != lineTemplate {
				t.Errorf("'%v' appears to not be using the hardcoded lineTemplate", ts.name)
			}
			if templateData.(lineTemplateData).ChartType != "line" {
				t.Errorf("'%v' appears to not be returning a line chart", ts.name)
			}
			if templateData.(lineTemplateData).Title != ts.title {
				t.Errorf("'%v' did not use the specified title", ts.name)
			}
			if len(templateData.(lineTemplateData).Datasets) == 0 {
				t.Errorf("'%v' dataset is empty", ts.name)
			}
			ss := strings.Split(templateData.(lineTemplateData).Labels, ",")
			if len(ts.sss) != len(ss) {
				t.Errorf("'%v' is returning less labels than specified: %v instead of expected %v", ts.name, len(ss), len(ts.sss))
			}
		}
	}
}
