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

func DInt(field string, razr int) (int64, bool) {
	// return bool as special value
	var result int64
	if len(field) != razr {
		log.Fatalf("WRONG INFO FOR INT decoding bits: %v\n", razr)
	}
	_result, err := strconv.ParseUint(field[1:], 2, 64)
	Chk(err)

	if string(field[0]) == "1" {
		// Check for special value
		// -0 - special value
		if _result == 0 {
			return 0, true
		}
		result = - int64(_result)
	} else {
		result = int64(_result)
	}
	return result, false
}

//=====================================
func EUint (num uint64, numbits int) string {
	result:=fmt.Sprintf("%0" + fmt.Sprintf("%d", numbits) + "b", num)
	if len(result) > numbits {
		log.Fatal("Wrong info for encode Uint %v \n", numbits)
	}
	return result
}

func EInt (num int64, numbits int, special bool) string {
	var result string
	if special {
		result = "1" + fmt.Sprintf("%0"+fmt.Sprintf("%d",numbits-1) +"b",0)
	} else {
		if num < 0 {
			result = "1" + fmt.Sprintf("%0"+fmt.Sprintf("%d", numbits-1)+"b", -num)
		} else {
			result = fmt.Sprintf("%0"+fmt.Sprintf("%d", numbits)+"b", num)
		}
		if len(result) > numbits {
			log.Fatal("Wrong info for encode Uint %v \n", numbits)
		}
	}
	return result
}


//===========1012 HEADER==========
func D_DF001 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF001 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

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

func D_DF004 (field string) uint64 {
	// 30bits
	return DUint(field, 30)
}

