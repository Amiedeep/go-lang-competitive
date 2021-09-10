package main

func merge(input1, input2 []int) []int {
	i, j, k := 0, 0, 0
	var output []int

	for i < len(input1) && j < len(input2) {
		if input1[i] < input2[j] {
			output = append(output, input1[i])
			i++
		} else {
			output = append(output, input2[j])
			j++
		}
		k++
	}

	for f := i; f < len(input1); f++ {
		output = append(output, input1[f])
		k++
	}

	for f := j; f < len(input2); f++ {
		output = append(output, input2[f])
		k++
	}

	return output
}

func mergeSort(input []int, startIndex, endIndex int) []int {
	if startIndex == endIndex {
		return input[startIndex : startIndex+1]
	}

	mid := (startIndex + endIndex) / 2

	// fmt.Printf("%d:%d:%d:%d", len(input), mid, startIndex, endIndex)
	// fmt.Println()

	return merge(mergeSort(input, startIndex, mid), mergeSort(input, mid+1, endIndex))
}

// func main() {

// 	input := []int{12, 11, 13, 5, 6, 7}

// 	output := mergeSort(input, 0, len(input)-1)

// 	utils.PrintSlice(output)
// }
