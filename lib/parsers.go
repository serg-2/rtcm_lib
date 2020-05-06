package rtcmlib

import (
	"fmt"
	"math"
	"strconv"
)

func Parse_1012(info Type1012) Type1012Satellite {
	var result Type1012Satellite

	result.Ident = info.SSN
	// Satellite Frequency Channel Indicator|| No. of channel|| Nominal value of frequency in L1 Band, MHz || Nominal value of frequency in L2 Band, MHz
	//0 -07 1598.0625 1242.9375
	//1 -06 1598.6250 1243.3750
	//2 -05 1599.1875 1243.8125
	//3 -04 1599.7500 1244.2500
	//4 -03 1600.3125 1244.6875
	//5 -02 1600.8750 1245.1250
	//6 -01 1601.4375 1245.5625
	//7	 00 1602.0 	  1246.0
	//8  01 1602.5625 1246.4375
	//9  02 1603.125  1246.875
	//10 03 1603.6875 1247.3125
	//11 04 1604.25   1247.75
	//12 05 1604.8125 1248.1875
	//13 06 1605.375  1248.625
	//14 07 1605.9375 1249.0625
	//15 08 1606.5 	  1249.5
	//16 09 1607.0625 1249.9375
	//17 10 1607.625  1250.375
	//18 11 1608.1875 1250.8125
	//19 12 1608.75   1251.25
	//20 13 1609.3125 1251.6875
	result.Channel = int(info.SFCN) - 7

	// 0 - C/A Code
	// 1 - P Code
	result.L1.Ind = info.L1ci


	//The GLONASS L1 Pseudorange field provides the raw L1
	//pseudorange measurement at the reference station in meters, modulo
	//two light-milliseconds (599,584.916 meters). The L1 pseudorange
	//measurement is reconstructed by the user receiver from the L1
	//pseudorange field by:
	//(L1 pseudorange measurement) = (L1 pseudorange field) modulo
	//(599,584.916 m) + integer as determined from the user receiver's
	//estimate of the reference station range, or as provided by the extended
	//data set.
	result.L1.Prange = float64(info.L1PR) * 0.02


	result.L1.Delta.Value=float64(info.L1dPR.Value) * 0.0005
	result.L1.Delta.Special=info.L1dPR.Special

	//Lock Time Indicator provides a measure of the
	//amount of time that has elapsed during which the Reference Station
	//receiver has maintained continuous lock on that satellite signal. If a
	//cycle slip occurs during the previous measurement cycle, the lock
	//indicator will be reset to zero.
	// Indicator(i) || Minimum Lock Time (s) || Range of Indicated Lock Times
	//0-23		i		 		0 < lock time < 24
	//24-47		i*2-24 			24 ≤ lock time < 72
	//48-71 	i*4-120 		72 ≤ lock time < 168
	//72-95 	i*8-408 		168 ≤ lock time < 360
	//96-119 	i*16-1176 		360 ≤ lock time < 744
	//120-126 	i*32-3096 		744 ≤ lock time < 937
	//127 		--- 			lock time >= 937
	result.L1.Lockt=info.L1Lt

	//Integer L1 Pseudorange Modulus Ambiguity
	//represents the integer number of full pseudorange modulus divisions
	//(599,584.916 m) of the raw L1 pseudorange measurement
	//
	// In gpsdecode:
	// "amb":info.L1MA,
	//result.L1.Amb=float64(info.L1MA)*599584.916 // for Meters
	result.L1.Amb=info.L1MA

	// L1 CNR measurements provide the reference station's
	//estimate of the carrier-to-noise ratio of the satellite’s signal in dB-Hz.
	// 0 - the CNR measurement is not computed.
	result.L1.CNR=float64(info.L1CNR)*0.25


	// The GLONASS L2 Code Indicator depicts which L2 code is processed
	//by the reference station.
	//0 - C/A code
	//1- P code
	//2 - Reserved
	//3 - Reserved
	result.L2.Ind= info.L2CI

	//The GLONASS L2-L1 Pseudorange Difference field is utilized, rather
	//than the full L2 pseudorange, in order to reduce the message length.
	//The receiver must reconstruct the L2 code phase pseudorange by using
	//the following formula:
	// (GLONASS L2 pseudorange measurement) = (L1 pseudorange as
	//reconstructed from L1 pseudorange field) + (L2-L1 pseudorange field)
	// 200h (-163.84) – there is no valid L2 code available, or the value
	//exceeds the allowed range.
	result.L2.Prange.Value=float64(info.L2PR.Value)*0.02
	result.L2.Prange.Special=info.L2PR.Special

	//L2 PhaseRange - L1 Pseudorange
	result.L2.Delta.Value =float64(info.L2dPR.Value) * 0.0005
	result.L2.Delta.Special =info.L2dPR.Special

	// L2 Lock Time Indicator provides a measure of the
	//amount of time that has elapsed during which the Reference Station
	//receiver has maintained continuous lock on that satellite signal. If a
	//cycle slip occurs during the previous measurement cycle, the lock
	//indicator will be reset to zero.
	result.L2.Lockt=info.L2Lt

	// L2 CNR measurements provide the reference station's
	//estimate of the carrier-to-noise ratio of the satellite’s signal in dB-Hz.
	//0 – The CNR measurement is not computed.
	result.L2.CNR =float64(info.L2CNR)*0.25

	return result
}

