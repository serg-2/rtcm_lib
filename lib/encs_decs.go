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

func Endode_1004(info Type1004) string{
	var result string
	result += E_DF009(info.SSN) + E_DF010(info.L1ci) + E_DF011(info.L1PR)
	result += E_DF012(info.L1dPR)
	result += E_DF013(info.L1Lt) + E_DF014(info.L1MA) +	E_DF015(info.L1CNR)
	result += E_DF016(info.L2CI)

	result += E_DF017(info.L2PR)
	result += E_DF018(info.L2dPR)

	result += E_DF019(info.L2Lt) + E_DF020(info.L2CNR)
	return result
}

func Endode_1012_Header(info Type1012Header) string{
	return  E_DF003(info.StationId) + E_DF034(info.Epoch) +
		E_DF005(info.SmoothFlag) + E_DF035(info.NoSatellites) +
		E_DF036(info.DIndicator) + E_DF037(info.SmoothInterval)
}

func Endode_1004_Header(info Type1004Header) string{
	return  E_DF003(info.StationId) + E_DF004(info.Epoch) +
		E_DF005(info.SmoothFlag) + E_DF006(info.NoSatellites) +
		E_DF007(info.DIndicator) + E_DF008(info.SmoothInterval)
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

func Decode_1004(message string) Type1004 {
	var info Type1004
	info.SSN = D_DF009(message[:6])
	message = message[6:]
	info.L1ci = D_DF010(message[:1])
	message = message[1:]
	info.L1PR = D_DF011(message[:24])
	message = message[24:]
	info.L1dPR = D_DF012(message[:20])
	message = message[20:]
	info.L1Lt = D_DF013(message[:7])
	message = message[7:]

	info.L1MA = D_DF014(message[:8])
	message = message[8:]

	info.L1CNR = D_DF015(message[:8])
	message = message[8:]
	info.L2CI = D_DF016(message[:2])
	message = message[2:]
	info.L2PR = D_DF017(message[:14])
	message = message[14:]
	info.L2dPR = D_DF018(message[:20])
	message = message[20:]
	info.L2Lt = D_DF019(message[:7])
	message = message[7:]
	info.L2CNR = D_DF020(message[:8])
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
	info.SmoothFlag = D_DF005(message[:1])
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

func Decode_1004_header(message string) Type1004Header {
	var info Type1004Header
	info.StationId = D_DF003(message[:12])
	message = message[12:]
	info.Epoch = D_DF004(message[:30])
	message = message[30:]
	info.SmoothFlag = D_DF005(message[:1])
	message = message[1:]
	info.NoSatellites = D_DF006(message[:5])
	message = message[5:]
	info.DIndicator = D_DF007(message[:1])
	message = message[1:]
	info.SmoothInterval = D_DF008(message[:3])
	message = message[3:]
	if len(message) > 0 {
		fmt.Println("Bad header received for 1004 message type")
	}
	return info
}

//=== COMMON===================================

func Decode_CommonHeader(message string) int {
	return int (D_DF002(message[:12]))
}

//=====================================
func Encode_CommonHeader(id int) string {
	return E_DF002(uint64(id))
}



