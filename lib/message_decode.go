package rtcmlib

import (
	"log"
	"strconv"
)

func DecodeMessage (message string) interface{} {

	var result interface{}

	crcReceived := message[len(message)-24:]
	// REMOVE CRC
	message = message[:len(message)-24]
	crcCalculated := GenerateCRC(message)
	if crcReceived != crcCalculated {
		log.Fatalln("BAD CRC!!!!!")
	}
	//fmt.Printf("Received   CRC: %v\n", crcReceived)
	//fmt.Printf("Calculated CRC: %v\n", crcCalculated)

	// REMOVE HEADER
	if message[:14] != "11010011" + "000000" {
		log.Fatal("Wrong message header")
	}
	message = message[14:]

	// REMOVE LENGTH
	length, _ := strconv.ParseUint(message[:10], 2, 64)
	message = message[10:]

	if uint64(len(message)/8) != length {
		log.Fatalln("Wrong message length received")
	}

	// Type of message
	messageType := Decode_CommonHeader(message[:12])
	message=message[12:]

	switch messageType{
	case 1012:
		var messageDict Type1012Parsed
		var bodyDict Type1012Satellite

		// Const fields
		messageDict.Class = "RTCM3"
		messageDict.Device = "stdin"

		// Fill already known values
		messageDict.Length = length
		messageDict.Type = messageType

		//truncate 1012 header
		headerDict := Decode_1012_header(message[:49])
		Parse_1012_header(&messageDict, headerDict)
		//Print_json(headerDict)
		message=message[49:]

		// Add satellites
		for i := 0; i < int(headerDict.NoSatellites); i++ {
			satellite := Decode_1012(message[i*130 : i*130+130])
			bodyDict = Parse_1012(satellite)

			messageDict.Satellites = append(messageDict.Satellites, bodyDict)

		}
		result = messageDict

	case 1004:
		var messageDict Type1004Parsed
		var bodyDict Type1004Satellite

		// Const fields
		messageDict.Class = "RTCM3"
		messageDict.Device = "stdin"

		// Fill already known values
		messageDict.Length = length
		messageDict.Type = messageType

		//truncate 1004 header
		headerDict := Decode_1004_header(message[:52])
		Parse_1004_header(&messageDict, headerDict)
		//Print_json(headerDict)

		message=message[52:]

		// Add satellites
		for i := 0; i < int(headerDict.NoSatellites); i++ {

			satellite := Decode_1004(message[i*125 : i*125+125])
			bodyDict = Parse_1004(satellite)

			messageDict.Satellites = append(messageDict.Satellites, bodyDict)

		}

		result = messageDict

	}

	return result
}