func parse_1087Satellite(info Type1087Satellite) Type1087SatelliteParsed {
	var result Type1087SatelliteParsed
	var RoughRangeInt specialUint64
	// RoughRangeInt
	// DF397
	// Rough range can be used to restore complete observables for a given
	//satellite. Rough range takes 18 bits which are split between two fields
	//(DF397 and DF398). This field contains the integer number of
	//milliseconds in the satellite rough range.
	//A bit pattern equivalent to FFh (255 ms) indicates invalid value.
	if info.RoughRangeInt == 255 {
		RoughRangeInt.Special = true
		RoughRangeInt.Value = 0
	} else {
		RoughRangeInt.Special = false
		RoughRangeInt.Value = info.RoughRangeInt
	}

	// Info
	// The GLONASS Satellite Frequency Channel Number identifies the
	//frequency of the GLONASS satellite.
	// DF Value || No of channel ||Nominal value of frequency in L1 Band, MHz || Nominal value of frequency in L2 Band, MHz
	//0 -7 1598.0625 	1242.9375
	//1 -6 1598.6250 	1243.3750
	//2 -5 1599.1875 	1243.8125
	//3 -4 1599.7500 	1244.2500
	//4 -3 1600.3125	1244.6875
	//5 -2 1600.8750 	1245.1250
	//6 -1 1601.4375 	1245.5625
	//7	 0 1602.0 		1246.0
	//8	 1 1602.5625 	1246.4375
	//9	 2 1603.125 	1246.875
	//10 3 1603.6875 	1247.3125
	//11 4 1604.25 		1247.75
	//12 5 1604.8125 	1248.1875
	//13 6 1605.375 	1248.625
	result.Info = int(info.Info) - 7

	// DF398
	//RoughRangeRemainder
	//See the note above. Allows restoring full rough range with accuracy 1/1024 ms (about 300 m).

	RoughRangeRemainder := float64(info.RoughRangeRemainder) * math.Pow(2,-10)

	//Calculating RoughRange
	if RoughRangeInt.Special {
		result.RoughRange.Special = true
		result.RoughRange.Value = 0
	} else {
		result.RoughRange.Special = false
		result.RoughRange.Value = float64(RoughRangeInt.Value) + RoughRangeRemainder
	}

	//RoughPhaseRange
	// DF399
	result.RatePhaseRangeInt = info.RatePhaseRangeInt

	// SatelliteNumber
	result.SatelliteNumber = info.SatelliteNumber

	// Signal mapping
	// Signal ID DF395 || Frequency Band || Signal
	// 2		G1		C/A
	// 3		G1		P
	// 8		G2		C/A
	// 9		G2		P
	for _, signal := range info.Signals {
		if signal.SignalNumber == 2 || signal.SignalNumber == 3 {
			result.L1 = parse_1087SignalsL1(signal)
		}
		if signal.SignalNumber == 8 || signal.SignalNumber == 9 {
			result.L2 = parse_1087SignalsL2(signal)
		}
	}

	return result
}


