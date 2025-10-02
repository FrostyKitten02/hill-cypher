package cipher

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
