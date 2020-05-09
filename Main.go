package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

const layoutISO = "2006-01-02"

type Holiday struct {
	Date      string
	LocalName string
	Name      string
	Fixed     bool
	DayOfYear int
	Type      string
}

func handleError(err error) {
	if err != nil {
		log.Println("Smth went wrong ", err)
		os.Exit(1)
	}
}
func main() {
	now := time.Now().YearDay()
	var holidays []Holiday
	resp, err := http.Get("https://date.nager.at/api/v2/PublicHolidays/2020/UA")
	handleError(err)
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	err = json.Unmarshal(all, &holidays)
	handleError(err)
	var holiday Holiday
	var next bool
	for i := 0; i < len(holidays); i++ {
		holidays[i].calcDayOfYear()
		if now == holidays[i].DayOfYear {
			holiday = holidays[i]
			break
		} else if now < holiday.DayOfYear {
			holiday = holidays[i]
			next = true
			break
		}
	}

	printResult(next, holiday)
	fmt.Scanf("h")
}
func printResult(next bool, holiday Holiday) {
	var b bytes.Buffer
	b.WriteString("Today ")
	if next {
		b.WriteString("no holiday next near holiday ")
	}

	date, err := time.Parse(layoutISO, holiday.Date)
	handleError(err)
	b.WriteString(holiday.Name)
	b.WriteString(" ")
	b.WriteString(date.Month().String())
	b.WriteString(" ")
	b.WriteString(strconv.Itoa(date.Day()))
	var range1 time.Time
	var range2 time.Time
	switch date.Weekday() {
	case time.Friday:
		range1 = date
		range2 = date.Add(24 * 2 * time.Hour)
	case time.Saturday:
		range1 = date
		range2 = date.Add(24 * 2 * time.Hour)
	case time.Sunday:
		range1 = date.Add(-24 * time.Hour)
		range2 = date.Add(24 * 2 * time.Hour)
	case time.Monday:
		range1 = date.Add(-24 * 3 * time.Hour)
		range2 = date
	}
	if &range1 != nil && &range2 != nil {
		b.WriteString(" holidays will last 3 days: ")
		b.WriteString(range1.Month().String())
		b.WriteString(" ")
		b.WriteString(strconv.Itoa(range1.Day()))
		b.WriteString(" - ")
		b.WriteString(range2.Month().String())
		b.WriteString(" ")
		b.WriteString(strconv.Itoa(range2.Day()))
	}
	log.Println(b.String())
}

func (h *Holiday) calcDayOfYear() {
	date, err := time.Parse(layoutISO, h.Date)
	handleError(err)
	h.DayOfYear = date.YearDay()
}
