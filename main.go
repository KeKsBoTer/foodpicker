package main

import (
	"os"
	"log"
	"hash/fnv"
	"net/http"
	"text/template"
	"time"
)

// getRandomRestaurant returns a arbitrary Restaurant from the list based on the date
// offset can be seen as a kind of seed
// in practice the date is hashed and used as index
func getRandomRestaurant(t time.Time, restaurants []Restaurant, offset int) Restaurant {
	dateString := t.Format("01-02-2006")
	h := fnv.New32a()
	h.Write([]byte(dateString))
	randIndex := (int(h.Sum32()) + offset) % len(restaurants)
	return restaurants[randIndex]
}

// isList check if given item is within the given list
func inList(list []Pick, item Restaurant) bool {
	for _, s := range list {
		if s.Restaurant == item {
			return true
		}
	}
	return false
}

// Pick is a Restaurant pick for a date
type Pick struct {
	Date string
	Restaurant
}

func getWeekday(day time.Weekday) string {
	names := []string{"Sonntag", "Montag", "Dienstag", "Mittwoch", "Donnerstag", "Freitag", "Samstag"}
	return names[int(day)]
}

func generateRestaurantForWeek(date time.Time, Restaurants []Restaurant) []Pick {
	weekday := int(date.Weekday())
	picks := make([]Pick, weekday+1)
	for i := range picks {
		offset := weekday - i
		date := date.AddDate(0, 0, -offset)
		for j := 0; ; j++ {
			p := getRandomRestaurant(date, Restaurants, j)
			if !inList(picks[:i], p) {
				picks[i] = Pick{Date: getWeekday(date.Weekday()), Restaurant: p}
				break
			}
		}
	}
	return picks
}

func main() {
	apiKey := os.Getenv("MAPS_API_KEY")
	restaurants, err := getRestaurants(apiKey)
	if err != nil {
		panic(err)
	}

	if len(restaurants) < 7 {
		panic("need at least 7 different entries")
	}

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}

	randomRestaurant := func(w http.ResponseWriter, r *http.Request) {
		dt := time.Now()
		picks := generateRestaurantForWeek(dt, restaurants)
		if len(picks) > 1 {
			picks = picks[1:]
		}
		err := tmpl.Execute(w, picks)
		if err != nil{
			log.Println(err)
		}
	}

	http.HandleFunc("/", randomRestaurant)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
