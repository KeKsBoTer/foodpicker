package main

import (
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
	"time"
)

// getRandomFood returns a arbitrary Food from the list based on the date
// offset can be seen as a kind of seed
// in practice the date is hashed and used as index
func getRandomFood(t time.Time, foods []string, offset int) string {
	dateString := t.Format("01-02-2006")
	h := fnv.New32a()
	h.Write([]byte(dateString))
	randIndex := (int(h.Sum32()) + offset) % len(foods)
	return foods[randIndex]
}

// isList check if given item is within the given list
func inList(list []Pick, item string) bool {
	for _, s := range list {
		if s.Food == item {
			return true
		}
	}
	return false
}

// Pick is a food pick for a date
type Pick struct {
	Date string
	Food string
}

func getWeekday(day time.Weekday) string {
	names := []string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"}
	return names[int(day)]
}

func generateFoodForWeek(date time.Time, foods []string) []Pick {
	weekday := int(date.Weekday())
	picks := make([]Pick, weekday+1)
	for i := range picks {
		offset := weekday - i
		date := date.AddDate(0, 0, -offset)
		for j := 0; ; j++ {
			p := getRandomFood(date, foods, j)
			if !inList(picks[:i], p) {
				picks[i] = Pick{Date: getWeekday(date.Weekday()), Food: p}
				break
			}
		}
	}
	return picks
}

func main() {
	b, err := ioutil.ReadFile("food.txt")
	if err != nil {
		panic(err)
	}

	restaurants := strings.Split(string(b), "\n")

	if len(restaurants) < 7 {
		panic("need at least 7 different entries")
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	randomFood := func(w http.ResponseWriter, r *http.Request) {
		dt := time.Now()
		picks := generateFoodForWeek(dt, restaurants)
		if len(picks) > 1 {
			picks = picks[1:]
		}
		tmpl.Execute(w, picks)
	}

	http.HandleFunc("/", randomFood)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
