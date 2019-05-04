package leetcode

import "testing"

func Test890(t *testing.T) {
	s := "3[a]2[bc]"
	t.Logf("Input %s Output: %s", s, decodeString(s))

	s2 := "3[a2[c]]"
	t.Logf("Input %s Output: %s", s2, decodeString(s2))

	s3 := "2[abc]3[cd]ef"
	t.Logf("Input %s Output: %s", s3, decodeString(s3))

	s4 := ""
	t.Logf("Input %s Output: %s", s4, decodeString(s4))

	s5 := "10[a]"
	t.Logf("Input %s Output: %s", s5, decodeString(s5))

	s6 := "10[a21[b]10[c10[d]]]"
	t.Logf("Input %s Output: %s", s6, decodeString(s6))
}
