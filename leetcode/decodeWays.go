package main

var (
	inputMap = map[string]string{"1": "A", "2": "B", "3": "C", "4": "D", "5": "E", "6": "F", "7": "G", "8": "H", "9": "I", "10": "J", "11": "K", "12": "L", "13": "M", "14": "N", "15": "O", "16": "P", "17": "Q", "18": "R", "19": "S", "20": "T", "21": "U", "22": "V", "23": "W", "24": "X", "25": "Y", "26": "Z"}
)

func decode(s string, index int) int {
	if index >= len(s) {
		return 1
	}

	if _, ok := inputMap[string(s[index])]; !ok {
		return 0
	}

	if index < len(s)-1 {
		if _, ok := inputMap[s[index:index+2]]; ok {
			return decode(s, index+2) + decode(s, index+1)
		}
	}
	return decode(s, index+1)
}

func numDecodings(s string) int {
	return decode(s, 0)
}

// func main() {
// 	s := "06"

// 	fmt.Println(numDecodings(s))
// }