func parse_1087SignalsL1(info Type1087Signal) Type1087L1 {
	var result Type1087L1

	//RangeInt
	//DF405
	if info.PseudoRangeCorrection.Special {
		result.PseudoRangeCorrection.Special = true
		result.PseudoRangeCorrection.Value = 0
	} else {
		result.PseudoRangeCorrection.Special = false
		result.PseudoRangeCorrection.Value = float64(info.PseudoRangeCorrection.Value) * math.Pow(2,-29)
	}

	//PhaseRange
	//DF406
	if info.PhaseRangeCorrection.Special {
		result.PhaseRangeCorrection.Special = true
		result.PhaseRangeCorrection.Value = 0
	} else {
		result.PhaseRangeCorrection.Special = false
		result.PhaseRangeCorrection.Value = float64(info.PhaseRangeCorrection.Value) * math.Pow(2,-31)
	}

	//PhaseRangeTI
	//DF407
	// Indicator (i) || Minimum Lock Time (s) || Range of Indicated Lock Times
	//0-23 		i 					0 < lock time < 24
	//24-47 	i * 2 - 24 			24 ≤ lock time < 72
	//48-71 	i * 4 - 120 		72 ≤ lock time < 168
	//72-95 	i * 8 - 408 		168 ≤ lock time < 360
	//96-119 	i * 16 - 1176 		360 ≤ lock time < 744
	//120-126 	i * 32 - 3096 		744 ≤ lock time < 937
	//127 		--- 				lock time >= 937
	result.PhaseRangeTI = info.PhaseRangeTI

	//AI
	//DF420
	// 0 – No half-cycle ambiguity.
	//1 – Half-cycle ambiguity
	//When transmitting PhaseRange with unresolved polarity encoding
	//software shall set this bit to 1. Receiving software that is not capable of
	//handling half-cycle ambiguities shall skip such PhaseRange
	//observables. If polarity resolution forced PhaseRange to be corrected by
	//half-a-cycle, then the associated GNSS PhaseRange Lock Time
	//Indicator (DF402, DF407) must be reset to zero, indicating that despite
	//continuous tracking the final PhaseRange experienced non-continuity.

	switch info.AI{
	case 0:
		result.AI = "No half-cycle ambiguity."
	case 1:
		result.AI = "Half-cycle ambiguity"
	}

	//CNR
	//DF408
	// A value “0” indicates that the CNR measurement has not been
	//computed, or is not available.
	//Availability or unavailability of the CNR does not affect validity of
	//other observables.
	result.CNR = float64(info.CNR) * math.Pow(2,-4)

	//PhaseRangeRate
	//DF404
	//Fine Phase Range Rate for a given signal. Full Phase Range Rate is the
	//sum of this field and the Satellite Rough Phase Range Rate (DF399).
	//A bit pattern equivalent to 4000h (–1.6384 m/s) indicates invalid value
	// m/s
	if info.RatePhaseRangeRemainder.Special {
		result.RatePhaseRangeRemainder.Special = true
		result.RatePhaseRangeRemainder.Value = 0
	} else {
		result.RatePhaseRangeRemainder.Special = false
		result.RatePhaseRangeRemainder.Value = float64(info.RatePhaseRangeRemainder.Value) * 0.0001
	}


	switch info.SignalNumber{
	case 2:
		result.Signal = "C/A"
	case 3:
		result.Signal = "P"
	}

	return result

}


