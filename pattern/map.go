package main

import (
	"log"
	"strings"
)

func MapStringToString(strArr []string, fn func(str string) string) []string {
	var newArr = []string{}
	for _, s := range strArr {
		newArr = append(newArr, fn(s))
	}
	return newArr
}

func Reduce(strArr []int, fn func(str int) int) int {
	sum := 0
	for _, a := range strArr {
		sum += fn(a)
	}
	return sum
}

func Filter(arr []int, fn func(n int) bool) []int {
	var newArray = []int{}
	for _, it := range arr {
		if fn(it) {
			newArray = append(newArray, it)
		}
	}
	return newArray
}

func main() {
	arr := []string{"a", "b", "c"}
	toString := MapStringToString(arr, func(str string) string {
		return strings.ToUpper(str)
	})
	log.Println(toString)

	intArr := []int{1, 2, 3}
	reduceResult := Reduce(intArr, func(str int) int {
		return str
	})
	log.Println(reduceResult)

}
