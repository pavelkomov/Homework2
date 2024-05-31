package main

import (
	"fmt"
	"sort"
)

type Candidates interface {
}

func main() {
	//1
	fmt.Println("  //1")
	sl1 := []int{1, 2, 2, 2, 2, 2}
	sl2 := []int{0, 1, 2, 2}

	res, ok := intersectionSlices(&sl1, &sl2)
	if ok {
		fmt.Println(*res)
	}
	//2
	fmt.Println("  //2")
	arr := [...]string{"Ann", "Ann", "Ann", "Kate", "Kate", "John", "Kate", "AAA", "AAA", "AAA", "AAA"}
	votes := arr[:]
	fmt.Println(CountVotes(votes))

	//3-7
	ControlWork()

	//8
	fmt.Println("  //8")
	arr1 := []int{1, 2, 32, 64, 256, 16}
	arr2 := []int{1, 2, 64, 16, 256, 1, 32}
	fmt.Println(isEqualarrays(arr1, arr2))

	//9
	fmt.Println("  //9")
	vector1 := Vector[int]{1, 2, 3, 4}
	vector2 := Vector[int]{1, 2, 3}
	fmt.Print("vector1: ")
	fmt.Println(vector1)
	fmt.Print("vector2: ")
	fmt.Println(vector2)
	fmt.Print("is equal? ")
	fmt.Println(isEqual(vector1, vector2))

	fmt.Print("sum = ")
	fmt.Println(sum(vector1))
	i, ok := isElement(3, vector1)
	if ok {
		fmt.Printf("3 is element? %t index = %d\n", ok, i)
	}
	v, _ := deleteByIndex(1, vector1)
	fmt.Println(v)

	v1, _ := deleteByValue(3, vector2)
	fmt.Println(v1)

}

func intersectionSlices(in ...*[]int) (*[]int, bool) {

	el := make(map[int]int)

	for _, slice := range in {

		if *slice == nil {
			return nil, false
		}
		if len(*slice) == 0 {
			result := make([]int, 0)
			return &result, true
		}

		dedupedSlice := Dedupe(slice)

		for _, num := range *dedupedSlice {

			count, ok := el[num]

			if !ok {
				el[num] = 1
				continue
			}

			el[num] = count + 1

		}
	}

	result := make([]int, 0, len(el))
	for num, k := range el {
		if k == len(in) {
			result = append(result, num)
		}

	}
	sort.Ints(result)

	return &result, true
}

type Candidate struct {
	Name  string
	Votes int
}
