package main

import (
	"fmt"

	"github.com/thearyanahmed/codes-of-binarythreads/algorithms/sorting"
	"github.com/thearyanahmed/codes-of-binarythreads/utils"
)

const MAX = 100

func main() {
	array := utils.RandomArray(MAX)

	sorting.SelectionSort(array, len(array))

	fmt.Println(array)
}
