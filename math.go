/*
@Time : 2020/12/17 17:47
@Author : xuhanlin
@File : math.go
@Description :
*/

package util

// CeilToPowerOfTwo返回大于等于n的最小2的倍数
func CeilToPowerOfTwo(n int) int {
	if n&MaxIntHeadBit != 0 && n > MaxIntHeadBit {
		panic("argument is too large")
	}
	if n <= 2 {
		return 2
	}
	n--
	n = fillBits(n)
	n++
	return n
}

// fillBits将最高位1以下的0全部填充为1
func fillBits(n int) int {
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	return n
}