func parse_1087SignalsL2(info Type1087Signal) Type1087L2 {
	var result Type1087L2

	//RangeInt
	//DF405
	if info.PseudoRangeCorrection.Special {
		result.PseudoRangeCorrection.Special = true
		result.PseudoRangeCorrection.Value = 0
	} else {
		result.PseudoRangeCorrection.Special = false
		result.PseudoRangeCorrection.Value = float64(info.PseudoRangeCorrection.Value) * math.Pow(2,-29)
	}

	//PhaseRange
	//DF406
	if info.PhaseRangeCorrection.Special {
		result.PhaseRangeCorrection.Special = true
		result.PhaseRangeCorrection.Value = 0
	} else {
		result.PhaseRangeCorrection.Special = false
		result.PhaseRangeCorrection.Value = float64(info.PhaseRangeCorrection.Value) * math.Pow(2,-31)
	}

	//PhaseRangeTI
	//DF407
	// Indicator (i) || Minimum Lock Time (s) || Range of Indicated Lock Times
	//0-23 		i 					0 < lock time < 24
	//24-47 	i * 2 - 24 			24 ≤ lock time < 72
	//48-71 	i * 4 - 120 		72 ≤ lock time < 168
	//72-95 	i * 8 - 408 		168 ≤ lock time < 360
	//96-119 	i * 16 - 1176 		360 ≤ lock time < 744
	//120-126 	i * 32 - 3096 		744 ≤ lock time < 937
	//127 		--- 				lock time >= 937
	result.PhaseRangeTI = info.PhaseRangeTI

	//AI
	//DF420
	// 0 – No half-cycle ambiguity.
	//1 – Half-cycle ambiguity
	//When transmitting PhaseRange with unresolved polarity encoding
	//software shall set this bit to 1. Receiving software that is not capable of
	//handling half-cycle ambiguities shall skip such PhaseRange
	//observables. If polarity resolution forced PhaseRange to be corrected by
	//half-a-cycle, then the associated GNSS PhaseRange Lock Time
	//Indicator (DF402, DF407) must be reset to zero, indicating that despite
	//continuous tracking the final PhaseRange experienced non-continuity.

	switch info.AI{
	case 0:
		result.AI = "No half-cycle ambiguity."
	case 1:
		result.AI = "Half-cycle ambiguity"
	}

	//CNR
	//DF408
	// A value “0” indicates that the CNR measurement has not been
	//computed, or is not available.
	//Availability or unavailability of the CNR does not affect validity of
	//other observables.
	result.CNR = float64(info.CNR) * math.Pow(2,-4)

	//PhaseRangeRate
	//DF404
	//Fine Phase Range Rate for a given signal. Full Phase Range Rate is the
	//sum of this field and the Satellite Rough Phase Range Rate (DF399).
	//A bit pattern equivalent to 4000h (–1.6384 m/s) indicates invalid value
	// m/s
	if info.RatePhaseRangeRemainder.Special {
		result.RatePhaseRangeRemainder.Special = true
		result.RatePhaseRangeRemainder.Value = 0
	} else {
		result.RatePhaseRangeRemainder.Special = false
		result.RatePhaseRangeRemainder.Value = float64(info.RatePhaseRangeRemainder.Value) * 0.0001
	}


	switch info.SignalNumber{
	case 8:
		result.Signal = "C/A"
	case 9:
		result.Signal = "P"
	}

	return result

}


