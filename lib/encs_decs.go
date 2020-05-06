package rtcmlib

import (
	"fmt"
)

func Endode_1012(info Type1012) string {
	var result string
	result = E_DF038(info.SSN) + E_DF039(info.L1ci)
	result += E_DF040(info.SFCN) + E_DF041(info.L1PR)
	result += E_DF042(info.L1dPR)
	result += E_DF043(info.L1Lt)
	result += E_DF044(info.L1MA) + E_DF045(info.L1CNR)
	result += E_DF046(info.L2CI) + E_DF047(info.L2PR)
	result += E_DF048(info.L2dPR) + E_DF049(info.L2Lt)
	result += E_DF050(info.L2CNR)

	return result
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

func Decode_1087_header(info *Type1087, message string) uint64 {

	info.StationId = D_DF003(message[:12])
	message = message[12:]

	info.Day = D_DF416(message[:3])
	message = message[3:]

	info.Epoch = D_DF034(message[:27])
	message = message[27:]

	info.MMB = D_DF393(message[:1])
	message = message[1:]

	info.IODS = D_DF409(message[:3])
	message = message[3:]

	//Reserved
	message = message[7:]

	info.CSI = D_DF411(message[:2])
	message = message[2:]

	info.ECI = D_DF412(message[:2])
	message = message[2:]

	info.SIndi = D_DF417(message[:1])
	message = message[1:]

	info.SInter = D_DF418(message[:3])
	message = message[3:]

	//info.SatMask = D_DF394(message[:64])
	//info.SatNumber = OnesCount64(info.SatMask)
	info.SatMask = message[:64]
	info.SatNumber = OnesCount(info.SatMask)
	message = message[64:]

	//info.SignalMask = D_DF395(message[:32])
	//info.SignalNumber = OnesCount32(uint32(info.SignalMask))
	info.SignalMask = message[:32]
	info.SignalNumber = OnesCount(info.SignalMask)

	message = message[32:]

	// GNSS Cell Mask  DF396 bit(X) X (X≤64)
	// info.SatNumber*info.SignalNumber
	//info.SatSignalTable = D_DF396(message[:info.SatNumber*info.SignalNumber])
	info.SatSignalTable = message[:info.SatNumber*info.SignalNumber]

	return 157 + uint64(info.SatNumber*info.SignalNumber)
}

func Decode_1087_satellite(message string, currentNumber int, quantSatellites int) Type1087Satellite {
	var info Type1087Satellite
	var shift int

	// The number of integer milliseconds in GNSS Satellite rough ranges
	info.RoughRangeInt = D_DF397(message[currentNumber*8:currentNumber*8+8])
	shift += 8*quantSatellites

	//Extended Satellite Information
	//GLONASS Satellite Frequency Channel Number (DF419) is used as extended satellite information
	// in the header of MSM7
	info.Info = D_DF419(message[shift+currentNumber*4:shift+currentNumber*4+4])
	shift += 4*quantSatellites

	// GNSS Satellite rough ranges modulo 1 millisecond
	info.RoughRangeRemainder = D_DF398(message[shift+currentNumber*10:shift+currentNumber*10+10])
	shift += 10*quantSatellites

	//GNSS Satellite rough PhaseRangeRates
	info.RatePhaseRangeInt = D_DF399(message[shift+currentNumber*14:shift+currentNumber*14+14])

	return info
}

func Decode_1087_signal(message string, currentNumber int, quantSignals int) Type1087Signal {
	var info Type1087Signal
	var shift int

	// GNSS signal fine Pseudoranges with extended resolution
	info.PseudoRangeCorrection = D_DF405(message[currentNumber*20:currentNumber*20+20])
	shift += 20*quantSignals

	//GNSS signal fine PhaseRange data with extended resolution
	info.PhaseRangeCorrection = D_DF406(message[shift+currentNumber*24:shift+currentNumber*24+24])
	shift += 24*quantSignals

	// GNSS PhaseRange Lock Time Indicator with extended range and resolution.
	info.PhaseRangeTI = D_DF407(message[shift+currentNumber*10:shift+currentNumber*10+10])
	shift += 10*quantSignals

	//Half-cycle ambiguity indicator
	info.AI = D_DF420(message[shift+currentNumber*1:shift+currentNumber*1+1])
	shift += 1*quantSignals

	//GNSS signal CNRs with extended resolution
	info.CNR = D_DF408(message[shift+currentNumber*10:shift+currentNumber*10+10])
	shift += 10*quantSignals

	//GNSS signal fine PhaseRangeRates
	info.RatePhaseRangeRemainder = D_DF404(message[shift+currentNumber*15:shift+currentNumber*15+15])

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



