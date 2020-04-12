package main

import (
	"./lib"
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	// Read from FILE
	content, err := ioutil.ReadFile("./rtcm_1004.dump")
	rtcmlib.Chk(err)

	// Convert byte array to string binary
	var messToDecode string
	for i:=0;i<len(content);i++ {
		messToDecode += fmt.Sprintf("%08b", content[i])
	}


	decodedMessage := rtcmlib.DecodeMessage(messToDecode)


	//rtcmlib.Print_json(decodedMessage)

	//======ENCODING ===========================

	/*
	// 1012 message
	encodedMessage := rtcmlib.EncodeMessage(1012, decodedMessage)

	//fmt.Println(encodedMessage)

	// Convert string binary to byte array
	var reply []byte
	for i:=0;i<len(encodedMessage);i+=8 {
		value, _ := strconv.ParseUint(encodedMessage[i:i+8], 2, 64)
		reply = append(reply, byte(value))
	}

	// WRITE TO FILE
	err = ioutil.WriteFile("./myoutput.dump", reply, 0644)
	rtcmlib.Chk(err)
	*/

	// 1004 message
	encodedMessage := rtcmlib.EncodeMessage(1004, decodedMessage)

	//fmt.Println(encodedMessage)

	// Convert string binary to byte array
	var reply []byte
	for i:=0;i<len(encodedMessage);i+=8 {
		value, _ := strconv.ParseUint(encodedMessage[i:i+8], 2, 64)
		reply = append(reply, byte(value))
	}

	// WRITE TO FILE
	err = ioutil.WriteFile("./myoutput.dump", reply, 0644)
	rtcmlib.Chk(err)





}
