package logreg

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "strings"
  "strconv"
  "testing"
)

func TestLR(t *testing.T) {
  log.SetFlags(log.LstdFlags | log.Lshortfile)

  lr, err := NewLogReg("data/horseColicTraining.txt")
  if err != nil {
    log.Fatalf(err.Error())
    return
  }
  fmt.Println("Start training...")
  lr.Train()
  fmt.Println("Training done...\n")

  fmt.Println("Start testing...")

  f, err := os.Open("data/horseColicTest.txt")
  if err != nil {
    log.Fatalf(err.Error())
    return
  }
  defer f.Close()

  r := bufio.NewReader(f)
  var line []byte
  ntest := 0
  nacc := 0
  for {
    if line, err = r.ReadSlice('\n'); err != nil {
      break
    }

    ntest++
    var feat []float64
    var item []string = strings.Split(string(line), " ")
    for i := 0; i < len(item) - 1; i++ {
      if x, err := strconv.ParseFloat(item[i], 64); err == nil {
        feat = append(feat, x)
      } else {
        log.Fatalf(err.Error())
      }
    }

    ystr := item[len(item) - 1]
    ystr = strings.Trim(ystr, "\n")
    ystr = strings.Trim(ystr, "\r")
    if y, err := strconv.ParseFloat(ystr, 64); err == nil {
      if int(y) == lr.Predict(feat) {
        nacc++
      }
    } else {
      log.Fatalf(err.Error())
    }
  }

  fmt.Printf("ntest = %d, nacc = %d, acc = %.1f%%\n", ntest, nacc, 100.0 * float64(nacc) / float64(ntest))
  fmt.Println("Testing done...\n")
}
