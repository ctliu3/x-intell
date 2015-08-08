package ap

import (
  "testing"
  "log"
  "os"
  "bufio"
  "fmt"
  "strings"
  "strconv"
)

func TestAP(t *testing.T) {
  log.SetFlags(log.LstdFlags | log.Lshortfile)
  ap := new(AP)

  var file string = "data/testSet.txt"
  f, err := os.Open(file)
  if err != nil {
    log.Fatalf("Can not open file %d\n", file)
    return
  }
  defer f.Close()

  var feats []*Feat
  scanner := bufio.NewScanner(f)
  scanner.Split(bufio.ScanLines)
  for scanner.Scan() {
    var items []string = strings.Split(scanner.Text(), " ")
    var feat Feat
    for _, item := range items {
      val, err := strconv.ParseFloat(string(item), 64)
      if err != nil {
        log.Fatalf("Can not parse string %s\n", string(item))
        return
      }
      feat.x = append(feat.x, val)
    }
    feats = append(feats, &feat)
  }
  fmt.Printf("There are %d features.\n", len(feats))

  fmt.Println("Start clustring ...")
  var centers []int = ap.cluster(feats)
  fmt.Println("Clustring done ...\n")

  fmt.Printf("#Cluster: %d\n", len(centers))
  fmt.Println(centers)
}
