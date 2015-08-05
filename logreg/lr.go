package logreg

import (
  //"fmt"
  "math/rand"
  "math"
)

type LogReg struct {
  theta []float64 // weight
  lr_ float64 // learning rate
  decay float64
  dataset []*RowLabeledDate
  nrow int
  ndim int
}

func NewLogReg(file string) (*LogReg, error) {
  lr := &LogReg{}
  var err error
  if lr.dataset, err = ReadData(file); err != nil {
    return nil, err
  }

  lr.nrow = len(lr.dataset)
  lr.ndim = len(lr.dataset[0].Feat)
  lr.lr_ = 0.001
  lr.decay = 0.5
  lr.theta = make([]float64, lr.ndim)
  return lr, nil
}

func (lr *LogReg) Train() {
  maxItr := 5000

  // Init theta parameter
  for j := 0; j < lr.ndim; j++ {
    // rand.Float64 generates random number in range [0.0, 1.0)
    lr.theta[j] = rand.Float64()
  }

  // Using batch gradient descent to learn the parameters.
  nextTheta := make([]float64, lr.ndim)
  itr := 0
  for itr < maxItr {
    for j := 0; j < lr.ndim; j++ {
      var grad float64 = 0
      for i := 0; i < lr.nrow; i++ { // go through the training set
        grad += (lr.hypoth(lr.dataset[i].Feat) - float64(lr.dataset[i].Label)) * lr.dataset[i].Feat[j]
      }
      grad = 1.0 * grad / float64(lr.nrow)
      grad += lr.decay * lr.theta[j] / float64(lr.nrow)
      nextTheta[j] = lr.theta[j] - lr.lr_ * grad
    }

    for j := 0; j < lr.ndim; j++ {
      lr.theta[j] = nextTheta[j]
    }
    itr++
    //fmt.Printf("itr = %d, cost = %f\n", itr, lr.cost())
  }
}

func (lr *LogReg) Predict(feat []float64) int {
  h := lr.hypoth(feat)
  if h > 0 {
    return 1
  } else {
    return 0
  }
}

func (lr *LogReg) cost() float64 {
  var cost float64 = 0

  var h float64
  for i := 0; i < lr.nrow; i++ {
    h = lr.hypoth(lr.dataset[i].Feat)
    cost += float64(lr.dataset[i].Label) * math.Log(h)
    cost += (1.0 - float64(lr.dataset[i].Label)) * math.Log(1.0 - h)
  }
  cost *= -1.0 / float64(lr.nrow)

  var penalty float64 = 0
  for j := 1; j < lr.ndim; j++ {
    penalty += lr.theta[j] * lr.theta[j]
  }
  penalty *= lr.decay / (2 * float64(lr.nrow))

  return cost + penalty
}

func (lr *LogReg) hypoth(feat []float64) float64 {
  var g float64 = 0
  for j := 0; j < lr.ndim; j++ {
    g += lr.theta[j] * feat[j];
  }
  return 1.0 / (1 + math.Exp(-g))
}
