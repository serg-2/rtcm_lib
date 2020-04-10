package rtcmlib

import (
	"fmt"
	"log"
	"strconv"
)

//=====================================
func DUint (field string, razr int) uint64 {
	if len(field) != razr {
		log.Fatalf("WRONG INFO FOR UINT decoding bits: %v\n", razr)
	}
	result, err := strconv.ParseUint(field, 2, 64)
	Chk(err)
	return result
}

func DInt(field string, razr int) int64 {
	var result int64
	if len(field) != razr {
		log.Fatalf("WRONG INFO FOR INT decoding bits: %v\n", razr)
	}
	_result, err := strconv.ParseUint(field[1:], 2, 64)
	Chk(err)

	if string(field[0]) == "1" {
		result = - int64(_result)
	} else {
		result = int64(_result)
	}
	return result
}

//=====================================
func EUint (num uint64, numbits int) string {
	result:=fmt.Sprintf("%0" + fmt.Sprintf("%d", numbits) + "b", num)
	if len(result) > numbits {
		log.Fatal("Wrong info for encode Uint %v \n", numbits)
	}
	return result
}

func EInt (num int64, numbits int) string {
	var result string
	if num < 0 {
		result = "1" + fmt.Sprintf("%0" + fmt.Sprintf("%d", numbits-1) + "b", -num)
	} else {
		result = fmt.Sprintf("%0"+fmt.Sprintf("%d", numbits)+"b", num)
	}
	if len(result) > numbits {
		log.Fatal("Wrong info for encode Uint %v \n", numbits)
	}
	return result
}


//===========1012 HEADER==========
func D_DF002 (field string) uint64 {
	// 12bits
	return DUint(field, 12)
}

func E_DF002 (num uint64) string {
	// 12bits
	return EUint(num, 12)
}

func D_DF003 (field string) uint64 {
	// 12bits
	return DUint(field, 12)
}

func E_DF003 (num uint64) string {
	// 12bits
	return EUint(num, 12)
}

func D_DF034 (field string) uint64 {
	// 27bits
	return DUint(field, 27)
}

func E_DF034 (num uint64) string {
	// 27bits
	return EUint(num, 27)
}

func D_DF005 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF005 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF035 (field string) uint64 {
	// 5bits
	return DUint(field, 5)
}

func E_DF035 (num uint64) string {
	// 5bits
	return EUint(num, 5)
}

func D_DF036 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF036 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF037 (field string) uint64 {
	// 3bit
	return DUint(field, 3)
}

func E_DF037 (num uint64) string {
	// 3bit
	return EUint(num, 3)
}

//===========1012 Body============

func D_DF038 (field string) uint64 {
	// 6bits
	// 0 – The slot number is unknown
	// 1 to 24 – Slot number of the GLONASS satellite
	// >32 – Reserved for Satellite-Based Augmentation Systems (SBAS)
	// 40-58 for SBAS (Add 80 for SBAS PRN)
	return DUint(field, 6)
}

func E_DF038 (num uint64) string {
	// 6bits
	return EUint(num, 6)
}

func D_DF039 (field string) uint64 {
	//return D1bit(field)
	return DUint(field, 1)
}

func E_DF039 (num uint64) string {
	return EUint(num, 1)
}

func D_DF040 (field string) uint64 {
	// 5bits
	return DUint(field, 5)
}

func E_DF040 (num uint64) string {
	// 5bits
	return EUint(num, 5)
}

func D_DF041 (field string) uint64 {
	// 25bits
	return DUint(field, 25)
}

func E_DF041 (num uint64) string {
	// 25bits
	return EUint(num, 25)
}

func D_DF042 (field string) int64 {
	// 20bits
	return DInt(field, 20)
}

func E_DF042 (num int64) string {
	// 20bits
	return EInt(num, 20)
}

func D_DF043 (field string) uint64 {
	// 7bits
	return DUint(field, 7)
}

func E_DF043 (num uint64) string {
	// 7bits
	return EUint(num, 7)
}

func D_DF044 (field string) uint64 {
	// 7bits
	return DUint(field, 7)
}

func E_DF044 (num uint64) string {
	// 7bits
	return EUint(num, 7)
}

func D_DF045 (field string) uint64 {
	// 8bits
	return DUint(field, 8)
}

func E_DF045 (num uint64) string {
	// 8bits
	return EUint(num, 8)
}

func D_DF046 (field string) uint64 {
	// 2bits
	return DUint(field, 2)
}

func E_DF046 (num uint64) string {
	// 2bits
	return EUint(num, 2)
}

func D_DF047 (field string) uint64 {
	// 14bits
	return DUint(field, 14)
}

func E_DF047 (num uint64) string {
	// 14bits
	return EUint(num, 14)
}

func D_DF048 (field string) int64 {
	// 20bits
	return DInt(field, 20)
}

func E_DF048 (num int64) string {
	// 20bits
	return EInt(num, 20)
}

func D_DF049 (field string) uint64 {
	// 7bits
	return DUint(field, 7)
}

func E_DF049 (num uint64) string {
	// 7bits
	return EUint(num, 7)
}

func D_DF050 (field string) uint64 {
	// 8bits
	return DUint(field, 8)
}

func E_DF050 (num uint64) string {
	// 8bits
	return EUint(num, 8)
}

