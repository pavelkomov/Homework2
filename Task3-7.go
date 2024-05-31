package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Grade int    `json:"grade"`
}

type Object struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Result struct {
	ObjectID  int `json:"object_id"`
	StudentID int `json:"student_id"`
	Result    int `json:"result"`
}

type Data struct {
	Students []Student `json:"students"`
	Objects  []Object  `json:"objects"`
	Results  []Result  `json:"results"`
}

func ControlWork() {

	file, err := os.ReadFile("dz3.json")

	if err != nil {
		log.Fatal(err)
	}

	var jsonData Data
	err = json.Unmarshal(file, &jsonData)
	if err != nil {
		log.Fatal(err)
	}

	makeMap(jsonData)
}
func makeMap(data Data) {

	stMap := make(map[int]Student)
	objMap := make(map[int]Object)

	for _, st := range data.Students {
		stMap[st.ID] = st
	}

	for _, obj := range data.Objects {
		objMap[obj.ID] = obj
	}
	printResult(stMap, objMap, data.Results) //3
	printMean(stMap, objMap, data.Results)   //4

	printResultCache(data)            //5
	printMeanFunc(data, stMap)        //6
	findExcellentStudent(data, stMap) //7

}

// 3
func printResult(stMap map[int]Student, objMap map[int]Object, results []Result) {

	fmt.Println("  //3")
	fmt.Printf("___________________________________________\n")
	fmt.Printf("|%-14s|%-6s|%-10s|%-8s| \n", "Student name", "Grade", "Object", "Result")
	fmt.Printf("___________________________________________\n")

	for _, res := range results {

		fmt.Printf("|%-14s|%-6d|%-10s|%-8d| \n", stMap[res.StudentID].Name, stMap[res.StudentID].Grade, objMap[res.ObjectID].Name, res.Result)

	}

	fmt.Printf("___________________________________________\n")
}

// 4
func printMean(stMap map[int]Student, objMap map[int]Object, results []Result) {

	var count9, count10, count11 int
	var sum9, sum10, sum11 int

	for _, res := range results {

		if res.ObjectID != 1 {
			continue
		}

		switch stMap[res.StudentID].Grade {
		case 9:
			sum9 = sum9 + res.Result
			count9 += 1
		case 10:
			sum10 = sum10 + res.Result
			count10 += 1
		case 11:
			sum11 = sum11 + res.Result
			count11 += 1
		}

	}
	fmt.Println("  //4")
	fmt.Printf("___________________\n")
	fmt.Printf("|%-10s|%-6s|\n", "Math", "Mean")
	fmt.Printf("___________________\n")

	fmt.Printf("|%-10s|%-6.2f|\n", "9 grade", float32(sum9)/float32(count9))
	fmt.Printf("|%-10s|%-6.2f|\n", "10 grade", float32(sum10)/float32(count10))
	fmt.Printf("|%-10s|%-6.2f|\n", "11 grade", float32(sum11)/float32(count11))
	fmt.Printf("|%-10s|%-6.2f|\n", "mean", (float32(sum9)+float32(sum10)+float32(sum11))/(float32(count9)+float32(count10)+float32(count11)))
	fmt.Printf("___________________\n")

	for _, res := range results {

		if res.ObjectID != 2 {
			continue
		}

		switch stMap[res.StudentID].Grade {
		case 9:
			sum9 = sum9 + res.Result
			count9 += 1
		case 10:
			sum10 = sum10 + res.Result
			count10 += 1
		case 11:
			sum11 = sum11 + res.Result
			count11 += 1
		}

	}

	fmt.Printf("___________________\n")
	fmt.Printf("|%-10s|%-6s|\n", "Biology", "Mean")
	fmt.Printf("___________________\n")

	fmt.Printf("|%-10s|%-6.2f|\n", "9 grade", float32(sum9)/float32(count9))
	fmt.Printf("|%-10s|%-6.2f|\n", "10 grade", float32(sum10)/float32(count10))
	fmt.Printf("|%-10s|%-6.2f|\n", "11 grade", float32(sum11)/float32(count11))
	fmt.Printf("|%-10s|%-6.2f|\n", "mean", (float32(sum9)+float32(sum10)+float32(sum11))/(float32(count9)+float32(count10)+float32(count11)))
	fmt.Printf("___________________\n")

	for _, res := range results {

		if res.ObjectID != 3 {
			continue
		}

		switch stMap[res.StudentID].Grade {
		case 9:
			sum9 = sum9 + res.Result
			count9 += 1
		case 10:
			sum10 = sum10 + res.Result
			count10 += 1
		case 11:
			sum11 = sum11 + res.Result
			count11 += 1
		}

	}

	fmt.Printf("___________________\n")
	fmt.Printf("|%-10s|%-6s|\n", "Geology", "Mean")
	fmt.Printf("___________________\n")

	fmt.Printf("|%-10s|%-6.2f|\n", "9 grade", float32(sum9)/float32(count9))
	fmt.Printf("|%-10s|%-6.2f|\n", "10 grade", float32(sum10)/float32(count10))
	fmt.Printf("|%-10s|%-6.2f|\n", "11 grade", float32(sum11)/float32(count11))
	fmt.Printf("|%-10s|%-6.2f|\n", "mean", (float32(sum9)+float32(sum10)+float32(sum11))/(float32(count9)+float32(count10)+float32(count11)))
	fmt.Printf("___________________\n")

}

