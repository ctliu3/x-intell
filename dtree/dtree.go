package dtree

import (
  "math"
  "fmt"
)

type TreeNode struct {
  node []*TreeNode
  split int
  label string
}

type Feat struct {
  feat []string
  label string
}

// Decision Tree
type DTree struct {
  root *TreeNode
  data []*Feat
}

func NewDTree() (dtree *DTree) {
  tree := new(DTree)
  return tree
}

func (dt *DTree) Create(data []*Feat) {
  fmt.Println("Creating the decision tree...")
  dt.root = new(TreeNode)
  create(dt.root, data)
}

//func (dt *DecisionTree) Predict(feat *Feat) (label string) {
//}

// Split the tree using ID3 algo.
func create(cur *TreeNode, data []*Feat) {
  // If the samples have the same labels.
  if isSameLabel(data) {
    cur.label = data[0].label
    return
  }
  // No property to split, vote for the label
  if len(data[0].feat) == 0 {
    cur.label = vote(data)
    return
  }

  fmt.Println("Get the info gain.")
  labels := make([]string, 0)
  for _, feat := range data {
    labels = append(labels, feat.label)
  }
  maxInfoGain := math.Inf(-1)
  baseEntropy := entropy(labels)
  nfeat := len(data[0].feat)
  idx := 0
  for j := 0; j < nfeat; j++ {
    // ig := infoGain(i, data)
    ig := baseEntropy - splitEntropy(j, data)
    if ig > maxInfoGain {
      maxInfoGain = ig
      idx = j
    }
  }
  fmt.Printf("idx = %d\n", idx)

  // Create the decision recursively
  cur.split = idx
  cur.node = make([]*TreeNode, len(labels))
  subtree := make(map[string][]*Feat)
  for _, item := range data {
    x := item.feat[idx]
    _, ok := subtree[x]
    if ok == false {
      subtree[x] = make([]*Feat, 0)
    }
    newitem := new(Feat)
    newitem.label = item.label
    for j := 0; j < nfeat; j++ {
      if j == idx {
        continue
      }
      newitem.feat = append(newitem.feat, item.feat[j])
    }
    subtree[x] = append(subtree[x], newitem)
  }

  fmt.Printf("\nThe following branches are: \n[ ")
  for x, _ := range subtree {
    fmt.Printf("%s ", x)
  }
  fmt.Printf("]\n")

  branch := 0
  for _, subdata := range subtree {
    cur.node[branch] = new(TreeNode)
    create(cur.node[branch], subdata)
    branch += 1
  }
  fmt.Printf("\n#Branch = %d\n", branch)
}

func isSameLabel(data []*Feat) (isSame bool) {
  for _, item := range data {
    if item.label != data[0].label {
      return false
    }
  }
  return true
}

func vote(data []*Feat) (label string) {
  c := make(map[string]int)
  for _, item := range data {
    count, ok := c[item.label]
    if ok {
      c[item.label] = count + 1
    } else {
      c[item.label] = 1
    }
  }

  n := 0
  var maxlabel string
  for label, count := range c {
    if count > n {
      maxlabel = label
      n = count
    }
  }

  return maxlabel
}

func entropy(labels []string) (entropy float64) {
  c := make(map[string]int)

  for _, label := range labels {
    count, ok := c[label]
    if ok {
      c[label] = count + 1
    } else {
      c[label] = 1
    }
  }

  entro := 0.0
  n := len(labels)
  for _, count := range c {
    p := float64(count) / float64(n)
    entro -= p * math.Log2(p)
  }

  return entro
}

func splitEntropy(j int, data []*Feat) (sum float64) {
  mpLabels := make(map[string][]string)

  for _, item := range data {
    x := item.feat[j]
    _, ok := mpLabels[x]
    if ok == false {
      mpLabels[x] = make([]string, 0)
    }
    mpLabels[x] = append(mpLabels[x], item.label)
  }

  entropySum := 0.0
  for _, arr := range mpLabels {
    entropySum += (float64(len(arr)) / float64(len(data))) * entropy(arr)
  }

  return entropySum
}
