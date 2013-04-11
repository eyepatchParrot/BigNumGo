package main

func UIntToHexString(val uint32, includeZero bool) (ret string) {
	for i := 0; i < 8; i++ {
		curDigit := byte(7 - i)
		digit := byte((val & (0xF << (4 * curDigit))) >> (4 * curDigit))
		if includeZero || digit != 0 || curDigit == 0 {
			var digitChar byte
			if digit > 9 {
				digitChar = 'A' + (digit - 10)
			} else {
				digitChar = '0' + digit
			}
			ret = ret + string(digitChar)
			includeZero = true
		}
	}
	return
}
