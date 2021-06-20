package main

func merge(left []int, right []int) []int {
	var temp []int
	leftCount := 0
	rightCount := 0

	for leftCount < len(left) && rightCount < len(right) {
		if right[rightCount] > left[leftCount] {
			temp = append(temp, left[leftCount])
			leftCount++
		} else {
			temp = append(temp, right[rightCount])
			rightCount++
		}
	}

	for len(right) > rightCount {
		temp = append(temp, right[rightCount])
		rightCount++
	}

	for len(left) > leftCount {
		temp = append(temp, left[leftCount])
		leftCount++
	}

	return temp
}

func mergeSort(numbers []int) []int {
	if len(numbers) == 1 {
		return numbers
	} else {
		mid := len(numbers) / 2
		left := mergeSort(numbers[0:mid])
		right := mergeSort(numbers[mid :])

		return merge(left, right)
	}
}

func main() {

}
