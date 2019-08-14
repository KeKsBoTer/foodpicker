package main

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func hasDuplicate(list []string) bool {
	sort.Strings(list)
	for i := 0; i < len(list)-1; i++ {
		if list[i] == list[i+1] {
			return true
		}
	}
	return false
}

func Test_generateFoodForWeek(t *testing.T) {
	type args struct {
		date  time.Time
		foods []string
	}
	// generating 1000 sample dates and check if duplicates are generated
	tests := make([]time.Time, 1000)
	for i := range tests {
		tests[i] = time.Now().AddDate(0, 0, rand.Intn(int(time.Now().Unix()/(60*60*24))))
	}
	food := []string{"1", "2", "3", "4", "5", "6", "7"}
	t.Parallel()
	for i, tt := range tests {
		t.Run("test "+string(i), func(t *testing.T) {
			if got := generateFoodForWeek(tt, food); hasDuplicate(got) {
				t.Errorf("generateFoodForWeek(), found duplicate: %v", got)
			}
		})
	}
}
