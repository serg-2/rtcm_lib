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
	case 1077:
		//log.Println("Received message type: 1077")

	case 1087:
		var messageDict Type1087
		var satelliteDict Type1087Satellite
		var signalDict Type1087Signal
		var signalTable []int

		//log.Printf("1087 type received")
		//log.Printf("Length: %v\n", length)
		// Header part
		lengthToCut := Decode_1087_header(&messageDict, message)
		message=message[lengthToCut:]
		//Print_json(messageDict)

		// Satellite part (36bit * Number of Satellites)
		for i:=0; i < int(messageDict.SatNumber); i++ {
			satelliteDict = Decode_1087_satellite(message, i, int(messageDict.SatNumber))
			//Print_json(satelliteDict)


			// Editing satellites
			satFound := 0
			currentSatNumber := 0
			for j:= 0; j < 64; j++ {
				if string(messageDict.SatMask[j]) == "0" {continue}
				currentSatNumber = j + 1
				satFound++
				if satFound == i + 1 {
					break
				}
			}
			satelliteDict.SatelliteNumber = currentSatNumber

			messageDict.Satellites = append(messageDict.Satellites, satelliteDict)

		}
		//Print_json(messageDict)
		message=message[36*messageDict.SatNumber:]

		/*
		// Output Signal to Satellite table
		for k:= 0; k < messageDict.SignalNumber; k++ {
			fmt.Printf("For signal number %v Satellite table: %v\n", k+1, messageDict.SatSignalTable[messageDict.SatNumber*k:messageDict.SatNumber*k+messageDict.SatNumber])
		}
		*/

		// Generate signal table
		for k:=0; k < len(messageDict.SignalMask); k++ {
			if string(messageDict.SignalMask[k]) == "1" {
				signalTable = append(signalTable, k+1)
			}
		}
		//log.Printf("Signal Mask: %v\n", messageDict.SignalMask)
		//log.Printf("Signal table: %v\n", signalTable)

		// Signal part ()
		for i:=0; i < int(messageDict.SignalNumber*messageDict.SatNumber); i++ {
			// Skip empty signal
			if string(messageDict.SatSignalTable[i]) == "0" {continue}
			signalIndex := i/messageDict.SatNumber
			satelliteIndex := i-(i/messageDict.SatNumber)*messageDict.SatNumber
			//fmt.Printf("Signal index: %v Satellite index: %v\n",signalIndex, satelliteIndex)

			signalDict = Decode_1087_signal(message, i, int(messageDict.SatNumber*messageDict.SignalNumber))
			// Update signal with signal number
			signalDict.SignalNumber = signalTable[signalIndex]

			// Put signal to satellite
			messageDict.Satellites[satelliteIndex].Signals = append(messageDict.Satellites[satelliteIndex].Signals, signalDict)

		}
		message=message[80*messageDict.SignalNumber*messageDict.SatNumber:]

		// Final check
		if len(message) >=8 {
			log.Println("Bad message filling to end of byte received. Message Type 1087")
			log.Println(message)
		}

		result2 := Parse_1087(messageDict)

		// Finish
		result = result2
	default:
		log.Printf("Received message type: %v\n", messageType)

	}

	return result
}
