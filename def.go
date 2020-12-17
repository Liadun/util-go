/*
@Time : 2020/12/17 20:41
@Author : xuhanlin
@File : def.go
@Description :
*/

package util

const (
	BitSize = 32 << (^uint(0) >> 63)
	MaxIntHeadBit = 1 << (BitSize - 2)
	IntMax = int(^uint(0) >> 1)
	IntMin = ^IntMax
	UIntMax = ^uint(0)
	UIntMin = 0
)
