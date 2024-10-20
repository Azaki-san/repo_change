package main

import (
	"fmt"
	"math/rand"
	"time"
)

// MonkeySort (BogoSort) - сортировка обезьяны
// Перемешивает массив до тех пор, пока он не отсортируется, это неэффективный алгоритм
func MonkeySort(arr []int) []int {
	for !isSorted(arr) {
		rand.Shuffle(len(arr), func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
	}
	return arr
}

// QuickSort - быстрая сортировка
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := rand.Int() % len(arr)

	// Поменять местами опорный элемент с последним
	arr[pivot], arr[right] = arr[right], arr[pivot]

	// Разбиение
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Поменять опорный элемент на место
	arr[left], arr[right] = arr[right], arr[left]

	QuickSort(arr[:left])
	QuickSort(arr[left+1:])

	return arr
}

// RandomizedSort - рандомная сортировка (используется как развлекательная сортировка)
// Перемешивает массив случайным образом до тех пор, пока он не отсортируется
func RandomizedSort(arr []int) []int {
	for !isSorted(arr) {
		rand.Shuffle(len(arr), func(i, j int) {
			arr[i], arr[j] = arr[j], arr[i]
		})
	}
	return arr
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i-1] > arr[i] {
			return false
		}
	}
	return true
}

// MergeSort - сортировка слиянием
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := []int{38, 27, 43, 3, 9, 82, 10}

	fmt.Println("Original:", arr)
	fmt.Println("QuickSort:", QuickSort(append([]int(nil), arr...)))
	fmt.Println("RandomizedSort:", RandomizedSort(append([]int(nil), arr...)))
	fmt.Println("MergeSort:", MergeSort(append([]int(nil), arr...)))
}
