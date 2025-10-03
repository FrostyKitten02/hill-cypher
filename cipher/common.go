package cipher

import (
	"gonum.org/v1/gonum/mat"
	"strings"
)

// 257 for ASCII support, 1 more to get prime number of possible chars and this allows us to use a filler char
const CHARSET = 389                         //number of all characters
const MAX_CHAR_VAL = CHARSET - 1            //last valid char value, also used for filler char
const MAX_VALID_CHAR_VAL = MAX_CHAR_VAL - 1 //last valid char we can use from input string, ASCII

func intMatrixToFloat64(intMatrix [][]int) [][]float64 {
	floatMatrix := make([][]float64, len(intMatrix))

	for i, row := range intMatrix {
		floatMatrix[i] = make([]float64, len(row))
		for j, val := range row {
			floatMatrix[i][j] = float64(val)
		}
	}

	return floatMatrix
}

func flattenIntMatrixToFloat64(intMatrix [][]int) []float64 {
	var flat []float64

	for _, row := range intMatrix {
		for _, val := range row {
			flat = append(flat, float64(val))
		}
	}

	return flat
}

func intSliceToFloat64(ints []int) []float64 {
	floats := make([]float64, len(ints))
	for i, v := range ints {
		floats[i] = float64(v)
	}
	return floats
}

func applyMod(vals []int, mod int) {
	for i := 0; i < len(vals); i++ {
		vals[i] = vals[i] % mod
	}
}

// for matrices with one column
func matToSlice(dense *mat.Dense) []int {
	r, c := dense.Dims()
	if c > 1 {
		panic("too many dimensions")
	}

	result := make([]int, r)
	for i := 0; i < r; i++ {
		result[i] = int(dense.At(i, 0))
	}

	return result
}

func StrToInt(str string, clean bool) []int {
	result := make([]int, 0)
	for _, r := range str {
		//clean string of invalid characters
		if clean && int(r) > MAX_CHAR_VAL {
			continue
		}
		result = append(result, int(r))
	}

	return result
}

func IntsToStr(ints []int, clean bool) string {
	var builder strings.Builder
	for _, val := range ints {
		if clean && val >= MAX_CHAR_VAL {
			continue
		}
		builder.WriteRune(rune(val))
	}

	return builder.String()
}

func fillSlice(slice []int, val int) {
	for i := 0; i < len(slice); i++ {
		slice[i] = val
	}
}
