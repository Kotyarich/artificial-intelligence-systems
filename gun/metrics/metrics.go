package metrics

import (
	"ais.com/m/model"
	"math"
)

func Euclidean(gun1, gun2 model.Gun, weights []float32) float64 {
	vector1, vector2 := GunsToVectors(gun1, gun2)

	var sum float64
	for i := range vector1 {
		sum += math.Pow(float64((vector1[i]-vector2[i]) * weights[i]), 2)
	}

	return math.Sqrt(sum)
}

func L1Distance(gun1, gun2 model.Gun, weights []float32) float64 {
	vector1, vector2 := GunsToVectors(gun1, gun2)

	var sum float64
	for i := range vector1 {
		sum += math.Abs(float64((vector1[i]-vector2[i]) * weights[i]))
	}

	return sum
}

func mean(vector []float32, weights []float32) float64 {
	s := 0.
	sw := 0.

	for i, v := range vector {
		s += float64(v) * float64(weights[i])
		sw += float64(weights[i])
	}

	return s / sw
}

func PearsonCorrelation(gun1, gun2 model.Gun, weights []float32) float64 {
	vector1, vector2 := GunsToVectors(gun1, gun2)

	xu := mean(vector1, weights)
	yu := mean(vector2, weights)

	var (
		s             float64
		sxx           float64
		syy           float64
	)

	for i, xv := range vector1 {
		yv := vector2[i]
		xd := float64(xv * weights[i]) - xu
		yd := float64(yv * weights[i]) - yu
		sxx += xd * xd
		syy += yd * yd
		s += xd * yd
	}

	return s / math.Sqrt(sxx*syy)
}

func TreeMetric(gun1, gun2 model.Gun) float64 {
	pathSum1 := TreePathSum(gun1)
	pathSum2 := TreePathSum(gun2)

	return math.Abs(pathSum2 - pathSum1)
}