package rtcmlib

import (
	"fmt"
	"log"
)

// Change info to universal
func EncodeMessage (messageType int, messagePassed interface{})string {
	var encoded, body string

	// Generate Header
	encoded = "11010011" + "000000"

	switch messageType{

	case 1012:
		message := messagePassed.(Type1012Parsed)
		// Prepare satellites
		for i:=0; i< len(message.Satellites); i++ {

			satellite := Prepare_1012(message.Satellites[i])
			_body := Endode_1012(satellite)
			if len(_body) != 130 {
				log.Fatalf("Bad info for encoding 1012 type message\n")
			}
			body += _body
		}
		// Add Header 1012
		preparedHeader := Prepare_1012_header(message)
		header_1012 := Encode_CommonHeader(1012) + Endode_1012_Header(preparedHeader)

		body = header_1012 + body

		// Fill Zeroes to end of byte
		for len(body) % 8 != 0 {
			body += "0"
		}
	case 1004:
		message := messagePassed.(Type1004Parsed)
		// Prepare satellites
		for i:=0; i< len(message.Satellites); i++ {

			satellite := Prepare_1004(message.Satellites[i])
			_body := Endode_1004(satellite)
			if len(_body) != 125 {
				log.Fatalf("Bad info for encoding 1004 type message\n")
			}
			body += _body
		}
		// Add Header 1004
		preparedHeader := Prepare_1004_header(message)
		header_1004 := Encode_CommonHeader(1004) + Endode_1004_Header(preparedHeader)

		body = header_1004 + body

		// Fill Zeroes to end of byte
		for len(body) % 8 != 0 {
			body += "0"
		}

	}
	// Check message too big
	if len(body) > 1023*8{
		log.Fatalf("Bad message length!\n")
	}
	length := fmt.Sprintf("%010b", len(body) / 8)

	encoded = encoded + length + body

	crc := GenerateCRC(encoded)

	encoded += crc

	return encoded
}
