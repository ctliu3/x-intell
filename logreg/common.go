package logreg

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

type RowLabeledDate struct {
  Feat []float64
  Label int
}

// Read training data from text file.
func ReadData(file string) ([]*RowLabeledDate, error) {
  f, err := os.Open(file)
  if err != nil {
    return nil, fmt.Errorf("Cannot open file %s\n", f)
  }
  defer f.Close()

  r := bufio.NewReader(f)
  var line []byte
  var dataset []*RowLabeledDate
  c := 0
  for {
    if line, err = r.ReadSlice('\n'); err != nil {
      break
    }

    c++
    var item []string = strings.Split(string(line), " ")
    var rowdata RowLabeledDate
    for i := 0; i < len(item) - 1; i++ {
      if x, err := strconv.ParseFloat(item[i], 64); err == nil {
        rowdata.Feat = append(rowdata.Feat, x)
      } else {
        return nil, err
      }
    }

    ystr := item[len(item) - 1]
    ystr = strings.Trim(ystr, "\n")
    ystr = strings.Trim(ystr, "\r")
    if y, err := strconv.ParseFloat(ystr, 64); err == nil {
      rowdata.Label = int(y)
    } else {
      return nil, err
    }
    dataset = append(dataset, &rowdata)
  }

  return dataset, nil
}
