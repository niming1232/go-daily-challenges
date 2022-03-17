package sreprep

func MergeSort(array []int) []int {
	// Base case
	if len(array) == 1 {
		return array
	}

	firstHalf := MergeSort(array[:len(array)/2])
	secondHalf := MergeSort(array[len(array)/2:])

	i, j := 0, 0
	mergedArray := make([]int, len(array))
	for i != len(firstHalf) && j != len(secondHalf) {
		if firstHalf[i] <= secondHalf[j] {
			mergedArray = append(mergedArray, firstHalf[i])
			i++
		} else {
			mergedArray = append(mergedArray, firstHalf[j])
			j++
		}
	}
	if i < len(firstHalf) {
		mergedArray = append(mergedArray, firstHalf[i:]...)
	} else if j < len(secondHalf) {
		mergedArray = append(mergedArray, secondHalf[j:]...)
	}
	return mergedArray
}
