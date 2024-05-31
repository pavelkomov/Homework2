package main

import "sort"

func Dedupe(slice *[]int) *[]int {
	m, uniq := make(map[int]struct{}), make([]int, 0, len(*slice))
	for _, v := range *slice {
		if _, ok := m[v]; !ok {
			m[v], uniq = struct{}{}, append(uniq, v)
		}
	}
	return &uniq
}

func CountVotes(votes []string) []Candidate {

	map1 := make(map[string]int)
	for _, vote := range votes {
		count, ok := map1[vote]

		if !ok {
			map1[vote] = 1
			continue
		}

		map1[vote] = count + 1

	}

	resultSlice := make([]Candidate, 0, len(map1))
	for key, value := range map1 {
		resultSlice = append(resultSlice, Candidate{key, value})
	}

	sort.Slice(resultSlice, func(i, j int) bool {
		return resultSlice[i].Votes > resultSlice[j].Votes
	})

	return resultSlice

}