//5

type Cache[K int, V any] struct {
	m map[K]V
}

func (c *Cache[K, V]) Init() {
	c.m = make(map[K]V)
}

func (c *Cache[K, V]) Get(key K) (V, bool) {
	k, ok := c.m[key]
	return k, ok
}

func (c *Cache[K, V]) Set(key K, value V) {
	c.m[key] = value
}

func printResultCache(data Data) {

	var students Cache[int, Student]
	var objects Cache[int, Object]

	students.Init()
	objects.Init()

	for _, st := range data.Students {
		students.Set(st.ID, st)
	}

	for _, ob := range data.Objects {
		objects.Set(ob.ID, ob)
	}

	fmt.Println("  //5")
	fmt.Printf("___________________________________________\n")
	fmt.Printf("|%-14s|%-6s|%-10s|%-8s| \n", "Student name", "Grade", "Object", "Result")
	fmt.Printf("___________________________________________\n")

	for _, res := range data.Results {

		st, ok1 := students.Get(res.StudentID)
		obj, ok2 := objects.Get(res.ObjectID)

		if !(ok1 && ok2) {
			continue
		}
		fmt.Printf("|%-14s|%-6d|%-10s|%-8d| \n", st.Name, st.Grade, obj.Name, res.Result)

	}

	fmt.Printf("___________________________________________\n")
}

