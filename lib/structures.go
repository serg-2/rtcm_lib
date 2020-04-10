package rtcmlib

type TypeCommonHeader struct {
	// Message Number (“1009”=0011 1111 0001)
	MessageNumber	uint64
}

type Type1012Header struct {
	//Reference Station ID DF003 uint12 12
	StationId		uint64
	//GLONASS Epoch Time (t k ) DF034 uint27 27
	Epoch			uint64
	//Synchronous GNSS Flag DF005 bit(1) 1
	SFlag			uint64
	//No. of GLONASS Satellite Signals Processed DF035 uint5 5
	NoSatellites	uint64
	//GLONASS Divergence-free Smoothing Indicator DF036 bit(1) 1
	DIndicator		uint64
	//GLONASS Smoothing Interval DF037 bit(3) 3
	SmoothInterval	uint64
}

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
	L1dPR		int64
	//GLONASS L1 Lock time Indicator DF043 uint7 7
	L1Lt		uint64
	//GLONASS Integer L1 Pseudorange Modulus Ambiguity DF044 uint7 7
	L1MA		uint64
	//GLONASS L1 CNR DF045 uint8 8
	L1CNR		uint64
	//GLONASS L2 Code Indicator DF046 bit(2) 2
	L2CI		uint64
	//GLONASS L2-L1 Pseudorange Difference DF047 uint14 14
	L2PR		uint64
	//GLONASS L2 PhaseRange – L1 Pseudorange DF048 int20 20
	L2dPR		int64
	//GLONASS L2 Lock time Indicator DF049 uint7 7
	L2Lt		uint64
	//GLONASS L2 CNR DF050 uint8 8
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
	Delta		float64
	//GLONASS L1 Lock time Indicator DF043 uint7 7
	Lockt		uint64
	//GLONASS Integer L1 Pseudorange Modulus Ambiguity DF044 uint7 7
	Amb			float64
	//GLONASS L1 CNR DF045 uint8 8
	CNR			float64
}

type Type1012L2 struct {
	//GLONASS L2 Code Indicator DF046 bit(2) 2
	Ind			uint64
	//GLONASS L2-L1 Pseudorange Difference DF047 uint14 14
	Prange		float64
	//GLONASS L2 PhaseRange – L1 Pseudorange DF048 int20 20
	Delta		float64
	//GLONASS L2 Lock time Indicator DF049 uint7 7
	Lockt		uint64
	//GLONASS L2 CNR DF050 uint8 8
	CNR			float64
}