func E_DF004 (num uint64) string {
	// 30bits
	return EUint(num, 30)
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

func D_DF006 (field string) uint64 {
	// 5bits
	return DUint(field, 5)
}

func E_DF006 (num uint64) string {
	// 5bits
	return EUint(num, 5)
}

func D_DF007 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF007 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF009 (field string) uint64 {
	// 6bit
	return DUint(field, 6)
}

func E_DF009 (num uint64) string {
	// 6bit
	return EUint(num, 6)
}

func D_DF010 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF010 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF011 (field string) uint64 {
	// 24bit
	return DUint(field, 24)
}

func E_DF011 (num uint64) string {
	// 24bit
	return EUint(num, 24)
}

func D_DF021 (field string) uint64 {
	// 6bit
	return DUint(field, 6)
}

func E_DF021 (num uint64) string {
	// 6bit
	return EUint(num, 6)
}

func D_DF022 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF022 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF023 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF023 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF024 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF024 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF025 (field string) specialInt64 {
	// 38bits
	result, special :=  DInt(field, 38)
	return specialInt64{result,special}
}

func E_DF025 (num specialInt64) string {
	// 38bits
	return EInt(num.Value, 38, num.Special)
}

func D_DF026 (field string) specialInt64 {
	// 38bits
	result, special :=  DInt(field, 38)
	return specialInt64{result,special}
}

func E_DF026 (num specialInt64) string {
	// 38bits
	return EInt(num.Value, 38, num.Special)
}

func D_DF027 (field string) specialInt64 {
	// 38bits
	result, special :=  DInt(field, 38)
	return specialInt64{result,special}
}

func E_DF027 (num specialInt64) string {
	// 38bits
	return EInt(num.Value, 38, num.Special)
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

func D_DF008 (field string) uint64 {
	// 3bit
	return DUint(field, 3)
}

func E_DF008 (num uint64) string {
	// 3bit
	return EUint(num, 3)
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

func D_DF042 (field string) specialInt64 {
	// 20bits
	result, special :=  DInt(field, 20)
	return specialInt64{result,special}
}

func E_DF042 (num specialInt64) string {
	// 20bits
	return EInt(num.Value, 20, num.Special)
}

func D_DF012 (field string) specialInt64 {
	// 20bits
	result, special :=  DInt(field, 20)
	return specialInt64{result,special}
}

func E_DF012 (num specialInt64) string {
	// 20bits
	return EInt(num.Value, 20, num.Special)
}

func D_DF043 (field string) uint64 {
	// 7bits
	return DUint(field, 7)
}

func E_DF043 (num uint64) string {
	// 7bits
	return EUint(num, 7)
}

func D_DF013 (field string) uint64 {
	// 7bits
	return DUint(field, 7)
}

func E_DF013 (num uint64) string {
	// 7bits
	return EUint(num, 7)
}

func D_DF014 (field string) uint64 {
	// 8bits
	return DUint(field, 8)
}

func E_DF014 (num uint64) string {
	// 8bits
	return EUint(num, 8)
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

func D_DF015 (field string) uint64 {
	// 8bits
	return DUint(field, 8)
}

func E_DF015 (num uint64) string {
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

func D_DF364 (field string) uint64 {
	// 2bits
	return DUint(field, 2)
}

func E_DF364 (num uint64) string {
	// 2bits
	return EUint(num, 2)
}

func D_DF016 (field string) uint64 {
	// 2bits
	return DUint(field, 2)
}

func E_DF016 (num uint64) string {
	// 2bits
	return EUint(num, 2)
}

func D_DF047 (field string) specialInt64 {
	// 14bits
	result, special :=  DInt(field, 14)
	return specialInt64{result,special}
}

func E_DF047 (num specialInt64) string {
	// 14bits
	return EInt(num.Value, 14, num.Special)
}

func D_DF017 (field string) specialInt64 {
	// 14bits
	result, special :=  DInt(field, 14)
	return specialInt64{result,special}
}

func E_DF017 (num specialInt64) string {
	// 14bits
	return EInt(num.Value, 14, num.Special)
}

func D_DF141 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF141 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF142 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF142 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF048 (field string) specialInt64 {
	// 20bits
	result, special :=  DInt(field, 20)
	return specialInt64{result,special}
}

func E_DF048 (num specialInt64) string {
	// 20bits
	return EInt(num.Value, 20, num.Special)
}

func D_DF018 (field string) specialInt64 {
	// 20bits
	result, special :=  DInt(field, 20)
	return specialInt64{result,special}
}

func E_DF018 (num specialInt64) string {
	// 20bits
	return EInt(num.Value, 20, num.Special)
}

func D_DF049 (field string) uint64 {
	// 7bits
	return DUint(field, 7)
}

func E_DF049 (num uint64) string {
	// 7bits
	return EUint(num, 7)
}

func D_DF019 (field string) uint64 {
	// 7bits
	return DUint(field, 7)
}

func E_DF019 (num uint64) string {
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

func D_DF020 (field string) uint64 {
	// 8bits
	return DUint(field, 8)
}

func E_DF020 (num uint64) string {
	// 8bits
	return EUint(num, 8)
}

func D_DF393 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF393 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}

func D_DF394 (field string) uint64 {
	// 64bit
	return DUint(field, 64)
}

func E_DF394 (num uint64) string {
	// 64bit
	return EUint(num, 64)
}

func D_DF395 (field string) uint64 {
	// 32bit
	return DUint(field, 32)
}

func E_DF395 (num uint64) string {
	// 32bit
	return EUint(num, 32)
}

func D_DF396 (field string) uint64 {
	// NSat*Nsig bits
	return DUint(field, len(field))
}

func E_DF396 (num uint64, tableSize int) string {
	// NSat*Nsig bits
	return EUint(num, tableSize)
}

func D_DF397 (field string) uint64 {
	// 8bit
	return DUint(field, 8)
}

func E_DF397 (num uint64) string {
	// 8bit
	return EUint(num, 8)
}

func D_DF398 (field string) uint64 {
	// 10bit
	return DUint(field, 10)
}

func E_DF398 (num uint64) string {
	// 10bit
	return EUint(num, 10)
}

func D_DF399 (field string) specialInt64 {
	// 14bits
	result, special :=  DInt(field, 14)
	return specialInt64{result,special}
}

func E_DF399 (num specialInt64) string {
	// 14bits
	return EInt(num.Value, 14, num.Special)
}

func D_DF404 (field string) specialInt64 {
	// 15bits
	result, special :=  DInt(field, 15)
	return specialInt64{result,special}
}

func E_DF404 (num specialInt64) string {
	// 15bits
	return EInt(num.Value, 15, num.Special)
}


func D_DF405 (field string) specialInt64 {
	// 20bits
	result, special :=  DInt(field, 20)
	return specialInt64{result,special}
}

func E_DF405 (num specialInt64) string {
	// 20bits
	return EInt(num.Value, 20, num.Special)
}

func D_DF406 (field string) specialInt64 {
	// 24bits
	result, special :=  DInt(field, 24)
	return specialInt64{result,special}
}

func E_DF406 (num specialInt64) string {
	// 24bits
	return EInt(num.Value, 24, num.Special)
}

func D_DF407 (field string) uint64 {
	// 10bit
	return DUint(field, 10)
}

func E_DF407 (num uint64) string {
	// 10bit
	return EUint(num, 10)
}

func D_DF408 (field string) uint64 {
	// 10bit
	return DUint(field, 10)
}

func E_DF408 (num uint64) string {
	// 10bit
	return EUint(num, 10)
}


func D_DF409 (field string) uint64 {
	// 3bit
	return DUint(field, 3)
}

func E_DF409 (num uint64) string {
	// 3bit
	return EUint(num, 3)
}

func D_DF411 (field string) uint64 {
	// 2bit
	return DUint(field, 2)
}

func E_DF411 (num uint64) string {
	// 2bit
	return EUint(num, 2)
}

func D_DF412 (field string) uint64 {
	// 2bit
	return DUint(field, 2)
}

func E_DF412 (num uint64) string {
	// 2bit
	return EUint(num, 2)
}

func D_DF416 (field string) uint64 {
	// 3bit
	return DUint(field, 3)
}

func E_DF416 (num uint64) string {
	// 3bit
	return EUint(num, 3)
}

func D_DF417 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF417 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}
func D_DF418 (field string) uint64 {
	// 3bit
	return DUint(field, 3)
}

func E_DF418 (num uint64) string {
	// 3bit
	return EUint(num, 3)
}

func D_DF419 (field string) uint64 {
	// 4bit
	return DUint(field, 4)
}

func E_DF419 (num uint64) string {
	// 4bit
	return EUint(num, 4)
}

func D_DF420 (field string) uint64 {
	// 1bit
	return DUint(field, 1)
}

func E_DF420 (num uint64) string {
	// 1bit
	return EUint(num, 1)
}
