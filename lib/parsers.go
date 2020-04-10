package rtcmlib

import (
	"fmt"
	"math"
	"strconv"
)

func Parse_2012(info Type1012) Type1012Satellite {
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

	result.L1 = Type1012L1{
		// 0 - C/A Code
		// 1 - P Code
		Ind: info.L1ci,
		//The GLONASS L1 Pseudorange field provides the raw L1
		//pseudorange measurement at the reference station in meters, modulo
		//two light-milliseconds (599,584.916 meters). The L1 pseudorange
		//measurement is reconstructed by the user receiver from the L1
		//pseudorange field by:
		//(L1 pseudorange measurement) = (L1 pseudorange field) modulo
		//(599,584.916 m) + integer as determined from the user receiver's
		//estimate of the reference station range, or as provided by the extended
		//data set.
		Prange:float64(info.L1PR) * 0.02,

		//L1 PhaseRange – L1 Pseudorange
		Delta:float64(info.L1dPR) * 0.0005,

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
		Lockt:info.L1Lt,

		//Integer L1 Pseudorange Modulus Ambiguity
		//represents the integer number of full pseudorange modulus divisions
		//(599,584.916 m) of the raw L1 pseudorange measurement
		//
		// In gpsdecode:
		// "amb":info.L1MA,
		Amb:float64(info.L1MA)*599584.916, // for Meters

		// L1 CNR measurements provide the reference station's
		//estimate of the carrier-to-noise ratio of the satellite’s signal in dB-Hz.
		// 0 - the CNR measurement is not computed.
		CNR:float64(info.L1CNR)*0.25,
	}
	result.L2 = Type1012L2{
		// The GLONASS L2 Code Indicator depicts which L2 code is processed
		//by the reference station.
		//0 - C/A code
		//1- P code
		//2 - Reserved
		//3 - Reserved
		Ind: info.L2CI,

		//The GLONASS L2-L1 Pseudorange Difference field is utilized, rather
		//than the full L2 pseudorange, in order to reduce the message length.
		//The receiver must reconstruct the L2 code phase pseudorange by using
		//the following formula:
		// (GLONASS L2 pseudorange measurement) = (L1 pseudorange as
		//reconstructed from L1 pseudorange field) + (L2-L1 pseudorange field)
		// 200h (-163.84) – there is no valid L2 code available, or the value
		//exceeds the allowed range.
		Prange:float64(info.L2PR)*0.02,

		//L2 PhaseRange - L1 Pseudorange
		Delta:float64(info.L2dPR) * 0.0005,

		// L2 Lock Time Indicator provides a measure of the
		//amount of time that has elapsed during which the Reference Station
		//receiver has maintained continuous lock on that satellite signal. If a
		//cycle slip occurs during the previous measurement cycle, the lock
		//indicator will be reset to zero.
		Lockt:info.L2Lt,

		// L2 CNR measurements provide the reference station's
		//estimate of the carrier-to-noise ratio of the satellite’s signal in dB-Hz.
		//0 – The CNR measurement is not computed.
		CNR:float64(info.L2CNR)*0.25,
	}
	return result
}

func Prepare_1012(info Type1012Satellite) Type1012 {
	var result Type1012

	result.SSN = info.Ident

	result.SFCN  = uint64(info.Channel + 7)

	result.L1ci = info.L1.Ind

	result.L1PR = uint64(math.Round(info.L1.Prange / 0.02))

	result.L1dPR = int64(math.Round(info.L1.Delta / 0.0005))

	result.L1Lt = info.L1.Lockt

	result.L1MA = uint64(math.Round(info.L1.Amb / 599584.916))

	result.L1CNR = uint64(math.Round(info.L1.CNR / 0.25))

	result.L2CI = info.L2.Ind

	result.L2PR = uint64(math.Round(info.L2.Prange /0.02))

	result.L2dPR = int64(math.Round(info.L2.Delta / 0.0005))

	result.L2Lt = info.L2.Lockt

	result.L2CNR = uint64(math.Round(info.L2.CNR / 0.25))

	return result
}

func Prepare_1012_header(info Type1012Parsed) Type1012Header {
	var result Type1012Header

	result.StationId = info.Station_id

	result.Epoch = info.Tow

	if info.Smoothing {
		result.SFlag = 1
	} else {
		result.SFlag = 0
	}

	if info.Sync {
		result.DIndicator = 1
	} else {
		result.DIndicator = 0
	}

	result.SmoothInterval,_ = strconv.ParseUint(info.Interval, 10, 64)

	result.NoSatellites = uint64(len(info.Satellites))

	return result
}

func Parse_2012_header(result *Type1012Parsed, info Type1012Header) {
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
	result.Smoothing	= !(info.SFlag == 0)

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

