package encoder

type _Digit struct {
	len   uint8
	digit [7]byte
}

const (
	_MIN_NUMBER = -99_999
	_MAX_NUMBER = 100_000
	_TAB_OFFSET = 99_999
	_TAB_LENGTH = _MAX_NUMBER - _MIN_NUMBER + 1
	_MAX_DIGITS = 6
)

// _DigitTab is a 1.5 MB table cache for -99_999 ~ +99_999.
var _DigitTab [_TAB_LENGTH]_Digit

func init() {
	count := func(number int) int {
		if number == 0 {
			return 1
		}
		cnt := 0
		for number > 0 {
			cnt++
			number /= 10
		}
		return cnt
	}

	for number := _MIN_NUMBER; number <= _MAX_NUMBER; number++ {
		abs := number
		sgn := 0
		if number < 0 {
			abs = -number
			sgn = 1
		}

		cnt := count(abs) + sgn
		if cnt > _MAX_DIGITS {
			panic("number is too big")
		}

		entry    := &_DigitTab[number + _TAB_OFFSET]
		entry.len = uint8(cnt)
		digits   := &entry.digit
		digits[0] = '-'
	
		// write digits from right to left
		for i := cnt - 1; i >= sgn; i-- {
			digits[i] = byte((abs % 10) + int('0'))
			abs      /= 10
		}
	}
}