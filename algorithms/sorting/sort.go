package sorting

// time: O(n^2)
// space: O(1)
// article: https://binarythreads.com/s/
func SelectionSort(array []int, n int) {
	for i := 0; i < n; i++ {
		low := i

		for j := i + 1; j < n; j++ {
			if array[j] < array[low] {
				low = j
			}
		}

		if array[i] > array[low] {
			array[i], array[low] = array[low], array[i]
		}
	}
}
