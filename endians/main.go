package main

import (
	"encoding/binary"
	"fmt"
)

func main() {
	var num uint16 = 65534 // 0xFFFE

	// Create a 2-byte buffer
	buf := make([]byte, 2)

	// Big Endian
	binary.BigEndian.PutUint16(buf, num)
	fmt.Printf("Big Endian:    % X\n", buf)

	// Little Endian
	binary.LittleEndian.PutUint16(buf, num)
	fmt.Printf("Little Endian: % X\n", buf)
}