func Parse_1087(info Type1087) Type1087Parsed {
	var result Type1087Parsed

	//The Reference Station ID is determined by the service provider. Its
	//primary purpose is to link all message data to their unique source.
	result.StationId = info.StationId

	//0 – Sunday
	//1 – Monday
	//2 – Tuesday
	//3 – Wednesday
	//4 – Thursday
	//5 – Friday
	//6 – Saturday
	//7 – The day of week is not known
	result.Day = info.Day

	//GLONASS Epoch Time of measurement is defined by the GLONASS
	//ICD as UTC(SU) + 3.0 hours. It rolls over at 86,400 seconds for
	//GLONASS, except for the leap second, where it rolls over at 86,401.
	// in ms.
	result.Epoch = info.Epoch

	//CSI
	//0 – clock steering is not applied
	//In this case receiver clock must be kept in the range of ± 1 ms
	//(approximately ± 300 km)
	//1 – clock steering has been applied
	//In this case receiver clock must be kept in the range of ± 1 microsecond
	//(approximately ± 300 meters).
	//2 – unknown clock steering status
	//3 – reserved
	switch info.CSI {
	case 0:
		result.CSI = "clock steering is not applied"
	case 1:
		result.CSI = "clock steering has been applied"
	case 2:
		result.CSI = "unknown clock steering status"
	case 3:
		result.CSI = "clock steering RESERVED status"
	}

	//ECI
	//0 – internal clock is used
	//1 – external clock is used, clock status is “locked”
	//2 – external clock is used, clock status is “not locked”, which may
	//indicate external clock failure and that the transmitted data may not be
	//reliable.
	//3 – unknown clock is used
	switch info.ECI {
	case 0:
		result.ECI = "internal clock is used"
	case 1:
		result.ECI = "external clock is used, clock status is “locked”"
	case 2:
		result.ECI = "external clock is used, clock status is “not locked”"
	case 3:
		result.ECI = "unknown clock is used"
	}

	// Sindi
	// 1 – Divergence-free smoothing is used
	// 0 – Other type of smoothing is used
	switch info.SIndi {
	case 0:
		result.SIndi = "Other type of smoothing is used"
	case 1:
		result.SIndi = "Divergence-free smoothing is used"
	}

	// SInter
	// The GNSS Smoothing Interval is the integration period over which the
	//pseudorange code phase measurements are averaged using carrier phase
	//information.
	//Divergence-free smoothing may be continuous over the entire period
	//for which the satellite is visible.
	//Notice: A value of zero indicates no smoothing is used.
	//
	/*
	switch info.SInter {
	case 0:
		result.SInter = "No smoothing is used"
	case 1:
		result.SInter = "Smoothing < 30 s"
	case 2:
		result.SInter = "Smoothing 30-60 s"
	case 3:
		result.SInter = "Smoothing 1-2 min"
	case 4:
		result.SInter = "Smoothing 2-4 min"
	case 5:
		result.SInter = "Smoothing 4-8 min"
	case 6:
		result.SInter = "Smoothing >8 min"
	case 7:
		result.SInter = "Unlimited smoothing interval"
	} */
	result.SInter = info.SInter

	for _, satellite := range info.Satellites {
		result.Satellites = append(result.Satellites, parse_1087Satellite(satellite))
	}

	return result
}

