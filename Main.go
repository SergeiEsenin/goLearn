package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

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
	if next {
		log.Println("Today no holiday next near holiday", holiday)
		os.Exit(0)
	}

	log.Println("Today", holiday.LocalName)
	os.Exit(0)
}

func (h *Holiday) calcDayOfYear() {
	const layoutISO = "2006-01-02"
	num, err := time.Parse(layoutISO, h.Date)
	handleError(err)
	h.DayOfYear = num.YearDay()
}
