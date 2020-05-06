package rtcmlib

type TypeCommonHeader struct {
	// Message Number (“1009”=0011 1111 0001)
	MessageNumber	uint64
}

//=================HEADERS=================

type Type1012Header struct {
	//Reference Station ID DF003 uint12 12
	StationId		uint64
	//GLONASS Epoch Time (t k ) DF034 uint27 27
	Epoch			uint64
	//Synchronous GNSS Flag DF005 bit(1) 1
	SmoothFlag		uint64
	//No. of GLONASS Satellite Signals Processed DF035 uint5 5
	NoSatellites	uint64
	//GLONASS Divergence-free Smoothing Indicator DF036 bit(1) 1
	DIndicator		uint64
	//GLONASS Smoothing Interval DF037 bit(3) 3
	SmoothInterval	uint64
}

type Type1004Header struct {
	//Reference Station ID DF003 uint12 12
	StationId		uint64
	//GPS Epoch Time (TOW) DF004 uint30 30
	Epoch			uint64
	//Synchronous GNSS Flag DF005 bit(1) 1
	SmoothFlag		uint64
	//No. of GPS Satellite Signals Processed DF006 uint5 5
	NoSatellites	uint64
	//GPS Divergence-free Smoothing Indicator DF007 bit(1) 1
	DIndicator		uint64
	//GPS Smoothing Interval DF008 bit(3) 3
	SmoothInterval	uint64
}

// ========= BODY================

type Type1012 struct {
	//GLONASS Satellite ID (Satellite Slot Number) DF038 uint6 6
	SSN			uint64
	//GLONASS L1 Code Indicator DF039 bit(1) 1
	L1ci		uint64
	//GLONASS Satellite Frequency Channel Number DF040 uint5 5
	SFCN		uint64
	//GLONASS L1 Pseudorange DF041 uint25 25
	L1PR		uint64
	//GLONASS L1 PhaseRange – L1 Pseudorange DF042 int20 20
	L1dPR		specialInt64
	//GLONASS L1 Lock time Indicator DF043 uint7 7
	L1Lt		uint64
	//GLONASS Integer L1 Pseudorange Modulus Ambiguity DF044 uint7 7
	L1MA		uint64
	//GLONASS L1 CNR DF045 uint8 8
	L1CNR		uint64
	//GLONASS L2 Code Indicator DF046 bit(2) 2
	L2CI		uint64
	//GLONASS L2-L1 Pseudorange Difference DF047 uint14 14
	L2PR		specialInt64
	//GLONASS L2 PhaseRange – L1 Pseudorange DF048 int20 20
	L2dPR		specialInt64
	//GLONASS L2 Lock time Indicator DF049 uint7 7
	L2Lt		uint64
	//GLONASS L2 CNR DF050 uint8 8
	L2CNR		uint64
}

type specialInt64 struct {
	Value		int64
	Special		bool
}

type specialUint64 struct {
	Value		uint64
	Special		bool
}


type Specialfloat64 struct {
	Value 		float64
	Special 	bool
}

type Type1004 struct {
	//GPS Satellite ID DF009 uint6 6
	SSN			uint64
	//GPS L1 Code Indicator DF010 bit(1) 1
	L1ci		uint64
	//GPS L1 Pseudorange DF011 uint24 24
	L1PR		uint64
	//GPS L1 PhaseRange – L1 Pseudorange DF012 int20 20
	L1dPR		specialInt64
	//GPS L1 Lock time Indicator DF013 uint7 7
	L1Lt		uint64
	//GPS Integer L1 Pseudorange Modulus Ambiguity DF014 uint8 8
	L1MA		uint64
	//GPS L1 CNR DF015 uint8 8
	L1CNR		uint64
	//GPS L2 Code Indicator DF016 bit(2) 2
	L2CI		uint64
	//DF 17
	L2PR		specialInt64
	// DF 18
	L2dPR		specialInt64
	//DF 19
	L2Lt		uint64
	//GPS L2 CNR DF020 uint8 8
	L2CNR		uint64
}

type Type1012Parsed struct {
	//from gpsdecode
	Class			string
	Device			string
	Length			uint64
	Type			int
	//header
	Station_id		uint64
	Tow				uint64
	Smoothing		bool
	//NoSatellites	uint64
	Sync			bool
	Interval		string
	//body
	Satellites		[]Type1012Satellite

}

type Type1004Parsed struct {
	//from gpsdecode
	Class			string
	Device			string
	Length			uint64
	Type			int
	//header
	Station_id		uint64
	Tow				uint64
	Smoothing		bool
	//NoSatellites	uint64
	Sync			bool
	Interval		string
	//body
	Satellites		[]Type1004Satellite

}

type Type1087 struct {
	//header part ------------------------
	//Reference Station ID DF003 uint12 12
	StationId		uint64
	//GLONASS Day Of Week DF416 uint3 3
	Day				uint64
	//GLONASS Epoch Time (t k ) DF034 uint27 27
	Epoch			uint64
	//Multiple Message Bit DF393 bit(1) 1
	MMB				uint64
	//IODS – Issue of Data Station DF409 uint3 3
	//This field is reserved to be used to link MSM with future site-
	//description (receiver, antenna description, etc.) messages.
	//A value of “0” indicates that this field is not utilized.
	IODS			uint64
	//Reserved DF001 bit(7) 7 (may be GNSS specific)

	//Clock Steering Indicator DF411 uint2 2
	CSI				uint64
	// External Clock Indicator DF412 uint2 2
	ECI				uint64
	//GNSS Divergence-free Smoothing Indicator DF417 bit(1) 1
	SIndi 			uint64
	// GNSS Smoothing Interval DF418 bit(3) 3
	SInter			uint64
	//GNSS Satellite Mask DF394 bit(64) 64
	SatMask			string
	SatNumber		int
	//GNSS Signal Mask DF395 bit(32) 32
	SignalMask		string
	SignalNumber	int
	// GNSS Cell Mask DF396 bit(X) X (X≤64)
	SatSignalTable	string

	// satellites part ----------------------
	Satellites		[]Type1087Satellite

}