func Parse_1004(info Type1004) Type1004Satellite {
	var result Type1004Satellite

	//A GPS Satellite ID number from 1 to 32 refers to the PRN code of the
	//GPS satellite. Satellite ID’s higher than 32 are reserved for satellite
	//signals from Satellite-Based Augmentation Systems (SBAS’s) such as
	//the FAA’s Wide-Area Augmentation System (WAAS). SBAS PRN
	//codes cover the range 120-138. The Satellite ID’s reserved for SBAS
	//satellites are 40-58, so that the SBAS PRN codes are derived from the
	//Version 3 Satellite ID codes by adding 80.

	result.Ident = info.SSN



	//The GPS L1 Code Indicator identifies the code being tracked by the
	//reference station. Civil receivers can track the C/A code, and
	//optionally the P code, while military receivers can track C/A, and can
	//also track P and Y code, whichever is broadcast by the satellite
	//0 - C/A Code
	//1 - P(Y) Code Direct
	result.L1.Ind= info.L1ci

	//The GPS L1 Pseudorange field provides the raw L1 pseudorange
	//measurement at the reference station in meters, modulo one light-
	//millisecond (299,792.458 meters). The GPS L1 pseudorange
	//measurement is reconstructed by the user receiver from the L1
	//pseudorange field by:
	//(GPS L1 pseudorange measurement) = (GPS L1 pseudorange field)
	//modulo (299,792.458 m) + integer as determined from the user
	//receiver's estimate of the reference station range, or as provided by the
	//extended data set.
	//80000h - invalid L1 pseudorange; used only in the calculation of L2
	//measurements.
	result.L1.Prange.Value=float64(info.L1PR) * 0.02

	//L1 PhaseRange – L1 Pseudorange
	result.L1.Delta.Value=float64(info.L1dPR.Value) * 0.0005
	result.L1.Delta.Special=info.L1dPR.Special

	//The GPS L1 Lock Time Indicator provides a measure of the amount of
	//time that has elapsed during which the Reference Station receiver has
	//maintained continuous lock on that satellite signal. If a cycle slip
	//occurs during the previous measurement cycle, the lock indicator will
	//be reset to zero.
	// Indicator(i) || Minimum Lock Time (s) || Range of Indicated Lock Times
	//0-23		i		 		0 < lock time < 24
	//24-47		i*2-24 			24 ≤ lock time < 72
	//48-71 	i*4-120 		72 ≤ lock time < 168
	//72-95 	i*8-408 		168 ≤ lock time < 360
	//96-119 	i*16-1176 		360 ≤ lock time < 744
	//120-126 	i*32-3096 		744 ≤ lock time < 937
	//127 		--- 			lock time >= 937
	result.L1.Lockt=info.L1Lt

	//The GPS Integer L1 Pseudorange Modulus Ambiguity represents the
	//integer number of full pseudorange modulus divisions (299,792.458 m)
	//of the raw L1 pseudorange measurement.
	//
	// In gpsdecode:
	// "amb":info.L1MA,
	result.L1.Amb=float64(info.L1MA)*299792.458 // for Meters

	// The GPS L1 CNR measurements provide the reference station's
	//estimate of the carrier-to-noise ratio of the satellite’s signal in dB-Hz.
	//0 - the CNR measurement is not computed.
	result.L1.CNR=float64(info.L1CNR)*0.25


	// The GPS L2 Code Indicator depicts which L2 code is processed by the
	//reference station, and how it is processed.
	//0 - C/A or L2C code
	//1 - P(Y) code direct
	//2 - P(Y) code cross-correlated
	//3 - Correlated P/Y
	//The GPS L2 Code Indicator refers to the method used by the GPS
	//reference station receiver to recover the L2 pseudorange. The GPS L2
	//Code Indicator should be set to “0” (C/A or L2C code) for any of the
	//L2 civil codes. It is assumed here that a satellite will not transmit both
	//C/A code and L2C code signals on L2 simultaneously, so that the
	//reference station and user receivers will always utilize the same signal.
	//The code indicator should be set to “1” if the satellite’s signal is
	//correlated directly, i.e., either P code or Y code depending whether
	//anti-spoofing (AS) is switched off or on. The code indicator should be
	//set to “2” when the reference station receiver L2 pseudorange
	//measurement is derived by adding a cross-correlated pseudorange
	//measurement (Y2-Y1) to the measured L1 C/A code. The code
	//indicator should be set to 3 when the GPS reference station receiver is
	//using a proprietary method that uses only the L2 P(Y) code signal to
	//derive L2 pseudorange
	result.L2.Ind= info.L2CI

	//The GPS L2-L1 Pseudorange Difference field is utilized, rather than
	//the full L2 pseudorange, in order to reduce the message length. The
	//receiver must reconstruct the L2 code phase pseudorange by using the
	//following formula:
	//(GPS L2 pseudorange measurement) =
	//(GPS L1 pseudorange as reconstructed from L1 pseudorange field) +
	//(GPS L2-L1 pseudorange field)
	//2000h (-163.84m) - no valid L2 code available, or that the value
	//exceeds the allowed range.
	result.L2.Prange.Value=float64(info.L2PR.Value)*0.02
	result.L2.Prange.Special=info.L2PR.Special

	//L2 PhaseRange - L1 Pseudorange
	result.L2.Delta.Value=float64(info.L2dPR.Value) * 0.0005
	result.L2.Delta.Special=info.L2dPR.Special

	//The GPS L2 Lock Time Indicator provides a measure of the amount of
	//time that has elapsed during which the Reference Station receiver has
	//maintained continuous lock on that satellite signal. If a cycle slip
	//occurs during the previous measurement cycle, the lock indicator will
	//be reset to zero.
	result.L2.Lockt=info.L2Lt

	//The GPS L2 CNR measurements provide the reference station's
	//estimate of the carrier-to-noise ratio of the satellite’s signal in dB-Hz.
	//0 - the CNR measurement is not computed
	result.L2.CNR=float64(info.L2CNR)*0.25

	return result
}


