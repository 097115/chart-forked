# chart [![Build Status](https://img.shields.io/travis/marianogappa/chart.svg)](https://travis-ci.org/marianogappa/chart) [![Coverage Status](https://coveralls.io/repos/github/MarianoGappa/chart/badge.svg?branch=master&nocache=1)](https://coveralls.io/github/MarianoGappa/chart?branch=master) [![GitHub license](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/marianogappa/chart/master/LICENSE)

Quick & smart charting for STDIN

[Blogpost](https://movio.co/en/blog/improving-with-sql-and-charts/)

## Learn by example!

[Cheatsheet](https://marianogappa.github.io/chart/)

## Syntax

```
chart [options]
```

- `pie`: render a pie chart
- `bar`: render a bar chart
- `line`: render a line chart
- `scatter`: render a scatter plot chart
- `log`: use logarithmic scale (bar chart only)
- `' '|';'|','|'\t'`: this character separates columns on each line (\t = default)
- `-t|--title`: title for the chart
- `-x`: label for the x axis
- `-y`: label for the y axis
- `--date-format`: Sets the date format, according to [https://golang.org/src/time/format.go](https://golang.org/src/time/format.go)
- `--debug`: Use to make sure to double-check the chart is showing what you expect.
- `-h|--help`: Show help
- `--zero-based`: Makes y-axis begin at zero

## Installation

```
go get -u github.com/marianogappa/chart
```

or get the latest binary for your OS in the [Releases section](https://github.com/MarianoGappa/chart/releases).

## Example use cases

- Pie chart of your most used terminal commands
```
history | awk '{print $2}' | chart
```

![Pie chart of your most used terminal commands](img/pie.png?v=1)

- Bar chart of today's currency value against USD, in logarithmic scale
```
curl -s http://api.fixer.io/latest?base=USD | jq -r ".rates | to_entries| \
    map(\"\(.key)\t\(.value|tostring)\")|.[]" | chart bar log -t "Currency value against USD"
```

![Bar chart of today's currency value against USD, in logarithmic scale](img/bar-log.png?v=1)

- Bar chart of a Github user's lines of code per language (requires setting up an Access Token)
```
USER=???
ACCESS_TOKEN=???
curl -u $USER:$ACCESS_TOKEN -s "https://api.github.com/user/repos" | \
    jq -r 'map(.languages_url) | .[]' | xargs curl -s -u $USER:$ACCESS_TOKEN | \
    jq -r '. as $in| keys[] | [.+ " "]+[$in[.] | tostring] | add' | \
    awk '{arr[$1]+=$2} END {for (i in arr) {print i,arr[i]}}' | \
    awk '{print $2 "\t" $1}' | sort -nr | chart bar
```

![Bar chart of a Github user's lines of code per language (requires setting up an Access Token)](img/bar.png?v=1)

- Line chart of the stargazers of this repo over time up to Jan 2017 (received some attention after the publication of [this](https://movio.co/blog/migrate-Scala-to-Go/) blogpost)
```
curl -s "https://api.github.com/repos/marianogappa/chart/stargazers?page=1&per_page=100" \
-H"Accept: application/vnd.github.v3.star+json" | \
jq --raw-output 'map(.starred_at) | .[]' | awk '{print NR "\t" $0}' | \
chart line --date-format 2006-01-02T15:04:05Z
```

![Line chart of Github stargazers of this repo over time](img/line.png?v-1)

## Charting MySQL output

`chart` works great with [sql](https://github.com/MarianoGappa/sql), or with any `mysql -Nsre '...'` query.

## I don't trust the chart is correct

Me neither. Add `--debug` to double-check (e.g. some rows could be being ignored due to parse failures, separator could be incorrect, column types could be inferred wrongly).

```
$ cat /tmp/c | ./chart bar --debug
Lines read  3
Line format inferred    ff
Lines used  3
Float column count  2
String column count 0
Date/Time column count  0
Chart type  bar
Scale type  linear
Separator   [tab]
```

## Details

- `chart` is still experimental.
- it infers STDIN format by analysing line format on each line (doesn't infer separator though; defaults to `\t` and accepts user overrides) and computing the winner format.
- it uses the awesome [ChartJS](http://www.chartjs.org/) library to plot the charts.
- when input data is string-only, `chart` infers a "word frequency pie chart" use case.
- should work on Linux/Mac/Windows thanks to [open-golang](https://github.com/skratchdot/open-golang).

## Known issues

- Javascript's floating point messes up y-axis https://github.com/marianogappa/chart/issues/15

## Contribute

PRs are greatly appreciated and are currently [being merged](https://github.com/marianogappa/chart/pull/3).
If you have a use case that is not supported by `chart`, [I'd love to hear about it](https://github.com/marianogappa/chart/issues), but if it's too complex I'd recommend you to try [gnuplot](http://www.gnuplot.info/).
