package jan2021

// func main() {
// 	var t, n int
// 	var input string

// 	fmt.Scan(&t)

// 	o := make([]string, t)
// 	chars := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p'}

// 	for i := 0; i < t; i++ {
// 		fmt.Scanln(&n)
// 		fmt.Scanln(&input)

// 		o[i] = decode(input, chars)
// 	}

// 	for i := 0; i < len(o); i++ {
// 		fmt.Println(o[i])
// 	}
// }

func decode(input string, chars []byte) string {

	var result string

	for i := 0; i < len(input); {
		decodedChars := make([]byte, len(chars))
		copy(decodedChars, chars)
		for j := 0; j < 4; j++ {
			decodedChars = decodeBit(input[i+j], decodedChars)
		}
		i += 4
		result += string(decodedChars)
	}

	return result
}

func decodeBit(bit byte, chars []byte) []byte {

	if bit == '0' {
		return chars[:len(chars)/2]
	} else {
		return chars[len(chars)/2:]
	}
}
