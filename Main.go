package main

import (
	"encoding/json"
	"fmt"
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
	now := time.Now()
	fmt.Println(now)
	var holidays []Holiday
	resp, err := http.Get("https://date.nager.at/api/v2/PublicHolidays/2017/UA")
	handleError(err)
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	handleError(err)

	err = json.Unmarshal(all, &holidays)
	handleError(err)
	for i := 0; i < len(holidays); i++ {
		holidays[i].calcDayOfYear()
	}
	log.Println(holidays)
}

func (h *Holiday) calcDayOfYear() {
	const layoutISO = "2006-01-02"
	num, err := time.Parse(layoutISO, h.Date)
	handleError(err)
	h.DayOfYear = num.YearDay()
}
