package main

import "fmt"

func swap(src string, idxFrom int, idxTo int) (dst string) {
	src[idxFrom] = src[idxTo]
	return
}

func main() {
	str := "hola"

	fmt.Println(string(str[1]))
	fmt.Println(string(str[2]))
	fmt.Printf("%c\n", str[1])
	fmt.Printf("%c\n", str[2])
}
