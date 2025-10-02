package cipher

import "gonum.org/v1/gonum/mat"

func encrypt(key [][]int, plaintext []int) []int {
	blockLen := len(key)
	textLen := len(plaintext)
	chunks := (blockLen / textLen) + 1 //go rounds down so we need to plus 1
	keyMat := mat.NewDense(blockLen, blockLen, flattenIntMatrixToFloat64(key))

	for blockIndex := 0; blockIndex < chunks; blockIndex++ {
		var dataBlock []int
		if blockIndex == chunks {
			start := blockIndex * blockLen
			finish := textLen + 1
			dataBlock = plaintext[start:finish]
		} else {
			start := blockIndex * blockLen
			finish := (blockIndex + 1) * blockLen
			dataBlock = plaintext[start:finish]
		}

		blockMat := mat.NewDense(blockLen, 1, intSliceToFloat64(dataBlock))
		blockResultMat := mat.NewDense(blockLen, 1, nil)
		blockResultMat.Mul(keyMat, blockMat)
		
	}
}
