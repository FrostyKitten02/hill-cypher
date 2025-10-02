package cipher

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func Decrypt(key [][]int, plaintext string) string {
	result := decrypt(key, StrToInt(plaintext, false))
	return IntsToStr(result, true)
}

func decrypt(key [][]int, plaintext []int) []int {
	blockLen := len(key)
	textLen := len(plaintext)
	chunks := textLen / blockLen
	originalKeyMat := mat.NewDense(blockLen, blockLen, flattenIntMatrixToFloat64(key))

	keyMat, keyErr := inverseMatrixMod(originalKeyMat, CHARSET)
	if keyErr != nil {
		fmt.Println("Error in key matrix compute")
		panic(keyErr)
	}

	result := make([]int, chunks*blockLen)
	for blockIndex := 0; blockIndex < chunks; blockIndex++ {
		var dataBlock []int
		start := blockIndex * blockLen
		finish := (blockIndex + 1) * blockLen
		dataBlock = plaintext[start:finish]

		blockMat := mat.NewDense(blockLen, 1, intSliceToFloat64(dataBlock))
		blockResultMat := mat.NewDense(blockLen, 1, nil)
		blockResultMat.Mul(keyMat, blockMat)
		blockResultArr := matToSlice(blockResultMat)
		if blockIndex == chunks {

		}
		applyMod(blockResultArr, CHARSET)
		copy(result[start:], blockResultArr)
	}

	return result
}
