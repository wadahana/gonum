package gonum

import (
    "fmt"
    "time"
    "os"
    "io"
    "strings"
    "strconv"
    "encoding/csv"
    "testing"
)

// func Test_Xts(t *testing.T) {
//     ll := make([]time.Time, 0);
//     data := make([]float64, 0);
//     fmt.Printf("Raw:\n");
//
//     for i := 0; i < 10; i++ {
//         data = append(data, col1[i]);
//         data = append(data, col2[i]);
//         data = append(data, col3[i]);
//
//         date, _ := time.Parse("2006-01-02", dates[i]);
//         ll = append(ll, date);
//
//         fmt.Printf("%v, %0.4f, %0.4f, %0.4f\n", date.Format("2006-01-02 15:04:05"), col1[i], col2[i], col3[i]);
//     }
//     m := NewMatrix(10,3,data);
//
//     xts := NewXts(ll, m);
//     fmt.Printf("Xts:\n%v\n", xts);
//
//     names := xts.GetNames();
//     fmt.Printf("names:\n%v\n", names);
//     xts.SetNames("One", "Two");
//     names = xts.GetNames();
//     fmt.Printf("names:\n%v\n", names);
// }


func Test_Xts(t *testing.T) {
    xts1 := newXtsFromCSV("./sh600688.csv");
    xts2 := newXtsFromCSV("./sh601988.csv");

//    fmt.Printf("xts1: \n%v", xts1);
//    fmt.Printf("xts2: \n%v", xts2);

    fmt.Printf("%d, %d\n", xts1.GetRowNum(), xts2.GetRowNum());

    sub1 := xts1.GetColumes([]int{1,2,3});
    fmt.Printf("sub1: \n%v", sub1);
}

func newXtsFromCSV(filename string) *Xts {
    inp, err := os.Open(filename);
    if err != nil {
        panic(err);
    }
    defer inp.Close();

    reader := csv.NewReader(inp);
    reader.Comment = '#';
    reader.Comma = ',';
    reader.TrimLeadingSpace = true;
    reader.FieldsPerRecord = 8;

    ts := make([]time.Time, 0);
    data := make([]float64, 0);
    names := make([]string, 0);

    // first line is name of colume;
    record, err := reader.Read();
    if err == nil {
        for _, n := range record {
            names = append(names, n);
        }
    }

    // read time and data
    for {
        record, err := reader.Read();
        if err == io.EOF {
            fmt.Println("io.EOF\n");
            break;
        } else if err != nil {
            fmt.Printf("Error: %v\n", err);
            continue;
        }
        date, _ := time.Parse("2006-01-02", strings.TrimSpace(record[0]));
        ts = append(ts, date);
        for i := 0; i < 7; i++ {
            v, err := strconv.ParseFloat(strings.TrimSpace(record[i + 1]), 64);
            if err != nil {
                return nil
            }
            data = append(data, v);
        }
    }

    if len(data) != len(ts) * 7 {
        return nil
    }

    m := NewMatrix(len(ts), 7, data);
    xts := NewXts(ts, m);
    xts.SetNames(names...);
    return xts;
}
