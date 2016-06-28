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
    "math"
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

    names := xts.GetNames();
    if strings.Compare(names[0], "Colume.1") != 0 ||
        strings.Compare(names[1], "Colume.2") != 0 ||
        strings.Compare(names[2], "Colume.3") != 0 {
        t.Errorf("xts default colume name test fail.");
    }
    fmt.Printf("default names:\n%v\n", names);

    xts.SetNames("First", "Second", "Third");
    names = xts.GetNames();
    if strings.Compare(names[0], "First") != 0 ||
        strings.Compare(names[1], "Second") != 0 ||
        strings.Compare(names[2], "Third") != 0 {
        t.Errorf("xts setname / getname test fail.");
    }
    fmt.Printf("names:\n%v\n", names);
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

/*
    m<-c(1,2,3,4,5,6,7,8,9,10,112.1,933.1,403.5,334.4,677.1,500.6,253.3,744.3,801.8,108.2,10,8,3,2,5,4,1,6,7,9)
*/
func Test_Xts_Math(t *testing.T) {
    m := testcase_NewMatrix();
    times := make([]time.Time, 0);
    for _, s := range dates {
        d , _ := time.Parse("2006-01-02", s);
        times = append(times, d);
    }
    xts := NewXts(times, m);

    d1 := xts.Diff(1);
    d3 := xts.Diff(3);

    t.Logf("\n%v[%d,%d]\n", d1, d1.GetRowNum(), d1.GetColumeNum());
    t.Logf("\n%v[%d,%d]\n", d3, d3.GetRowNum(), d3.GetColumeNum());

    d1_ok := [30]float64 {0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0, 821.0, -529.6, -69.1, 342.7, -176.5, -247.3, 491.0, 57.5, -693.6, 0, -2, -5, -1, 3, -1, -3, 5, 1, 2};
    d1_ok[0] = math.NaN();
    d1_ok[10] = math.NaN();
    d1_ok[20] = math.NaN();
    data1 := d1.GetData();

    for i, v := range data1 {
        if math.Abs(d1_ok[i] - v) > 0.001 {
            t.Errorf("after xts.Diff, data error");
            return;
        }
    }

    d3_ok := [30]float64 {0, 0, 0, 3, 3, 3, 3, 3, 3, 3, 0, 0, 0, 222.3, -256.0, 97.1, -81.1, 67.2, 301.2, -145.1, 0, 0, 0, -8, -3, 1, -1, 1, 3, 8};
    for i := 0; i < 3; i++ {
        d3_ok[0 + i] = math.NaN();
        d3_ok[10 + i] = math.NaN();
        d3_ok[20 + i] = math.NaN();
    }
    data3 := d3.GetData();

    for i, v := range data3 {
        if math.Abs(d3_ok[i] - v) > 0.001 {
            t.Errorf("after xts.Diff, data error");
            return;
        }
    }
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
