package dtree

import (
  "fmt"
  "os"
  "bufio"
  "strings"

  "testing"
)

func TestDTree(t *testing.T) {
  filename := "data/lenses.txt"

  f, err := os.Open(filename)
  if err != nil {
    t.Fatalf("Open %s file failed.\n", filename)
    return
  }

  r := bufio.NewReader(f)
  data := make([]*Feat, 0)
  for {
    line, _, err := r.ReadLine()
    if err != nil {
      break
    }
    s := strings.Split(string(line), " ")
    item := &Feat{}
    for i := 0; i < 4; i++ {
      item.feat = append(item.feat, s[i])
    }
    item.label = s[4]
    if len(s) > 5 {
      item.label += " " + s[5]
    }
    data = append(data, item)
  }

  tree := NewDTree()
  tree.Create(data)
  fmt.Println("Done")
}
