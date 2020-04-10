package rtcmlib

import (
	"fmt"
	"strconv"
)

const CRCSEED=0
const CRCPOLY=0x1864CFB

func crc24q_hash (data []byte) uint {

	var crc uint
	crc = 0

	crc24q := crcInit()

	for i := 0; i < len(data); i++ {
		crc1 := crc << 8
		crc2 := crc >> 16
		crc = crc1 ^ crc24q[data[i] ^ byte(crc2)]
	}

	crc = crc & 0x00ffffff

	return crc
}

func generateCRC(string1 string) string {
	var message []byte


	// Converting to byte array
	for i:=0;i<len(string1);i+=8 {
		_mess,_ := strconv.ParseUint(string1[i:i+8], 2, 64)
		message = append(message, byte(_mess))
	}

	// CRC32Q calculating
	//fmt.Println("HASH:")
	//fmt.Printf("%024b\n", crc24q_hash(message))

	return fmt.Sprintf("%024b", crc24q_hash(message))
}

func crcInit() [256]uint {
	var h uint
	var table [256] uint

	table[0] = CRCSEED
	table[1] = CRCPOLY
	h = CRCPOLY

	for i := 2; i < 256; i *= 2 {
		if h = h << 1; (h & 0x1000000) != 0 {
			h ^= CRCPOLY
		}
		for j := 0; j < i; j++ {
			table[i + j] = table[j] ^ h
		}
	}
	return table
}

// UNUSED. For DEBUG
func printCRC32QTable () {
	crc24q := crcInit()
	for i := 0; i < 256; i++ {
		fmt.Printf("0x%08X, ", crc24q[i])
		if (i % 4) == 3 {
			fmt.Printf("\n")
		}
	}
}


