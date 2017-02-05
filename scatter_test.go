package main

import (
	"testing"
	"time"
)

func TestScatter(t *testing.T) {
	tests := []struct {
		name      string
		fss       [][]float64
		sss       [][]string
		tss       [][]time.Time
		minFSS    []float64
		maxFSS    []float64
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
			name:   "basic working example",
			fss:    [][]float64{[]float64{1}, []float64{2}, []float64{3}},
			sss:    nil,
			minFSS: []float64{1},
			maxFSS: []float64{3},
			title:  "Basic working example",
		},
	}

	for _, ts := range tests {
		templateData, resultScatterTemplate, err := setupScatter(ts.fss, ts.sss, ts.tss, ts.minFSS, ts.maxFSS, ts.title, linear, "", "", false)
		if ts.fails && err == nil {
			t.Errorf("'%v' should have failed", ts.name)
		}

		if !ts.fails {
			if err != nil {
				t.Errorf("'%v' shouldn't have failed, but did with %v", ts.name, err)
			}
			if resultScatterTemplate != scatterTemplate {
				t.Errorf("'%v' appears to not be using the hardcoded scatterTemplate", ts.name)
			}
			if templateData.(scatterTemplateData).Title != ts.title {
				t.Errorf("'%v' did not use the specified title", ts.name)
			}
			if len(templateData.(scatterTemplateData).FSS) == 0 {
				t.Errorf("'%v' dataset is empty", ts.name)
			}
		}
	}
}
