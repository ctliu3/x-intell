package ap

import (
  "math"
  "fmt"
  "sort"
)

// Affinity Propagtaion
type AP struct {
  r [][]float64 // respsonsibility
  a [][]float64 // availabity
  s [][]float64 // similarity
  damping float64 // damping factor, range in [0.5, 1)
}

func (ap *AP) cluster(data []*Feat) []int {
  ap.damping = 0.9
  var n int = len(data)
  var maxIter int = 500

  Init2DArray(&ap.s, n)
  Init2DArray(&ap.r, n)
  Init2DArray(&ap.a, n)

  sim := make([]float64, n * (n - 1) / 2)
  c := 0
  for i := range ap.s {
    ap.s[i][i] = 0.0
    for j := range ap.s[i] {
      ap.a[i][j] = 0.0
      ap.r[i][j] = 0.0
      if i < j {
        ap.s[i][j] = -EuclideanDist(data[i], data[j])
        ap.s[j][i] = ap.s[i][j]
        sim[c] = ap.s[i][j]
        c++
      }
    }
  }

  sort.Float64s(sim)
  nsim := len(sim)
  median := sim[int(nsim / 2)]
  if nsim % 2 == 0 {
    median = (median + sim[int(nsim / 2) - 1]) / 2
  }
  fmt.Println("meadian = ", median)

  for i := 0; i < n; i++ {
    ap.s[i][i] = median
  }

  var itr int = 0
  var nr [][]float64 // used for keep the info for the next iteration.
  var na [][]float64
  Init2DArray(&nr, n)
  Init2DArray(&na, n)

  for itr < maxIter {
    // Update r.
    for i := 0; i < n; i++ {
      for k := 0; k < n; k++ {
        maxval := math.Inf(-1)
        for j := 0; j < n; j++ {
          if j != k {
            maxval = math.Max(maxval, ap.a[i][j] + ap.s[i][j])
          }
        }
        nr[i][k] = ap.s[i][k] - maxval
      }
    }

    // Update a.
    for i := 0; i < n; i++ {
      for k := 0; k < n; k++ {
        if i != k {
          na[i][k] = ap.r[k][k]
          for j := 0; j < n; j++ {
            if j != i && j != k {
              na[i][k] += math.Max(0.0, ap.r[j][k])
            }
          }
          na[i][k] = math.Min(0.0, na[i][k])
        }
      }
    }
    for k := 0; k < n; k++ {
      na[k][k] = 0
      for j := 0; j < n; j++ {
        if j != k {
          na[k][k] += math.Max(0.0, ap.r[j][k])
        }
      }
    }

    for i := 0; i < n; i++ {
      for k := 0; k < n; k++ {
        ap.r[i][k] = (1.0 - ap.damping) * nr[i][k] + ap.damping * ap.r[i][k]
        ap.a[i][k] = (1.0 - ap.damping) * na[i][k] + ap.damping * ap.a[i][k]
      }
    }
    itr++
  }

  var centers []int
  for i := 0; i < n; i++ {
    var val float64 = math.Inf(-1)
    var idx int = 0
    for k := 0; k < n; k++ {
      if ap.a[i][k] + ap.r[i][k] > val {
        val = ap.a[i][k] + ap.r[i][k]
        idx = k
      }
    }
    var isCenter bool = true
    for _, center := range centers {
      if center == idx {
        isCenter = false
        break
      }
    }
    if isCenter {
      centers = append(centers, idx)
    }
  }

  return centers
}
