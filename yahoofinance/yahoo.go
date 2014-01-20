package yahoofinance

import (
	"fmt"
	"github.com/aktau/gofinance/fquery"
	"strconv"
	"strings"
	"time"
)

const (
	HistoricalUrl = "http://ichart.finance.yahoo.com/table.csv"
)

const (
	TypeCsv = iota
	TypeYql
)

type Source struct {
	srcType int
}

func NewCvs() fquery.Source {
	return &Source{srcType: TypeCsv}
}

func NewYql() fquery.Source {
	return &Source{srcType: TypeYql}
}

func (s *Source) Fetch(symbols []string) ([]fquery.Result, error) {
	switch s.srcType {
	case TypeCsv:
		return csvQuotes(symbols)
	case TypeYql:
		return yqlQuotes(symbols)
	}

	return nil, fmt.Errorf("yahoo finance: unknown backend type: %v", s.srcType)
}

func (s *Source) Hist(symbols []string) (map[string]fquery.Hist, error) {
	return yqlHist(symbols, nil, nil)
}

func (s *Source) HistLimit(symbols []string, start time.Time, end time.Time) (map[string]fquery.Hist, error) {
	return yqlHist(symbols, &start, &end)
}

func (s *Source) String() string {
	return "Yahoo Finance (YQL)"
}

/* completes data */
func (q *YqlJsonQuote) Process() {
	/* day and year range */
	pc := strings.Split(q.DaysRangeRaw, " - ")
	q.DayLow, _ = strconv.ParseFloat(pc[0], 64)
	q.DayHigh, _ = strconv.ParseFloat(pc[1], 64)

	pc = strings.Split(q.YearRangeRaw, " - ")
	q.YearLow, _ = strconv.ParseFloat(pc[0], 64)
	q.YearHigh, _ = strconv.ParseFloat(pc[1], 64)
}