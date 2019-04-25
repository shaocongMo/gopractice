package m_math_test

import (
	m_math "practice/go2"
	"testing"
)

func TestSum(t *testing.T) {
	sum := m_math.Sum(1, 1)
	if sum != 2 {
		t.Logf("testsum m_math.Sum(1, 1) return %d", sum)
		t.Fail()
	}
}

func BenchmarkSum(t *testing.B) {
	sum := m_math.Sum(1, 1)
	if sum != 2 {
		t.Logf("benchmarksum m_math.Sum(1, 1) return %d", sum)
		t.Fail()
	}
}

func TestDivisor1(t *testing.T) {
	result, err := m_math.Divisor(6, 2)
	if err != nil {
		t.Logf("error: %s\n", err)
		t.Fail()
	} else {
		if result != 3 {
			t.Errorf("6 divide 2 return %d", result)
		}
	}
}

func TestDivisor2(t *testing.T) {
	result, err := m_math.Divisor(6, 0)
	if err == nil {
		t.Errorf("6 divide 0 return %d", result)
	}
}
