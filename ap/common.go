package ap

import (
  "math"
)

type Feat struct {
  x []float64
}

func EuclideanDist(feat1 *Feat, feat2 *Feat) float64 {
  n := len(feat1.x)
  var dist float64 = 0
  for i := 0; i < n; i++ {
    dist += (feat1.x[i] - feat2.x[i]) * (feat1.x[i] - feat2.x[i])
  }
  return math.Sqrt(dist)
}

func Init2DArray(a *[][]float64, n int) {
  *a = make([][]float64, n)
  for i := 0; i < n; i ++ {
    (*a)[i] = make([]float64, n)
  }
}
