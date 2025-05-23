package bookstore

import (
	"fmt"
	"strings"
)

var discounts = map[int]float64{
	1: 0.00,
	2: 0.05,
	3: 0.10,
	4: 0.20,
	5: 0.25,
}

func Cost(books []int) int {
	counts := make([]int, 5)
	for _, b := range books {
		counts[b-1]++
	}
	memo := map[string]int{}
	return minCost(counts, memo)
}

func minCost(counts []int, memo map[string]int) int {
	key := countsKey(counts)
	if val, ok := memo[key]; ok {
		return val
	}

	empty := true
	for _, c := range counts {
		if c > 0 {
			empty = false
			break
		}
	}
	if empty {
		return 0
	}

	minTotal := 1 << 30 // MaxInt

	for mask := 1; mask < 32; mask++ {
		group := []int{}
		valid := true
		for i := 0; i < 5; i++ {
			if (mask>>i)&1 == 1 {
				if counts[i] == 0 {
					valid = false
					break
				}
				group = append(group, i)
			}
		}
		if !valid {
			continue
		}

		newCounts := make([]int, 5)
		copy(newCounts, counts)
		for _, idx := range group {
			newCounts[idx]--
		}

		groupSize := len(group)
		groupCost := int(float64(groupSize*800) * (1 - discounts[groupSize]))
		total := groupCost + minCost(newCounts, memo)
		if total < minTotal {
			minTotal = total
		}
	}

	memo[key] = minTotal
	return minTotal
}

func countsKey(counts []int) string {
	sb := &strings.Builder{}
	for _, c := range counts {
		fmt.Fprintf(sb, "%d,", c)
	}
	return sb.String()
}
