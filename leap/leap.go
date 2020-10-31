package leap

// IsLeapYear check if a year is divisible by 4 or by 4, 100 and 400
func IsLeapYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || (year%4 == 0 && year%100 == 0 && year%400 == 0) {
		return true
	}
	return false
}
