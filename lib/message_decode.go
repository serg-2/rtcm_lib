package rtcmlib

import (
	"log"
	"strconv"
)

func DecodeMessage (message string) interface{} {

	//var headerDict map[string]interface{}
	var messageDict Type1012Parsed
	var bodyDict Type1012Satellite

	// Const fields
	messageDict.Class = "RTCM3"
	messageDict.Device = "stdin"

	crcReceived := message[len(message)-24:]
	// REMOVE CRC
	message = message[:len(message)-24]
	crcCalculated := generateCRC(message)
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
	messageDict.Length = length
	message = message[10:]

	if uint64(len(message)/8) != length {
		log.Fatalln("Wrong message length received")
	}

	messageType := Decode_CommonHeader(message[:12])
	messageDict.Type = messageType
	message=message[12:]

	switch messageType{
	case 1012:
		//truncate 1012 header
		header := Decode_1012_header(message[:49])
		//headerDict = Parse_2012_header(header)
		Parse_2012_header(&messageDict, header)
		//Print_json(headerDict)
		message=message[49:]

		// Add satellites
		for i := 0; i < int(header.NoSatellites); i++ {
			satellite := Decode_1012(message[i*130 : i*130+130])
			bodyDict = Parse_2012(satellite)

			messageDict.Satellites = append(messageDict.Satellites, bodyDict)

		}


	case 1004:
		/*
			//truncate 1004 header
			header := Decode_1004_header(message[:52])
			message=message[52:]
			for i := 0; i < int(header.NoSatellites); i++ {
				// Change to 1004
				info = Decode_1012(message[i*125 : i*125+125])
				fmt.Printf("%v\n", info)
			}
		*/
	}

	return messageDict
}
