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

func Test_NewXts(t *testing.T) {
    m := testcase_NewMatrix();
    times := make([]time.Time, 0);
    for _, s := range dates {
        d , _ := time.Parse("2006-01-02", s);
        times = append(times, d);
    }
    xts := NewXts(times, m);
    data := xts.GetColumeData(0);
    for i, v := range data {
        if int(v) != i + 1 {
            t.Errorf("NewXts not sort by time.");
            return;
        }
    }
    t.Logf("\n%v[%d,%d]\n", xts, xts.GetRowNum(), xts.GetColumeNum());
}

func Test_Xts_RBind(t *testing.T) {
    l_sub1 := []int{1,3,5,7,9};
    l_sub2 := []int {2,4,6,8,0};
    l_sub3 := []int {1,4,6,8,0};

    data1 := make([]float64, 5*3, 5*3);
    time1 := make([]time.Time, 5, 5);
    for i,ii := range l_sub1 {
        data1[i + 0] = col1[ii];
        data1[i + 5] = col2[ii];
        data1[i + 10] = col3[ii];
        time1[i], _ = time.Parse("2006-01-02", dates[ii]);
    }
    m1 := NewMatrix(5,3,data1);
    xts1 := NewXts(time1, m1);

    data2 := make([]float64, 5*3, 5*3);
    time2 := make([]time.Time, 5, 5);
    for i,ii := range l_sub2 {
        data2[i + 0] = col1[ii];
        data2[i + 5] = col2[ii];
        data2[i + 10] = col3[ii];
        time2[i], _ = time.Parse("2006-01-02", dates[ii]);
    }
    m2 := NewMatrix(5,3,data2);
    xts2 := NewXts(time2, m2);

    data3 := make([]float64, 5*3, 5*3);
    time3 := make([]time.Time, 5, 5);
    for i,ii := range l_sub3 {
        data3[i + 0] = col1[ii];
        data3[i + 5] = col2[ii];
        data3[i + 10] = col3[ii];
        time3[i], _ = time.Parse("2006-01-02", dates[ii]);
    }
    m3 := NewMatrix(5,3,data3);
    xts3 := NewXts(time3, m3);

    xts4 := xts2.RBind(xts1);
    xts5 := xts1.RBind(xts3);

    t.Logf("\n%v[%d,%d]\n",xts1, xts1.GetRowNum(), xts1.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n",xts2, xts2.GetRowNum(), xts2.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n",xts3, xts3.GetRowNum(), xts3.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n",xts4, xts4.GetRowNum(), xts4.GetColumeNum());
    if xts5 != nil {
        t.Errorf("two xts have same index data could not cbind..");
        return;
    } else {
        t.Logf("\n two xts have same index data could not cbind.");
    }
    return;
}

func Test_Xts_CBind(t *testing.T) {
    l_sub1 := []int{1,3,5,7,9};
    l_sub2 := []int {2,4,6,8,0};
    l_sub3 := []int {1,4,5,8,0};

    data1 := make([]float64, 5, 5);
    time1 := make([]time.Time, 5, 5);
    for i,ii := range l_sub1 {
        data1[i + 0] = col1[ii];
        time1[i], _ = time.Parse("2006-01-02", dates[ii]);
    }
    m1 := NewMatrix(5,1,data1);
    xts1 := NewXts(time1, m1);

    data2 := make([]float64, 5, 5);
    time2 := make([]time.Time, 5, 5);
    for i,ii := range l_sub2 {
        data2[i + 0] = col1[ii];
        time2[i], _ = time.Parse("2006-01-02", dates[ii]);
    }
    m2 := NewMatrix(5,1,data2);
    xts2 := NewXts(time2, m2);

    data3 := make([]float64, 5, 5);
    time3 := make([]time.Time, 5, 5);
    for i,ii := range l_sub3 {
        data3[i + 0] = col1[ii];
        time3[i], _ = time.Parse("2006-01-02", dates[ii]);
    }
    m3 := NewMatrix(5,1,data3);
    xts3 := NewXts(time3, m3);

    xts4 := xts1.CBind(xts2);
    xts5 := xts1.CBind(xts3);

    t.Logf("\n%v[%d,%d]\n", xts1, xts1.GetRowNum(), xts1.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n", xts2, xts2.GetRowNum(), xts2.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n", xts3, xts3.GetRowNum(), xts3.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n", xts4, xts4.GetRowNum(), xts4.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n", xts5, xts5.GetRowNum(), xts5.GetColumeNum());
}
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
//    xts1 := newXtsFromCSV("./sh600688.csv");
//    xts2 := newXtsFromCSV("./sh601988.csv");

//    fmt.Printf("xts1: \n%v", xts1);
//    fmt.Printf("xts2: \n%v", xts2);

//    fmt.Printf("%d, %d\n", xts1.GetRowNum(), xts2.GetRowNum());

//    sub1 := xts1.GetColumes([]int{1,2,3});
//    fmt.Printf("sub1: \n%v", sub1);
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