// 6
func Filter[T any](s []T, f func(T) bool) []T {
	var r []T
	for _, v := range s {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
func Map[T1, T2 any](s []T1, f func(T1) T2) []T2 {
	r := make([]T2, len(s))
	for i, v := range s {
		r[i] = f(v)
	}
	return r
}
func Reduce[T1, T2 any](s []T1, init T2, f func(T1, T2) T2) T2 {
	r := init
	for _, v := range s {
		r = f(v, r)
	}
	return r
}
func findMean(r []Result) float32 {

	marks := Map(r, func(a Result) int {
		return a.Result
	})

	sum := Reduce(marks, float32(0), func(a int, b float32) float32 {
		return float32(a) + b
	})
	return float32(sum) / float32(len(r))
}

func printMeanFunc(data Data, m map[int]Student) {

	filtredMath := Filter(data.Results, func(result Result) bool {
		return result.ObjectID == 1
	})

	filtredBio := Filter(data.Results, func(result Result) bool {
		return result.ObjectID == 2
	})

	filtredGeo := Filter(data.Results, func(result Result) bool {
		return result.ObjectID == 3
	})

	//math
	filtredMathGrade9 := Filter(filtredMath, func(result Result) bool {

		return m[result.StudentID].Grade == 9

	})

	filtredMathGrade10 := Filter(filtredMath, func(result Result) bool {

		return m[result.StudentID].Grade == 10

	})
	filtredMathGrade11 := Filter(filtredMath, func(result Result) bool {

		return m[result.StudentID].Grade == 11

	})
	fmt.Println("  //6")
	fmt.Printf("___________________\n")
	fmt.Printf("|%-10s|%-6s|\n", "Math", "Mean")
	fmt.Printf("___________________\n")

	fmt.Printf("|%-10s|%-6.2f|\n", "9 grade", findMean(filtredMathGrade9))
	fmt.Printf("|%-10s|%-6.2f|\n", "10 grade", findMean(filtredMathGrade10))
	fmt.Printf("|%-10s|%-6.2f|\n", "11 grade", findMean(filtredMathGrade11))
	fmt.Printf("|%-10s|%-6.2f|\n", "mean", findMean(filtredMath))
	fmt.Printf("___________________\n")

	//Biology
	filtredBioGrade9 := Filter(filtredBio, func(result Result) bool {

		return m[result.StudentID].Grade == 9

	})

	filtredBioGrade10 := Filter(filtredBio, func(result Result) bool {

		return m[result.StudentID].Grade == 10

	})
	filtredBioGrade11 := Filter(filtredBio, func(result Result) bool {

		return m[result.StudentID].Grade == 11

	})

	fmt.Printf("___________________\n")
	fmt.Printf("|%-10s|%-6s|\n", "Math", "Mean")
	fmt.Printf("___________________\n")

	fmt.Printf("|%-10s|%-6.2f|\n", "9 grade", findMean(filtredBioGrade9))
	fmt.Printf("|%-10s|%-6.2f|\n", "10 grade", findMean(filtredBioGrade10))
	fmt.Printf("|%-10s|%-6.2f|\n", "11 grade", findMean(filtredBioGrade11))
	fmt.Printf("|%-10s|%-6.2f|\n", "mean", findMean(filtredBio))
	fmt.Printf("___________________\n")

	filtredGeoGrade9 := Filter(filtredGeo, func(result Result) bool {

		return m[result.StudentID].Grade == 9

	})

	filtredGeoGrade10 := Filter(filtredGeo, func(result Result) bool {

		return m[result.StudentID].Grade == 10

	})
	filtredGeoGrade11 := Filter(filtredGeo, func(result Result) bool {

		return m[result.StudentID].Grade == 11

	})

	fmt.Printf("___________________\n")
	fmt.Printf("|%-10s|%-6s|\n", "Math", "Mean")
	fmt.Printf("___________________\n")

	fmt.Printf("|%-10s|%-6.2f|\n", "9 grade", findMean(filtredGeoGrade9))
	fmt.Printf("|%-10s|%-6.2f|\n", "10 grade", findMean(filtredGeoGrade10))
	fmt.Printf("|%-10s|%-6.2f|\n", "11 grade", findMean(filtredGeoGrade11))
	fmt.Printf("|%-10s|%-6.2f|\n", "mean", findMean(filtredGeo))
	fmt.Printf("___________________\n")

}

/* 7. Выведите в консоль круглых отличников из числа студентов, используя функцию Filter.
Вывод реализуйте как в задаче #3.*/

func findExcellentStudent(data Data, stMap map[int]Student) {

	filtredExcellent := Filter(data.Results, func(result Result) bool {
		return result.Result == 5
	})

	m := make(map[Student]int)

	for _, r := range filtredExcellent {

		st, ok := stMap[r.StudentID]
		if !ok {
			continue
		}
		m[st] += 1

	}

	fmt.Println("  //7")
	for k, v := range stMap {
		if k == 3 {
			fmt.Printf("%s ", v.Name)
		}
	}

	fmt.Println()

}
