package rtcmlib

import (
	"fmt"
)

func Endode_1012(info Type1012) string{
	return E_DF038(info.SSN) + E_DF039(info.L1ci) +
		E_DF040(info.SFCN) + E_DF041(info.L1PR) +
		E_DF042(info.L1dPR) + E_DF043(info.L1Lt) +
		E_DF044(info.L1MA) + E_DF045(info.L1CNR) +
		E_DF046(info.L2CI) + E_DF047(info.L2PR) +
		E_DF048(info.L2dPR) + E_DF049(info.L2Lt) +
		E_DF050(info.L2CNR)
}

func Endode_1012_Header(info Type1012Header) string{
	return  E_DF003(info.StationId) + E_DF034(info.Epoch) +
	 	E_DF005(info.SFlag) + E_DF035(info.NoSatellites) +
		E_DF036(info.DIndicator) + E_DF037(info.SmoothInterval)
}

//=====================================
func Decode_1012(message string) Type1012 {
	var info Type1012
	info.SSN = D_DF038(message[:6])
	message = message[6:]
	info.L1ci = D_DF039(message[:1])
	message = message[1:]
	info.SFCN = D_DF040(message[:5])
	message = message[5:]
	info.L1PR = D_DF041(message[:25])
	message = message[25:]
	info.L1dPR = D_DF042(message[:20])
	message = message[20:]
	info.L1Lt = D_DF043(message[:7])
	message = message[7:]
	info.L1MA = D_DF044(message[:7])
	message = message[7:]
	info.L1CNR = D_DF045(message[:8])
	message = message[8:]
	info.L2CI = D_DF046(message[:2])
	message = message[2:]
	info.L2PR = D_DF047(message[:14])
	message = message[14:]
	info.L2dPR = D_DF048(message[:20])
	message = message[20:]
	info.L2Lt = D_DF049(message[:7])
	message = message[7:]
	info.L2CNR = D_DF050(message[:8])
	message = message[8:]

	if len(message) >=8 {
		fmt.Println("Bad message filling to end of byte received")
		fmt.Println(message)
	}
	return info
}

func Decode_1012_header(message string) Type1012Header {
	var info Type1012Header
	info.StationId = D_DF003(message[:12])
	message = message[12:]
	info.Epoch = D_DF034(message[:27])
	message = message[27:]
	info.SFlag = D_DF005(message[:1])
	message = message[1:]
	info.NoSatellites = D_DF035(message[:5])
	message = message[5:]
	info.DIndicator = D_DF036(message[:1])
	message = message[1:]
	info.SmoothInterval = D_DF037(message[:3])
	message = message[3:]
	if len(message) > 0 {
		fmt.Println("Bad header received for 1012 message type")
	}
	return info
}


func Decode_CommonHeader(message string) int {
	return int (D_DF002(message[:12]))
}

//=====================================
func Encode_CommonHeader(id int) string {
	return E_DF002(uint64(id))
}