func Prepare_1012(info Type1012Satellite) Type1012 {
	var result Type1012

	result.SSN = info.Ident

	result.SFCN  = uint64(info.Channel + 7)

	result.L1ci = info.L1.Ind

	result.L1PR = uint64(math.Round(info.L1.Prange / 0.02))

	result.L1dPR.Special = info.L1.Delta.Special
	result.L1dPR.Value = int64(math.Round(info.L1.Delta.Value / 0.0005))

	result.L1Lt = info.L1.Lockt

	//result.L1MA = uint64(math.Round(info.L1.Amb / 599584.916))
	result.L1MA = info.L1.Amb

	result.L1CNR = uint64(math.Round(info.L1.CNR / 0.25))

	result.L2CI = info.L2.Ind

	result.L2PR.Special = info.L2.Prange.Special
	result.L2PR.Value = int64(math.Round(info.L2.Prange.Value /0.02))

	result.L2dPR.Special = info.L2.Delta.Special
	result.L2dPR.Value = int64(math.Round(info.L2.Delta.Value / 0.0005))

	result.L2Lt = info.L2.Lockt

	result.L2CNR = uint64(math.Round(info.L2.CNR / 0.25))

	return result
}

func Prepare_1004(info Type1004Satellite) Type1004 {
	var result Type1004

	result.SSN = info.Ident

	result.L1ci = info.L1.Ind

	if info.L1.Prange.Special {
		result.L1PR = 0x800000
	} else {
		result.L1PR = uint64(math.Round(info.L1.Prange.Value / 0.02))
	}

	result.L1dPR.Special = info.L1.Delta.Special
	result.L1dPR.Value = int64(math.Round(info.L1.Delta.Value / 0.0005))

	result.L1Lt = info.L1.Lockt

	result.L1MA = uint64(math.Round(info.L1.Amb / 299792.458 ))

	result.L1CNR = uint64(math.Round(info.L1.CNR / 0.25))

	result.L2CI = info.L2.Ind

	result.L2PR.Special = info.L2.Prange.Special
	result.L2PR.Value = int64(math.Round(info.L2.Prange.Value /0.02))

	result.L2dPR.Value = int64(math.Round(info.L2.Delta.Value / 0.0005))
	result.L2dPR.Special = info.L2.Delta.Special

	result.L2Lt = info.L2.Lockt

	result.L2CNR = uint64(math.Round(info.L2.CNR / 0.25))

	return result
}

func Prepare_1012_header(info Type1012Parsed) Type1012Header {
	var result Type1012Header

	result.StationId = info.Station_id

	result.Epoch = info.Tow

	if info.Sync {
		result.SmoothFlag = 1
	} else {
		result.SmoothFlag = 0
	}

	if info.Smoothing {
		result.DIndicator = 1
	} else {
		result.DIndicator = 0
	}

	result.SmoothInterval,_ = strconv.ParseUint(info.Interval, 10, 64)

	result.NoSatellites = uint64(len(info.Satellites))

	return result
}

