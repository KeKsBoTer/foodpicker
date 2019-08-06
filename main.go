package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"hash/fnv"
	"time"
	"fmt"
)

func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}


func main() {
	b, err := ioutil.ReadFile("food.txt")
    if err != nil {
        panic(err)
    }

	restaurants := strings.Split(string(b),"\n")
	fmt.Println(len(restaurants))

	randomFood := func(w http.ResponseWriter, r *http.Request) {
		dt := time.Now()
		today := dt.Format("01-02-2006")
		randIndex := hash(today)%len(restaurants)
		fmt.Println(randIndex,restaurants[randIndex],restaurants)
		w.Write([]byte("Today we eat at: "+restaurants[randIndex]))
	}

	http.HandleFunc("/", randomFood)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}