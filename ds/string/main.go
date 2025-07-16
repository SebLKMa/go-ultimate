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

func main() {
	//str := "123456"
	str := "ahola"
	fmt.Println(str)
	str2 := reverse(str)
	fmt.Println(str2)
}