func Prepare_1004_header(info Type1004Parsed) Type1004Header {
	var result Type1004Header

	result.StationId = info.Station_id

	result.Epoch = info.Tow

	if info.Sync {
		result.SmoothFlag = 1
	} else {
		result.SmoothFlag = 0
	}

	if info.Smoothing {
		result.DIndicator = 1
	} else {
		result.DIndicator = 0
	}

	result.SmoothInterval,_ = strconv.ParseUint(info.Interval, 10, 64)

	result.NoSatellites = uint64(len(info.Satellites))

	return result
}


func Parse_1012_header(result *Type1012Parsed, info Type1012Header) {
	//var result Type1012Parsed

	// The Reference Station ID is determined by the service provider. Its
	//primary purpose is to link all message data to their unique source.
	result.Station_id = info.StationId

	//Epoch Time of measurement is defined by the GLONASS
	//ICD as UTC(SU) + 3.0 hours. It rolls over at 86,400 seconds for
	//GLONASS, except for the leap second, where it rolls over at 86,401.
	result.Tow = info.Epoch

	//0 - No further GNSS observables referenced to the same Epoch Time
	//will be transmitted. This enables the receiver to begin processing
	//the data immediately after decoding the message.
	//1 - The next message will contain observables of another GNSS
	//source referenced to the same Epoch Time.
	//Note: “Synchronous" here means that the measurements are taken
	//within one microsecond of each other
	result.Smoothing	= !(info.SmoothFlag == 0)

	//0 - Divergence-free smoothing not used
	//1 - Divergence-free smoothing used
	result.Sync	= !(info.DIndicator == 0)

	//Smoothing Interval is the integration period over
	//which reference station pseudorange code phase measurements are
	//averaged using carrier phase information. Divergence-free smoothing
	//may be continuous over the entire period the satellite is visible.
	//Indicator || Smoothing Interval
	//000 (0) 		No smoothing
	//001 (1) 		< 30 s
	//010 (2) 		30-60 s
	//011 (3) 		1-2 min
	//100 (4) 		2-4 min
	//101 (5) 		4-8 min
	//110 (6) 		>8 min
	//111 (7) 		Unlimited smoothing interval

	result.Interval = fmt.Sprintf("%d", info.SmoothInterval)

	// Number of satellites
	// Not Used
	// info.NoSatellites
	// DF035

	return
}

func Parse_1004_header(result *Type1004Parsed, info Type1004Header) {
	//var result Type1004Parsed

	// The Reference Station ID is determined by the service provider. Its
	//primary purpose is to link all message data to their unique source.
	result.Station_id = info.StationId

	//GPS Epoch Time is provided in milliseconds from the beginning of the GPS week,
	//which begins at midnight GMT on Saturday night/Sunday morning,
	//measured in GPS time (as opposed to UTC).
	// 1 ms
	result.Tow = info.Epoch

	//0 - No further GNSS observables referenced to the same Epoch Time
	//will be transmitted. This enables the receiver to begin processing
	//the data immediately after decoding the message.
	//1 - The next message will contain observables of another GNSS
	//source referenced to the same Epoch Time.
	//Note: “Synchronous" here means that the measurements are taken
	//within one microsecond of each other
	result.Sync = !(info.SmoothFlag == 0)

	//0 - Divergence-free smoothing not used
	//1 - Divergence-free smoothing used
	result.Smoothing	= !(info.DIndicator == 0)

	//The GPS Smoothing Interval is the integration period over
	//which reference station pseudorange code phase measurements
	//are averaged using carrier phase information.
	//Divergence-free smoothing may be continuous over
	//the entire period the satellite is visible.
	//Indicator || Smoothing Interval
	//000 (0) 		No smoothing
	//001 (1) 		< 30 s
	//010 (2) 		30-60 s
	//011 (3) 		1-2 min
	//100 (4) 		2-4 min
	//101 (5) 		4-8 min
	//110 (6) 		>8 min
	//111 (7) 		Unlimited smoothing interval

	result.Interval = fmt.Sprintf("%d", info.SmoothInterval)

	// Number of satellites
	// Not Used
	// info.NoSatellites
	// DF006

	return
}
