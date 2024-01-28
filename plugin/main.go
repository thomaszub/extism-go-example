package main

import (
	"encoding/binary"
	"math"

	"github.com/extism/go-pdk"
	"github.com/thomaszub/extism-tg/plugin/calc"
)

//export mean
func mean() int32 {
	nums := bytesToFloats(pdk.Input())
	std := calc.Mean(nums)
	bytes := floatToBytes(std)
	pdk.Output(bytes)
	return 0
}

//export stdDev
func stdDev() int32 {
	nums := bytesToFloats(pdk.Input())
	std := calc.StdDev(nums)
	bytes := floatToBytes(std)
	pdk.Output(bytes)
	return 0
}

func bytesToFloats(input []byte) []float64 {
	num := len(input) / 8
	nums := make([]float64, num)
	for i := 0; i < num; i++ {
		bytes := input[8*i : (8 * (i + 1))]
		bits := binary.LittleEndian.Uint64(bytes)
		float := math.Float64frombits(bits)
		nums[i] = float
	}
	return nums
}

func floatToBytes(num float64) []byte {
	bytes := make([]byte, 8)
	bits := math.Float64bits(num)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func main() {}
