package main

import "sort"

/*
 * @lc app=leetcode id=332 lang=golang
 *
 * [332] Reconstruct Itinerary
 */

// @lc code=start
type solver struct {
	table map[string]map[string]int
	route []string
}

func (s *solver) visit(v string) {
	table := s.table

	keys := []string{}

	for key, count := range table[v] {
		if count > 0 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	for _, key := range keys {
		if table[v][key] > 0 {
			table[v][key]--
			s.visit(key)
		}
	}
	s.route = append(s.route, v)
}

func findItinerary(tickets [][]string) []string {
	table := map[string]map[string]int{}

	for _, ticket := range tickets {
		from, to := ticket[0], ticket[1]

		if _, ok := table[from]; !ok {
			table[from] = map[string]int{}
		}
		table[from][to]++
	}

	s := solver{
		table: table,
	}

	s.visit("JFK")

	route := s.route

	for i := 0; i < len(route)/2; i++ {
		route[i], route[len(route)-i-1] = route[len(route)-i-1], route[i]
	}

	return route
}

// @lc code=end
