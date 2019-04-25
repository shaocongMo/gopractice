package m_math

import "errors"

func Sum(a int, b int) int {
	return a + b
}

func Divisor(dividend int, divisor int) (result int, err error) {
	if divisor == 0 {
		err = errors.New("division by zero")
		return
	}
	result = dividend / divisor
	return
}
