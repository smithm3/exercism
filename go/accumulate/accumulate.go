package accumulate

const testVersion = 1

func Accumulate(strArr []string, operation func(string) string) []string {

	for index, element := range strArr {
		strArr[index] = operation(element)
	}

	return strArr
}