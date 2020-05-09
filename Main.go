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
	if err != nil {
		log.Println("Connection failed with ", err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Failed to copy req resp with ", err)
		os.Exit(1)
	}
	err = json.Unmarshal(all, &holidays)
	if err != nil {
		log.Println("Failed to unmarshall response ", err)
		os.Exit(1)
	}
	log.Println(holidays)
	//log.Println(string(all))
	const layoutISO = "2006-01-02"
	date := "2020-05-09"
	t, err := time.Parse(layoutISO, date)
	log.Println(t.Weekday().String())
	handleError(err)
	log.Println(t.YearDay())
}
func (h *Holiday) calcDayOfYear() {
	const layoutISO = "2006-01-02"

}
