package chi_distribution

import "math"

func ChiDist(dof float64, x float64) float64 {
	return 1 / (math.Pow(2.0, dof/2.0) * math.Gamma(dof/2.0)) * math.Pow(x, dof/2.0-1.0) * math.Pow(math.E, -x/2.0)
}
