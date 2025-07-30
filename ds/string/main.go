package main

import (
	"fmt"
)

// swaps data[idxFrom] with data[idxTo]
/* Example:
str := "hola"
bstr := []byte(str)
swap(&bstr, 2, 1)
str2 := string(bstr)
*/
func swap(data *[]byte, idxFrom int, idxTo int) {
	keep := (*data)[idxFrom]
	(*data)[idxFrom] = (*data)[idxTo]
	(*data)[idxTo] = keep
}

func reverse(src string) (dst string) {
	bsrc := []byte(src)
	max := len(bsrc)
	mid := max / 2
	for i := 0; i < mid; i++ {
		swap(&bsrc, i, max-i-1)
	}
	return string(bsrc)
}

func removeDups(src string) string {
	// Using rune instead of byte for unicode string
	rsrc := []rune(src)
	max := len(rsrc)

	stopper := 1 // stopper starts at second position
	for currentIndex := 0; currentIndex < max; currentIndex++ {
		var runnerIndex int // runner always compares from beginning
		for runnerIndex = 0; runnerIndex < stopper; runnerIndex++ {
			if rsrc[runnerIndex] == rsrc[currentIndex] {
				break // duplicate found
			}
		}

		if runnerIndex == stopper {
			// runner ran until stop point means no duplicate
			rsrc[stopper] = rsrc[currentIndex]
			stopper++
		}
	}

	// If duplicates have been removed, remove the remaining "residue"
	/* slower option
	fmt.Printf("%d\n", stopper)
	for runnerIndex := stopper; runnerIndex < max; runnerIndex++ {
		(*rsrc)[runnerIndex] = 0
	}
	fmt.Printf("%d\n", len(*rsrc))
	*/
	if stopper != max {
		return string(rsrc)[:stopper]
	}

	return string(rsrc)
}

func main() {
	//str := "123456"
	str := "ahola"
	fmt.Println(str)
	str2 := reverse(str)
	fmt.Println(str2)

	str3 := "1223345566"
	fmt.Printf("%s len:%d\n", str3, len(str3))
	str4 := removeDups(str3)
	fmt.Printf("%s len:%d\n", str4, len(str4))
	str5 := removeDups(str4)
	fmt.Printf("%s len:%d verified\n", str5, len(str5))
}
