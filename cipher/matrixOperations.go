package cipher

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math"
)

func mod(a, m int) int {
	a %= m
	if a < 0 {
		a += m
	}
	return a
}

func modInverse(a, m int) (int, error) {
	a = mod(a, m)
	if a == 0 {
		return 0, fmt.Errorf("no inverse for 0")
	}
	t, newT := 0, 1
	r, newR := m, a
	for newR != 0 {
		q := r / newR
		t, newT = newT, t-q*newT
		r, newR = newR, r-q*newR
	}
	if r != 1 {
		return 0, fmt.Errorf("no modular inverse exists (gcd != 1)")
	}
	if t < 0 {
		t += m
	}
	return t, nil
}

func detFloatToInt(d float64) int {
	return int(math.Round(d))
}

func buildMinorDense(src *mat.Dense, rRem, cRem int) *mat.Dense {
	n, _ := src.Dims()
	if n <= 1 {
		return mat.NewDense(0, 0, nil)
	}
	data := make([]float64, 0, (n-1)*(n-1))
	for i := 0; i < n; i++ {
		if i == rRem {
			continue
		}
		for j := 0; j < n; j++ {
			if j == cRem {
				continue
			}
			data = append(data, src.At(i, j))
		}
	}
	return mat.NewDense(n-1, n-1, data)
}

func inverseMatrixMod(A *mat.Dense, modNum int) (*mat.Dense, error) {
	n, mcols := A.Dims()
	if n != mcols {
		return nil, fmt.Errorf("matrix must be square")
	}
	// 1) Determinant of A (using Gonum), round to nearest int, then mod
	detF := mat.Det(A)
	det := detFloatToInt(detF)
	detMod := mod(det, modNum)
	if detMod == 0 {
		return nil, fmt.Errorf("determinant %d ≡ 0 (mod %d): not invertible", det, modNum)
	}

	// 2) modular inverse of determinant
	detInv, err := modInverse(detMod, modNum)
	if err != nil {
		return nil, fmt.Errorf("determinant has no modular inverse: %w", err)
	}

	// 3) Build cofactor matrix (each entry = (-1)^{i+j} * det(minor(i,j)) mod modNum)
	cof := make([][]int, n)
	for i := 0; i < n; i++ {
		cof[i] = make([]int, n)
		for j := 0; j < n; j++ {
			minor := buildMinorDense(A, i, j)
			var minorDetInt int
			mr, _ := minor.Dims()
			if mr == 0 { // 1x1 original -> minor is empty
				minorDetInt = 1
			} else {
				minorDetF := mat.Det(minor)
				minorDetInt = detFloatToInt(minorDetF)
			}
			sign := 1
			if (i+j)%2 != 0 {
				sign = -1
			}
			cof[i][j] = mod(sign*minorDetInt, modNum)
		}
	}

	// 4) Adjugate = transpose of cofactor matrix
	adj := make([][]int, n)
	for i := 0; i < n; i++ {
		adj[i] = make([]int, n)
		for j := 0; j < n; j++ {
			adj[i][j] = cof[j][i]
		}
	}

	// 5) Multiply adjugate by detInv (mod modNum) to get inverse matrix entries
	invInt := make([][]int, n)
	for i := 0; i < n; i++ {
		invInt[i] = make([]int, n)
		for j := 0; j < n; j++ {
			invInt[i][j] = mod(adj[i][j]*detInv, modNum)
		}
	}

	// 6) Convert to *mat.Dense with float64 entries (values are small integers)
	data := make([]float64, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			data = append(data, float64(invInt[i][j]))
		}
	}
	invDense := mat.NewDense(n, n, data)
	return invDense, nil
}
