package leap

const testVersion = 3

func IsLeapYear(year int) bool {

	isLeapYear := (year % 4) == 0

	isLeapYear = isLeapYear && (year % 100) != 0

	return isLeapYear || (year % 400) == 0
}
