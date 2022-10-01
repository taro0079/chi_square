package chi_distribution

import "math"

func ChiDist(dof float64, x float64) float64 {
	return 1 / (math.Pow(2.0, dof/2.0) * math.Gamma(dof/2.0)) * math.Pow(x, dof/2.0-1.0) * math.Pow(math.E, -x/2.0)
}

func xGen(start float64, end float64, num int) []float64 {
	result := make([]float64, num)
	dx := (end - start) / float64(num-1)
	for i := 1; i < num; i++ {
		result[i] = result[i-1] + dx
	}
	result[0] = start
	result[num-1] = end

	return result
}
