package uintvar

import "strconv"

// Parse takes in a byte of array and puts out the uinit, the number of bytes used and an error
func Parse(in []byte) (u uint64, n int, err error) {
	values := []byte{}
	cont := true
	c := 0

	for cont {
		values = append(values, in[c])
		cont = hasContinueByte(in[c])
		c++
	}

	n = len(values)

	byteString := ""

	for _, val := range values {
		binVal := strconv.FormatInt(int64(val), 2)
		if len(binVal) == 8 {
			byteString += binVal[1:]
		} else {
			byteString += prefixWithZeros(binVal, 7)
		}
	}

	for i := 0; i < n; i++ {
		byteString = "0" + byteString
	}

	u, err = strconv.ParseUint(byteString, 2, 64)

	return
}

func hasContinueByte(in byte) bool {
	in &= 0x80
	return in == 0x80
}

func prefixWithZeros(s string, l int) string {
	toAdd := l - len(s)
	for i := 0; i < toAdd; i++ {
		s = "0" + s
	}

	return s
}
