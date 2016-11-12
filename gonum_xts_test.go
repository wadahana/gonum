package gonum


import (
//    "fmt"
    "time"
    // "os"
    // "io"
    // "strings"
    // "strconv"
    // "encoding/csv"
     "testing"
    // "math"
)

var xts0_data = [] float64 {
    100.1, 100.2, 100.3, 100.4, 100.5,
    101.1, 101.2, 101.3, 101.4, 101.5,
    102.1, 102.2, 102.3, 102.4, 102.5,
    103.1, 102.2, 103.3, 103.4, 103.5,
};

var xts0_time = [] time.Time {
    time.Date(2001, 05, 01, 0, 0, 0, 0, time.Local),
    time.Date(2001, 03, 01, 0, 0, 0, 0, time.Local),
    time.Date(2001, 04, 01, 0, 0, 0, 0, time.Local),
    time.Date(2001, 01, 01, 0, 0, 0, 0, time.Local),
    time.Date(2001, 02, 01, 0, 0, 0, 0, time.Local),
}

var xts1_data = [] float64 {
    10.1, 10.2, 10.3, 10.4, 10.5, 10.6,
    11.1, 11.2, 11.3, 11.4, 11.5, 11.6,
    12.1, 12.2, 12.3, 12.4, 12.5, 12.6,
    13.1, 13.2, 13.3, 13.4, 13.5, 13.6,
    14.1, 14.2, 14.3, 14.4, 14.5, 14.6,
};

var xts1_time = [] time.Time {
    time.Date(2001, 05, 02, 0, 0, 0, 0, time.Local),
    time.Date(2001, 03, 01, 0, 0, 0, 0, time.Local),
    time.Date(2001, 04, 01, 0, 0, 0, 0, time.Local),
    time.Date(2001, 01, 02, 0, 0, 0, 0, time.Local),
    time.Date(2001, 02, 01, 0, 0, 0, 0, time.Local),
    time.Date(2001, 07, 01, 0, 0, 0, 0, time.Local),
}

var xts2_data = [] float64 {
    20.1, 20.2, 20.3, 20.4,
    21.1, 21.2, 21.3, 21.4,
    22.1, 22.2, 22.3, 22.4,
    23.1, 22.2, 23.3, 23.4,
};

var xts2_time = [] time.Time {
    time.Date(2001, 05, 02, 0, 0, 0, 0, time.Local),
    time.Date(2001, 03, 02, 0, 0, 0, 0, time.Local),
    time.Date(2001, 04, 02, 0, 0, 0, 0, time.Local),
    time.Date(2001, 01, 02, 0, 0, 0, 0, time.Local),
}

func Test_Xts(t *testing.T) {
    XtsDemo(t);
//    XtsTest(t);
}

func XtsTest (t *testing.T) {

}

func XtsDemo(t *testing.T) {

    var m0 *Matrix = nil;
    var x0 *Xts = nil;
    var err error = nil;

    t.Logf("Xts Setter/Getter Demo .....\n")
    m0, err = NewMatrixWithData(5, 4, xts0_data);
    if err == nil {
        x0, err = NewXts(xts0_time, m0);
    }
    if err != nil {
        t.Errorf("New Xts Test fail, Error: %v\n", err);
        return ;
    }
    x0.SetName("Col1", "Col2", "Col3", "Col4");
    t.Logf("x0:\n%v\n", x0);

    err = x0.Set(1,1, 1000.1);
    if err != nil {
        t.Errorf("Xts.Set Testfail, Error: %v\n", err);
        return;
    }
    v := x0.Get(1,1);
    if v != 1000.1 {
        t.Errorf("Xts.Set/Get Test fail\n");
        return;
    }

    t.Logf("Xts SetName()/GetNames() Demo .....\n")
    var m1 *Matrix = nil;
    var x1 *Xts = nil;

    m1, err = NewMatrixWithData(6, 5, xts1_data);
    if err == nil {
        x1, err = NewXts(xts1_time, m1);
    }
    if err != nil {
        t.Errorf("New Xts Test fail, Error: %v\n", err);
        return ;
    }
    err = x1.SetName("Col_A", "Col_B", "Col_C", "Col_D", "Col_E");
    if err != nil {
        t.Errorf("Xts SetName fail, Error: %v\n", err);
        return ;
    }
    t.Logf("x1:\n%v\n", x1);

    l_name := x1.GetName();
    var col_name string = "";
    for _, name := range l_name {
        col_name = col_name + ",  " + name ;
    }
    t.Logf("Colume Name: %v\n", col_name);

    t.Logf("Xts GetColumesByIndex()/GetColumesByName() Demo .....\n")
    sub1, err := x1.GetColumesByIndex(1,3,2,2);
    if err != nil {
        t.Errorf("Xts GetColumesByIndex() Test fail, Error: %v\n", err);
        return ;
    }
    t.Logf("\n%v\n", sub1);

    sub2, err := x1.GetColumesByName("Col_A", "Col_D", "Col_D", "Col_B");
    if err != nil {
        t.Errorf("Xts GetColumesByName() Test fail, Error: %v\n", err);
        return ;
    }
    t.Logf("\n%v\n", sub2);

    t.Logf("Xts CBind()/RBind() Demo .....\n")
    var m2 *Matrix = nil;
    var x2 *Xts = nil;

    m2, err = NewMatrixWithData(4, 4, xts2_data);
    if err == nil {
        x2, err = NewXts(xts2_time, m2);
    }
    if err != nil {
        t.Errorf("New Xts Test fail, Error: %v\n", err);
        return ;
    }
    x2.SetName("Col_I", "Col_II", "Col_III", "Col_IV");
    t.Logf("x2:\n%v\n", x2);

    cbx, err := x0.CBind(x1)
    if err != nil {
        t.Errorf("Xts.CBind Test fail, Error: %v\n", err);
        return ;
    }
    t.Logf("x0 colume bind with x1: \n%v\n", cbx);

    rbx, err := x0.RBind(x2)
    if err != nil {
        t.Errorf("Xts.RBind Test fail, Error: %v\n", err);
        return ;
    }
    t.Logf("x0 row bind with x2: \n%v\n", rbx);



}
/*
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
    t.Logf("default names:\n%v\n", names);

    xts.SetNames("First", "Second", "Third");
    names = xts.GetNames();
    if strings.Compare(names[0], "First") != 0 ||
        strings.Compare(names[1], "Second") != 0 ||
        strings.Compare(names[2], "Third") != 0 {
        t.Errorf("xts setname / getname test fail.");
    }
    t.Logf("names:\n%v\n", names);

    xts1 := NewXtsWithColumes(times, col1, col2);
    t.Logf("\n%v[%d,%d]\n", xts1, xts1.GetRowNum(), xts1.GetColumeNum());

    data1 := xts1.GetColumeData(0);
    for i, v := range data1 {
        if int(v) != i + 1 {
            t.Errorf("NewXtsWithColumes not sort by time.");
            return;
        }
    }
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
*/
/*
    m<-c(1,2,3,4,5,6,7,8,9,10,112.1,933.1,403.5,334.4,677.1,500.6,253.3,744.3,801.8,108.2,10,8,3,2,5,4,1,6,7,9)
*/
    /*
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
*/