type Type1087Parsed struct {

	//Reference Station ID DF003 uint12 12
	StationId		uint64
	//GLONASS Day Of Week DF416 uint3 3
	Day				uint64
	//GLONASS Epoch Time (t k ) DF034 uint27 27
	Epoch			uint64
	//Multiple Message Bit DF393 bit(1) 1
	// NOT USED
	//IODS – Issue of Data Station DF409 uint3 3
	// NOT USED
	//Reserved DF001 bit(7) 7 (may be GNSS specific)
	// NOT USED
	//Clock Steering Indicator DF411 uint2 2 CONVERTED TO MESSAGE
	CSI				string
	// External Clock Indicator DF412 uint2 2 CONVERTED TO MESSAGE
	ECI				string
	//GNSS Divergence-free Smoothing Indicator DF417 bit(1) 1 CONVERTED TO MESSAGE
	SIndi 			string
	// GNSS Smoothing Interval DF418 bit(3) 3
	SInter			uint64

	// satellites part ----------------------
	Satellites		[]Type1087SatelliteParsed

}

type Type1087SatelliteParsed struct {

	//DF397
	//The number of integer
	//milliseconds in GNSS
	//Satellite rough ranges
	//RoughRangeInt		specialUint64

	//DF398
	//GNSS Satellite rough ranges modulo 1 millisecond
	//RoughRangeRemainder	uint64

	// Specific for each GNSS
	// DF419 For GLONASS DECODED to INT

	Info				int

	// NEW VALUE As SUM of Int and remainder
	RoughRange			Specialfloat64

	// DF399
	// GNSS Satellite rough PhaseRangeRates CONVERTED from int
	RatePhaseRangeInt	specialInt64

	SatelliteNumber		int

	L1 					Type1087L1
	L2 					Type1087L2

	// Additional
	RealPseudorangeL1	Specialfloat64
	RealPseudorangeL2	Specialfloat64
	RealPhaserangeL1	Specialfloat64
	RealPhaserangeL2	Specialfloat64
}


type Type1004Satellite struct {
	Ident			uint64
	L1				Type1004L1
	L2				Type1004L2
}

type Type1087Satellite struct {
	RoughRangeInt			uint64
	Info					uint64
	RoughRangeRemainder		uint64
	RatePhaseRangeInt		specialInt64
	// External from mask
	SatelliteNumber			int
	Signals					[]Type1087Signal
}

type Type1087Signal struct {
	// GNSS signal fine Pseudoranges with extended resolution
	PseudoRangeCorrection	specialInt64
	PhaseRangeCorrection	specialInt64
	PhaseRangeTI			uint64
	AI 						uint64
	CNR						uint64
	RatePhaseRangeRemainder	specialInt64
	SignalNumber			int

}

type Type1087L1 struct {
	// GNSS signal fine Pseudoranges with extended resolution
	PseudoRangeCorrection	Specialfloat64
	PhaseRangeCorrection	Specialfloat64
	PhaseRangeTI			uint64
	AI 						string
	CNR						float64
	RatePhaseRangeRemainder	Specialfloat64
	// Added from UP
	Signal					string
}

type Type1087L2 struct {
	// GNSS signal fine Pseudoranges with extended resolution
	PseudoRangeCorrection	Specialfloat64
	PhaseRangeCorrection	Specialfloat64
	PhaseRangeTI			uint64
	AI 						string
	CNR						float64
	RatePhaseRangeRemainder	Specialfloat64
	// Added from UP
	Signal					string
}

type Type1012Satellite struct {
	//GLONASS Satellite ID (Satellite Slot Number) DF038 uint6 6
	Ident			uint64
	L1				Type1012L1
	L2				Type1012L2
	//GLONASS Satellite Frequency Channel Number DF040 uint5 5
	Channel			int
}

type Type1012L1 struct {
	//GLONASS L1 Code Indicator DF039 bit(1) 1
	Ind			uint64
	//GLONASS L1 Pseudorange DF041 uint25 25
	Prange		float64
	//GLONASS L1 PhaseRange – L1 Pseudorange DF042 int20 20
	Delta		Specialfloat64
	//GLONASS L1 Lock time Indicator DF043 uint7 7
	Lockt		uint64
	//GLONASS Integer L1 Pseudorange Modulus Ambiguity DF044 uint7 7
	Amb			uint64
	//GLONASS L1 CNR DF045 uint8 8
	CNR			float64
}

type Type1004L1 struct {
	//
	Ind			uint64
	//
	Prange		Specialfloat64
	//
	Delta		Specialfloat64
	//
	Lockt		uint64
	//
	Amb			float64
	//
	CNR			float64
}



type Type1012L2 struct {
	//GLONASS L2 Code Indicator DF046 bit(2) 2
	Ind			uint64
	//GLONASS L2-L1 Pseudorange Difference DF047 uint14 14
	Prange		Specialfloat64
	//GLONASS L2 PhaseRange – L1 Pseudorange DF048 int20 20
	Delta		Specialfloat64
	//GLONASS L2 Lock time Indicator DF049 uint7 7
	Lockt		uint64
	//GLONASS L2 CNR DF050 uint8 8
	CNR			float64
}

type Type1004L2 struct {
	//
	Ind			uint64
	//
	Prange		Specialfloat64
	//
	Delta		Specialfloat64
	//
	Lockt		uint64
	//
	CNR			float64
}
