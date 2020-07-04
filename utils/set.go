package utils

import "fmt"

//Set implements set like ds with the help of maps
type Set struct {
	KeySet map[int]struct{}
}

// type void struct{}

// func main() {
// 	// m := map[int]bool{}

// 	// m[1] = true
// 	// m[2] = true
// 	// m[3] = true

// 	// fmt.Println(unsafe.Sizeof(m))    //always prints 8
// 	// fmt.Println(unsafe.Sizeof(m[0])) //prints 1

// 	// m = make(map[int]bool)

// 	// m[1] = true
// 	// m[2] = true
// 	// m[3] = true
// 	// m[4] = true
// 	// m[5] = true

// 	// fmt.Println(unsafe.Sizeof(m))    //always prints 8
// 	// fmt.Println(unsafe.Sizeof(m[1])) //prints 1

// 	a := make(map[int]struct{})

// 	b := struct{}{}
// 	a[1] = b

// 	fmt.Println(unsafe.Sizeof(a))    //always prints 8 even with different datatypes of key because it stores one var
// 	fmt.Println(unsafe.Sizeof(a[1])) //prints 0

// 	// c := make([]bool, 1000)
// 	// fmt.Println(unsafe.Sizeof(c)) //always print 24 even with different datatypes of key because slice stores 3 vars

// 	// var e int
// 	// fmt.Println(unsafe.Sizeof(e)) prints 8
// }

//NewSet return set datastructure.
func NewSet() *Set {

	return &Set{
		KeySet: make(map[int]struct{}),
	}

}

//Add adds element to the set
func (set *Set) Add(element int) {

	set.KeySet[element] = struct{}{}
}

//Delete deletes element to the set
func (set *Set) Delete(element int) {
	delete(set.KeySet, element)
}

//Print prints element to the set
func (set *Set) Print() {
	for key := range set.KeySet {
		fmt.Println(key)
	}
}

//Contains check for element existance in the set
func (set *Set) Contains(value int) bool {

	_, ok := set.KeySet[value]

	return ok
}

//Size returns size of the set
func (set *Set) Size() int {
	return len(set.KeySet)
}
