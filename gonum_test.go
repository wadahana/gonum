package gonum

import (
    "fmt"
    "time"
    "testing"
)


var dates = [10]string {
    "2015-10-01", //7
    "2014-01-01", //4
    "2012-11-04", //3
    "2015-07-05", //6
    "2014-01-31", //5
    "2015-10-12", //8
    "2016-01-18", //9
    "2011-04-04", //2
    "2016-01-28", //10
    "2010-12-31", //1
};

var col1 = [] float64 {
    7.0,
    4.0,
    3.0,
    6.0,
    5.0,
    8.0,
    9.0,
    2.0,
    10.0,
    1.0,
};

var col2 = [] float64 {
    253.3,
    334.4,
    43.5,
    50.6,
    6.17,
    74.3,
    8.8,
    9.901,
    10.2334,
    11004.1234,
};

var col3 = [] float64 {
    1.3,
    2.4,
    3.5,
    4.6,
    5.17,
    6.3,
    7.8,
    8.901,
    9.2334,
    10.1234,
};

func Test_Matrix(t *testing.T) {
    data := make([]float64, 0);
    for i := 0; i < 10; i++ {
        data = append(data, col1[i]);
        data = append(data, col2[i]);
        data = append(data, col3[i]);
    }
    m := NewMatrix(10,3,data);
    fmt.Printf("Raw:\n");
    fmt.Printf("%v\n", m);
    m.SwapRow(1,9);
    fmt.Printf("After Swap(1,9):\n");
    fmt.Printf("%v\n", m);
}

func Test_Xts(t *testing.T) {
    ll := make([]time.Time, 0);
    data := make([]float64, 0);
    fmt.Printf("Raw:\n");
    for i := 0; i < 10; i++ {
        data = append(data, col1[i]);
        data = append(data, col2[i]);
        data = append(data, col3[i]);

        date, _ := time.Parse("2006-01-02", dates[i]);
        ll = append(ll, date);

        fmt.Printf("%v, %0.4f, %0.4f, %0.4f\n", date.Format("2006-01-02 15:04:05"), col1[i], col2[i], col3[i]);
    }
    m := NewMatrix(10,3,data);

    xts := NewXts(ll, m);
    fmt.Printf("Xts:\n%v\n", xts);

    names := xts.GetNames();
    fmt.Printf("names:\n%v\n", names);
    xts.SetNames("One", "Two");
    names = xts.GetNames();
    fmt.Printf("names:\n%v\n", names);
}
